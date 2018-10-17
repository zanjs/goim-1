package controller

import (
	"goim/logic/entity"
	"goim/logic/service"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/user")
	g.GET("/group/:id", UserController{}.ListGroupByUserId)
	g.POST("", UserController{}.Regist)
	g.POST("/signin", UserController{}.SignIn)
}

type UserController struct{}

// ListByUserId 获取用户群组
func (UserController) ListGroupByUserId(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewBadRequst(err))
	}

	groups, err := service.GroupService.ListByUserId(Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
	}
	c.JSON(OK, NewSuccess(groups))
}

// Regist 用户注册
func (UserController) Regist(c *gin.Context) {
	var user entity.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	id, err := service.UserService.Regist(Context(), user)
	if err != nil {
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(id))
}

// SignIn 用户登录
func (UserController) SignIn(c *gin.Context) {
	var signIn entity.SignIn
	err := c.ShouldBindJSON(&signIn)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	err = service.UserService.SignIn(Context(), signIn)
	if err != nil {
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(nil))
}
