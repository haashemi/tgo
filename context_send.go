package tgo

func (ctx *UpdateContext) Send(text string, optionalParams *SendMessageOptions) (*Message, error) {
	return ctx.bot.SendMessage(ChatID(ctx.ChatID()), text, optionalParams)
}

func (ctx *UpdateContext) SendPhoto(photo InputFile, optionalParams *SendPhotoOptions) (*Message, error) {
	return ctx.bot.SendPhoto(ChatID(ctx.ChatID()), photo, optionalParams)
}

func (ctx *UpdateContext) SendAudio(audio InputFile, optionalParams *SendAudioOptions) (*Message, error) {
	return ctx.bot.SendAudio(ChatID(ctx.ChatID()), audio, optionalParams)
}

func (ctx *UpdateContext) SendDocument(document InputFile, optionalParams *SendDocumentOptions) (*Message, error) {
	return ctx.bot.SendDocument(ChatID(ctx.ChatID()), document, optionalParams)
}

func (ctx *UpdateContext) SendVideo(video InputFile, optionalParams *SendVideoOptions) (*Message, error) {
	return ctx.bot.SendVideo(ChatID(ctx.ChatID()), video, optionalParams)
}

func (ctx *UpdateContext) SendAnimation(animation InputFile, optionalParams *SendAnimationOptions) (*Message, error) {
	return ctx.bot.SendAnimation(ChatID(ctx.ChatID()), animation, optionalParams)
}

func (ctx *UpdateContext) SendVoice(voice InputFile, optionalParams *SendVoiceOptions) (*Message, error) {
	return ctx.bot.SendVoice(ChatID(ctx.ChatID()), voice, optionalParams)
}

func (ctx *UpdateContext) SendVideoNote(videoNote InputFile, optionalParams *SendVideoNoteOptions) (*Message, error) {
	return ctx.bot.SendVideoNote(ChatID(ctx.ChatID()), videoNote, optionalParams)
}

func (ctx *UpdateContext) SendLocation(latitude, longitude float64, optionalParams *SendLocationOptions) (*Message, error) {
	return ctx.bot.SendLocation(ChatID(ctx.ChatID()), latitude, longitude, optionalParams)
}

func (ctx *UpdateContext) SendVenue(latitude, longitude float64, title, address string, optionalParams *SendVenueOptions) (*Message, error) {
	return ctx.bot.SendVenue(ChatID(ctx.ChatID()), latitude, longitude, title, address, optionalParams)
}

func (ctx *UpdateContext) SendContact(phoneNumber, firstName string, optionalParams *SendContactOptions) (*Message, error) {
	return ctx.bot.SendContact(ChatID(ctx.ChatID()), phoneNumber, firstName, optionalParams)
}

func (ctx *UpdateContext) SendPoll(question string, options []string, optionalParams *SendPollOptions) (*Message, error) {
	return ctx.bot.SendPoll(ChatID(ctx.ChatID()), question, options, optionalParams)
}

func (ctx *UpdateContext) SendDice(optionalParams *SendDiceOptions) (*Message, error) {
	return ctx.bot.SendDice(ChatID(ctx.ChatID()), optionalParams)
}

func (ctx *UpdateContext) SendSticker(sticker InputFile, optionalParams *SendStickerOptions) (*Message, error) {
	return ctx.bot.SendSticker(ChatID(ctx.ChatID()), sticker, optionalParams)
}

func (ctx *UpdateContext) SendInvoice(title, description, payload, providerToken, currency string, prices []*LabeledPrice, optionalParams *SendInvoiceOptions) (*Message, error) {
	return ctx.bot.SendInvoice(ChatID(ctx.ChatID()), title, description, payload, providerToken, currency, prices, optionalParams)
}

func (ctx *UpdateContext) SendGame(gameShortName string, optionalParams *SendGameOptions) (*Message, error) {
	return ctx.bot.SendGame(ctx.ChatID(), gameShortName, optionalParams)
}

// ToDo: implement SendMediaGroup

// ToDo: create an enum for the actions
func (ctx *UpdateContext) SendChatAction(action string, optionalParams *SendChatActionOptions) error {
	_, err := ctx.bot.SendChatAction(ctx.ChatID(), action, optionalParams)
	return err
}
