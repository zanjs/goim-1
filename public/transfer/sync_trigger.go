package transfer

// 同步消息触发
type SyncTrigger struct {
	DeviceId int64 // 设备id
	UserId   int64 // 用户id
	Sequence int64 // 已经同步的消息序列号
}
