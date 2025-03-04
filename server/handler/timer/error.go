package timer

type ierror struct {
	error
	Code    ErrorCode
	Message string
}

func (i *ierror) GetCode() int {
	return int(i.Code)
}
func (i *ierror) GetMessage() string {
	return i.Message
}
