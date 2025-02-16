package server

import (
	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/config"
	"github.com/fh-x4/littletool/server/handler/hello"
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
}
