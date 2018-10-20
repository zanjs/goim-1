package service

import (
	"goim/logic/dao"
	"goim/public/context"
	"log"
)

type userRequenceService struct{}

var UserRequenceService = new(userRequenceService)

// GetNext 获取下一个序列
func (*userRequenceService) GetNext(ctx *context.Context, userId int64) (int64, error) {
	err := ctx.Session.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	ctx.Session.Rollback()

	err = dao.UserSequenceDao.Increase(ctx, userId)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	sequence, err := dao.UserSequenceDao.GetSequence(ctx, userId)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return sequence, nil
}
