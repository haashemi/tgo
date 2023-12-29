package callback

import (
	"sync"
	"time"

	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/routers/message"
)

type Context struct {
	// CallbackQuery contains the raw received query
	*tgo.CallbackQuery

	// Bot is the bot instance which got the update.
	Bot *tgo.Bot

	// Storage contains an in-context storage used for middlewares to pass some data
	// to the next middleware or even the handler.
	Storage sync.Map
}

// Session returns the user's session storage.
// it will return the chat's session if user-id is zero.
func (ctx *Context) Session() *sync.Map {
	return ctx.Bot.GetSession(ctx.From.Id)
}

// Send sends a message into the message's chat if exist, otherwise sends in the sender's chat, with the preferred ParseMode.
// It will set the target ChatId if not set.
func (ctx *Context) Send(msg tgo.Sendable) (*tgo.Message, error) {
	if msg.GetChatID() == nil {
		if ctx.Message != nil {
			switch ctxMsg := ctx.Message.(type) {
			case *tgo.InaccessibleMessage:
				msg.SetChatID(ctxMsg.Chat.Id)
			case *tgo.Message:
				msg.SetChatID(ctxMsg.Chat.Id)
			}
		} else {
			msg.SetChatID(ctx.From.Id)
		}
	}

	return ctx.Bot.Send(msg)
}

// Ask asks a question from the callback query sender and waits for the passed timeout for their response.
func (ctx *Context) Ask(msg tgo.Sendable, timeout time.Duration) (question, answer *message.Context, err error) {
	chatID := ctx.From.Id
	if ctx.Message != nil {
		switch ctxMsg := ctx.Message.(type) {
		case *tgo.InaccessibleMessage:
			chatID = ctxMsg.Chat.Id
		case *tgo.Message:
			chatID = ctxMsg.Chat.Id
		}
	}

	rawQuestion, rawAnswer, err := ctx.Bot.Ask(chatID, ctx.From.Id, msg, timeout)

	question = &message.Context{Message: rawQuestion, Bot: ctx.Bot}
	answer = &message.Context{Message: rawAnswer, Bot: ctx.Bot}

	return question, answer, err
}

// Answer answers to the sent callback query.
// it fills the CallbackQueryId field by default.
func (ctx *Context) Answer(options *tgo.AnswerCallbackQuery) error {
	options.CallbackQueryId = ctx.CallbackQuery.Id

	_, err := ctx.Bot.AnswerCallbackQuery(options)
	return err
}
