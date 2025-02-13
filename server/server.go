package server

import (
	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/server/handler/hello"
)

func RunServer() {
	engine := httpserver.NewServer()
	engine.GET("/hello", httpserver.CreateHandler(&hello.HandlerGen{}))
	engine.Run(":8080")
}
