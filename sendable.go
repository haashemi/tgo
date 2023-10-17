package tgo

type Sendable interface {
	SetChatID(id int64)
	Send(api *API) (*Message, error)
}

type Replyable interface {
	Sendable
	SetReplyToMessageId(id int64)
}

type ParseModeSettable interface {
	GetParseMode() ParseMode
	SetParseMode(mode ParseMode)
}

func (x *SendAnimation) SetChatID(id int64) { x.ChatId = ID(id) }
func (x *SendAudio) SetChatID(id int64)     { x.ChatId = ID(id) }
func (x *SendContact) SetChatID(id int64)   { x.ChatId = ID(id) }
func (x *SendDice) SetChatID(id int64)      { x.ChatId = ID(id) }
func (x *SendDocument) SetChatID(id int64)  { x.ChatId = ID(id) }
func (x *SendGame) SetChatID(id int64)      { x.ChatId = id }
func (x *SendInvoice) SetChatID(id int64)   { x.ChatId = ID(id) }
func (x *SendLocation) SetChatID(id int64)  { x.ChatId = ID(id) }
func (x *SendMessage) SetChatID(id int64)   { x.ChatId = ID(id) }
func (x *SendPhoto) SetChatID(id int64)     { x.ChatId = ID(id) }
func (x *SendPoll) SetChatID(id int64)      { x.ChatId = ID(id) }
func (x *SendSticker) SetChatID(id int64)   { x.ChatId = ID(id) }
func (x *SendVenue) SetChatID(id int64)     { x.ChatId = ID(id) }
func (x *SendVideo) SetChatID(id int64)     { x.ChatId = ID(id) }
func (x *SendVideoNote) SetChatID(id int64) { x.ChatId = ID(id) }
func (x *SendVoice) SetChatID(id int64)     { x.ChatId = ID(id) }

func (x *SendAnimation) Send(api *API) (*Message, error) { return api.SendAnimation(x) }
func (x *SendAudio) Send(api *API) (*Message, error)     { return api.SendAudio(x) }
func (x *SendContact) Send(api *API) (*Message, error)   { return api.SendContact(x) }
func (x *SendDice) Send(api *API) (*Message, error)      { return api.SendDice(x) }
func (x *SendDocument) Send(api *API) (*Message, error)  { return api.SendDocument(x) }
func (x *SendGame) Send(api *API)                        { api.SendGame(x) }
func (x *SendInvoice) Send(api *API) (*Message, error)   { return api.SendInvoice(x) }
func (x *SendLocation) Send(api *API) (*Message, error)  { return api.SendLocation(x) }
func (x *SendMessage) Send(api *API) (*Message, error)   { return api.SendMessage(x) }
func (x *SendPhoto) Send(api *API) (*Message, error)     { return api.SendPhoto(x) }
func (x *SendPoll) Send(api *API) (*Message, error)      { return api.SendPoll(x) }
func (x *SendSticker) Send(api *API) (*Message, error)   { return api.SendSticker(x) }
func (x *SendVenue) Send(api *API) (*Message, error)     { return api.SendVenue(x) }
func (x *SendVideo) Send(api *API) (*Message, error)     { return api.SendVideo(x) }
func (x *SendVideoNote) Send(api *API) (*Message, error) { return api.SendVideoNote(x) }
func (x *SendVoice) Send(api *API) (*Message, error)     { return api.SendVoice(x) }

func (x *SendAnimation) GetParseMode() ParseMode { return x.ParseMode }
func (x *SendAudio) GetParseMode() ParseMode     { return x.ParseMode }
func (x *SendDocument) GetParseMode() ParseMode  { return x.ParseMode }
func (x *SendMessage) GetParseMode() ParseMode   { return x.ParseMode }
func (x *SendPhoto) GetParseMode() ParseMode     { return x.ParseMode }
func (x *SendVideo) GetParseMode() ParseMode     { return x.ParseMode }
func (x *SendVoice) GetParseMode() ParseMode     { return x.ParseMode }

func (x *SendAnimation) SetParseMode(mode ParseMode) { x.ParseMode = mode }
func (x *SendAudio) SetParseMode(mode ParseMode)     { x.ParseMode = mode }
func (x *SendDocument) SetParseMode(mode ParseMode)  { x.ParseMode = mode }
func (x *SendMessage) SetParseMode(mode ParseMode)   { x.ParseMode = mode }
func (x *SendPhoto) SetParseMode(mode ParseMode)     { x.ParseMode = mode }
func (x *SendVideo) SetParseMode(mode ParseMode)     { x.ParseMode = mode }
func (x *SendVoice) SetParseMode(mode ParseMode)     { x.ParseMode = mode }

func (b *Bot) Send(msg Sendable) (*Message, error) {
	if x, ok := msg.(ParseModeSettable); ok {
		if x.GetParseMode() != ParseModeNone {
			x.SetParseMode(b.defaultParseMode)
		}
	}

	return msg.Send(b.API)
}
