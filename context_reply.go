package tgo

func (ctx *UpdateContext) Reply(text string, optionalParams *SendMessageOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendMessageOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.Send(text, optionalParams)
}

func (ctx *UpdateContext) ReplyPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendPhotoOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendPhoto(photo, optionalParams)
}

func (ctx *UpdateContext) ReplyAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendAudioOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendAudio(audio, optionalParams)
}

func (ctx *UpdateContext) ReplyDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendDocumentOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendDocument(document, optionalParams)
}

func (ctx *UpdateContext) ReplyVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVideoOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVideo(video, optionalParams)
}

func (ctx *UpdateContext) ReplyAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendAnimationOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendAnimation(animation, optionalParams)
}

func (ctx *UpdateContext) ReplyVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVoiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVoice(voice, optionalParams)
}

func (ctx *UpdateContext) ReplyVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVideoNoteOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVideoNote(videoNote, optionalParams)
}

func (ctx *UpdateContext) ReplyLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendLocationOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendLocation(latitude, longitude, optionalParams)
}

func (ctx *UpdateContext) ReplyVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendVenueOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendVenue(latitude, longitude, title, address, optionalParams)
}

func (ctx *UpdateContext) ReplyContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendContactOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendContact(phoneNumber, firstName, optionalParams)
}

func (ctx *UpdateContext) ReplyPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendPollOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendPoll(question, options, optionalParams)
}

func (ctx *UpdateContext) ReplyDice(optionalParams *SendDiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendDiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendDice(optionalParams)
}

func (ctx *UpdateContext) ReplySticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendStickerOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendSticker(sticker, optionalParams)
}

func (ctx *UpdateContext) ReplyInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendInvoiceOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendInvoice(title, description, payload, providerToken, currency, prices, optionalParams)
}

func (ctx *UpdateContext) ReplyGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	if optionalParams == nil {
		optionalParams = &SendGameOptions{}
	}
	optionalParams.ReplyToMessageId = ctx.MessageID()
	return ctx.SendGame(gameShortName, optionalParams)
}

// ToDo: ReplyMediaGroup
// ToDo: ReplyChatAction
