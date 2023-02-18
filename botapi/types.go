package botapi

import (
	"encoding/json"
	"fmt"
	"reflect"
)

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

type InlineKeyboard interface{}

type Params map[string]string

func NewParams() Params { return Params{} }

func (p Params) Add(key string, value any) { p[key] = fmt.Sprint(value) }

func (p Params) AddOptional(key string, value any) {
	if !reflect.ValueOf(value).IsZero() {
		p[key] = fmt.Sprint(value)
	}
}

func (p Params) AddOptionalJSON(key string, value any) {
	if value != nil {
		raw, _ := json.Marshal(value)
		p[key] = string(raw)
	}
}
