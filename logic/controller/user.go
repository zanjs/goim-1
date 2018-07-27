package controller

import (
	"goim/logic/entity"
	"goim/logic/service"

	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/user")
	g.POST("", UserController{}.Regist)
	g.PUT("/signin", UserController{}.SignIn)
}

type UserController struct{}

// Regist 用户注册
func (UserController) Regist(c *gin.Context) {
	var user entity.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	id, err := service.NewUserService().Regist(user)
	if err != nil {
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(id))
}

// Regist 用户注册
func (UserController) SignIn(c *gin.Context) {
	var signIn entity.SignIn
	err := c.ShouldBindJSON(&signIn)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	err = service.NewUserService().SignIn(signIn)
	if err != nil {
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(nil))
}
