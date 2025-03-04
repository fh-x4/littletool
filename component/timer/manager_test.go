package timer

import (
	"os"
	"runtime/pprof"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/fh-x4/littletool/component/logger"
)

type mockAction struct {
	key  string
	typ  string
	call bool
}

func (m *mockAction) Key() string {
	return m.key
}
func (m *mockAction) Type() string {
	return m.typ
}
func (m *mockAction) Call() error {
	m.call = true
	return nil
}

func TestSetTimer(t *testing.T) {
	NewProducer()
	initLogger()
	go func() {
		for {
			select {
			case entity := <-ch:
				entity.Call()
			}
		}
	}()
	action := &mockAction{key: "testKey", typ: "testType"}
	SetTimer(action.Key(), 1*time.Second, action)

	time.Sleep(2 * time.Second)

	assert.True(t, action.call, "Expected action to be called")
}

func TestDeleteTimer(t *testing.T) {
	NewProducer()
	initLogger()
	go func() {
		for {
			select {
			case entity := <-ch:
				entity.Call()
			}
		}
	}()
	action := &mockAction{key: "testKey", typ: "testType"}
	SetTimer(action.Key(), 1*time.Second, action)
	DeleteTimer(action.Key())

	time.Sleep(2 * time.Second)

	assert.False(t, action.call, "Expected action not to be called")
}

func BenchmarkSetTimer(b *testing.B) {
	NewProducer()
	initLogger()
	go func() {
		for {
			select {
			case entity := <-ch:
				entity.Call()
			}
		}
	}()

	cpuProfile, err := os.Create("cpu_profile.prof")
	if err != nil {
		b.Fatal(err)
	}
	defer cpuProfile.Close()
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	memProfile, err := os.Create("mem_profile.prof")
	if err != nil {
		b.Fatal(err)
	}
	defer memProfile.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			action := &mockAction{key: "testKey", typ: "testType"}
			SetTimer(action.Key(), 1*time.Second, action)
		}
	})

	pprof.WriteHeapProfile(memProfile)
}

func initLogger() {
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Out = os.Stdout
	logger.SetLogger(logrus.NewEntry(log))
}
