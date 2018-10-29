package controller

import (
	"goim/logic/model"
	"goim/logic/service"
	"goim/public/imerror"
	"strconv"
)

func init() {
	g := Engine.Group("/friend")
	g.GET("/:user_id", handler(FriendControlelr{}.Friends))
	g.POST("", handler(FriendControlelr{}.Add))
	g.DELETE("", handler(FriendControlelr{}.Delete))
}

type FriendControlelr struct{}

// Friend 好友
func (FriendControlelr) Friends(c *context) {
	idStr := c.Param("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.response(nil, imerror.ErrBadRequest)
		return
	}
	c.response(service.FriendService.ListUserFriend(Context(), id))
}

func (FriendControlelr) Add(c *context) {
	var friendAdd model.FriendAdd
	if c.bindJson(&friendAdd) != nil {
		return
	}
	c.response(nil, service.FriendService.Add(Context(), friendAdd))
}

func (FriendControlelr) Delete(c *context) {
	var json struct {
		UserId int `json:"user_id"`
		Friend int `json:"friend"`
	}
	if c.bindJson(&json) != nil {
		return
	}
	c.response(nil, service.FriendService.Delete(Context(), json.UserId, json.Friend))
}
