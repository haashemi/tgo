package tgo

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"syscall"
	"time"
)

type Bot struct {
	*User          // embedding all bot information directly to the Bot
	*API           // embedding the api to add all api methods to the bot
	*DefaultRouter // embedding a default router to the Bot

	defaultParseMode ParseMode

	asks   map[int64]chan<- Context
	askMut sync.RWMutex

	routers []Router

	// contains user-ids with their session
	sessions sync.Map
}

type Options struct {
	Host             string
	Client           *http.Client
	DefaultParseMode ParseMode
}

func NewBot(token string, opts Options) (bot *Bot, err error) {
	api := NewAPI(token, opts.Host, opts.Client)
	me, err := api.GetMe()
	if err != nil {
		return nil, err
	}

	defaultRouter := NewDefaultRouter()

	bot = &Bot{
		User:          me,
		API:           api,
		DefaultRouter: defaultRouter,

		defaultParseMode: opts.DefaultParseMode,

		asks: make(map[int64]chan<- Context),

		routers: []Router{defaultRouter},
	}

	return bot, nil
}

// GetSession returns the stored session as a sync.Map.
// it creates a new session if session id didn't exists.
func (bot *Bot) GetSession(sessionID int64) *sync.Map {
	result, ok := bot.sessions.Load(sessionID)
	if ok {
		return result.(*sync.Map)
	}

	session := &sync.Map{}
	bot.sessions.Store(sessionID, session)
	return session
}

func (bot *Bot) AddRouter(router Router) error {
	if err := router.Setup(bot); err != nil {
		return err
	}

	bot.routers = append(bot.routers, router)
	return nil
}

func (bot *Bot) StartPolling() error {
	var offset int64

	for {
		data, err := bot.GetUpdates(&GetUpdates{
			// ToDo: I have no idea if I'm getting the offset in the right way or not. But it works!
			Offset: offset,

			// ToDo: decreasing a second is kinda risky... what if the timeout be a second?... 0?
			Timeout: int64(bot.client.Timeout.Seconds()) - 1,

			// ToDo: remove this line after supporting all update types.
			AllowedUpdates: []string{"message", "edited_message", "channel_post", "edited_channel_post", "callback_query"},
		})
		if err != nil {
			if errors.Is(err, syscall.ECONNRESET) {
				time.Sleep(time.Second / 2)
				continue
			}
			return err
		}

		for _, update := range data {
			offset = update.UpdateId + 1

			go func(update *Update) {
				ctx := NewContext(update, bot)

				if update.Message != nil && bot.sendAnswerIfAsked(ctx) {
					return
				}

				for _, router := range bot.routers {
					if used := router.HandleUpdate(ctx); used {
						return
					}
				}
			}(update)
		}
	}
}

func (bot *Bot) sendAnswerIfAsked(ctx Context) (sent bool) {
	bot.askMut.RLock()
	receiver, ok := bot.asks[ctx.ChatID()]
	bot.askMut.RUnlock()

	if ok {
		receiver <- ctx
		return true
	}

	return false
}

func (bot *Bot) waitForAnswer(question *Message, timeout time.Duration) (Context, error) {
	chatID := question.ChatID()
	waiter := make(chan Context, 1)

	bot.askMut.Lock()
	bot.asks[chatID] = waiter
	bot.askMut.Unlock()

	defer func() {
		bot.askMut.Lock()
		delete(bot.asks, chatID)
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
