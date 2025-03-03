package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/fh-x4/littletool/component/httpserver"
	"github.com/fh-x4/littletool/component/logger"
	"github.com/fh-x4/littletool/component/runner"
	"github.com/fh-x4/littletool/config"
	"github.com/fh-x4/littletool/server/handler/hbr"
	"github.com/fh-x4/littletool/server/handler/mytool/aes_ecb"
	"github.com/fh-x4/littletool/server/handler/mytool/hash"
	"github.com/fh-x4/littletool/server/handler/mytool/hello"
	"github.com/fh-x4/littletool/server/handler/timer"
	"github.com/gin-gonic/gin"
)

type server struct{}

func GerRunner() runner.Task {
	return &server{}
}

func (s *server) GetName() string {
	return "web_server"
}

func (s *server) Run(ctx context.Context) {
	conf := config.Get()
	engine := httpserver.NewServer()

	route(engine)

	srv := http.Server{
		Handler: engine,
		Addr:    conf.HttpServe,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.GetLogger().Errorf("http server start failed:%v", err)
		}
		logger.GetLogger().Info("http server closed")
	}()

	<-ctx.Done()
	if err := srv.Shutdown(ctx); err != nil {
		logger.GetLogger().Errorf("http shutdown failed:%v", err)
	}
}

func route(e *gin.Engine) {
	e.GET("/hello", httpserver.CreateHandler(&hello.HandlerGen{}))
	e.POST("/mytool/hash", httpserver.CreateHandler(&hash.HandlerGen{}))
	e.POST("/mytool/aes_encrypt", httpserver.CreateHandler(&aes_ecb.AesEncryptGen{}))
	e.POST("/mytool/aes_decrypt", httpserver.CreateHandler(&aes_ecb.AesDecryptGen{}))

	e.POST("/hbr/damage_caculate", httpserver.CreateHandler(&hbr.HandlerGen{}))

	e.POST("/timer/set_timer", httpserver.CreateHandler(&timer.SetTimerGen{}))
}
