package discogs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DiscogsError interface {
	error
	Status() int
	MessageText() string
	Details() interface{}
}

type APIError struct {
	StatusCode int         `json:"-"`
	Message    string      `json:"message"`
	ErrorText  string      `json:"error"`
	Detail     interface{} `json:"detail,omitempty"`
	Raw        string      `json:"-"`
}

func (e *APIError) Error() string {
	msg := e.Message
	if msg == "" {
		msg = e.ErrorText
	}

	if e.Detail != nil {
		detailJSON, _ := json.MarshalIndent(e.Detail, "", "  ")
		return fmt.Sprintf(
			"discogs error: status=%d message=%s\n%s",
			e.StatusCode,
			msg,
			string(detailJSON),
		)
	}

	if msg != "" {
		return fmt.Sprintf(
			"discogs error: status=%d message=%s",
			e.StatusCode,
			msg,
		)
	}

	return fmt.Sprintf(
		"discogs error: status=%d body=%s",
		e.StatusCode,
		e.Raw,
	)
}

// 👇 Interface methods

func (e *APIError) Status() int {
	return e.StatusCode
}

func (e *APIError) MessageText() string {
	if e.Message != "" {
		return e.Message
	}
	return e.ErrorText
}

func (e *APIError) Details() interface{} {
	return e.Detail
}

// 👇 parser

func parseError(resp *http.Response) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("discogs error: status=%d (failed to read body)", resp.StatusCode)
	}

	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Raw:        string(bodyBytes),
	}

	if err := json.Unmarshal(bodyBytes, apiErr); err != nil {
		return apiErr
	}

	return apiErr
}
