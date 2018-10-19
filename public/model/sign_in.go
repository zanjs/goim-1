package model

// SignIn 设备登录
type SignIn struct {
	DeviceId uint64 // 设备id
	UserId   uint64 // 用户id
	Token    string // token
}

//  SignInACK 设备登录回执
type SignInACK struct {
	Code    int    // 设备id
	Message string // 用户id
}
