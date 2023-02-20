package tgo

import "fmt"

type ChatID string

func NewChatID(id any) ChatID {
	if val, ok := id.(string); ok {
		return ChatID(val)
	}

	return ChatID(fmt.Sprint(id))
}

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
