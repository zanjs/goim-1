package entity

import (
	"time"
)

// Group 群组
type Group struct {
	Id         int       `json:"id"`          // 群组id
	Name       int       `json:"name"`        // 组名
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
