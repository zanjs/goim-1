package dao

import (
	"goim/logic/model"
	"goim/public/context"
	"goim/public/logger"
)

type groupDao struct{}

var GroupDao = new(groupDao)

// Get 获取群组信息
func (*groupDao) Get(ctx *context.Context, id int) (*model.Group, error) {
	row := ctx.Session.QueryRow("select name from t_group where id = ?", id)
	group := new(model.Group)
	err := row.Scan(&group.Name)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	return group, nil
}

// Insert 插入一条群组信息
func (*groupDao) Add(ctx *context.Context, name string) (int64, error) {
	result, err := ctx.Session.Exec("insert into t_group(name) value(?)", name)
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}
	return id, nil
}
