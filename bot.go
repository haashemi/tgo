package tgo

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Bot struct {
	*API
	*User
	*MessageParty

	askMut sync.RWMutex
	asks   map[string]chan<- MessageContext

	sessionsMut sync.Mutex
	sessions    map[int64]*Session
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
		API:          api,
		User:         me,
		MessageParty: &MessageParty{onMessage: make(map[Filter][]MessageHandler)},
		asks:         make(map[string]chan<- MessageContext),
		sessions:     make(map[int64]*Session),
	}

	return bot, nil
}

func (bot *Bot) StartPolling() error {
	var offset int64

	for {
		data, err := bot.GetUpdates(&GetUpdatesOptions{
			Offset:  offset,
			Timeout: bot.TimeoutSeconds() - 1,
		})
		if err != nil {
			return err
		}

		for _, update := range data {
			offset = update.UpdateId + 1

			go func(update *Update) {
				ctx := newContext(bot, update)

				switch ctx.UpdateType() {

				case "Message":
					bot.askMut.RLock()
					uid := fmt.Sprintf("%d-%d", ctx.ChatID(), ctx.SenderID())
					receiver, ok := bot.asks[uid]
					bot.askMut.RUnlock()

					if ok {
						receiver <- ctx.(MessageContext)
						return
					}
					bot.handleOnMessage(ctx.(MessageContext), update)

				case "EditedMessage":
					return
				case "ChannelPost":
					return
				case "EditedChannelPost":
					return
				case "CallbackQuery":
					return
				}
			}(update)
		}
	}
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
