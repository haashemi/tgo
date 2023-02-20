package tgo

type Context = *context

type context struct {
	bot     *Bot
	data    *Update
	stopped bool
}

func newContext(bot *Bot, update *Update) *context {
	return &context{
		bot:  bot,
		data: update,
	}
}

func (ctx *context) API() *API {
	return ctx.bot.api
}

func (ctx *context) Bot() *Bot {
	return ctx.bot
}

func (ctx *context) Session() *Session {
	return ctx.bot.GetSession(ctx.SenderChatID())
}

// ChatID returns the chat id of where the update sent from
//
// TODO: Handle other type of updates like Callback
func (ctx *context) ChatID() int64 {
	return ctx.data.Message.Chat.Id
}

// ChatID returns the chat id of who send the update
//
// TODO: Handle other type of updates and manage optional fields
func (ctx *context) SenderChatID() int64 {
	if ctx.data.CallbackQuery != nil {
		return ctx.data.CallbackQuery.From.Id
	} else if ctx.data.Message != nil {
		return ctx.data.Message.SenderChat.Id
	} else if ctx.data.ChannelPost != nil {
		return ctx.data.ChannelPost.SenderChat.Id
	} else if ctx.data.EditedMessage != nil {
		return ctx.data.EditedMessage.SenderChat.Id
	}

	return -1
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

func (ctx *context) Send(sendable SendableMessage) (*Message, error) {
	return sendable.SetChatID(NewChatID(ctx.ChatID())).
		SetThreadID(ctx.ThreadID()).
		Send(ctx.bot.api)
}

// RawUpdate returns the update came from telegram directly.
func (ctx *context) RawUpdate() *Update {
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
