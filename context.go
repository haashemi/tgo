package tgo

import (
	"sync"
	"time"
)

type Contextable interface {
	ChatID() int64
	SenderID() int64
	ThreadID() int64
	MessageID() int64
}

type Context = *botContext

type botContext struct {
	Contextable

	bot     *Bot
	update  *Update
	stopped bool
}

// Stop stops the context.
// It will only be used for middlewares and NOT the main handlers.
func (ctx *botContext) Stop() { ctx.stopped = true }

func (ctx *botContext) ResetStopped() { ctx.stopped = false }

func (ctx *botContext) IsStopped() bool { return ctx.stopped }

// Bot returns the bot
func (ctx *botContext) Bot() *Bot { return ctx.bot }

// Session returns the user's session storage.
// it will return the chat's session if user-id is zero.
func (ctx *botContext) Session() *sync.Map {
	id := ctx.SenderID()
	if id == 0 {
		id = ctx.ChatID()
	}

	return ctx.bot.GetSession(id)
}

func (ctx *botContext) Update() *Update {
	return ctx.update
}

func (ctx *botContext) Message() *Message {
	if ctx.update.Message != nil {
		return ctx.update.Message
	} else if ctx.update.EditedMessage != nil {
		return ctx.update.EditedMessage
	} else if ctx.update.ChannelPost != nil {
		return ctx.update.ChannelPost
	}
	return ctx.update.EditedChannelPost
}

func (ctx *botContext) InlineQuery() *InlineQuery {
	return ctx.update.InlineQuery
}

func (ctx *botContext) ChosenInlineResult() *ChosenInlineResult {
	return ctx.update.ChosenInlineResult
}

func (ctx *botContext) CallbackQuery() *CallbackQuery {
	return ctx.update.CallbackQuery
}

func (ctx *botContext) ShippingQuery() *ShippingQuery {
	return ctx.update.ShippingQuery
}

func (ctx *botContext) PreCheckoutQuery() *PreCheckoutQuery {
	return ctx.update.PreCheckoutQuery
}

func (ctx *botContext) Poll() *Poll {
	return ctx.update.Poll
}

func (ctx *botContext) PollAnswer() *PollAnswer {
	return ctx.update.PollAnswer
}

func (ctx *botContext) ChatMemberUpdated() *ChatMemberUpdated {
	if ctx.update.MyChatMember != nil {
		return ctx.update.MyChatMember
	}

	return ctx.update.ChatMember
}

func (ctx *botContext) ChatJoinRequest() *ChatJoinRequest {
	return ctx.update.ChatJoinRequest
}

// Text returns the message's text or media caption or callback query's data
func (ctx *botContext) Text() string {
	if msg := ctx.Message(); msg != nil {
		return msg.String()
	} else if query := ctx.CallbackQuery(); query != nil {
		return query.Data
	}

	return ""
}

func (ctx *botContext) Send(text string, optionalParams *SendMessageOptions) (*Message, error) {
	return ctx.bot.SendMessage(NewChatID(ctx.ChatID()), text, optionalParams)
}

func (ctx *botContext) SendPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	return ctx.bot.SendPhoto(NewChatID(ctx.ChatID()), photo, optionalParams)
}

func (ctx *botContext) SendAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	return ctx.bot.SendAudio(NewChatID(ctx.ChatID()), audio, optionalParams)
}

func (ctx *botContext) SendDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	return ctx.bot.SendDocument(NewChatID(ctx.ChatID()), document, optionalParams)
}

func (ctx *botContext) SendVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	return ctx.bot.SendVideo(NewChatID(ctx.ChatID()), video, optionalParams)
}

func (ctx *botContext) SendAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	return ctx.bot.SendAnimation(NewChatID(ctx.ChatID()), animation, optionalParams)
}

func (ctx *botContext) SendVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	return ctx.bot.SendVoice(NewChatID(ctx.ChatID()), voice, optionalParams)
}

