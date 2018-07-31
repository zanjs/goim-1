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

type GroupAdd struct {
	Name    string `json:"name"`     // 群组名称
	UserIds []int  `json:"user_ids"` // 群组成员
}

type GroupUserUpdate struct {
	GroupId int   `json:"name"`     // 群组名称
	UserIds []int `json:"user_ids"` // 群组成员
}
