package tgo

import "github.com/haashemi/tgo/tg"

type (
	SendAnimation tg.SendAnimation // SendAnimation is the extended version of `tg.SendAnimation`
	SendAudio     tg.SendAudio     // SendAudio is the extended version of `tg.SendAudio`
	SendContact   tg.SendContact   // SendContact is the extended version of `tg.SendContact`
	SendDice      tg.SendDice      // SendDice is the extended version of `tg.SendDice`
	SendDocument  tg.SendDocument  // SendDocument is the extended version of `tg.SendDocument`
	SendGame      tg.SendGame      // SendGame is the extended version of `tg.SendGame`
	SendInvoice   tg.SendInvoice   // SendInvoice is the extended version of `tg.SendInvoice`
	SendLocation  tg.SendLocation  // SendLocation is the extended version of `tg.SendLocation`
	SendMessage   tg.SendMessage   // SendMessage is the extended version of `tg.SendMessage`
	SendPhoto     tg.SendPhoto     // SendPhoto is the extended version of `tg.SendPhoto`
	SendPoll      tg.SendPoll      // SendPoll is the extended version of `tg.SendPoll`
	SendSticker   tg.SendSticker   // SendSticker is the extended version of `tg.SendSticker`
	SendVenue     tg.SendVenue     // SendVenue is the extended version of `tg.SendVenue`
	SendVideo     tg.SendVideo     // SendVideo is the extended version of `tg.SendVideo`
	SendVideoNote tg.SendVideoNote // SendVideoNote is the extended version of `tg.SendVideoNote`
	SendVoice     tg.SendVoice     // SendVoice is the extended version of `tg.SendVoice`
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

// Replyable is an interface that represents any object that can be replied to.
type Replyable interface {
	Sendable
	SetReplyToMessageId(id int64)
}

// ParseModeSettable is an interface that represents any object that can have its ParseMode set
// Or in other words, messages with captions.
type ParseModeSettable interface {
	Sendable

	// GetParseMode returns the current set parse mode.
	GetParseMode() tg.ParseMode

	// SetParseMode updates the parse mode.
	SetParseMode(mode tg.ParseMode)
}

func (x *SendAnimation) GetChatID() tg.ChatID { return x.ChatId }
func (x *SendAudio) GetChatID() tg.ChatID     { return x.ChatId }
func (x *SendContact) GetChatID() tg.ChatID   { return x.ChatId }
func (x *SendDice) GetChatID() tg.ChatID      { return x.ChatId }
func (x *SendDocument) GetChatID() tg.ChatID  { return x.ChatId }
func (x *SendGame) GetChatID() tg.ChatID      { return tg.ID(x.ChatId) }
func (x *SendInvoice) GetChatID() tg.ChatID   { return x.ChatId }
func (x *SendLocation) GetChatID() tg.ChatID  { return x.ChatId }
func (x *SendMessage) GetChatID() tg.ChatID   { return x.ChatId }
func (x *SendPhoto) GetChatID() tg.ChatID     { return x.ChatId }
func (x *SendPoll) GetChatID() tg.ChatID      { return x.ChatId }
func (x *SendSticker) GetChatID() tg.ChatID   { return x.ChatId }
func (x *SendVenue) GetChatID() tg.ChatID     { return x.ChatId }
func (x *SendVideo) GetChatID() tg.ChatID     { return x.ChatId }
func (x *SendVideoNote) GetChatID() tg.ChatID { return x.ChatId }
func (x *SendVoice) GetChatID() tg.ChatID     { return x.ChatId }

func (x *SendAnimation) SetChatID(id int64) { x.ChatId = tg.ID(id) }
func (x *SendAudio) SetChatID(id int64)     { x.ChatId = tg.ID(id) }
func (x *SendContact) SetChatID(id int64)   { x.ChatId = tg.ID(id) }
func (x *SendDice) SetChatID(id int64)      { x.ChatId = tg.ID(id) }
func (x *SendDocument) SetChatID(id int64)  { x.ChatId = tg.ID(id) }
func (x *SendGame) SetChatID(id int64)      { x.ChatId = id }
func (x *SendInvoice) SetChatID(id int64)   { x.ChatId = tg.ID(id) }
func (x *SendLocation) SetChatID(id int64)  { x.ChatId = tg.ID(id) }
func (x *SendMessage) SetChatID(id int64)   { x.ChatId = tg.ID(id) }
func (x *SendPhoto) SetChatID(id int64)     { x.ChatId = tg.ID(id) }
func (x *SendPoll) SetChatID(id int64)      { x.ChatId = tg.ID(id) }
func (x *SendSticker) SetChatID(id int64)   { x.ChatId = tg.ID(id) }
func (x *SendVenue) SetChatID(id int64)     { x.ChatId = tg.ID(id) }
func (x *SendVideo) SetChatID(id int64)     { x.ChatId = tg.ID(id) }
func (x *SendVideoNote) SetChatID(id int64) { x.ChatId = tg.ID(id) }
func (x *SendVoice) SetChatID(id int64)     { x.ChatId = tg.ID(id) }

func (x *SendAnimation) GetParseMode() tg.ParseMode { return x.ParseMode }
func (x *SendAudio) GetParseMode() tg.ParseMode     { return x.ParseMode }
func (x *SendDocument) GetParseMode() tg.ParseMode  { return x.ParseMode }
func (x *SendMessage) GetParseMode() tg.ParseMode   { return x.ParseMode }
func (x *SendPhoto) GetParseMode() tg.ParseMode     { return x.ParseMode }
func (x *SendVideo) GetParseMode() tg.ParseMode     { return x.ParseMode }
func (x *SendVoice) GetParseMode() tg.ParseMode     { return x.ParseMode }

func (x *SendAnimation) SetParseMode(mode tg.ParseMode) { x.ParseMode = mode }
func (x *SendAudio) SetParseMode(mode tg.ParseMode)     { x.ParseMode = mode }
func (x *SendDocument) SetParseMode(mode tg.ParseMode)  { x.ParseMode = mode }
func (x *SendMessage) SetParseMode(mode tg.ParseMode)   { x.ParseMode = mode }
func (x *SendPhoto) SetParseMode(mode tg.ParseMode)     { x.ParseMode = mode }
func (x *SendVideo) SetParseMode(mode tg.ParseMode)     { x.ParseMode = mode }
func (x *SendVoice) SetParseMode(mode tg.ParseMode)     { x.ParseMode = mode }

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

func (x *SendAnimation) Send(api *tg.API) (*tg.Message, error) {
	return api.SendAnimation((*tg.SendAnimation)(x))
}

func (x *SendAudio) Send(api *tg.API) (*tg.Message, error) {
	return api.SendAudio((*tg.SendAudio)(x))
}

func (x *SendContact) Send(api *tg.API) (*tg.Message, error) {
	return api.SendContact((*tg.SendContact)(x))
}

func (x *SendDice) Send(api *tg.API) (*tg.Message, error) {
	return api.SendDice((*tg.SendDice)(x))
}

func (x *SendDocument) Send(api *tg.API) (*tg.Message, error) {
	return api.SendDocument((*tg.SendDocument)(x))
}

func (x *SendGame) Send(api *tg.API) (*tg.Message, error) {
	return api.SendGame((*tg.SendGame)(x))
}

func (x *SendInvoice) Send(api *tg.API) (*tg.Message, error) {
	return api.SendInvoice((*tg.SendInvoice)(x))
}

func (x *SendLocation) Send(api *tg.API) (*tg.Message, error) {
	return api.SendLocation((*tg.SendLocation)(x))
}

func (x *SendMessage) Send(api *tg.API) (*tg.Message, error) {
	return api.SendMessage((*tg.SendMessage)(x))
}

func (x *SendPhoto) Send(api *tg.API) (*tg.Message, error) {
	return api.SendPhoto((*tg.SendPhoto)(x))
}

func (x *SendPoll) Send(api *tg.API) (*tg.Message, error) {
	return api.SendPoll((*tg.SendPoll)(x))
}

func (x *SendSticker) Send(api *tg.API) (*tg.Message, error) {
	return api.SendSticker((*tg.SendSticker)(x))
}

func (x *SendVenue) Send(api *tg.API) (*tg.Message, error) {
	return api.SendVenue((*tg.SendVenue)(x))
}

func (x *SendVideo) Send(api *tg.API) (*tg.Message, error) {
	return api.SendVideo((*tg.SendVideo)(x))
}

func (x *SendVideoNote) Send(api *tg.API) (*tg.Message, error) {
	return api.SendVideoNote((*tg.SendVideoNote)(x))
}

func (x *SendVoice) Send(api *tg.API) (*tg.Message, error) {
	return api.SendVoice((*tg.SendVoice)(x))
}
