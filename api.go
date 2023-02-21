package tgo

import (
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

	return &API{url: CreateApiURL(host, token), client: client}
}

func (api *API) GetHTTPTimeout() int64 {
	return int64(api.client.Timeout.Seconds())
}

func CreateApiURL(host, token string) string {
	return host + "/bot" + token + "/"
}
