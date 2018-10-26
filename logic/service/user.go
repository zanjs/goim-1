package service

import (
	"database/sql"
	"errors"
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/context"
	"goim/public/logger"
)

type userService struct{}

var UserService = new(userService)

var ErrNumberExist = errors.New("user number exist")

func (*userService) Regist(ctx *context.Context, user model.User) (int64, error) {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}
	defer ctx.Session.Rollback()

	id, err := dao.UserDao.Add(ctx, user)
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}

	if id == 0 {
		return 0, ErrNumberExist
	}

	err = dao.DeviceSyncSequenceDao.Add(ctx, id, 0)
	if err != nil {
		logger.Sugaer.Error(err)
		return 0, err
	}
	ctx.Session.Commit()
	return id, nil
}

var (
	ErrDeviceNotFound = errors.New("device not found")
	ErrToken          = errors.New("error token")
	ErrUserNotFound   = errors.New("user not found")
	ErrPassword       = errors.New("error password")
)

// SignIn 登录
func (*userService) SignIn(ctx *context.Context, signIn model.SignIn) error {
	token, err := dao.DeviceDao.GetToken(ctx, signIn.DeviceId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrDeviceNotFound
		}
		logger.Sugaer.Error(err)
		return err
	}

	if signIn.Token != token {
		return ErrToken
	}

	password, err := dao.UserDao.GetPassword(ctx, signIn.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		logger.Sugaer.Error(err)
		return err
	}

	if signIn.Password != password {
		return ErrPassword
	}

	err = dao.DeviceDao.UpdateUserId(ctx, signIn.DeviceId, signIn.UserId)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	return nil
}
