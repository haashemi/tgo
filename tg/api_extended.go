package tg

// ParseMode is a type that represents the parse mode of a message.
// eg. Markdown, HTML, etc.
type ParseMode string

const (
	// does not parse the message
	ParseModeNone ParseMode = ""
	// parses the message as Markdown
	ParseModeMarkdown ParseMode = "Markdown"
	// parses the message as Markdown but using telegram's V2 markdown
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	// parses the message as HTML
	ParseModeHTML ParseMode = "HTML"
)

// Username implements a string ChatID.
type Username string

// IsChatID does nothing but implements ChatID interface.
func (Username) IsChatID() {}

// ID implements a int64 ChatID.
type ID int64

// IsChatID does nothing but implements ChatID interface.
func (ID) IsChatID() {}

// GetChatID implements Sendable's GetChatID method.
func (x *SendAnimation) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendAudio) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendContact) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendDice) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendDocument) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendGame) GetChatID() ChatID { return ID(x.ChatId) }

// GetChatID implements Sendable's GetChatID method.
func (x *SendInvoice) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendLocation) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendMessage) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendPhoto) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendPoll) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendSticker) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendVenue) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendVideo) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendVideoNote) GetChatID() ChatID { return x.ChatId }

// GetChatID implements Sendable's GetChatID method.
func (x *SendVoice) GetChatID() ChatID { return x.ChatId }

// SetChatID implements Sendable's SetChatID method.
func (x *SendAnimation) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendAudio) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendContact) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendDice) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendDocument) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendGame) SetChatID(id int64) { x.ChatId = id }

// SetChatID implements Sendable's SetChatID method.
func (x *SendInvoice) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendLocation) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendMessage) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendPhoto) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendPoll) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendSticker) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendVenue) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendVideo) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendVideoNote) SetChatID(id int64) { x.ChatId = ID(id) }

// SetChatID implements Sendable's SetChatID method.
func (x *SendVoice) SetChatID(id int64) { x.ChatId = ID(id) }

// Send implements Sendable's Send method.
func (x *SendAnimation) Send(api *API) (*Message, error) { return api.SendAnimation(x) }

// Send implements Sendable's Send method.
func (x *SendAudio) Send(api *API) (*Message, error) { return api.SendAudio(x) }

// Send implements Sendable's Send method.
func (x *SendContact) Send(api *API) (*Message, error) { return api.SendContact(x) }

// Send implements Sendable's Send method.
func (x *SendDice) Send(api *API) (*Message, error) { return api.SendDice(x) }

// Send implements Sendable's Send method.
func (x *SendDocument) Send(api *API) (*Message, error) { return api.SendDocument(x) }

// Send implements Sendable's Send method.
func (x *SendGame) Send(api *API) (*Message, error) { return api.SendGame(x) }

// Send implements Sendable's Send method.
func (x *SendInvoice) Send(api *API) (*Message, error) { return api.SendInvoice(x) }

// Send implements Sendable's Send method.
func (x *SendLocation) Send(api *API) (*Message, error) { return api.SendLocation(x) }

// Send implements Sendable's Send method.
func (x *SendMessage) Send(api *API) (*Message, error) { return api.SendMessage(x) }

// Send implements Sendable's Send method.
func (x *SendPhoto) Send(api *API) (*Message, error) { return api.SendPhoto(x) }

// Send implements Sendable's Send method.
func (x *SendPoll) Send(api *API) (*Message, error) { return api.SendPoll(x) }

// Send implements Sendable's Send method.
func (x *SendSticker) Send(api *API) (*Message, error) { return api.SendSticker(x) }

// Send implements Sendable's Send method.
func (x *SendVenue) Send(api *API) (*Message, error) { return api.SendVenue(x) }

// Send implements Sendable's Send method.
func (x *SendVideo) Send(api *API) (*Message, error) { return api.SendVideo(x) }

// Send implements Sendable's Send method.
func (x *SendVideoNote) Send(api *API) (*Message, error) { return api.SendVideoNote(x) }

// Send implements Sendable's Send method.
func (x *SendVoice) Send(api *API) (*Message, error) { return api.SendVoice(x) }

// GetParseMode implements ParseModeSettable's GetParseMode method.
func (x *SendAnimation) GetParseMode() ParseMode { return x.ParseMode }

// SetParseMode implements ParseModeSettable's SetParseMode method.
func (x *SendAnimation) SetParseMode(mode ParseMode) { x.ParseMode = mode }

// GetParseMode implements ParseModeSettable's GetParseMode method.
func (x *SendAudio) GetParseMode() ParseMode { return x.ParseMode }

// SetParseMode implements ParseModeSettable's SetParseMode method.
func (x *SendAudio) SetParseMode(mode ParseMode) { x.ParseMode = mode }

// GetParseMode implements ParseModeSettable's GetParseMode method.
func (x *SendDocument) GetParseMode() ParseMode { return x.ParseMode }

// SetParseMode implements ParseModeSettable's SetParseMode method.
func (x *SendDocument) SetParseMode(mode ParseMode) { x.ParseMode = mode }

// GetParseMode implements ParseModeSettable's GetParseMode method.
func (x *SendMessage) GetParseMode() ParseMode { return x.ParseMode }

// SetParseMode implements ParseModeSettable's SetParseMode method.
func (x *SendMessage) SetParseMode(mode ParseMode) { x.ParseMode = mode }

// GetParseMode implements ParseModeSettable's GetParseMode method.
func (x *SendPhoto) GetParseMode() ParseMode { return x.ParseMode }

// SetParseMode implements ParseModeSettable's SetParseMode method.
func (x *SendPhoto) SetParseMode(mode ParseMode) { x.ParseMode = mode }

// GetParseMode implements ParseModeSettable's GetParseMode method.
func (x *SendVideo) GetParseMode() ParseMode { return x.ParseMode }

// SetParseMode implements ParseModeSettable's SetParseMode method.
func (x *SendVideo) SetParseMode(mode ParseMode) { x.ParseMode = mode }

// GetParseMode implements ParseModeSettable's GetParseMode method.
func (x *SendVoice) GetParseMode() ParseMode { return x.ParseMode }

// SetParseMode implements ParseModeSettable's SetParseMode method.
func (x *SendVoice) SetParseMode(mode ParseMode) { x.ParseMode = mode }

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendAnimation) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendAudio) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendContact) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendDice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendDocument) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendGame) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendInvoice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendLocation) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendMessage) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendPhoto) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendPoll) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendSticker) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendVenue) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendVideo) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendVideoNote) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}

// SetReplyToMessageId implements Replyable's SetReplyToMessageId method.
func (x *SendVoice) SetReplyToMessageId(id int64) {
	x.ReplyParameters = &ReplyParameters{MessageId: id, ChatId: x.GetChatID()}
}
