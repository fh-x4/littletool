package implement

import "github.com/fh-x4/littletool/component/timer"

var generateFunc map[string]func(string) (timer.IAction, error)

const (
	MethodDemo = "demo"
	MethodHttp = "http"
)

func init() {
	generateFunc = make(map[string]func(string) (timer.IAction, error))
	generateFunc[MethodHttp] = GenHttpCallback
	generateFunc[MethodDemo] = GenDemoCallback
}

func GetGenerateFunc(types string) (func(string) (timer.IAction, error), bool) {
	fn, ok := generateFunc[types]
	return fn, ok
}
