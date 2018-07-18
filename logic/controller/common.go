package controller

const OK = 200

// 请求返回值
const (
	success             = 0 // 成功返回
	unauthorized        = 1 // 需要认证
	badRequest          = 2 // 客户端请求错误
	internalServerError = 3 // 服务器内部错误
)

// Response 用户响应数据
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	// 用户没有登录
	Unauthorized = &Response{Code: unauthorized, Message: "unauthorized"}
	// 服务器内部错误
	InternalServerError = &Response{Code: internalServerError, Message: "internal server error"}
	// 客户端请求参数错误
	BadRequest = &Response{Code: badRequest, Message: "bad request"}
)

func NewSuccess(data interface{}) *Response {
	return &Response{
		Code:    success,
		Message: "success",
		Data:    data,
	}
}
