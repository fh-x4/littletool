package hello

import (
	"arbiter_littletool/component/httpserver"
	"context"
)

type HelloHandler struct {
	req *helloReq
	rsp *helloRsp
}
type helloReq struct {
}
type helloRsp struct {
	Hello string `json:"hello"`
}
type IE struct {
	error
	Code    int
	Message string
}

func (i *IE) GetCode() int {
	return i.Code
}
func (i *IE) GetMessage() string {
	return i.Message
}
func (h *HelloHandler) GetRequest() interface{} {
	return h.req
}
func (h *HelloHandler) GetRespond() interface{} {
	return h.rsp
}

func (h *HelloHandler) Call(ctx context.Context) httpserver.IError {
	h.rsp.Hello = "hello world"
	return nil
}

type HandlerGen struct{}

func (hg *HandlerGen) GenHandler() httpserver.IHandler {
	return &HelloHandler{
		req: &helloReq{},
		rsp: &helloRsp{},
	}
}
