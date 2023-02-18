package tgo

import "github.com/haashemi/tgo/botapi"

type Context = *context

type context struct {
	bot     *Bot
	data    *botapi.Update
	stopped bool
}

func newContext(bot *Bot, update *botapi.Update) *context {
	return &context{
		bot:  bot,
		data: update,
	}
}

func (ctx *context) API() *botapi.API {
	return ctx.bot.api
}

func (ctx *context) Bot() *Bot {
	return ctx.bot
}

// ChatID returns the chat id of where the update sent from
//
// TODO: Handle other type of updates like Callback
func (ctx *context) ChatID() int64 {
	return ctx.data.Message.Chat.Id
}

// ThreadID returns id of the topic where the update came from, if any.
// it returns 0 if there is no thread.
func (ctx *context) ThreadID() int64 {
	if ctx.data.Message != nil {
		return ctx.data.Message.MessageThreadId
	} else if ctx.data.EditedMessage != nil {
		return ctx.data.EditedMessage.MessageThreadId
	} else if ctx.data.ChannelPost != nil {
		return ctx.data.ChannelPost.MessageThreadId
	}
	return 0
}

func (ctx *context) Send(sendable SendableMessage) (*botapi.Message, error) {
	return sendable.SetChatID(NewChatID(ctx.ChatID())).
		SetThreadID(ctx.ThreadID()).
		Send(ctx.bot.api)
}

// RawUpdate returns the update came from telegram directly.
func (ctx *context) RawUpdate() *botapi.Update {
	return ctx.data
}

// Stop stops the context.
// It will only be used for middlewares and NOT the main handlers.
func (ctx *context) Stop() {
	ctx.stopped = true
}

func (ctx *context) isStopped() bool {
	return ctx.stopped
}

func (ctx *context) clone() *context {
	return &context{bot: ctx.bot, data: ctx.data, stopped: ctx.stopped}
}
