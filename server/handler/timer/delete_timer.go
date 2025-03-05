package timer

import (
	"context"

	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/component/timer"
)

type deleteTimerHandler struct {
	req *deleteTimerReq
	rsp *deleteTimerRsp
}
type deleteTimerReq struct {
	Key string `json:"key"`
}
type deleteTimerRsp struct {
}

func (h *deleteTimerHandler) GetRequest() interface{} {
	return h.req
}
func (h *deleteTimerHandler) GetRespond() interface{} {
	return h.rsp
}
func (h *deleteTimerHandler) Call(ctx context.Context) httpserver.IError {
	timer.DeleteTimer(h.req.Key)
	return nil
}

type DeleteTimerGen struct{}

func (hg *DeleteTimerGen) GenHandler() httpserver.IHandler {
	return &deleteTimerHandler{
		req: &deleteTimerReq{},
		rsp: &deleteTimerRsp{},
	}
}
