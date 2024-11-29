package tg

import "io"

// Username implements a string ChatID.
type Username string

// IsChatID does nothing but implements ChatID interface.
func (Username) IsChatID() {}

// ID implements a int64 ChatID.
type ID int64

// IsChatID does nothing but implements ChatID interface.
func (ID) IsChatID() {}

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

// InputFile represents a file that can be used as input.
type InputFile struct {
	Value  string    // The value of the file (e.g., file ID, URL, or path)
	Reader io.Reader // The reader for the file content
}

// IsInputFile is a marker method to indicate that the struct is an InputFile.
func (InputFile) IsInputFile() {}

// IsUploadable checks if the InputFile is uploadable.
// An InputFile is considered uploadable if it has a non-nil Reader.
func (ifu *InputFile) IsUploadable() bool {
	return ifu.Reader != nil
}

// MarshalJSON converts the InputFile to JSON.
// If the InputFile has a non-nil Reader, it returns a JSON string with an attachment reference.
// Otherwise, it returns a JSON string with the file value.
func (ifu *InputFile) MarshalJSON() ([]byte, error) {
	if ifu.Reader != nil {
		return []byte(`"attach://` + ifu.Value + `"`), nil
	}
	return []byte(`"` + ifu.Value + `"`), nil
}
