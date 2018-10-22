package transfer

type Message struct {
	DeviceId int64
	Messages []MessageItem
}

// 单条消息投递
type MessageItem struct {
	SenderType     int    // 发送者类型
	SenderId       int64  // 发送者id
	SenderDeviceId int64  // 发送者设备id
	ReceiverType   int    // 接收者类型
	ReceiverId     int64  // 接收者id
	Type           int    // 消息类型
	Content        string // 消息内容
	Sequence       int64  // 消息序列
}

type MessageACK struct {
	DeviceId     int64 // 设备id
	UserId       int64 // 用户id
	SyncSequence int64 // 消息序列
}
