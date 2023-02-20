package tgo

import (
	"net/http"
	"sync"
)

type Bot struct {
	*Party

	me  *User
	api *API

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

			go bot.handleUpdate(newContext(bot, update))
		}
	}
}
