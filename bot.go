package tgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
)

//go:generate go run ./internal/codegen

const TelegramHost = "https://api.telegram.org"

type Bot struct {
	*User  // embedding all bot information directly to the Bot
	*party // embedding all party methods directly to the Bot

	url    string
	client *http.Client

	asks   map[string]chan<- Context
	askMut sync.RWMutex

	// contains user-ids with their session
	sessions sync.Map
}

type Options struct {
	Host   string
	Client *http.Client
}

func NewBot(token string, opts Options) (bot *Bot, err error) {
	if opts.Host == "" {
		opts.Host = TelegramHost
	}

	if opts.Client == nil {
		opts.Client = &http.Client{Timeout: 30 * time.Second}
	}

	bot = &Bot{
		party: &party{},

		url:    opts.Host + "/bot" + token + "/",
		client: opts.Client,

		asks: make(map[string]chan<- Context),
	}

	bot.User, err = bot.GetMe()

	return bot, err
}

func (bot *Bot) StartPolling() error {
	var offset int64

	for {
		data, err := bot.GetUpdates(&GetUpdatesOptions{
			// ToDo: I have no idea if I'm getting the offset in the right way or not.
			// But it works!
			Offset: offset,

			// ToDo: decreasing a second is kinda risky... what if the timeout be a second?... 0?
			Timeout: int64(bot.client.Timeout.Seconds()) - 1,

			// ToDo: support all type of updates, then remove this line.
			//
			// remaining:
			// 	inline_query, chosen_inline_result, shipping_query, pre_checkout_query, poll, poll_answer, my_chat_member, chat_member, chat_join_request
			AllowedUpdates: []string{"message", "edited_message", "channel_post", "edited_channel_post", "callback_query"},
		})
		if err != nil {
			return err
		}

		for _, update := range data {
			offset = update.UpdateId + 1

			go func(update *Update) {
				ctx := &botContext{bot: bot, update: update}

				switch {
				case update.Message != nil:
					ctx.Contextable = update.Message
					if bot.sendAnswerIfAsked(ctx) {
						return
					}
				case update.EditedMessage != nil:
					ctx.Contextable = update.EditedMessage
				case update.ChannelPost != nil:
					ctx.Contextable = update.ChannelPost
				case update.EditedChannelPost != nil:
					ctx.Contextable = update.EditedChannelPost
				case update.CallbackQuery != nil:
					ctx.Contextable = update.CallbackQuery
				}

				bot.handleUpdate(ctx)
			}(update)
		}
	}
}

// GetSession returns the stored session as a sync.Map.
// it creates a new session if session id didn't exists.
func (bot *Bot) GetSession(sessionID int64) *sync.Map {
	result, ok := bot.sessions.Load(sessionID)
	if !ok {
		return result.(*sync.Map)
	}

	session := &sync.Map{}
	bot.sessions.Store(sessionID, session)
	return session
}

func (bot *Bot) sendAnswerIfAsked(ctx Context) (sent bool) {
	bot.askMut.RLock()
	uid := fmt.Sprintf("%d-%d", ctx.ChatID(), ctx.SenderID())
	receiver, ok := bot.asks[uid]
	bot.askMut.RUnlock()

	if ok {
		receiver <- ctx
		return true
	}

	return false
}

func (bot *Bot) waitForAnswer(question *Message, timeout time.Duration) (Context, error) {
	uid := fmt.Sprintf("%d-%d", question.ChatID(), question.SenderID())
	waiter := make(chan Context, 1)

	bot.askMut.Lock()
	bot.asks[uid] = waiter
	bot.askMut.Unlock()

	defer func() {
		bot.askMut.Lock()
		delete(bot.asks, uid)
		bot.askMut.Unlock()

		close(waiter)
	}()

	aCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case answer := <-waiter:
		return answer, nil

	case <-aCtx.Done():
		return nil, aCtx.Err()
	}
}

type ChatID string

func NewChatID(id any) ChatID {
	if val, ok := id.(string); ok {
		return ChatID(val)
	}

	return ChatID(fmt.Sprint(id))
}

type ParseMode string

const (
	ParseModeNone       ParseMode = ""
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
)

type PollType string

const (
	PollTypeAny     PollType = ""        // If this gets passed, the user will be allowed to create a poll of any type.
	PollTypeQuiz    PollType = "quiz"    // if this gets passed, the user will be allowed to create only polls in the quiz mode.
	PollTypeRegular PollType = "regular" // If this gets passed, only regular polls will be allowed.
)

type httpResponse[T any] struct {
	OK     bool `json:"ok"`
	Result T    `json:"result,omitempty"`
	*Error
}

type multipartForm interface{ HasUploadable() bool }

func doHTTP[T any](client *http.Client, baseURL, method string, rawData any) (data T, err error) {
	var url = baseURL + method

	var resp *http.Response

	if rawData == nil {
		if resp, err = client.Get(url); err != nil {
			return
		}
	} else if body, ok := rawData.(multipartForm); ok && body.HasUploadable() {
		r, w := io.Pipe()
		defer r.Close()

		m := multipart.NewWriter(w)

		go func() {
			defer w.Close()
			defer m.Close()

			params, files := getParamsAndFiles(body)
			for key, val := range params {
				m.WriteField(key, val)
			}

			for key, file := range files {
				ww, err := m.CreateFormFile(key, file.Name)
				if err != nil {
					w.CloseWithError(err)
					return
				} else if _, err = io.Copy(ww, file.Reader); err != nil {
					w.CloseWithError(err)
					return
				}
			}
		}()

		if resp, err = client.Post(url, m.FormDataContentType(), r); err != nil {
			return
		}
	} else {
		body := bytes.NewBuffer(nil)
		if err = json.NewEncoder(body).Encode(rawData); err != nil {
			return
		}

		if resp, err = client.Post(url, "application/json", body); err != nil {
			return
		}
	}

	defer resp.Body.Close()

	response := &httpResponse[T]{}
	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return
	} else if !response.OK {
		err = response.Error
		return
	}
	return response.Result, nil
}

func getParamsAndFiles(d any) (params map[string]string, files map[string]*InputFileUploadable) {
	params = make(map[string]string)
	files = make(map[string]*InputFileUploadable)

	v := reflect.ValueOf(d).Elem()
	vType := reflect.TypeOf(d).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fType := vType.Field(i)

		if field.IsZero() {
			continue
		}

		data := field.Interface()
		tag := strings.TrimSuffix(fType.Tag.Get("json"), ",omitempty")

		if tag == "" && fType.Anonymous {
			newParams, newFiles := getParamsAndFiles(data)

			for k, v := range newParams {
				params[k] = v
			}
			for k, v := range newFiles {
				files[k] = v
			}

			continue
		}

		if xx, ok := data.(InputFile); ok {
			if uplodable, ok := xx.(*InputFileUploadable); ok {
				files[tag] = uplodable
			} else {
				params[tag] = string(xx.(InputFileNotUploadable))
			}
		} else if kind := field.Type().Kind(); kind == reflect.Struct || kind == reflect.Interface {
			raw, _ := json.Marshal(data)
			params[tag] = string(raw)
		} else {
			params[tag] = fmt.Sprint(data)
		}
	}

	return params, files
}
