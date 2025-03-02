package timer

import (
	"context"
	"sync"

	"github.com/fh-x4/littletool/component/timer"
)

type task struct {
	childWorker int
	c           <-chan timer.IAction
}

func NewTimer(workers int, c chan timer.IAction) *task {
	return &task{
		childWorker: workers,
		c:           c,
	}
}

func (t *task) GetName() string {
	return "task_timer"
}

func (t *task) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}
	for range t.childWorker {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case ac := <-t.c:
					ac.Call()
				}
			}
		}()
	}
	wg.Wait()
}
