package error

type Error struct {
	Code    int
	Message string
	Data    interface{}
}

func NewError(code int, message string, data interface{}) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
