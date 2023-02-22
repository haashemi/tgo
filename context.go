package tgo

import (
	"time"
)

type Contextable interface {
	ChatID() int64
	SenderID() int64
	ThreadID() int64
	MessageID() int64
}

type Context interface {
	Contextable

	Stop()
	ResetStopped()
	IsStopped() bool
	UpdateType() string

	Bot() *Bot
	Session() *Session
	RawUpdate() *Update

	Send(text string, optionalParams *SendMessageOptions) (*Message, error)
	SendPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error)
	SendAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error)
	SendDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error)
	SendVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error)
	SendAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error)
	SendVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error)
	SendVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error)
	SendLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error)
	SendVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error)
	SendContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error)
	SendPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error)
	SendDice(optionalParams *SendDiceOptions) (*Message, error)
	SendSticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error)
	SendInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error)
	SendGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error)

	Reply(text string, optionalParams *SendMessageOptions) (*Message, error)
	ReplyPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error)
	ReplyAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error)
	ReplyDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error)
	ReplyVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error)
	ReplyAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error)
	ReplyVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error)
	ReplyVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error)
	ReplyLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error)
	ReplyVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error)
	ReplyContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error)
	ReplyPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error)
	ReplyDice(optionalParams *SendDiceOptions) (*Message, error)
	ReplySticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error)
	ReplyInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error)
	ReplyGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error)

	Ask(text string, optionalParams *SendMessageOptions, timeout time.Duration) (question *Message, answer *messageContext, err error)
}

type baseContext struct {
	Contextable

	bot        *Bot
	update     *Update
	updateType string
	stopped    bool
}

func newContext(bot *Bot, update *Update) Context {
	base := &baseContext{bot: bot, update: update}

	switch {
	case update.Message != nil:
		base.Contextable = update.Message
		base.updateType = "Message"
		return &messageContext{Context: base, Message: update.Message}

	case update.EditedMessage != nil:
		base.updateType = "EditedMessage"
		base.Contextable = update.EditedMessage
		return &messageContext{Context: base, Message: update.EditedMessage}

	case update.ChannelPost != nil:
		base.updateType = "ChannelPost"
		base.Contextable = update.ChannelPost
		return &messageContext{Context: base, Message: update.ChannelPost}

	case update.EditedChannelPost != nil:
		base.updateType = "EditedChannelPost"
		base.Contextable = update.EditedChannelPost
		return &messageContext{Context: base, Message: update.EditedChannelPost}

	// ToDO:
	case update.CallbackQuery != nil:
		return nil
	}

	return nil
}

// Stop stops the context.
// It will only be used for middlewares and NOT the main handlers.
func (ctx *baseContext) Stop() { ctx.stopped = true }

func (ctx *baseContext) ResetStopped() { ctx.stopped = false }

func (ctx *baseContext) IsStopped() bool { return ctx.stopped }

func (ctx *baseContext) UpdateType() string { return ctx.updateType }

// Bot returns the bot
func (ctx *baseContext) Bot() *Bot { return ctx.bot }

// Session returns the user's session storage.
// it will return the chat's session if user-id is zero.
func (ctx *baseContext) Session() *Session {
	id := ctx.SenderID()
	if id == 0 {
		id = ctx.ChatID()
	}

	return ctx.bot.GetSession(id)
}

func (ctx *baseContext) RawUpdate() *Update { return ctx.update }

func (ctx *baseContext) Send(text string, optionalParams *SendMessageOptions) (*Message, error) {
	return ctx.bot.SendMessage(NewChatID(ctx.ChatID()), text, optionalParams)
}

func (ctx *baseContext) SendPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	return ctx.bot.SendPhoto(NewChatID(ctx.ChatID()), photo, optionalParams)
}

func (ctx *baseContext) SendAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	return ctx.bot.SendAudio(NewChatID(ctx.ChatID()), audio, optionalParams)
}

func (ctx *baseContext) SendDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	return ctx.bot.SendDocument(NewChatID(ctx.ChatID()), document, optionalParams)
}

func (ctx *baseContext) SendVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	return ctx.bot.SendVideo(NewChatID(ctx.ChatID()), video, optionalParams)
}

func (ctx *baseContext) SendAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	return ctx.bot.SendAnimation(NewChatID(ctx.ChatID()), animation, optionalParams)
}

func (ctx *baseContext) SendVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	return ctx.bot.SendVoice(NewChatID(ctx.ChatID()), voice, optionalParams)
}

func (ctx *baseContext) SendVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	return ctx.bot.SendVideoNote(NewChatID(ctx.ChatID()), videoNote, optionalParams)
}

