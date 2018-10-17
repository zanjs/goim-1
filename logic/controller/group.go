package controller

import (
	"goim/logic/entity"
	"goim/logic/service"
	"log"

	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/group")
	g.GET("/:id", GroupController{}.Get)
	g.POST("", GroupController{}.CreateAndAddUser)
	g.POST("/user", GroupController{}.AddUser)
	g.DELETE("/user", GroupController{}.DeleteUser)
	g.PUT("/user/label", GroupController{}.UpdateLabel)
}

type GroupController struct{}

// Get 获取群组信息
func (GroupController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewBadRequst(err))
		return
	}

	group, err := service.GroupService.Get(Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(group))
}

// CreateAndAddUser 创建群组并且添加成员
func (GroupController) CreateAndAddUser(c *gin.Context) {
	var add entity.GroupAdd
	err := c.ShouldBindJSON(&add)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	id, err := service.GroupService.CreateAndAddUser(Context(), add)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
		return
	}

	c.JSON(OK, NewSuccess(id))
}

// AddUser 给群组添加用户
func (GroupController) AddUser(c *gin.Context) {
	var update entity.GroupUserUpdate
	err := c.ShouldBindJSON(&update)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	err = service.GroupService.AddUser(Context(), update)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
		return
	}

	c.JSON(OK, NewSuccess(nil))
}

// DeleteUser 从群组删除成员
func (GroupController) DeleteUser(c *gin.Context) {
	var update entity.GroupUserUpdate
	err := c.ShouldBindJSON(&update)
	if err != nil {
		c.JSON(OK, NewBadRequst(err))
		return
	}
	err = service.GroupService.DeleteUser(Context(), update)
	if err != nil {
		log.Println(err)
		c.JSON(OK, NewError(err))
		return
	}

	c.JSON(OK, NewSuccess(nil))
}

// UpdateLabel 更新用户群组备注
func (GroupController) UpdateLabel(c *gin.Context) {
	var json struct {
		GroupId int    `json:"group_id"`
		UserId  int    `json:"user_id"`
		Label   string `json:"label"`
	}
	c.ShouldBindJSON(&json)
	err := service.GroupService.UpdateLabel(Context(), json.GroupId, json.UserId, json.Label)
	if err != nil {
		c.JSON(OK, NewError(err))
		return
	}
	c.JSON(OK, NewSuccess(nil))
}
