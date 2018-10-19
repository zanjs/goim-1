package model

import (
	"time"
)

// Friend 好友关系
type Friend struct {
	Id         int64     `json:"id"`          // 自增主键
	UserId     int64     `json:"user_id"`     // 账户id
	FriendId   int64     `json:"friend_id"`   // 好友账户id
	Label      string    `json:"label"`       // 备注，标签
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

type FriendAdd struct {
	UserId      int64  `json:"user_id"`      // 用户id
	UserLabel   string `json:"user_label"`   // 用户对好友的标签
	Friend      int64  `json:"friend"`       // 好友id
	FriendLabel string `json:"friend_label"` // 好友对用户的标签
}

type FriendUser struct {
	UserId int64  `json:"user_id"` // 用户id
	Label  string `json:"lable"`   // 用户对好友的标签
	Number string `json:"number"`  // 手机号
	Name   string `json:"name"`    // 昵称
	Sex    int    `json:"sex"`     // 性别，1:男；2:女
	Img    string `json:"img"`     // 用户头像
}