func (ctx *baseContext) SendLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	return ctx.bot.SendLocation(NewChatID(ctx.ChatID()), latitude, longitude, optionalParams)
}

func (ctx *baseContext) SendVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error) {
	return ctx.bot.SendVenue(NewChatID(ctx.ChatID()), latitude, longitude, title, address, optionalParams)
}

func (ctx *baseContext) SendContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	return ctx.bot.SendContact(NewChatID(ctx.ChatID()), phoneNumber, firstName, optionalParams)
}

func (ctx *baseContext) SendPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	return ctx.bot.SendPoll(NewChatID(ctx.ChatID()), question, options, optionalParams)
}

func (ctx *baseContext) SendDice(optionalParams *SendDiceOptions) (*Message, error) {
	return ctx.bot.SendDice(NewChatID(ctx.ChatID()), optionalParams)
}

func (ctx *baseContext) SendSticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	return ctx.bot.SendSticker(NewChatID(ctx.ChatID()), sticker, optionalParams)
}

func (ctx *baseContext) SendInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	return ctx.bot.SendInvoice(NewChatID(ctx.ChatID()), title, description, payload, providerToken, currency, prices, optionalParams)
}

func (ctx *baseContext) SendGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	return ctx.bot.SendGame(ctx.ChatID(), gameShortName, optionalParams)
}

// ToDo: SendMediaGroup
// ToDo: SendChatAction

func (ctx *baseContext) Reply(text string, optionalParams *SendMessageOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendMessageOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.Send(text, optionalParams)
}

func (ctx *baseContext) ReplyPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendPhotoOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendPhoto(NewChatID(ctx.ChatID()), photo, optionalParams)
}

func (ctx *baseContext) ReplyAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendAudioOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendAudio(NewChatID(ctx.ChatID()), audio, optionalParams)
}

func (ctx *baseContext) ReplyDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendDocumentOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendDocument(NewChatID(ctx.ChatID()), document, optionalParams)
}

func (ctx *baseContext) ReplyVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVideoOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendVideo(NewChatID(ctx.ChatID()), video, optionalParams)
}

func (ctx *baseContext) ReplyAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendAnimationOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendAnimation(NewChatID(ctx.ChatID()), animation, optionalParams)
}

func (ctx *baseContext) ReplyVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVoiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendVoice(NewChatID(ctx.ChatID()), voice, optionalParams)
}

func (ctx *baseContext) ReplyVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVideoNoteOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendVideoNote(NewChatID(ctx.ChatID()), videoNote, optionalParams)
}

func (ctx *baseContext) ReplyLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendLocationOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendLocation(NewChatID(ctx.ChatID()), latitude, longitude, optionalParams)
}

func (ctx *baseContext) ReplyVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVenueOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendVenue(NewChatID(ctx.ChatID()), latitude, longitude, title, address, optionalParams)
}

func (ctx *baseContext) ReplyContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendContactOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendContact(NewChatID(ctx.ChatID()), phoneNumber, firstName, optionalParams)
}

func (ctx *baseContext) ReplyPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendPollOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendPoll(NewChatID(ctx.ChatID()), question, options, optionalParams)
}

func (ctx *baseContext) ReplyDice(optionalParams *SendDiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendDiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendDice(NewChatID(ctx.ChatID()), optionalParams)
}

func (ctx *baseContext) ReplySticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendStickerOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendSticker(NewChatID(ctx.ChatID()), sticker, optionalParams)
}

func (ctx *baseContext) ReplyInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendInvoiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendInvoice(NewChatID(ctx.ChatID()), title, description, payload, providerToken, currency, prices, optionalParams)
}

func (ctx *baseContext) ReplyGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendGameOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.bot.SendGame(ctx.ChatID(), gameShortName, optionalParams)
}

// ToDo: ReplyMediaGroup
// ToDo: ReplyChatAction

func (ctx *baseContext) Ask(text string, optionalParams *SendMessageOptions, timeout time.Duration) (question *Message, answer *messageContext, err error) {
	question, err = ctx.Send(text, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// ToDo: AskWithPhoto
// ToDo: AskWithAudio
// ToDo: AskWithDocument
// ToDo: AskWithVideo
// ToDo: AskWithAnimation
// ToDo: AskWithVoice
// ToDo: AskWithVideoNote
// ToDo: AskWithLocation
// ToDo: AskWithVenue
// ToDo: AskWithContact
// ToDo: AskWithPoll
// ToDo: AskWithDice
// ToDo: AskWithSticker
// ToDo: AskWithInvoice
// ToDo: AskWithGame
