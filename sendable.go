package tgo

type SendableMessage interface {
	SetChatID(id ChatID) *SendableText
	SetThreadID(id int64) *SendableText
	SetReplyToMessageID(id int64) *SendableText
	Send(api *API) (*Message, error)
}

// SendableText implements SendableMessage for api.SendMessage
type SendableText struct{ params SendMessageParams }

type TextOptions struct {
	Entities                 []*MessageEntity // List of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview    bool             // Disables link previews for links in this message
	DisableNotification      bool             // Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool             // Protects the contents of the sent message from forwarding and saving
	AllowSendingWithoutReply bool             // Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              ReplyMarkup      // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func NewText(text string, parseMode ...ParseMode) *SendableText {
	data := &SendableText{params: SendMessageParams{Text: text}}

	// Just use the latest passed ParseMode for whatever reason
	// as we want this method to be optional.
	for _, pm := range parseMode {
		data.params.ParseMode = pm
	}

	return data
}

func (d *SendableText) Options(opts TextOptions) *SendableText {
	d.params.Entities = opts.Entities
	d.params.DisableWebPagePreview = opts.DisableWebPagePreview
	d.params.DisableNotification = opts.DisableNotification
	d.params.ProtectContent = opts.ProtectContent
	d.params.AllowSendingWithoutReply = opts.AllowSendingWithoutReply
	d.params.ReplyMarkup = opts.ReplyMarkup

	return d
}

func (d *SendableText) SetChatID(id ChatID) *SendableText {
	d.params.ChatId = id
	return d
}

func (d *SendableText) SetThreadID(id int64) *SendableText {
	d.params.MessageThreadId = id
	return d
}

func (d *SendableText) SetReplyToMessageID(id int64) *SendableText {
	d.params.ReplyToMessageId = id
	return d
}

func (d *SendableText) Send(api *API) (*Message, error) {
	return api.SendMessage(d.params)
}
