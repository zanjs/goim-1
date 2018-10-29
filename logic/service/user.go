package service

import (
	"database/sql"
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/ctx"
	"goim/public/imerror"
	"goim/public/logger"
)

type userService struct{}

var UserService = new(userService)

func (*userService) Regist(ctx *ctx.Context, regist model.UserRegist) (*model.SignInResp, error) {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	defer ctx.Session.Rollback()

	// 设备验证
	device, err := dao.DeviceDao.Get(ctx, regist.DeviceId)
	if err == sql.ErrNoRows {
		return nil, imerror.ErrDeviceIdOrToken
	}
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	if device.Token != regist.Token {
		return nil, imerror.ErrDeviceIdOrToken
	}

	// 添加用户
	user := model.User{
		Number:   regist.Number,
		Nickname: regist.Nickname,
		Sex:      regist.Sex,
		Avatar:   regist.Avatar,
		Password: regist.Password,
	}
	userId, err := dao.UserDao.Add(ctx, user)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	if userId == 0 {
		return nil, imerror.ErrNumberUsed
	}

	err = dao.UserSequenceDao.Add(ctx, userId, 0)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	err = dao.DeviceDao.UpdateUserId(ctx, regist.DeviceId, userId)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	dao.DeviceSendSequenceDao.UpdateSendSequence(ctx, regist.DeviceId, 0)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	dao.DeviceSyncSequenceDao.UpdateSyncSequence(ctx, regist.DeviceId, 0)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	err = ctx.Session.Commit()
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	return &model.SignInResp{
		SendSequence: 0,
		SyncSequence: 0,
	}, nil
}

// SignIn 登录
func (*userService) SignIn(ctx *ctx.Context, signIn model.SignIn) (*model.SignInResp, error) {
	err := ctx.Session.Begin()
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	defer ctx.Session.Rollback()
	// 设备验证
	device, err := dao.DeviceDao.Get(ctx, signIn.DeviceId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, imerror.ErrDeviceIdOrToken
		}
		logger.Sugaer.Error(err)
		return nil, err
	}

	if signIn.Token != device.Token {
		return nil, imerror.ErrDeviceIdOrToken
	}

	// 用户验证
	user, err := dao.UserDao.GetByNumber(ctx, signIn.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, imerror.ErrNameOrPassword
		}
		logger.Sugaer.Error(err)
		return nil, err
	}
	if signIn.Password != user.Password {
		return nil, imerror.ErrNameOrPassword
	}

	err = dao.DeviceDao.UpdateUserId(ctx, signIn.DeviceId, user.Id)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	err = dao.DeviceSendSequenceDao.UpdateSendSequence(ctx, signIn.DeviceId, 0)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	maxSyncSequence, err := dao.DeviceSyncSequenceDao.GetMaxSyncSequenceByUserId(ctx, user.Id)
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}

	err = ctx.Session.Commit()
	if err != nil {
		logger.Sugaer.Error(err)
		return nil, err
	}
	return &model.SignInResp{
		SendSequence: 0,
		SyncSequence: maxSyncSequence,
	}, nil
}
