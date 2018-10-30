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
	c.response(service.UserService.Regist(Context(), c.deviceId, regist))
}

// SignIn 用户登录
func (UserController) SignIn(c *context) {
	var data struct {
		Number   string `json:"number"`
		Password string `json:"password"`
	}
	if c.bindJson(&data) != nil {
		return
	}
	c.response(service.UserService.SignIn(Context(), c.deviceId, data.Number, data.Password))
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
