package service

import (
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/ctx"
	"goim/public/logger"
)

type friendService struct{}

var FriendService = new(friendService)

// Add 添加好友关系
func (*friendService) Add(ctx *ctx.Context, add model.FriendAdd) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	defer ctx.Session.Rollback()

	friend1 := model.Friend{
		UserId:   add.UserId,
		FriendId: add.Friend,
		Label:    add.UserLabel,
	}
	err = dao.FriendDao.Add(ctx, friend1)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	friend2 := model.Friend{
		UserId:   add.Friend,
		FriendId: add.UserId,
		Label:    add.FriendLabel,
	}
	err = dao.FriendDao.Add(ctx, friend2)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	ctx.Session.Commit()
	return nil
}

// Delete 删除好友关系
func (*friendService) Delete(ctx *ctx.Context, userId, friend int) error {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	defer ctx.Session.Rollback()

	err = dao.FriendDao.Delete(ctx, userId, friend)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	err = dao.FriendDao.Delete(ctx, friend, userId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	ctx.Session.Commit()
	return nil
}
