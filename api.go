package tgo

import (
	"fmt"
	"net/http"
	"time"
)

//go:generate go run ./internal/codegen

const TelegramHost = "https://api.telegram.org"

type API struct {
	url    string
	client *http.Client
}

// API Implements pure bot-api methods
func NewAPI(token, host string, client *http.Client) *API {
	if host == "" {
		host = TelegramHost
	}
	if client == nil {
		client = &http.Client{Timeout: 30 * time.Second}
	}

	return &API{url: host + "/bot" + token + "/", client: client}
}

func (api *API) TimeoutSeconds() int64 {
	return int64(api.client.Timeout.Seconds())
}

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
