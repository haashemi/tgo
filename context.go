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

type Context struct {
	Contextable

	bot     *Bot
	update  *Update
	stopped bool
}

// Stop stops the context.
// It will only be used for middlewares and NOT the main handlers.
func (ctx *Context) Stop() { ctx.stopped = true }

func (ctx *Context) ResetStopped() { ctx.stopped = false }

func (ctx *Context) IsStopped() bool { return ctx.stopped }

// Bot returns the bot
func (ctx *Context) Bot() *Bot { return ctx.bot }

// Session returns the user's session storage.
// it will return the chat's session if user-id is zero.
func (ctx *Context) Session() *Session {
	id := ctx.SenderID()
	if id == 0 {
		id = ctx.ChatID()
	}

	return ctx.bot.GetSession(id)
}

func (ctx *Context) RawUpdate() *Update { return ctx.update }

func (ctx *Context) Send(text string, optionalParams *SendMessageOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendMessageOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendMessage(NewChatID(ctx.ChatID()), text, optionalParams)
}

func (ctx *Context) SendPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendPhotoOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendPhoto(NewChatID(ctx.ChatID()), photo, optionalParams)
}

func (ctx *Context) SendAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendAudioOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendAudio(NewChatID(ctx.ChatID()), audio, optionalParams)
}

func (ctx *Context) SendDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendDocumentOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendDocument(NewChatID(ctx.ChatID()), document, optionalParams)
}

func (ctx *Context) SendVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendVideoOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendVideo(NewChatID(ctx.ChatID()), video, optionalParams)
}

func (ctx *Context) SendAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendAnimationOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendAnimation(NewChatID(ctx.ChatID()), animation, optionalParams)
}

func (ctx *Context) SendVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendVoiceOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendVoice(NewChatID(ctx.ChatID()), voice, optionalParams)
}

func (ctx *Context) SendVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendVideoNoteOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendVideoNote(NewChatID(ctx.ChatID()), videoNote, optionalParams)
}

func (ctx *Context) SendLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendLocationOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendLocation(NewChatID(ctx.ChatID()), latitude, longitude, optionalParams)
}

func (ctx *Context) SendVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendVenueOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendVenue(NewChatID(ctx.ChatID()), latitude, longitude, title, address, optionalParams)
}

func (ctx *Context) SendContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendContactOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendContact(NewChatID(ctx.ChatID()), phoneNumber, firstName, optionalParams)
}

func (ctx *Context) SendPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendPollOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendPoll(NewChatID(ctx.ChatID()), question, options, optionalParams)
}

func (ctx *Context) SendDice(optionalParams *SendDiceOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendDiceOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendDice(NewChatID(ctx.ChatID()), optionalParams)
}

func (ctx *Context) SendSticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendStickerOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendSticker(NewChatID(ctx.ChatID()), sticker, optionalParams)
}

func (ctx *Context) SendInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendInvoiceOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendInvoice(NewChatID(ctx.ChatID()), title, description, payload, providerToken, currency, prices, optionalParams)
}

func (ctx *Context) SendGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	if threadID := ctx.ThreadID(); threadID != 0 {
		if optionalParams == nil {
			optionalParams = &SendGameOptions{}
		}
		optionalParams.MessageThreadId = threadID
	}
	return ctx.bot.SendGame(ctx.ChatID(), gameShortName, optionalParams)
}

// ToDo: SendMediaGroup
// ToDo: SendChatAction

func (ctx *Context) Reply(text string, optionalParams *SendMessageOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendMessageOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.Send(text, optionalParams)
}

func (ctx *Context) ReplyPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendPhotoOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendPhoto(photo, optionalParams)
}

func (ctx *Context) ReplyAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendAudioOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendAudio(audio, optionalParams)
}

func (ctx *Context) ReplyDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendDocumentOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendDocument(document, optionalParams)
}

func (ctx *Context) ReplyVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVideoOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVideo(video, optionalParams)
}

func (ctx *Context) ReplyAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendAnimationOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendAnimation(animation, optionalParams)
}

func (ctx *Context) ReplyVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVoiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVoice(voice, optionalParams)
}

func (ctx *Context) ReplyVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVideoNoteOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVideoNote(videoNote, optionalParams)
}

func (ctx *Context) ReplyLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendLocationOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendLocation(latitude, longitude, optionalParams)
}

func (ctx *Context) ReplyVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVenueOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVenue(latitude, longitude, title, address, optionalParams)
}

func (ctx *Context) ReplyContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendContactOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendContact(phoneNumber, firstName, optionalParams)
}

func (ctx *Context) ReplyPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendPollOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendPoll(question, options, optionalParams)
}

func (ctx *Context) ReplyDice(optionalParams *SendDiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendDiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendDice(optionalParams)
}

func (ctx *Context) ReplySticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendStickerOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendSticker(sticker, optionalParams)
}

func (ctx *Context) ReplyInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendInvoiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendInvoice(title, description, payload, providerToken, currency, prices, optionalParams)
}

func (ctx *Context) ReplyGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendGameOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendGame(gameShortName, optionalParams)
}

// ToDo: ReplyMediaGroup
// ToDo: ReplyChatAction

func (ctx *Context) Ask(text string, optionalParams *SendMessageOptions, timeout time.Duration) (question *Message, answer *messageContext, err error) {
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
func (ctx *Context) Delete() error {
	_, err := ctx.bot.DeleteMessage(NewChatID(ctx.ChatID()), ctx.MessageID())
	return err
}
