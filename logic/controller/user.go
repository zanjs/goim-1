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

// Regist 用户注册
func (UserController) Regist(c *context) {
	var regist model.UserRegist
	if c.bindJson(&regist) != nil {
		return
	}
	c.response(service.UserService.Regist(Context(), regist))
}

// SignIn 用户登录
func (UserController) SignIn(c *context) {
	var signIn model.SignIn
	if c.bindJson(&signIn) != nil {
		return
	}
	c.response(service.UserService.SignIn(Context(), signIn))
}

// ListByUserId 获取用户群组
func (UserController) ListGroupByUserId(c *context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.response(nil, imerror.ErrBadRequest)
	}
	c.response(service.GroupService.ListByUserId(Context(), id))
}
