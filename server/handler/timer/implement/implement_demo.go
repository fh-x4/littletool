package implement

import (
	"github.com/fh-x4/littletool/component/timer"
)

func GenDemoCallback(data string) (timer.IAction, error) {
	d := &demoTimer{
		data: data,
	}
	d.GenerateBaseInfo(MethodDemo)
	return d, nil
}

type demoTimer struct {
	base
	data string
}

func (h *demoTimer) Key() string {
	return h.key
}
func (h *demoTimer) Type() string {
	return h.types
}
func (h *demoTimer) Call() error {
	// do nothing
	return nil
}
