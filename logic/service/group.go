package service

import (
	"goim/lib/context"
	"goim/logic/dao"
	"goim/logic/entity"
	"log"
)

type groupService struct{}

var GroupService = new(groupService)

// ListByUserId 获取用户群组
func (*groupService) ListByUserId(ctx *context.Context, userId int) ([]*entity.Group, error) {
	ids, err := dao.GroupUserDao.ListbyUserId(ctx, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	groups := make([]*entity.Group, 0, 5)
	for i := range ids {
		group, err := GroupService.Get(ctx, ids[i])
		if err != nil {
			log.Println(err)
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// ListGroupUser 获取群组的用户信息
func (*groupService) Get(ctx *context.Context, id int) (*entity.Group, error) {
	group, err := dao.GroupUserDao.Get(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	group.GroupUser, err = dao.GroupUserDao.ListGroupUser(ctx, id)
	if err != nil {
		log.Println(err)
	}
	return group, err
}

// CreateAndAddUser 创建群组并且添加群成员
func (*groupService) CreateAndAddUser(ctx *context.Context, add entity.GroupAdd) (int64, error) {
	err := ctx.Session.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer ctx.Session.Rollback()

	id, err := dao.GroupDao.Add(ctx, add.Name)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	for i := range add.UserIds {
		err := dao.GroupUserDao.Add(ctx, id, add.UserIds[i])
		if err != nil {
			log.Println(err)
			return 0, err
		}
	}
	ctx.Session.Commit()
	return id, nil
}

// AddUser 给群组添加用户
func (*groupService) AddUser(ctx *context.Context, update entity.GroupUserUpdate) error {
	err := ctx.Session.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer ctx.Session.Rollback()

	for i := range update.UserIds {
		err := dao.GroupUserDao.Add(ctx, update.GroupId, update.UserIds[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	ctx.Session.Commit()
	return nil
}

// DeleteUser 从群组移除用户
func (*groupService) DeleteUser(ctx *context.Context, update entity.GroupUserUpdate) error {
	err := ctx.Session.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer ctx.Session.Rollback()

	for i := range update.UserIds {
		err := dao.GroupUserDao.Delete(ctx, update.GroupId, update.UserIds[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	ctx.Session.Commit()
	return nil
}

func (*groupService) UpdateLabel(ctx *context.Context, groupId int, userId int, label string) error {
	return dao.GroupUserDao.UpdateLabel(ctx, groupId, userId, label)
}
