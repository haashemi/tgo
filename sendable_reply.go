package tgo

import "github.com/haashemi/tgo/tg"

// Replyable is an interface that represents any object that can be replied to.
type Replyable interface {
	Sendable
	SetReplyToMessageId(id int64)
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendAnimation) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendAudio) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendContact) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendDice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendDocument) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendGame) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendInvoice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendLocation) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendMessage) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendPhoto) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendPoll) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendSticker) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendVenue) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendVideo) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendVideoNote) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendVoice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &tg.ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
