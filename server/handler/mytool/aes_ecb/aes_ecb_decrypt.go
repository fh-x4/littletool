package aes_ecb

import (
	"context"

	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/component/logger"
)

type aesEcbDecryptHandler struct {
	req *aesEcbDecryptReq
	rsp *aesEcbDecryptRsp
}
type aesEcbDecryptReq struct {
	Key    string `json:"key"`
	Cipher string `json:"cipher"`
}
type aesEcbDecryptRsp struct {
	Data interface{} `json:"data"`
}

type ierror struct {
	error
	Code    ErrorCode
	Message string
}

func (i *ierror) GetCode() int {
	return int(i.Code)
}
func (i *ierror) GetMessage() string {
	return i.Message
}
func (h *aesEcbDecryptHandler) GetRequest() interface{} {
	return h.req
}
func (h *aesEcbDecryptHandler) GetRespond() interface{} {
	return h.rsp
}

func (h *aesEcbDecryptHandler) Call(ctx context.Context) httpserver.IError {
	decrypted, err := aesDecrypt([]byte(h.req.Cipher), []byte(h.req.Key))
	if err != nil {
		logger.GetLogger().Errorf("decrypted failed:%v", err)
		return &ierror{
			error:   err,
			Code:    InternalServerError,
			Message: "未知错误",
		}
	}
	h.rsp.Data = string(decrypted)
	return nil
}

type AesDecryptGen struct{}

func (hg *AesDecryptGen) GenHandler() httpserver.IHandler {
	return &aesEcbDecryptHandler{
		req: &aesEcbDecryptReq{},
		rsp: &aesEcbDecryptRsp{},
	}
}
