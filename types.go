package tgo

import (
	"github.com/haashemi/tgo/botapi"
)

func NewChatID(id any) botapi.ChatID { return botapi.NewChatID(id) }

type ParseMode = botapi.ParseMode

const (
	ParseModeNone       ParseMode = botapi.ParseModeNone
	ParseModeMarkdown   ParseMode = botapi.ParseModeMarkdown
	ParseModeMarkdownV2 ParseMode = botapi.ParseModeMarkdownV2
	ParseModeHTML       ParseMode = botapi.ParseModeHTML
)
