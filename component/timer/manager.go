package timer

import (
	"sync"
	"time"
)

type timerEntity struct {
	entity     IAction
	controller chan int
}

type timerManager struct {
	sync.Mutex
	cm map[string]chan int
}

var tm timerManager

func init() {
	tm = timerManager{
		cm: make(map[string]chan int),
	}
}

func SetTimer(key string, trigerTime time.Duration, ia IAction) {
	t := time.NewTimer(trigerTime)
	ct := make(chan int)
	te := timerEntity{
		entity:     ia,
		controller: ct,
	}

	tm.Lock()
	defer tm.Unlock()
	tm.cm[ia.Key()] = ct
	go func() {
		select {
		case <-t.C:
			ch <- te.entity
		case <-te.controller:
			return
		}
	}()
}

func DeleteTimer(key string) {
	tm.Lock()
	defer tm.Unlock()

	tm.cm[key] <- 1
	delete(tm.cm, key)
}
