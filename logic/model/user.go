package model

import (
	"time"
)

// User 账户
type User struct {
	Id         int64     `json:"id"`          // 用户id
	Number     string    `json:"number"`      // 手机号
	Name       string    `json:"name"`        // 昵称
	Sex        int       `json:"sex"`         // 性别，1:男；2:女
	Avatar     string    `json:"avatar"`      // 用户头像
	Password   string    `json:"password"`    // 密码
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

// SignIn 登录结构体
type SignIn struct {
	DeviceId int64  `json:"device_id"`
	Token    string `json:"token"`
	UserId   int64  `json:"user_id"`
	Password string `json:"password"`
}
