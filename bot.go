package tgo

import (
	"net/http"
	"sync"
)

type Filter interface{ Check(update *Update) bool }

type Router interface {
	Setup(bot *Bot) error
	HandleUpdate(bot *Bot, upd *Update) (used bool)
}

type Bot struct {
	*API // embedding the api to add all api methods to the bot

	DefaultParseMode ParseMode

	asks   map[string]chan<- *Message
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

func NewBot(token string, opts Options) (bot *Bot) {
	api := NewAPI(token, opts.Host, opts.Client)

	return &Bot{
		API:              api,
		DefaultParseMode: opts.DefaultParseMode,
		asks:             make(map[string]chan<- *Message),
	}
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
