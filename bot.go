package tgo

import (
	"net/http"

	"github.com/haashemi/tgo/botapi"
)

type Bot struct {
	*Party
	*botSession

	api *botapi.API

	me *botapi.User
}

type Options struct {
	Host   string
	Client *http.Client
}

func NewBot(token string, opts Options) (*Bot, error) {
	api := botapi.NewAPI(token, opts.Host, opts.Client)

	me, err := api.GetMe()
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		Party:      &Party{},
		botSession: newBotSession(),
		api:        api,
		me:         me,
	}

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
