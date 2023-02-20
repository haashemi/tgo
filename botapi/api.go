package botapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	"time"
)

//go:generate go run ../internal/codegen

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

type multipartForm interface{ HasUploadable() bool }

type httpResponse[T any] struct {
	OK     bool `json:"ok"`
	Result T    `json:"result,omitempty"`
	*Error
}

type Error struct {
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

func (e Error) Error() string { return e.Description }

func doHTTP[T any](client *http.Client, baseURL, method string, rawData any) (data T, err error) {
	var url = baseURL + method

	var resp *http.Response

	if rawData == nil {
		if resp, err = client.Get(url); err != nil {
			return
		}
	} else if body, ok := rawData.(multipartForm); ok && body.HasUploadable() {
		r, w := io.Pipe()
		defer r.Close()

		m := multipart.NewWriter(w)

		go func() {
			defer w.Close()
			defer m.Close()

			params, files := getParamsAndFiles(body)
			for key, val := range params {
				m.WriteField(key, val)
			}

			for key, file := range files {
				ww, err := m.CreateFormFile(key, file.Name)
				if err != nil {
					w.CloseWithError(err)
					return
				} else if _, err = io.Copy(ww, file.Reader); err != nil {
					w.CloseWithError(err)
					return
				}
			}
		}()

		if resp, err = client.Post(url, m.FormDataContentType(), r); err != nil {
			return
		}
	} else {
		body := bytes.NewBuffer(nil)
		if err = json.NewEncoder(body).Encode(rawData); err != nil {
			return
		}

		if resp, err = client.Post(url, "application/json", body); err != nil {
			return
		}
	}

	defer resp.Body.Close()

	response := &httpResponse[T]{}
	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return
	} else if !response.OK {
		err = response.Error
		return
	}
	return response.Result, nil
}

func getParamsAndFiles(d any) (params map[string]string, files map[string]*InputFileUploadable) {
	params = make(map[string]string)
	files = make(map[string]*InputFileUploadable)

	v := reflect.ValueOf(d)
	vType := reflect.TypeOf(d)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.IsZero() {
			continue
		}

		data := field.Interface()
		tag := strings.TrimSuffix(vType.Field(i).Tag.Get("json"), ",omitempty")

		if xx, ok := data.(InputFile); ok {
			if xx.NeedsUpload() {
				files[tag] = xx.(*InputFileUploadable)
			} else {
				params[tag] = string(xx.(InputFileNotUploadable))
			}
		} else if field.Type().Kind() == reflect.Struct {
			raw, _ := json.Marshal(data)
			params[tag] = string(raw)
		} else {
			params[tag] = fmt.Sprint(data)
		}
	}

	return params, files
}
