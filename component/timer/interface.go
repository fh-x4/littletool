package timer

type IAction interface {
	Key() string
	Call() error
}

var ch chan IAction

func NewProducer() chan IAction {
	c := make(chan IAction, 100)
	ch = c
	return ch
}
