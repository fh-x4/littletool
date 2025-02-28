package runner

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/fh-x4/littletool/component/logger"
)

type Task interface {
	Run(ctx context.Context)
	GetName() string
}

var runner []Task

func RegisterTask(t Task) {
	runner = append(runner, t)
}

func Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		signal.Stop(q)
		close(q)
	}()
	go func() {
		<-q
		cancel()
	}()

	wg := &sync.WaitGroup{}
	for _, r := range runner {
		task := r
		go func() {
			defer func() {
				logger.GetLogger().Infof("task %s stop", task.GetName())
				wg.Done()
			}()
			logger.GetLogger().Infof("task %s start", task.GetName())
			task.Run(ctx)
		}()
		wg.Add(1)
	}
	wg.Wait()
}
