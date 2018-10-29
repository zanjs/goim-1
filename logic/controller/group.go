package controller

import (
	"goim/logic/model"
	"goim/logic/service"
	"goim/public/imerror"
	"strconv"
)

func init() {
	g := Engine.Group("/group")
	g.GET("/:id", handler(GroupController{}.Get))
	g.POST("", handler(GroupController{}.CreateAndAddUser))
	g.POST("/user", handler(GroupController{}.AddUser))
	g.DELETE("/user", handler(GroupController{}.DeleteUser))
	g.PUT("/user/label", handler(GroupController{}.UpdateLabel))
}

type GroupController struct{}

// Get 获取群组信息
func (GroupController) Get(c *context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.response(nil, imerror.ErrBadRequest)
		return
	}
	c.response(service.GroupService.Get(Context(), id))
}

// CreateAndAddUser 创建群组并且添加成员
func (GroupController) CreateAndAddUser(c *context) {
	var data = struct {
		Name    string  `json:"name"`     // 群组名称
		UserIds []int64 `json:"user_ids"` // 群组成员
	}{}
	if c.bindJson(&data) != nil {
		return
	}
	c.response(service.GroupService.CreateAndAddUser(Context(), data.Name, data.UserIds))
}

// AddUser 给群组添加用户
func (GroupController) AddUser(c *context) {
	var update model.GroupUserUpdate
	if c.bindJson(&update) != nil {
		return
	}
	c.response(nil, service.GroupService.AddUser(Context(), update))
}

// DeleteUser 从群组删除成员
func (GroupController) DeleteUser(c *context) {
	var update model.GroupUserUpdate
	if c.bindJson(&update) != nil {
		return
	}
	c.response(nil, service.GroupService.DeleteUser(Context(), update))
}

// UpdateLabel 更新用户群组备注
func (GroupController) UpdateLabel(c *context) {
	var json struct {
		GroupId int    `json:"group_id"`
		UserId  int    `json:"user_id"`
		Label   string `json:"label"`
	}
	if c.bindJson(&json) != nil {
		return
	}
	err := service.GroupService.UpdateLabel(Context(), json.GroupId, json.UserId, json.Label)
	c.response(nil, err)
}
