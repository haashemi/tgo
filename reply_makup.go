package tgo

type ReplyMarkup interface {
	// IsReplyMarkup does nothing and is only used to enforce type-safety
	IsReplyMarkup()
}

func NewInlineKeyboardMarkup(rows ...[]*InlineKeyboardButton) *InlineKeyboardMarkup {
	return &InlineKeyboardMarkup{InlineKeyboard: rows}
}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (InlineKeyboardMarkup) IsReplyMarkup() {}

func NewReplyKeyboardMarkup(rows ...[]*KeyboardButton) *ReplyKeyboardMarkup {
	return &ReplyKeyboardMarkup{Keyboard: rows}
}

func (keyboard *ReplyKeyboardMarkup) Persistent() *ReplyKeyboardMarkup {
	keyboard.IsPersistent = true
	return keyboard
}

func (keyboard *ReplyKeyboardMarkup) WithResizeKeyboard() *ReplyKeyboardMarkup {
	keyboard.ResizeKeyboard = true
	return keyboard
}

func (keyboard *ReplyKeyboardMarkup) WithOneTimeKeyboard() *ReplyKeyboardMarkup {
	keyboard.OneTimeKeyboard = true
	return keyboard
}

func (keyboard *ReplyKeyboardMarkup) WithInputFieldPlaceholder(placeholder string) *ReplyKeyboardMarkup {
	keyboard.InputFieldPlaceholder = placeholder
	return keyboard
}

func (keyboard *ReplyKeyboardMarkup) IsSelective() *ReplyKeyboardMarkup {
	keyboard.Selective = true
	return keyboard
}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ReplyKeyboardMarkup) IsReplyMarkup() {}

func NewReplyKeyboardRemove() *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{RemoveKeyboard: true}
}

func (data *ReplyKeyboardRemove) IsSelective() *ReplyKeyboardRemove {
	data.Selective = true
	return data
}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ReplyKeyboardRemove) IsReplyMarkup() {}

func NewForceReply() *ForceReply {
	return &ForceReply{ForceReply: true}
}

func (data *ForceReply) WithPlaceholder(placeholder string) *ForceReply {
	data.InputFieldPlaceholder = placeholder
	return data
}

func (data *ForceReply) IsSelective() *ForceReply {
	data.Selective = true
	return data
}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ForceReply) IsReplyMarkup() {}

//////////////////////////////////////////
//////////////////////////////////////////
//////////////////////////////////////////

func NewInlineKeyboardRow(buttons ...*InlineKeyboardButton) []*InlineKeyboardButton {
	return buttons
}

func NewInlineKeyboardButton(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text}
}

func (btn *InlineKeyboardButton) WithUrl(param string) *InlineKeyboardButton {
	btn.Url = param
	return btn
}

func (btn *InlineKeyboardButton) WithCallbackData(param string) *InlineKeyboardButton {
	btn.CallbackData = param
	return btn
}

func (btn *InlineKeyboardButton) WithWebApp(webAppURL string) *InlineKeyboardButton {
	btn.WebApp = &WebAppInfo{Url: webAppURL}
	return btn
}

func (btn *InlineKeyboardButton) WithLoginUrl(param LoginUrl) *InlineKeyboardButton {
	btn.LoginUrl = &param
	return btn
}

func (btn *InlineKeyboardButton) WithSwitchInlineQuery(param string) *InlineKeyboardButton {
	btn.SwitchInlineQuery = param
	return btn
}

func (btn *InlineKeyboardButton) WithSwitchInlineQueryCurrentChat(param string) *InlineKeyboardButton {
	btn.SwitchInlineQueryCurrentChat = param
	return btn
}

func (btn *InlineKeyboardButton) WithCallbackGame() *InlineKeyboardButton {
	btn.CallbackGame = &CallbackGame{}
	return btn
}

func (btn *InlineKeyboardButton) WithPay() *InlineKeyboardButton {
	btn.Pay = true
	return btn
}

func NewReplyKeyboardRow(buttons ...*KeyboardButton) []*KeyboardButton {
	return buttons
}

func NewKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{Text: text}
}

func (btn *KeyboardButton) WithRequestUser(param *KeyboardButtonRequestUser) *KeyboardButton {
	btn.RequestUser = param
	return btn
}

func (btn *KeyboardButton) WithRequestChat(param *KeyboardButtonRequestChat) *KeyboardButton {
	btn.RequestChat = param
	return btn
}

func (btn *KeyboardButton) WithRequestContact() *KeyboardButton {
	btn.RequestContact = true
	return btn
}

func (btn *KeyboardButton) WithRequestLocation() *KeyboardButton {
	btn.RequestLocation = true
	return btn
}

type PollType string

const (
	PollTypeAny     PollType = ""        // If this gets passed, the user will be allowed to create a poll of any type.
	PollTypeQuiz    PollType = "quiz"    // if this gets passed, the user will be allowed to create only polls in the quiz mode.
	PollTypeRegular PollType = "regular" // If this gets passed, only regular polls will be allowed.
)

func (btn *KeyboardButton) WithRequestPoll(pollType PollType) *KeyboardButton {
	btn.RequestPoll = &KeyboardButtonPollType{Type: string(pollType)}
	return btn
}

func (btn *KeyboardButton) WithWebApp(webAppURL string) *KeyboardButton {
	btn.WebApp = &WebAppInfo{Url: webAppURL}
	return btn
}
