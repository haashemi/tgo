package tgo

type ReplyMarkup interface {
	// IsReplyMarkup does nothing and is only used to enforce type-safety
	IsReplyMarkup()
}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (InlineKeyboardMarkup) IsReplyMarkup() {}

// NewInlineKeyboardMarkup returns an InlineKeyboardMarkup which represents
// an inline keyboard that appears right next to the message it belongs to.
func NewInlineKeyboardMarkup(rows ...[]*InlineKeyboardButton) ReplyMarkup {
	return &InlineKeyboardMarkup{InlineKeyboard: rows}
}

func NewInlineKeyboardRow(buttons ...*InlineKeyboardButton) []*InlineKeyboardButton {
	return buttons
}

func UrlInlineKeyboardButton(text string, url string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, Url: url}
}

func CallbackDataInlineKeyboardButton(text string, callbackData string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, CallbackData: callbackData}
}

func WebAppInlineKeyboardButton(text string, webAppURL string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, WebApp: &WebAppInfo{Url: webAppURL}}
}

func LoginUrlInlineKeyboardButton(text string, loginUrl *LoginUrl) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, LoginUrl: loginUrl}
}

func SwitchInlineQueryInlineKeyboardButton(text string, switchInlineQuery string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, SwitchInlineQuery: switchInlineQuery}
}

func SwitchInlineQueryCurrentChatInlineKeyboardButton(text string, switchInlineQueryCurrentChat string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, SwitchInlineQueryCurrentChat: switchInlineQueryCurrentChat}
}

func CallbackGameInlineKeyboardButton(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, CallbackGame: &CallbackGame{}}
}

func PayInlineKeyboardButton(text string) *InlineKeyboardButton {
	return &InlineKeyboardButton{Text: text, Pay: true}
}

type ReplyKeyboardOptionSetter func(*ReplyKeyboardMarkup)

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ReplyKeyboardMarkup) IsReplyMarkup() {}

func NewReplyKeyboardMarkup(rows [][]*KeyboardButton, options ...ReplyKeyboardOptionSetter) ReplyMarkup {
	keyboard := &ReplyKeyboardMarkup{Keyboard: rows}

	for _, setOption := range options {
		setOption(keyboard)
	}

	return keyboard
}

func KeyboardIsPersistent() ReplyKeyboardOptionSetter {
	return func(rkm *ReplyKeyboardMarkup) {
		rkm.IsPersistent = true
	}
}

func KeyboardIsResized() ReplyKeyboardOptionSetter {
	return func(rkm *ReplyKeyboardMarkup) {
		rkm.ResizeKeyboard = true
	}
}

func KeyboardIsOneTime() ReplyKeyboardOptionSetter {
	return func(rkm *ReplyKeyboardMarkup) {
		rkm.OneTimeKeyboard = true
	}
}

func KeyboardIsSelective() ReplyKeyboardOptionSetter {
	return func(rkm *ReplyKeyboardMarkup) {
		rkm.Selective = true
	}
}

func KeyboardPlaceholder(placeholder string) ReplyKeyboardOptionSetter {
	return func(rkm *ReplyKeyboardMarkup) {
		rkm.InputFieldPlaceholder = placeholder
	}
}

func NewReplyKeyboardRow(buttons ...*KeyboardButton) []*KeyboardButton {
	return buttons
}

func TextKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{Text: text}
}
func RequestUserKeyboardButton(text string, requestUser *KeyboardButtonRequestUser) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestUser: requestUser}
}
func RequestChatKeyboardButton(text string, requestChat *KeyboardButtonRequestChat) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestChat: requestChat}
}
func RequestContactKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestContact: true}
}
func RequestLocationKeyboardButton(text string) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestLocation: true}
}
func RequestPollKeyboardButton(text string, pollType PollType) *KeyboardButton {
	return &KeyboardButton{Text: text, RequestPoll: &KeyboardButtonPollType{Type: string(pollType)}}
}
func WebAppKeyboardButton(text string, webAppURL string) *KeyboardButton {
	return &KeyboardButton{Text: text, WebApp: &WebAppInfo{Url: webAppURL}}
}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ReplyKeyboardRemove) IsReplyMarkup() {}

func NewReplyKeyboardRemove(selective bool) *ReplyKeyboardRemove {
	return &ReplyKeyboardRemove{RemoveKeyboard: true, Selective: selective}
}

type ForceReplyOptionSetter func(*ForceReply)

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ForceReply) IsReplyMarkup() {}

// NewForceReply returns a ForceReply object which results Telegram clients to display
// a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces
// without having to sacrifice privacy mode.
func NewForceReply(options ...ForceReplyOptionSetter) ReplyMarkup {
	markup := &ForceReply{ForceReply: true}

	for _, setOption := range options {
		setOption(markup)
	}

	return markup
}

func ForceReplyPlaceholder(placeholder string) ForceReplyOptionSetter {
	return func(fr *ForceReply) {
		fr.InputFieldPlaceholder = placeholder
	}
}

func ForceReplyIsSelective() ForceReplyOptionSetter {
	return func(fr *ForceReply) {
		fr.Selective = true
	}
}
