package controller

import (
	"goim/logic/entity"
	"goim/logic/service"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/group")
	g.POST("", DeviceController{}.Regist)
}

type GroupController struct{}

func (GroupController) CreateAndAddUser(c *gin.Context) {
	var add entity.GroupAdd
	err := c.ShouldBindJSON(&add)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	id, err := service.NewGroupService().CreateAndAddUser(add)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
		return
	}

	c.JSON(OK, NewSuccess(id))
}

func (GroupController) AddUser(c *gin.Context) {
	var add entity.GroupUserUpdate
	err := c.ShouldBindJSON(&add)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	err = service.NewGroupService().AddUser(add)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
		return
	}

	c.JSON(OK, NewSuccess(nil))
}
