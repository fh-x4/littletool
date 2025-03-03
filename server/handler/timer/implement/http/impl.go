package http

import (
	"encoding/json"
	"errors"
	"time"

	resty "gopkg.in/resty.v1"
)

func init() {
	resty.SetTimeout(5 * time.Second)
}

func GenHttpCallback(data string) (*httpTimer, error) {
	h := &httpTimer{}
	if err := json.Unmarshal([]byte(data), h); err != nil {
		return nil, errors.New("data is not available")
	}
	return h, nil
}

type httpTimer struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
	Params  map[string]string `json:"params"`
}

func (h *httpTimer) Key() string {
	return ""
}
func (h *httpTimer) Call() error {
	_, err := resty.NewRequest().SetHeaders(h.Headers).SetBody(h.Body).SetQueryParams(h.Params).Execute(h.Method, h.URL)
	return err
}
