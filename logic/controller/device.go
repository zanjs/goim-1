package controller

import (
	"goim/logic/model"
	"goim/logic/service"
	"goim/public/errors"
	"goim/public/imerror"
)

func init() {
	g := Engine.Group("/device")
	g.POST("", handler(DeviceController{}.Regist))
}

type DeviceController struct{}

// Regist 设备注册
func (DeviceController) Regist(c *context) {
	var device model.Device
	if c.ShouldBindJSON(&device) != nil {
		return
	}

	if device.Type == 0 || device.Model == "" || device.Version == "" {
		c.response(nil, imerror.ErrBadRequest)
		return
	}

	id, token, err := service.DeviceService.Regist(Context(), device)
	c.response(map[string]interface{}{"id": id, "token": token}, err)
}
