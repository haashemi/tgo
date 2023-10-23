package message

import (
	"sync"
	"time"

	"github.com/haashemi/tgo"
)

type Context struct {
	// Message contains the raw received message
	*tgo.Message

	// Bot is the bot instance which got the update.
	Bot *tgo.Bot

	// Storage contains an in-context storage used for middlewares to pass some data
	// to the next middleware or even the handler.
	Storage sync.Map
}

// Session returns the user's session storage.
// it will return the chat's session if user-id is zero.
//
// ToDo: make sure that we are getting the chat id in the right way.
func (ctx *Context) Session() *sync.Map {
	var id int64

	if ctx.From != nil {
		id = ctx.From.Id
	} else if ctx.SenderChat != nil {
		id = ctx.SenderChat.Id
	} else {
		id = ctx.Chat.Id
	}

	return ctx.Bot.GetSession(id)
}

// String returns the message's text or media caption
func (m *Context) String() string {
	if m.Text != "" {
		return m.Text
	}

	return m.Caption
}

// Send sends a message into the current chat with the preferred ParseMode.
// It will set the target ChatId if not set.
func (ctx *Context) Send(msg tgo.Sendable) (*tgo.Message, error) {
	if msg.GetChatID() == nil {
		msg.SetChatID(ctx.Chat.Id)
	}

	return ctx.Bot.Send(msg)
}

// Reply replies to the current message with the preferred ParseMode.
// It will pass/override the ChatId and ReplyToMessageId field.
func (ctx *Context) Reply(msg tgo.Replyable) (*tgo.Message, error) {
	msg.SetChatID(ctx.Chat.Id)
	msg.SetReplyToMessageId(ctx.MessageId)

	return ctx.Bot.Send(msg)
}

// Ask asks a question from the message's sender and waits for the passed timeout for their response.
func (ctx *Context) Ask(msg tgo.Sendable, timeout time.Duration) (question, answer *tgo.Message, err error) {
	cid, sid := tgo.GetChatAndSenderID(ctx.Message)
	return ctx.Bot.Ask(cid, sid, msg, timeout)
}

// Delete deletes the received message.
func (ctx *Context) Delete() error {
	_, err := ctx.Bot.DeleteMessage(&tgo.DeleteMessage{ChatId: tgo.ID(ctx.Chat.Id), MessageId: ctx.MessageId})
	return err
}
