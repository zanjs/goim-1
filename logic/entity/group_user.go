package entity

import (
	"time"
)

// GroupUser 群组成员关系
type GroupUser struct {
	Id         int       `json:"id"`          // 自增主键
	GroupId    int       `json:"group_id"`    // 组id
	UserId     int       `json:"user_id"`     // 用户id
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
