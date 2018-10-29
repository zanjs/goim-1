package imerror

// BError 业务错误，属于正常预期的错误
type BError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *BError) Error() string {
	return e.Message
}

func NewBError(code int, message string) *BError {
	return &BError{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func WrapBErrorWithData(err *BError, data interface{}) *BError {
	return &BError{
		Code:    err.Code,
		Message: err.Message,
		Data:    data,
	}
}

var ErrUnauthorized = NewBError(1, "unauthorized") // 需要认证
var ErrBadRequest = NewBError(1, "bad request")    // 请求错误
var ErrUnknowError = NewBError(1, "unkown error")  // 未知错误
