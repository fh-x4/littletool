package aes_ecb

import (
	"context"

	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/component/logger"
)

type aesEcbEncryptHandler struct {
	req *aesEcbEncryptReq
	rsp *aesEcbEncryptRsp
}
type aesEcbEncryptReq struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}
type aesEcbEncryptRsp struct {
	Cipher string `json:"cipher"`
}

func (h *aesEcbEncryptHandler) GetRequest() interface{} {
	return h.req
}
func (h *aesEcbEncryptHandler) GetRespond() interface{} {
	return h.rsp
}

func (h *aesEcbEncryptHandler) Call(ctx context.Context) httpserver.IError {
	encrypt, err := aesEncrypt(h.req.Data, h.req.Key)
	if err != nil {
		logger.GetLogger().Errorf("decrypted failed:%v", err)
		return &ierror{
			error:   err,
			Code:    InternalServerError,
			Message: "未知错误",
		}
	}
	h.rsp.Cipher = string(encrypt)
	return nil
}

type AesEncryptGen struct{}

func (hg *AesEncryptGen) GenHandler() httpserver.IHandler {
	return &aesEcbEncryptHandler{
		req: &aesEcbEncryptReq{},
		rsp: &aesEcbEncryptRsp{},
	}
}
