package entity

import (
	"time"
)

// Friend 好友关系
type Friend struct {
	Id         int       `json:"id"`          // 自增主键
	UserId     int       `json:"user_id"`     // 账户id
	Friend     int       `json:"friend"`      // 好友账户id
	Label      string    `json:"label"`       // 备注，标签
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
