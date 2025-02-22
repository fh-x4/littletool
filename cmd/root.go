package cmd

import (
	"fmt"

	"github.com/fh-x4/littletool/component/logger"
	"github.com/fh-x4/littletool/config"
	"github.com/fh-x4/littletool/server"
	"github.com/spf13/cobra"
)

var (
	Version string
	Date    string
)

var configFile string

var root = cobra.Command{
	Use:     "little_tool",
	Short:   fmt.Sprintf("arbiter's little tool, build on %s", Date),
	Version: Version,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := config.Init(configFile); err != nil {
			panic(err)
		}
		if err := logger.Init(config.Get().Log); err != nil {
			panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer()
	},
}

func init() {
	root.PersistentFlags().StringVarP(&configFile, "config", "c", "/etc/config/config.json", "config file")
	// root.AddCommand()
}

func Execute() {
	root.Execute()
}
