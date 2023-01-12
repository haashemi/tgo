package tgo

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

const TelegramHost = "https://api.telegram.org"

type API struct {
	token string
	host  string
	http  *http.Client
}

// API Implements pure bot-api methods
func NewAPI(token string) *API {
	return &API{token: token, host: TelegramHost, http: &http.Client{}}
}

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

// ToDo: Better error handling
func unmarshal[T any](body io.ReadCloser) (x T, err error) {
	defer body.Close()

	data := &httpResponse[T]{}
	if err := json.NewDecoder(body).Decode(data); err != nil {
		return x, err
	} else if !data.OK {
		return x, data.Error
	}
	return data.Result, nil
}

type MultipartForm interface{ HasUploadable() bool }

func (c *API) doHTTP(method string, raw any) (*http.Response, error) {
	url := c.host + "/bot" + c.token + "/" + method

	if raw == nil {
		return c.http.Get(url)
	} else if data, ok := raw.(MultipartForm); ok && data.HasUploadable() {
		r, w := io.Pipe()
		defer r.Close()

		m := multipart.NewWriter(w)

		go func() {
			defer w.Close()
			defer m.Close()

			params, files := getParamsAndFiles(data)
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

		return c.http.Post(url, m.FormDataContentType(), r)
	}

	data := bytes.NewBuffer(nil)
	if err := json.NewEncoder(data).Encode(raw); err != nil {
		return nil, err
	}
	return c.http.Post(url, "application/json", data)
}

func getParamsAndFiles(d any) (params Params, files map[string]*InputFileWithUpload) {
	params = NewParams()
	files = make(map[string]*InputFileWithUpload)

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
				files[tag] = xx.(*InputFileWithUpload)
			} else {
				params.Add(tag, xx.(InputFileNoUpload))
			}
		} else if field.Type().Kind() == reflect.Struct {
			params.AddOptionalJSON(tag, data)
		} else {
			params.Add(tag, data)
		}
	}

	return params, files
}
