package httpserver

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	engine := gin.New()
	engine.Use(
		Recover(),
		LoggingRequest(),
		LoggingRespond(),
	)
	return engine
}
