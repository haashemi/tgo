package tgo

import (
	"errors"
	"net/http"
	"sync"
	"syscall"
	"time"

	"github.com/haashemi/tgo/tg"
)

// Sendable is an interface that represents any object that can be sent using an API client.
type Sendable interface {
	// GetChatID returns the chat ID associated with the sendable object.
	GetChatID() tg.ChatID

	// SetChatID sets the chat ID for the sendable object.
	SetChatID(id int64)

	// Send sends the sendable object using the provided API client.
	// It returns the sent message and any error encountered.
	Send(api *tg.API) (*tg.Message, error)
}

// ParseModeSettable is an interface that represents any object that can have its ParseMode set
// Or in other words, messages with captions.
type ParseModeSettable interface {
	GetParseMode() tg.ParseMode
	SetParseMode(mode tg.ParseMode)
}

// Replyable is an interface that represents any object that can be replied to.
type Replyable interface {
	Sendable
	SetReplyToMessageId(id int64)
}

type Filter interface{ Check(update *tg.Update) bool }

type Router interface {
	Setup(bot *Bot) error
	HandleUpdate(bot *Bot, upd *tg.Update) (used bool)
}

type Bot struct {
	*tg.API // embedding the api to add all api methods to the bot

	DefaultParseMode tg.ParseMode

	asks   map[string]chan<- *tg.Message
	askMut sync.RWMutex

	routers []Router

	// contains user-ids with their session
	sessions sync.Map
}

type Options struct {
	Host             string
	Client           *http.Client
	DefaultParseMode tg.ParseMode
}

func NewBot(token string, opts Options) (bot *Bot) {
	api := tg.NewAPI(token, opts.Host, opts.Client)

	return &Bot{
		API:              api,
		DefaultParseMode: opts.DefaultParseMode,
		asks:             make(map[string]chan<- *tg.Message),
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

// Send sends a message with the preferred ParseMode.
func (b *Bot) Send(msg Sendable) (*tg.Message, error) {
	if x, ok := msg.(ParseModeSettable); ok {
		if x.GetParseMode() == tg.ParseModeNone {
			x.SetParseMode(b.DefaultParseMode)
		}
	}

	return msg.Send(b.API)
}

// StartPolling does an infinite GetUpdates with the timeout of the passed timeoutSeconds.
// allowedUpdates by default passes nothing and uses the telegram's default.
//
// see tgo.GetUpdate for more detailed information.
func (bot *Bot) StartPolling(timeoutSeconds int64, allowedUpdates ...string) error {
	var offset int64

	for {
		data, err := bot.GetUpdates(&tg.GetUpdates{
			Offset:         offset, // Is there any better way to do this? open an issue/pull-request if you know. thx.
			Timeout:        timeoutSeconds,
			AllowedUpdates: allowedUpdates,
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

			go func(update *tg.Update) {
				if update.Message != nil && bot.sendAnswerIfAsked(update.Message) {
					return
				}

				for _, router := range bot.routers {
					if used := router.HandleUpdate(bot, update); used {
						return
					}
				}
			}(update)
		}
	}
}
