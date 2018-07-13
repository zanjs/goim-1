package entity

import "time"

// Device 设备
type Device struct {
	Id         int       `json:"id"`          // 设备id
	UserId     int       `json:"user_id"`     // 用户id
	Token      string    `json:"token"`       // 设备登录的token
	Type       int       `json:"device_type"` // 设备类型,0:Android；1:IOS；2：Windows;3:Web
	Model      int       `json:"model"`       // 机型
	Version    string    `json:"version"`     // 设备版本
	Status     int       `json:"state"`       // 在线状态，0：不在线；1：在线
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}
