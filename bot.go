package tgo

import (
	"fmt"
	"net/http"
	"sync"
)

type Bot struct {
	*Party

	api *API
	me  *User

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
		Party:    &Party{},
		api:      api,
		me:       me,
		asks:     make(map[string]chan<- *Context),
		sessions: make(map[int64]*Session),
	}

	return bot, nil
}

func (bot *Bot) Me() *User {
	return bot.me
}

func (bot *Bot) StartPolling() error {
	var offset int64

	for {
		data, err := bot.api.GetUpdates(GetUpdatesParams{
			Offset:  offset,
			Timeout: bot.api.GetHTTPTimeout() - 1,
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
					if receiver, ok := bot.asks[uid]; ok {
						receiver <- ctx
					}
					bot.askMut.RUnlock()

					return
				}

				bot.handleUpdate(ctx)
			}(update)
		}
	}
}
