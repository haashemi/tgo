package tgo

import (
	"context"
	"fmt"
	"time"
)

type Context struct {
	bot     *Bot
	data    *Update
	stopped bool
}

func newContext(bot *Bot, update *Update) *Context {
	return &Context{
		bot:  bot,
		data: update,
	}
}

func (ctx *Context) API() *API {
	return ctx.bot.api
}

func (ctx *Context) Bot() *Bot {
	return ctx.bot
}

func (ctx *Context) Session() *Session {
	return ctx.bot.GetSession(ctx.SenderChatID())
}

// ChatID returns the chat id of where the update sent from
//
// TODO: Handle other type of updates like Callback
func (ctx *Context) ChatID() int64 {
	return ctx.data.Message.Chat.Id
}

// ChatID returns the chat id of who send the update
//
// TODO: Handle other type of updates and manage optional fields
func (ctx *Context) SenderChatID() int64 {
	if ctx.data.CallbackQuery != nil {
		return ctx.data.CallbackQuery.From.Id
	} else if ctx.data.Message != nil {
		return ctx.data.Message.From.Id
		// return ctx.data.Message.SenderChat.Id
	} else if ctx.data.ChannelPost != nil {
		return ctx.data.ChannelPost.SenderChat.Id
	} else if ctx.data.EditedMessage != nil {
		return ctx.data.EditedMessage.SenderChat.Id
	}

	return -1
}

// Caption returns the sent message's text or media caption
//
// TODO: Support all media types on different updates
func (ctx *Context) Caption() string {
	return ctx.data.Message.Text
}

// ThreadID returns id of the topic where the update came from, if any.
// it returns 0 if there is no thread.
func (ctx *Context) ThreadID() int64 {
	if ctx.data.Message != nil {
		return ctx.data.Message.MessageThreadId
	} else if ctx.data.EditedMessage != nil {
		return ctx.data.EditedMessage.MessageThreadId
	} else if ctx.data.ChannelPost != nil {
		return ctx.data.ChannelPost.MessageThreadId
	}
	return 0
}

func (ctx *Context) Send(sendable SendableMessage) (*Message, error) {
	return sendable.SetChatID(NewChatID(ctx.ChatID())).
		SetThreadID(ctx.ThreadID()).
		Send(ctx.bot.api)
}

func (ctx *Context) Ask(sendable SendableMessage, timeout time.Duration) (question *Message, answer *Context, err error) {
	question, err = ctx.Send(sendable)
	if err != nil {
		return nil, nil, err
	}

	uid := fmt.Sprintf("%d-%d", ctx.ChatID(), ctx.SenderChatID())
	waiter := make(chan *Context, 1)

	ctx.bot.askMut.Lock()
	ctx.bot.asks[uid] = waiter
	ctx.bot.askMut.Unlock()

	defer func() {
		ctx.bot.askMut.Lock()
		delete(ctx.bot.asks, uid)
		ctx.bot.askMut.Unlock()

		close(waiter)
	}()

	aCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case answer = <-waiter:
		return

	case <-aCtx.Done():
		err = aCtx.Err()
		return
	}
}

// RawUpdate returns the update came from telegram directly.
func (ctx *Context) RawUpdate() *Update {
	return ctx.data
}

// Stop stops the context.
// It will only be used for middlewares and NOT the main handlers.
func (ctx *Context) Stop() {
	ctx.stopped = true
}

func (ctx *Context) isStopped() bool {
	return ctx.stopped
}

func (ctx *Context) clone() *Context {
	return &Context{bot: ctx.bot, data: ctx.data, stopped: ctx.stopped}
}
