package tgo

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

//go:generate go run ./cmd

const TelegramHost = "https://api.telegram.org"

type httpResponse[T any] struct {
	OK     bool `json:"ok"`
	Result T    `json:"result,omitempty"`
	*Error
}

// API is a telegram bot API client instance.
type API struct {
	host   string
	token  string
	client *http.Client
}

// NewAPI creates a new instance of the Telegram API client.
// It takes a token, an optional host (default is "https://api.telegram.org"),
// and an optional http.Client (default is a new http.Client instance).
// It returns a pointer to the API struct.
func NewAPI(token, host string, client *http.Client) *API {
	if host == "" {
		host = TelegramHost
	}

	if client == nil {
		client = &http.Client{}
	}

	return &API{
		host:   host,
		token:  token,
		client: client,
	}
}

// Download downloads a file from the Telegram server.
// It takes the filePath obtained from GetFile and returns an http.Response
// and an error. Please note that the filePath is not the same as the fileID.
func (api *API) Download(filePath string) (*http.Response, error) {
	return api.client.Get(api.host + "/file/bot" + api.token + "/" + filePath)
}

func callJson[T any](a *API, method string, rawData any) (T, error) {
	var response httpResponse[T]

	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(rawData); err != nil {
		return response.Result, err
	}

	resp, err := a.client.Post(a.host+"/bot"+a.token+"/"+method, "application/json", body)
	if err != nil {
		return response.Result, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response.Result, err
	} else if !response.OK {
		return response.Result, response.Error
	}

	return response.Result, nil
}

func callMultipart[T any](a *API, method string, params map[string]string, files map[string]*InputFile) (T, error) {
	r, w := io.Pipe()
	defer r.Close()

	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		for key, val := range params {
			m.WriteField(key, val)
		}

		for key, file := range files {
			ww, err := m.CreateFormFile(key, file.Value)
			if err != nil {
				w.CloseWithError(err)
				return
			} else if _, err = io.Copy(ww, file.Reader); err != nil {
				w.CloseWithError(err)
				return
			}
		}
	}()

	var response httpResponse[T]

	resp, err := a.client.Post(a.host+"/bot"+a.token+"/"+method, m.FormDataContentType(), r)
	if err != nil {
		return response.Result, nil
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response.Result, err
	} else if !response.OK {
		return response.Result, response.Error
	}

	return response.Result, nil
}
