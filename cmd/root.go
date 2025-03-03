package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/fh-x4/littletool/component/logger"
	"github.com/fh-x4/littletool/component/runner"
	ctimer "github.com/fh-x4/littletool/component/timer"
	"github.com/fh-x4/littletool/config"
	"github.com/fh-x4/littletool/server"
	"github.com/fh-x4/littletool/worker/timer"
)

var (
	Version string
	Commit  string
	Date    string
)

var configFile string

var root = cobra.Command{
	Use:     "little_tool",
	Short:   fmt.Sprintf("arbiter's little tool, build on %s, commit: %s", Date, Commit),
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
		ctx := context.Background()
		runner.RegisterTask(server.GerRunner())
		runner.RegisterTask(timer.NewTimer(1, ctimer.NewProducer()))
		runner.Run(ctx)
	},
}

func init() {
	root.PersistentFlags().StringVarP(&configFile, "config", "c", "/etc/config/config.json", "config file")
	// root.AddCommand()
}

func Execute() {
	root.Execute()
}
