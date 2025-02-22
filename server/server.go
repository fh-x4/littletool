package server

import (
	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/config"
	"github.com/fh-x4/littletool/server/handler/mytool/aes_ecb"
	"github.com/fh-x4/littletool/server/handler/mytool/hash"
	"github.com/fh-x4/littletool/server/handler/mytool/hello"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	conf := config.Get()
	engine := httpserver.NewServer()

	route(engine)

	engine.Run(conf.HttpServe)
}

func route(e *gin.Engine) {
	e.GET("/hello", httpserver.CreateHandler(&hello.HandlerGen{}))
	e.POST("/mytool/hash", httpserver.CreateHandler(&hash.HandlerGen{}))
	e.POST("/mytool/aes_encrypt", httpserver.CreateHandler(&aes_ecb.AesEncryptGen{}))
	e.POST("/mytool/aes_decrypt", httpserver.CreateHandler(&aes_ecb.AesDecryptGen{}))

	e.POST("/hbr/damage_caculate", nil)
}