func (ctx *botContext) SendVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	return ctx.bot.SendVideoNote(NewChatID(ctx.ChatID()), videoNote, optionalParams)
}

func (ctx *botContext) SendLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	return ctx.bot.SendLocation(NewChatID(ctx.ChatID()), latitude, longitude, optionalParams)
}

func (ctx *botContext) SendVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error) {
	return ctx.bot.SendVenue(NewChatID(ctx.ChatID()), latitude, longitude, title, address, optionalParams)
}

func (ctx *botContext) SendContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	return ctx.bot.SendContact(NewChatID(ctx.ChatID()), phoneNumber, firstName, optionalParams)
}

func (ctx *botContext) SendPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	return ctx.bot.SendPoll(NewChatID(ctx.ChatID()), question, options, optionalParams)
}

func (ctx *botContext) SendDice(optionalParams *SendDiceOptions) (*Message, error) {
	return ctx.bot.SendDice(NewChatID(ctx.ChatID()), optionalParams)
}

func (ctx *botContext) SendSticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	return ctx.bot.SendSticker(NewChatID(ctx.ChatID()), sticker, optionalParams)
}

func (ctx *botContext) SendInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	return ctx.bot.SendInvoice(NewChatID(ctx.ChatID()), title, description, payload, providerToken, currency, prices, optionalParams)
}

func (ctx *botContext) SendGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	return ctx.bot.SendGame(ctx.ChatID(), gameShortName, optionalParams)
}

// ToDo: SendMediaGroup
// ToDo: SendChatAction

func (ctx *botContext) Reply(text string, optionalParams *SendMessageOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendMessageOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.Send(text, optionalParams)
}

func (ctx *botContext) ReplyPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendPhotoOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendPhoto(photo, optionalParams)
}

func (ctx *botContext) ReplyAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendAudioOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendAudio(audio, optionalParams)
}

func (ctx *botContext) ReplyDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendDocumentOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendDocument(document, optionalParams)
}

func (ctx *botContext) ReplyVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVideoOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVideo(video, optionalParams)
}

func (ctx *botContext) ReplyAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendAnimationOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendAnimation(animation, optionalParams)
}

func (ctx *botContext) ReplyVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVoiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVoice(voice, optionalParams)
}

func (ctx *botContext) ReplyVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVideoNoteOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVideoNote(videoNote, optionalParams)
}

func (ctx *botContext) ReplyLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendLocationOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendLocation(latitude, longitude, optionalParams)
}

func (ctx *botContext) ReplyVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVenueOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVenue(latitude, longitude, title, address, optionalParams)
}

func (ctx *botContext) ReplyContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendContactOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendContact(phoneNumber, firstName, optionalParams)
}

func (ctx *botContext) ReplyPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendPollOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendPoll(question, options, optionalParams)
}

func (ctx *botContext) ReplyDice(optionalParams *SendDiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendDiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendDice(optionalParams)
}

func (ctx *botContext) ReplySticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendStickerOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendSticker(sticker, optionalParams)
}

func (ctx *botContext) ReplyInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendInvoiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendInvoice(title, description, payload, providerToken, currency, prices, optionalParams)
}

func (ctx *botContext) ReplyGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendGameOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendGame(gameShortName, optionalParams)
}

// ToDo: ReplyMediaGroup
// ToDo: ReplyChatAction

func (ctx *botContext) Ask(text string, optionalParams *SendMessageOptions, timeout time.Duration) (question *Message, answer *botContext, err error) {
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

// Delete deletes the message of that context
func (ctx *botContext) Delete() error {
	_, err := ctx.bot.DeleteMessage(NewChatID(ctx.ChatID()), ctx.MessageID())
	return err
}

// Answer answers to callback queries sent from inline keyboards
func (ctx *botContext) Answer(options *AnswerCallbackQueryOptions) error {
	_, err := ctx.bot.AnswerCallbackQuery(ctx.CallbackQuery().Id, options)
	return err
}
