package entity

import (
	"time"
)

// Message 消息
type Message struct {
	Id           int       `json:"id"`            // 自增主键
	UserId       int       `json:"user_id"`       // 用户id
	Sender       int       `json:"sender"`        // 发送者账户id
	ReceiverType int       `json:"receiver_type"` // 接收者账户id
	Receiver     int       `json:"receiver"`      // 接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id
	Type         int       `json:"type"`          // 消息类型,0：文本；1：语音；2：图片
	Content      string    `json:"content"`       // 内容
	SyncKey      int       `json:"sync_key"`      // 消息同步序列
	CreateTime   time.Time `json:"create_time"`   // 创建时间
	UpdateTime   time.Time `json:"update_time"`   // 更新时间
}
