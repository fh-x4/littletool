package cmd

import (
	"arbiter_littletool/component/httpserver"
	"arbiter_littletool/server/handler/hello"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var root = cobra.Command{
	Use: "小工具",
	Run: func(cmd *cobra.Command, args []string) {
		engine := gin.Default()
		engine.GET("/hello", httpserver.CreateHandler(&hello.HandlerGen{}))
		engine.Run(":8080")
	},
}

func Run() {
	root.Execute()
}
