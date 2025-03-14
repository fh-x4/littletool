package httpserver

import (
	"bytes"
	"io"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/fh-x4/littletool/component/logger"
	"github.com/fh-x4/littletool/config"
	"github.com/gin-gonic/gin"
)

func httpRecover() gin.HandlerFunc {
	log := logger.GetLogger()
	return func(c *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				log.Errorf("panic recover:%v, stack=%s", p, string(debug.Stack()))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func loggingRequest() gin.HandlerFunc {
	log := logger.GetLogger()
	maxLen := config.Get().MaxBodyLen
	return func(c *gin.Context) {
		request := c.Request
		body, err := io.ReadAll(request.Body)
		request.Body = io.NopCloser(bytes.NewReader(body))
		if err != nil {
			log.Warningf("read body failed:%v", err)
			c.Next()
			return
		}
		if len(body) > maxLen {
			body = body[:maxLen]
		}
		log.WithFields(map[string]interface{}{
			"method": request.Method,
			"url":    request.URL.RequestURI(),
			"ip":     request.RemoteAddr,
			"body":   string(body),
		}).Infof("incomming http request")
		c.Next()
	}
}

func loggingRespond() gin.HandlerFunc {
	log := logger.GetLogger()
	// maxLen := config.Get().MaxBodyLen
	return func(c *gin.Context) {
		current := time.Now()

		c.Next()

		status := c.Writer.Status()
		cost := time.Since(current)
		// c.Copy().Writer.
		log.WithFields(map[string]interface{}{
			"status_code": status,
			"status_text": http.StatusText(status),
			"cost":        cost,
			"body":        "",
		}).Infof("outgoing http respond")
	}
}
