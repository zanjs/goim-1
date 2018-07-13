package entity

import (
	"time"
)

// User 账户
type User struct {
	Id         int       `json:"id"`          // 用户id
	Number     string    `json:"number"`      // 手机号
	Name       string    `json:"nickname"`    // 昵称
	Password   string    `json:"password"`    // 密码
	Sex        int       `json:"sex"`         // 性别，1:男；2:女
	Img        string    `json:"img"`         // 用户头像
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
