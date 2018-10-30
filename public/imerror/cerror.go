package imerror

var (
	CCodeSuccess = 0 // 成功发送
)

// CError 接入层调用错误
type CError struct {
	Code    int
	Message string
}

func (e *CError) Error() string {
	return e.Message
}

func NewCError(code int, message string) *CError {
	return &CError{
		Code:    code,
		Message: message,
	}
}

var (
	CErrNotFriend  = NewCError(1, "not friend")
	CErrNotInGroup = NewCError(1, "not in group")
)
