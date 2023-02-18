package tgo

import "github.com/haashemi/tgo/botapi"

type SendableMessage interface {
	SetChatID(id botapi.ChatID) *SendableText
	SetThreadID(id int64) *SendableText
	SetReplyToMessageID(id int64) *SendableText
	Send(api *botapi.API) (*botapi.Message, error)
}

// SendableText implements SendableMessage for api.SendMessage
type SendableText struct{ params botapi.SendMessageParams }

type TextOptions struct {
	Entities                 []*botapi.MessageEntity `json:"entities,omitempty"`                    // Optional. A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	DisableWebPagePreview    bool                    `json:"disable_web_page_preview,omitempty"`    // Optional. Disables link previews for links in this message
	DisableNotification      bool                    `json:"disable_notification,omitempty"`        // Optional. Sends the message silently. Users will receive a notification with no sound.
	ProtectContent           bool                    `json:"protect_content,omitempty"`             // Optional. Protects the contents of the sent message from forwarding and saving
	AllowSendingWithoutReply bool                    `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified replied-to message is not found
	ReplyMarkup              botapi.InlineKeyboard   `json:"reply_markup,omitempty"`                // Optional. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

func NewText(text string, parseMode ...ParseMode) *SendableText {
	data := &SendableText{params: botapi.SendMessageParams{Text: text}}

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

func (d *SendableText) SetChatID(id botapi.ChatID) *SendableText {
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

func (d *SendableText) Send(api *botapi.API) (*botapi.Message, error) {
	return api.SendMessage(d.params)
}
