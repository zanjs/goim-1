package controller

import (
	"goim/logic/model"
	"goim/logic/service"
	"goim/public/imerror"
	"strconv"
)

func init() {
	g := Engine.Group("/user")
	g.GET("/group/:id", handler(UserController{}.ListGroupByUserId))
	g.POST("", handler(UserController{}.Regist))
	g.POST("/signin", handler(UserController{}.SignIn))
}

type UserController struct{}

// ListByUserId 获取用户群组
func (UserController) ListGroupByUserId(c *context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.response(nil, imerror.ErrBadRequest)
	}
	c.response(service.GroupService.ListByUserId(Context(), id))
}

// Regist 用户注册
func (UserController) Regist(c *context) {
	var user model.User
	if c.bindJson(&user) != nil {
		return
	}
	c.response(service.UserService.Regist(Context(), user))
}

// SignIn 用户登录
func (UserController) SignIn(c *context) {
	var signIn model.SignIn
	if c.bindJson(&signIn) != nil {
		return
	}
	c.response(nil, service.UserService.SignIn(Context(), signIn))
}
