package transfer

type MessageSend struct {
	SenderDeviceId int64  // 发送者设备id
	SenderUserId   int64  // 发送者用户id
	ReceiverType   int32  // 接收者类型，1：单发；2：群发
	ReceiverId     int64  // 接收者id
	Type           int32  // 消息类型
	Content        string // 消息内容
	SendSequence   int64  // 消息序列号
}

type MessageSendACK struct {
	DeviceId     int64 // 设备id
	SendSequence int64 // 消息序列号
}
