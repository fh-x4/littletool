package hbr

import (
	"context"

	"github.com/fh-x4/littletool/component/httpserver"
)

type damageCaculateHandler struct {
	req *damageCaculateReq
	rsp *damageCaculateRsp
}
type damageCaculateReq struct{}
type damageCaculateRsp struct{}

func (h *damageCaculateHandler) GetRequest() interface{} {
	return h.req
}
func (h *damageCaculateHandler) GetRespond() interface{} {
	return h.rsp
}
func (h *damageCaculateHandler) Call(ctx context.Context) httpserver.IError {
	return nil
}

type HandlerGen struct{}

func (hg *HandlerGen) GenHandler() httpserver.IHandler {
	return &damageCaculateHandler{
		req: &damageCaculateReq{},
		rsp: &damageCaculateRsp{},
	}
}
