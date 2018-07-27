package controller

const OK = 200

// 请求返回值
const (
	CodeSuccess      = 0 // 成功返回
	CodeUnauthorized = 1 // 需要认证
	CodeBadRequest   = 2 // 客户端请求错误
	CodeError        = 3 // 请求错误
)

// Response 用户响应数据
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	// 用户没有登录
	Unauthorized = &Response{Code: CodeUnauthorized, Message: "unauthorized"}
)

func NewSuccess(data interface{}) *Response {
	return &Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	}
}

func NewError(err error) *Response {
	return &Response{
		Code:    CodeError,
		Message: err.Error(),
	}
}

func NewBadRequst(err error) *Response {
	response := &Response{
		Code: CodeBadRequest,
	}
	if err != nil {
		response.Message = err.Error()
	}
	return response
}
