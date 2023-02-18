package tgo

import (
	"net/http"

	"github.com/haashemi/tgo/botapi"
)

type Bot struct {
	*Party

	api      *botapi.API
	sessions *botSession

	me *botapi.User
}

type Options struct {
	Host   string
	Client *http.Client
}

func NewBot(token string, opts Options) (*Bot, error) {
	bot := &Bot{
		Party:    &Party{},
		api:      botapi.NewAPI(token, opts.Host, opts.Client),
		sessions: newBotSession(),
	}

	me, err := bot.api.GetMe()
	if err != nil {
		return nil, err
	}
	bot.me = me

	return bot, nil
}

func (bot *Bot) Me() *botapi.User {
	return bot.me
}

func (bot *Bot) StartPolling() error {
	var offset int64

	for {
		data, err := bot.api.GetUpdates(botapi.GetUpdatesParams{
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
