package timer

type ErrorCode int

const (
	InternalServerError    ErrorCode = 20001
	TypeDataNotAvailable   ErrorCode = 20002
	MethodTypeNotAvailable ErrorCode = 20003
)
