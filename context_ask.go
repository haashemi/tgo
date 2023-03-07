// this file should be edited so much.
// Some of these methods should completely get removed.
// and some of them should be managed differently (eg: AskWithPoll, we should wait for vote, not message)
// So, please just use the Ask method, or others like AskWith Photo, Audio, Doc, and Video.

package tgo

import "time"

func (ctx *UpdateContext) Ask(text string, optionalParams *SendMessageOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.Send(text, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

func (ctx *UpdateContext) AskWithPhoto(photo InputFile, optionalParams *SendPhotoOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendPhoto(photo, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

func (ctx *UpdateContext) AskWithAudio(audio InputFile, optionalParams *SendAudioOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendAudio(audio, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

func (ctx *UpdateContext) AskWithDocument(document InputFile, optionalParams *SendDocumentOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendDocument(document, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

func (ctx *UpdateContext) AskWithVideo(video InputFile, optionalParams *SendVideoOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendVideo(video, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithAnimation(animation InputFile, optionalParams *SendAnimationOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendAnimation(animation, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

func (ctx *UpdateContext) AskWithVoice(voice InputFile, optionalParams *SendVoiceOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendVoice(voice, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendVideoNote(videoNote, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithLocation(latitude, longitude float64, optionalParams *SendLocationOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendLocation(latitude, longitude, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendVenue(latitude, longitude, title, address, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithContact(phoneNumber, firstName string, optionalParams *SendContactOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendContact(phoneNumber, firstName, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithPoll(questionText string, options []string, optionalParams *SendPollOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendPoll(questionText, options, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithDice(params *SendDiceOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendDice(params)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithSticker(sticker InputFile, optionalParams *SendStickerOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendSticker(sticker, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendInvoice(title, description, payload, providerToken, currency, prices, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}

// MAY GET REMOVED OR CHANGED ANYTIME
func (ctx *UpdateContext) AskWithGame(gameShortName string, optionalParams *SendGameOptions, timeout time.Duration) (question *Message, answer *UpdateContext, err error) {
	question, err = ctx.SendGame(gameShortName, optionalParams)
	if err != nil {
		return nil, nil, err
	}

	answer, err = ctx.bot.waitForAnswer(question, timeout)
	return
}
