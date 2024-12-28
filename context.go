package tgo

import (
	"errors"
	"sync"

	"github.com/haashemi/tgo/tg"
)

var (
	ErrNoChatToSend = errors.New("no chat to send the message")
)

type Context interface {
	// API returns the Bot's API instance.
	API() *tg.API

	// Bot returns the bot instance that got the update.
	Bot() *Bot

	// Message returns the received message.
	// Returns nil if the update type is not message.
	Message() *tg.Message

	// RawUpdate returns the raw received update.
	RawUpdate() *tg.Update

	// Update returns the actual received update (e.g., Message, CallbackQuery, etc.).
	// Returns nil if the update is unknown.
	//
	// Note: This method is intended to be used internally by the framework,
	// but it might be useful in some cases.
	Update() any

	// From returns sender of the update. Returns nil if the update doesn't have a sender.
	From() *tg.User

	// Chat returns the chat of the update. Returns nil if the update doesn't have a chat.
	Chat() *tg.Chat

	// Session returns the user's session storage.
	// Returns nil if the update doesn't have a sender.
	Session() *sync.Map

	Send(Sendable) (*tg.Message, error)
	Reply(replyable Replyable) (*tg.Message, error)
}

type BaseContext struct {
	bot    *Bot       // Bot instance that got the update
	update *tg.Update // Raw received update

	// Storage contains an in-context storage used for middlewares to pass some data
	// to the next middleware or even the handler.
	Storage sync.Map
}

func NewContext(bot *Bot, u *tg.Update) Context {
	return &BaseContext{bot: bot, update: u}
}

func (ctx *BaseContext) API() *tg.API {
	return ctx.bot.api
}

func (ctx *BaseContext) Bot() *Bot {
	return ctx.bot
}

func (ctx *BaseContext) Message() *tg.Message {
	switch {
	case ctx.update.Message != nil:
		return ctx.update.Message
	case ctx.update.EditedMessage != nil:
		return ctx.update.EditedMessage
	case ctx.update.ChannelPost != nil:
		return ctx.update.ChannelPost
	case ctx.update.EditedChannelPost != nil:
		return ctx.update.EditedChannelPost
	case ctx.update.BusinessMessage != nil:
		return ctx.update.BusinessMessage
	case ctx.update.EditedBusinessMessage != nil:
		return ctx.update.EditedBusinessMessage
	default:
		return nil
	}
}

func (ctx *BaseContext) RawUpdate() *tg.Update {
	return ctx.update
}

func (ctx *BaseContext) Update() any {
	switch {
	case ctx.update.Message != nil:
		return ctx.update.Message
	case ctx.update.EditedMessage != nil:
		return ctx.update.EditedMessage
	case ctx.update.ChannelPost != nil:
		return ctx.update.ChannelPost
	case ctx.update.EditedChannelPost != nil:
		return ctx.update.EditedChannelPost
	case ctx.update.BusinessConnection != nil:
		return ctx.update.BusinessConnection
	case ctx.update.BusinessMessage != nil:
		return ctx.update.BusinessMessage
	case ctx.update.EditedBusinessMessage != nil:
		return ctx.update.EditedBusinessMessage
	case ctx.update.DeletedBusinessMessages != nil:
		return ctx.update.DeletedBusinessMessages
	case ctx.update.MessageReaction != nil:
		return ctx.update.MessageReaction
	case ctx.update.MessageReactionCount != nil:
		return ctx.update.MessageReactionCount
	case ctx.update.InlineQuery != nil:
		return ctx.update.InlineQuery
	case ctx.update.ChosenInlineResult != nil:
		return ctx.update.ChosenInlineResult
	case ctx.update.CallbackQuery != nil:
		return ctx.update.CallbackQuery
	case ctx.update.ShippingQuery != nil:
		return ctx.update.ShippingQuery
	case ctx.update.PreCheckoutQuery != nil:
		return ctx.update.PreCheckoutQuery
	case ctx.update.PurchasedPaidMedia != nil:
		return ctx.update.PurchasedPaidMedia
	case ctx.update.Poll != nil:
		return ctx.update.Poll
	case ctx.update.PollAnswer != nil:
		return ctx.update.PollAnswer
	case ctx.update.MyChatMember != nil:
		return ctx.update.MyChatMember
	case ctx.update.ChatMember != nil:
		return ctx.update.ChatMember
	case ctx.update.ChatJoinRequest != nil:
		return ctx.update.ChatJoinRequest
	case ctx.update.ChatBoost != nil:
		return ctx.update.ChatBoost
	case ctx.update.RemovedChatBoost != nil:
		return ctx.update.RemovedChatBoost
	default:
		return nil
	}
}

func (ctx *BaseContext) From() *tg.User {
	switch upd := ctx.Update().(type) {
	case *tg.Message:
		return upd.From
	case *tg.BusinessConnection:
		return &upd.User
	case *tg.MessageReactionUpdated:
		return upd.User
	case *tg.InlineQuery:
		return &upd.From
	case *tg.ChosenInlineResult:
		return &upd.From
	case *tg.CallbackQuery:
		return &upd.From
	case *tg.ShippingQuery:
		return &upd.From
	case *tg.PreCheckoutQuery:
		return &upd.From
	case *tg.PaidMediaPurchased:
		return &upd.From
	case *tg.PollAnswer:
		return upd.User
	case *tg.ChatMemberUpdated:
		return &upd.From
	case *tg.ChatJoinRequest:
		return &upd.From
	default:
		return nil
	}
}

func (ctx *BaseContext) Chat() *tg.Chat {
	switch upd := ctx.Update().(type) {
	case *tg.Message:
		return &upd.Chat
	case *tg.BusinessMessagesDeleted:
		return &upd.Chat
	case *tg.MessageReactionUpdated:
		return &upd.Chat
	case *tg.MessageReactionCountUpdated:
		return &upd.Chat
	case *tg.CallbackQuery:
		switch msg := upd.Message.(type) {
		case *tg.Message:
			return &msg.Chat
		case *tg.InaccessibleMessage:
			return &msg.Chat
		}
	case *tg.ChatMemberUpdated:
		return &upd.Chat
	case *tg.ChatJoinRequest:
		return &upd.Chat
	case *tg.ChatBoostUpdated:
		return &upd.Chat
	case *tg.ChatBoostRemoved:
		return &upd.Chat
	}

	return nil
}

func (ctx *BaseContext) Session() *sync.Map {
	from := ctx.From()
	if from == nil {
		return nil
	}

	return ctx.bot.GetSession(from.Id)
}

func (ctx *BaseContext) Send(sendable Sendable) (*tg.Message, error) {
	chat := ctx.Chat()
	if chat == nil {
		return nil, ErrNoChatToSend
	}

	if sendable.GetChatID() == nil {
		sendable.SetChatID(chat.Id)
	}

	// TODO: Support MessageThreadID

	return ctx.bot.Send(sendable)
}

func (ctx *BaseContext) Reply(replyable Replyable) (*tg.Message, error) {
	chat := ctx.Chat()
	if chat == nil {
		return nil, ErrNoChatToSend
	}

	if replyable.GetChatID() == nil {
		replyable.SetChatID(chat.Id)
	}

	if rp := replyable.GetReplyParameters(); rp == nil || rp.MessageId == 0 {
		msg := ctx.Message()
		if msg != nil {
			replyable.SetReplyToMessageID(msg.MessageId)
		}
	}

	return ctx.bot.Send(replyable)
}
