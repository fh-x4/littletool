package httpserver

import (
	"context"
	"net/http"

	"github.com/fh-x4/littletool/component/logger"
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Call(ctx context.Context) IError
	GetRequest() interface{}
	GetRespond() interface{}
}

type iHandlerGen interface {
	GenHandler() IHandler
}

func CreateHandler(hg iHandlerGen) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := hg.GenHandler()
		ctx := c.Request.Context()
		req := h.GetRequest()
		if err := c.ShouldBind(req); err != nil {
			logger.GetLogger().Errorf("bind request failed:%v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1,
				"message": "参数异常",
			})
			return
		}
		if ie := h.Call(ctx); ie != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    ie.GetCode(),
				"message": ie.GetMessage(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "",
			"data":    h.GetRespond(),
		})
	}
}
