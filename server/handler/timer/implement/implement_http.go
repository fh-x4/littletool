package implement

import (
	"encoding/json"
	"errors"
	"time"

	resty "gopkg.in/resty.v1"

	"github.com/fh-x4/littletool/component/timer"
)

func GenHttpCallback(data string) (timer.IAction, error) {
	h := &httpTimer{}
	if err := json.Unmarshal([]byte(data), h); err != nil {
		return nil, errors.New("data is not available")
	}
	h.GenerateBaseInfo(MethodHttp)
	return h, nil
}

type httpTimer struct {
	base

	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
	Params  map[string]string `json:"params"`
}

func (h *httpTimer) Key() string {
	return h.key
}
func (h *httpTimer) Type() string {
	return h.types
}
func (h *httpTimer) Call() error {
	c := resty.DefaultClient.SetTimeout(5 * time.Second)
	_, err := c.NewRequest().SetHeaders(h.Headers).SetBody(h.Body).SetQueryParams(h.Params).Execute(h.Method, h.URL)
	return err
}
