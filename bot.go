package tgo

import (
	"fmt"
	"net/http"
	"sync"
)

type Bot struct {
	*API
	*User
	*Party

	askMut sync.RWMutex
	asks   map[string]chan<- *Context

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
		API:      api,
		User:     me,
		Party:    &Party{},
		asks:     make(map[string]chan<- *Context),
		sessions: make(map[int64]*Session),
	}

	return bot, nil
}

func (bot *Bot) StartPolling() error {
	var offset int64

	for {
		data, err := bot.GetUpdates(&GetUpdatesOptions{
			Offset:  offset,
			Timeout: bot.GetHTTPTimeout() - 1,
		})
		if err != nil {
			return err
		}

		for _, update := range data {
			offset = update.UpdateId + 1

			ctx := newContext(bot, update)

			go func(update *Update) {
				if update.Message != nil {
					uid := fmt.Sprintf("%d-%d", ctx.ChatID(), ctx.SenderChatID())

					bot.askMut.RLock()
					receiver, ok := bot.asks[uid]
					bot.askMut.RUnlock()

					if ok {
						receiver <- ctx
						return
					}
				}

				bot.handleUpdate(ctx)
			}(update)
		}
	}
}
