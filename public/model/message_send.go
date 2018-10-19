package model

type MessageSend struct {
	DeviceId     int64  // 设备id
	UserId       int64  // 用户id
	ReceiverType int32  // 接收者类型
	ReceiverId   int64  // 接收者id
	Type         int32  // 消息类型
	Content      string // 消息内容
	Sequence     int64  // 消息序列号
}

type MessageSendACK struct {
	Sequence int64 // 消息序列号
}
