package controller

import (
	"goim/logic/entity"
	"goim/logic/service"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/device")
	g.POST("", DeviceController{}.Regist)
}

type DeviceController struct{}

// Regist 设备注册
func (DeviceController) Regist(c *gin.Context) {
	var device entity.Device
	err := c.ShouldBindJSON(&device)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewBadRequst(err))
		return
	}

	if device.Type == 0 || device.Model == "" || device.Version == "" {
		log.Println(err)
		c.JSON(OK, NewBadRequst(nil))
		return
	}

	id, token, err := service.DeviceService.Regist(Context(), device)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
		return
	}

	result := make(map[string]interface{}, 2)
	result["id"] = id
	result["token"] = token

	c.JSON(OK, result)
}
