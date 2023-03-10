package tgo

import (
	"errors"
	"sync"
)

type Contextable interface {
	ChatID() int64
	SenderID() int64
	ThreadID() int64
	MessageID() int64
}

type Context = *UpdateContext

type UpdateContext struct {
	*Update
	Contextable

	bot *Bot

	// contains an in-context storage used for middlewares to pass some data
	// to the next middle or even the handler
	storage sync.Map
}

func NewContext(update *Update, bot *Bot) Context {
	var contextable Contextable

	switch {
	case update.Message != nil:
		contextable = update.Message
	case update.EditedMessage != nil:
		contextable = update.EditedMessage
	case update.ChannelPost != nil:
		contextable = update.ChannelPost
	case update.EditedChannelPost != nil:
		contextable = update.EditedChannelPost
	case update.CallbackQuery != nil:
		contextable = update.CallbackQuery
	}

	return &UpdateContext{Update: update, Contextable: contextable, bot: bot}
}

// Bot returns the bot
func (ctx *UpdateContext) Bot() *Bot { return ctx.bot }

// Session returns the user's session storage.
// it will return the chat's session if user-id is zero.
func (ctx *UpdateContext) Session() *sync.Map {
	id := ctx.SenderID()
	if id == 0 {
		id = ctx.ChatID()
	}

	return ctx.bot.GetSession(id)
}

// Storage returns the in-context storage.
// it's used to pass data from a middleware to the next middleware or handler.
func (ctx *UpdateContext) Storage() *sync.Map {
	return &ctx.storage
}

func (ctx *UpdateContext) GetMessage() *Message {
	if data, ok := ctx.Contextable.(*Message); ok {
		return data
	}
	return nil
}

// Text returns the message's text or media caption or callback query's data
func (ctx *UpdateContext) Text() string {
	switch data := ctx.Contextable.(type) {
	case *Message:
		return data.String()

	case *CallbackQuery:
		return data.Data

	default:
		return ""
	}
}

// Delete deletes the message of that context
func (ctx *UpdateContext) Delete() error {
	_, err := ctx.bot.DeleteMessage(ChatID(ctx.ChatID()), ctx.MessageID())
	return err
}

// Answer answers to callback queries sent from inline keyboards
func (ctx *UpdateContext) Answer(options *AnswerCallbackQueryOptions) error {
	if ctx.CallbackQuery == nil {
		return errors.New("context doesn't have CallbackQuery data")
	}

	_, err := ctx.bot.AnswerCallbackQuery(ctx.CallbackQuery.Id, options)
	return err
}
