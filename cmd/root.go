package cmd

import (
	"github.com/fh-x4/littletool/server"
	"github.com/spf13/cobra"
)

var (
	Version string
	Date    string
)

var root = cobra.Command{
	Use: "little_tool",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer()
	},
}

func Execute() {
	root.Execute()
}
