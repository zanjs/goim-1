package controller

import (
	"goim/logic/entity"

	"goim/logic/service"

	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/friend")
	g.GET("/:user_id", FriendControlelr{}.Friends)
	g.POST("", FriendControlelr{}.Add)
	g.DELETE("", FriendControlelr{}.Delete)
}

type FriendControlelr struct{}

// Friend 好友
func (FriendControlelr) Friends(c *gin.Context) {
	idStr := c.Param("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(OK, NewBadRequst(errors.New("error id")))
	}
	users, err := service.NewFriendService().ListUserFriend(id)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(users))

}

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

func (FriendControlelr) Delete(c *gin.Context) {
	var json struct {
		UserId int `json:"user_id"`
		Friend int `json:"friend"`
	}
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}

	err = service.NewFriendService().Delete(json.UserId, json.Friend)
	if err != nil {
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(nil))
}
