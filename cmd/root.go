package cmd

import (
	"github.com/fh-x4/littletool/server"
	"github.com/spf13/cobra"
)

var root = cobra.Command{
	Use: "小工具",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer()
	},
}

func Execute() {
	root.Execute()
}
