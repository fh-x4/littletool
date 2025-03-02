package timer

import (
	"context"
	"time"

	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/component/timer"
	"github.com/fh-x4/littletool/server/handler/timer/implement/http"
)

type setTimerHandler struct {
	req *setTimerReq
	rsp *setTimerRsp
}
type setTimerReq struct {
	CallbackType string `json:"callback_type"`
	CallbackData []byte `json:"callback_data"`
	Delay        int    `json:"delay"`
}
type setTimerRsp struct {
}

func (h *setTimerHandler) GetRequest() interface{} {
	return h.req
}
func (h *setTimerHandler) GetRespond() interface{} {
	return h.rsp
}
func (h *setTimerHandler) Call(ctx context.Context) httpserver.IError {
	var t timer.IAction
	var err error
	switch h.req.CallbackType {
	case "http":
		t, err = http.GenHttpCallback(h.req.CallbackData)
	}

	if err != nil {
		return &ierror{
			error: err,
			Code:  TypeDataNotAvailable,
		}
	}
	timer.SetTimer(t.Key(), time.Duration(h.req.Delay)*time.Second, t)
	return nil
}

type SetTimerGen struct{}

func (hg *SetTimerGen) GenHandler() httpserver.IHandler {
	return &setTimerHandler{
		req: &setTimerReq{},
		rsp: &setTimerRsp{},
	}
}
