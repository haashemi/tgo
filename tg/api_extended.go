package tg

// ParseMode is a type that represents the parse mode of a message.
// eg. Markdown, HTML, etc.
type ParseMode string

const (
	// does not parse the message
	ParseModeNone ParseMode = ""
	// parses the message as Markdown
	ParseModeMarkdown ParseMode = "Markdown"
	// parses the message as Markdown but using telegram's V2 markdown
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	// parses the message as HTML
	ParseModeHTML ParseMode = "HTML"
)

// Username implements a string ChatID.
type Username string

// IsChatID does nothing but implements ChatID interface.
func (Username) IsChatID() {}

// ID implements a int64 ChatID.
type ID int64

// IsChatID does nothing but implements ChatID interface.
func (ID) IsChatID() {}
