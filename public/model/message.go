package model

type Message struct {
	Messages []MessageItem
}

// 单条消息投递
type MessageItem struct {
	SenderType   int32  // 发送者类型
	SenderId     int64  // 发送者id
	DeviceId     int64  // 设备id
	ReceiverType int32  // 接收者类型
	ReceiverId   int64  // 接收者id
	Type         int32  // 消息类型
	Content      string // 消息内容
	Sequence     int64  // 消息序列
}

type MessageACK struct {
	DeviceId int64 // 设备id
	UserId   int64 // 用户id
	Sequence int64 // 消息序列
}
