package tgo

import (
	"fmt"
	"io"
)

// ChatID is just an any type.
// It should be a username (string) or a user-id (integer).
type ChatID any

type ParseMode string

const (
	ParseModeNone       ParseMode = ""
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
)

type ReplyMarkup interface {
	// IsReplyMarkup does nothing and is only used to enforce type-safety
	IsReplyMarkup()
}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ReplyKeyboardMarkup) IsReplyMarkup() {}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ReplyKeyboardRemove) IsReplyMarkup() {}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (InlineKeyboardMarkup) IsReplyMarkup() {}

// IsReplyMarkup is an empty function used to implement ReplyMarkup
func (ForceReply) IsReplyMarkup() {}

// SenderID returns the chat id of who sent the message
func (m *Message) SenderID() int64 {
	if m.From == nil {
		return 0
	}

	return m.From.Id
}

// ChatID returns the chat id of where the message is sent in
func (m *Message) ChatID() int64 {
	return m.Chat.Id
}

// ThreadID returns the chat's thread id of where the message is sent in
func (m *Message) ThreadID() int64 {
	return m.MessageThreadId
}

// MessageID returns ID of the sent message
func (m *Message) MessageID() int64 {
	return m.MessageId
}

// String returns the message's text or media caption
func (m *Message) String() string {
	if m.Text != "" {
		return m.Text
	}

	return m.Caption
}

// SenderID returns the chat id of who sent the query
func (q *CallbackQuery) SenderID() int64 {
	return q.From.Id
}

// ChatID returns the chat id of where the message of that callback button was in.
// it returns the sender's chat-id if the message couldn't be found.
func (q *CallbackQuery) ChatID() int64 {
	if q.Message == nil {
		return q.SenderID()
	}

	return q.Message.Chat.Id
}

// ThreadID returns the chat's thread id of where the message of that callback button was in.
// it returns zero if the message couldn't be found.
func (q *CallbackQuery) ThreadID() int64 {
	if q.Message == nil {
		return 0
	}

	return q.Message.MessageThreadId
}

// MessageID returns message ID of pressed callback button.
// it returns zero if the message couldn't be found.
func (q *CallbackQuery) MessageID() int64 {
	if q.Message == nil {
		return 0
	}

	return q.Message.MessageId
}

type InputFileNotUploadable string

func FileFromID(fileID string) InputFile {
	return InputFileNotUploadable(fileID)
}

func FileFromURL(url string) InputFile {
	return InputFileNotUploadable(url)
}

func (InputFileNotUploadable) IsInputFile() {}

type InputFileUploadable struct {
	Name   string
	Reader io.Reader
}

// MarshalJSON is a custom marshaller which be called with json.Marshal
func (iFile *InputFileUploadable) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"attach://%s"`, iFile.Name)), nil
}

func FileFromReader(name string, reader io.Reader) InputFile {
	return &InputFileUploadable{Name: name, Reader: reader}
}

func (InputFileUploadable) IsInputFile() {}
