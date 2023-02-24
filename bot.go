package tgo

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Bot struct {
	// ToDo: Maybe remove API and using Bot directly?

	*API   // embedding all API methods directly to the Bot (as they are customized enough)
	*User  // embedding all bot information directly to the Bot
	*party // embedding all party methods directly to the Bot

	asks   map[string]chan<- MessageContext
	askMut sync.RWMutex

	// this contains user-ids with their session
	sessions    map[int64]*Session
	sessionsMut sync.Mutex
}

type Options struct {
	Host   string
	Client *http.Client
}

func NewBot(token string, opts Options) (*Bot, error) {
	api := NewAPI(token, opts.Host, opts.Client)

	me, err := api.GetMe()
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		API:      api,
		User:     me,
		party:    &party{},
		asks:     make(map[string]chan<- MessageContext),
		sessions: make(map[int64]*Session),
	}

	return bot, nil
}

func (bot *Bot) StartPolling() error {
	var offset int64

	for {
		data, err := bot.GetUpdates(&GetUpdatesOptions{
			// ToDo: I have no idea if I'm getting the offset in the right way or not.
			Offset: offset,

			// ToDo: decreasing a second is kinda risky... what if the timeout be a second?... 0?
			Timeout: bot.TimeoutSeconds() - 1,

			// ToDo: support all type of updates, then remove this line.
			AllowedUpdates: []string{"message", "callback_query"},
		})
		if err != nil {
			return err
		}

		for _, update := range data {
			offset = update.UpdateId + 1

			go func(update *Update) {
				baseCtx := &Context{bot: bot, update: update}

				switch {
				case update.Message != nil:
					baseCtx.Contextable = update.Message
					ctx := &messageContext{Context: baseCtx}

					if bot.sendAnswerIfAsked(ctx) {
						return
					}
					bot.handleOnMessage(ctx)

				// ToDo: Support this
				// case update.EditedMessage != nil:
				// 	baseCtx.Contextable = update.EditedMessage

				// ToDo: Support this
				// case update.ChannelPost != nil:
				// 	baseCtx.Contextable = update.ChannelPost

				// ToDo: Support this
				// case update.EditedChannelPost != nil:
				// 	baseCtx.Contextable = update.EditedChannelPost

				case update.CallbackQuery != nil:
					baseCtx.Contextable = update.CallbackQuery
					bot.handleOnCallbackQuery(&callbackContext{Context: baseCtx})
				}
			}(update)
		}
	}
}

func (bot *Bot) sendAnswerIfAsked(ctx *messageContext) (sent bool) {
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

func (bot *Bot) waitForAnswer(question *Message, timeout time.Duration) (*messageContext, error) {
	uid := fmt.Sprintf("%d-%d", question.ChatID(), question.SenderID())
	waiter := make(chan MessageContext, 1)

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
