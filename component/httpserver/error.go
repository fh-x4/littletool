package httpserver

type IError interface {
	error
	GetCode() int
	GetMessage() string
}
