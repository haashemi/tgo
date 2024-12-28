package tgo

import (
	"errors"
	"net/http"
	"sync"
	"syscall"

	"github.com/haashemi/tgo/tg"
)

type Bot struct {
	api *tg.API
	Router

	// Default parse mode to be used on bot-specific methods.
	parseMode tg.ParseMode

	// Contains user-ids with their session
	sessions sync.Map
}

type Options struct {
	Host             string
	Client           *http.Client
	DefaultParseMode tg.ParseMode
}

func NewBot(token string, opts Options) (bot *Bot) {
	return &Bot{
		api:       tg.NewAPI(token, opts.Host, opts.Client),
		parseMode: opts.DefaultParseMode,
	}
}

// API returns the Bot's API instance.
func (bot *Bot) API() *tg.API {
	return bot.api
}

// GetSession returns the stored session as a sync.Map.
// it creates a new session if session id didn't exist.
func (bot *Bot) GetSession(sessionID int64) *sync.Map {
	result, ok := bot.sessions.Load(sessionID)
	if ok {
		return result.(*sync.Map)
	}

	session := &sync.Map{}
	bot.sessions.Store(sessionID, session)
	return session
}

// Send sends a message with the preferred ParseMode.
func (bot *Bot) Send(msg Sendable) (*tg.Message, error) {
	if x, ok := msg.(ParseModeSettable); ok {
		if x.GetParseMode() == tg.ParseModeNone {
			x.SetParseMode(bot.parseMode)
		}
	}

	return msg.Send(bot.api)
}

// StartPolling does an infinite GetUpdates with the timeout of the passed timeoutSeconds.
// allowedUpdates by default passes nothing and uses the telegram's default.
//
// see tgo.GetUpdate for more detailed information.
func (bot *Bot) StartPolling(timeoutSeconds int64, allowedUpdates ...string) error {
	var offset int64

	for {
		data, err := bot.api.GetUpdates(&tg.GetUpdates{
			Offset:         offset, // Is there any better way to do this? open an issue/pull-request if you know. thx.
			Timeout:        timeoutSeconds,
			AllowedUpdates: allowedUpdates,
		})
		if err != nil {
			if errors.Is(err, syscall.ECONNRESET) {
				continue
			}
			return err
		}

		for _, update := range data {
			offset = update.UpdateId + 1

			go func(update *tg.Update) { bot.HandleUpdate(bot, update) }(update)
		}
	}
}
