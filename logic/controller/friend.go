package controller

import (
	"goim/logic/entity"

	"goim/logic/service"

	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/friend")
	g.POST("", FriendControlelr{}.Add)
}

type FriendControlelr struct{}

func (FriendControlelr) Add(c *gin.Context) {
	var friendAdd entity.FriendAdd
	err := c.ShouldBindJSON(&friendAdd)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}

	err = service.NewFriendService().Add(friendAdd)
	if err != nil {
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(nil))
}
