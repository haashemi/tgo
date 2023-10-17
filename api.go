package tgo

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

//go:generate go run ./internal/gen

const TelegramHost = "https://api.telegram.org"

type httpResponse[T any] struct {
	OK     bool `json:"ok"`
	Result T    `json:"result,omitempty"`
	*Error
}

type API struct {
	url    string
	client *http.Client
}

func NewAPI(token, host string, client *http.Client) *API {
	if host == "" {
		host = TelegramHost
	}

	if client == nil {
		client = &http.Client{Timeout: 30 * time.Second}
	}

	return &API{
		url:    host + "/bot" + token + "/",
		client: client,
	}
}

func callJson[T any](a *API, method string, rawData any) (T, error) {
	var response httpResponse[T]

	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(rawData); err != nil {
		return response.Result, err
	}

	resp, err := a.client.Post(a.url+method, "application/json", body)
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

	resp, err := a.client.Post(a.url+method, m.FormDataContentType(), r)
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
