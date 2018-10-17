package service

import (
	"fmt"
	"goim/logic/dao"
	"goim/logic/entity"
	"goim/public/context"
	"log"
)

type friendService struct{}

var FriendService = new(friendService)

// Add 添加好友关系
func (*friendService) Add(ctx *context.Context, add entity.FriendAdd) error {
	err := ctx.Session.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer ctx.Session.Rollback()

	friend1 := entity.Friend{
		UserId:   add.UserId,
		FriendId: add.Friend,
		Label:    add.UserLabel,
	}
	err = dao.FriendDao.Add(ctx, friend1)
	if err != nil {
		fmt.Println(err)
		return err
	}

	friend2 := entity.Friend{
		UserId:   add.Friend,
		FriendId: add.UserId,
		Label:    add.FriendLabel,
	}
	err = dao.FriendDao.Add(ctx, friend2)
	if err != nil {
		fmt.Println(err)
		return err
	}

	ctx.Session.Commit()
	return nil
}

// Delete 删除好友关系
func (*friendService) Delete(ctx *context.Context, userId, friend int) error {
	err := ctx.Session.Begin()
	if err != nil {
		log.Println(err)
	}
	defer ctx.Session.Rollback()

	err = dao.FriendDao.Delete(ctx, userId, friend)
	if err != nil {
		log.Println(err)
		return err
	}

	err = dao.FriendDao.Delete(ctx, friend, userId)
	if err != nil {
		log.Println(err)
		return err
	}

	ctx.Session.Commit()
	return nil
}
