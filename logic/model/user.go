package model

import (
	"time"
)

// User 账户
type User struct {
	Id         int64     `json:"id"`          // 用户id
	Number     string    `json:"number"`      // 手机号
	Nickname   string    `json:"nickname"`    // 昵称
	Sex        int       `json:"sex"`         // 性别，1:男；2:女
	Avatar     string    `json:"avatar"`      // 用户头像
	Password   string    `json:"password"`    // 密码
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

// UserRegist 用户注册
type UserRegist struct {
	DeviceId int64  `json:"device_id"` // 设备id
	Token    string `json:"token"`     // 设备token
	Number   string `json:"number"`    // 手机号
	Nickname string `json:"nickname"`  // 昵称
	Sex      int    `json:"sex"`       // 性别，1:男；2:女
	Avatar   string `json:"avatar"`    // 用户头像
	Password string `json:"password"`  // 密码
}

// SignIn 登录结构体
type SignIn struct {
	DeviceId int64  `json:"device_id"`
	Token    string `json:"token"`
	Number   string `json:"number"`
	Password string `json:"password"`
}

// SignInResp 登录响应
type SignInResp struct {
	SendSequence int64 `json:"send_sequence"` // 发送序列号
	SyncSequence int64 `json:"sync_sequence"` // 同步序列号
}
