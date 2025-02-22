package hash

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/fh-x4/littletool/component/httpserver"
)

type HashHandler struct {
	req *hashReq
	rsp *hashRsp
}
type hashReq struct {
	Method HashMethod `json:"method"`
	Source string     `json:"source"`
}
type hashRsp struct {
	Encrypt string `json:"encrypt"`
}
type ierror struct {
	error
	Code    int
	Message string
}

func (i *ierror) GetCode() int {
	return i.Code
}
func (i *ierror) GetMessage() string {
	return i.Message
}
func (h *HashHandler) GetRequest() interface{} {
	return h.req
}
func (h *HashHandler) GetRespond() interface{} {
	return h.rsp
}

func (h *HashHandler) Call(ctx context.Context) httpserver.IError {
	switch h.req.Method {
	case MethodMd5:
		sum := md5.New()
		_, _ = sum.Write([]byte(h.req.Source))
		h.rsp.Encrypt = fmt.Sprintf("%X", sum.Sum(nil))
	case MethodBase64:
		h.rsp.Encrypt = base64.StdEncoding.EncodeToString([]byte(h.req.Source))
	case MethodSHA256:
		ha := sha256.New()
		h.rsp.Encrypt = fmt.Sprintf("%X", ha.Sum([]byte(h.req.Source)))
	}

	return nil
}

type HandlerGen struct{}

func (hg *HandlerGen) GenHandler() httpserver.IHandler {
	return &HashHandler{
		req: &hashReq{},
		rsp: &hashRsp{},
	}
}
