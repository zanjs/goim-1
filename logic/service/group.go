package service

import (
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/context"
	"goim/public/logger"
)

type groupService struct{}

var GroupService = new(groupService)

// ListByUserId 获取用户群组
func (*groupService) ListByUserId(ctx *context.Context, userId int) ([]*model.Group, error) {
	ids, err := dao.GroupUserDao.ListbyUserId(ctx, userId)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	groups := make([]*model.Group, 0, 5)
	for i := range ids {
		group, err := GroupService.Get(ctx, ids[i])
		if err != nil {
			logger.Sugaer.Error(err)
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// ListGroupUser 获取群组的用户信息
func (*groupService) Get(ctx *context.Context, id int64) (*model.Group, error) {
	group, err := dao.GroupUserDao.Get(ctx, id)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	group.GroupUser, err = dao.GroupUserDao.ListGroupUser(ctx, id)
	if err != nil {
		logger.Sugaer.Error(err)
	}
	return group, err
}

// CreateAndAddUser 创建群组并且添加群成员
func (*groupService) CreateAndAddUser(ctx *context.Context, groupName string, userIds []int64) (int64, error) {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}
	defer ctx.Session.Rollback()

	id, err := dao.GroupDao.Add(ctx, groupName)
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}

	for _, userId := range userIds {
		err := dao.GroupUserDao.Add(ctx, id, userId)
		if err != nil {
			logger.Sugaer.Error(err)
			return 0, err
		}
	}
	ctx.Session.Commit()
	return id, nil
}

// AddUser 给群组添加用户
func (*groupService) AddUser(ctx *context.Context, groupId int64, userIds []int64) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugaer.Error(err)
		return nil
	}
	defer ctx.Session.Rollback()

	for _, userId := range userIds {
		err := dao.GroupUserDao.Add(ctx, groupId, userId)
		if err != nil {
			logger.Sugaer.Error(err)
			return err
		}
	}
	ctx.Session.Commit()
	return nil
}

// DeleteUser 从群组移除用户
func (*groupService) DeleteUser(ctx *context.Context, groupId int64, userIds []int64) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugaer.Error(err)
		return nil
	}
	defer ctx.Session.Rollback()

	for _, userId := range userIds {
		err := dao.GroupUserDao.Delete(ctx, groupId, userId)
		if err != nil {
			logger.Sugaer.Error(err)
			return err
		}
	}
	ctx.Session.Commit()
	return nil
}

func (*groupService) UpdateLabel(ctx *context.Context, groupId int, userId int, label string) error {
	return dao.GroupUserDao.UpdateLabel(ctx, groupId, userId, label)
}
