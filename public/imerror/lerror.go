package imerror

// BError 业务错误，属于正常预期的错误
type LError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *LError) Error() string {
	return e.Message
}

func NewLError(code int, message string) *LError {
	return &LError{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func WrapLErrorWithData(err *LError, data interface{}) *LError {
	return &LError{
		Code:    err.Code,
		Message: err.Message,
		Data:    data,
	}
}

// 通用错误
var (
	ErrUnauthorized = NewLError(1, "unauthorized") // 需要认证
	ErrBadRequest   = NewLError(2, "bad request")  // 请求错误
	ErrUnknow       = NewLError(3, "unkown error") // 未知错误
)

// 业务错误
var (
	ErrDeviceIdOrToken = NewLError(1001, "error device token")       // 设备id或者token错误
	ErrNumberUsed      = NewLError(1002, "number has be used")       // 手机号码已经被使用
	ErrNameOrPassword  = NewLError(1003, "error number or password") // 用户名或者密码错误
)
