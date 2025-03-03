package timer

import (
	"context"
	"time"

	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/component/timer"
	"github.com/fh-x4/littletool/server/handler/timer/implement"
)

type setTimerHandler struct {
	req *setTimerReq
	rsp *setTimerRsp
}
type setTimerReq struct {
	CallbackType string `json:"callback_type"`
	CallbackData string `json:"callback_data"`
	Delay        int    `json:"delay"`
}
type setTimerRsp struct {
	Key string `json:"key"`
}

func (h *setTimerHandler) GetRequest() interface{} {
	return h.req
}
func (h *setTimerHandler) GetRespond() interface{} {
	return h.rsp
}
func (h *setTimerHandler) Call(ctx context.Context) httpserver.IError {
	fn, support := implement.GetGenerateFunc(h.req.CallbackType)
	if !support {
		return &ierror{
			Code:    MethodTypeNotAvailable,
			Message: "method not support",
		}
	}

	t, err := fn(h.req.CallbackData)
	if err != nil {
		return &ierror{
			error: err,
			Code:  TypeDataNotAvailable,
		}
	}
	timer.SetTimer(t.Key(), time.Duration(h.req.Delay)*time.Second, t)
	h.rsp.Key = t.Key()
	return nil
}

type SetTimerGen struct{}

func (hg *SetTimerGen) GenHandler() httpserver.IHandler {
	return &setTimerHandler{
		req: &setTimerReq{},
		rsp: &setTimerRsp{},
	}
}
