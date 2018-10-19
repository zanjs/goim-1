package service

import (
	"database/sql"
	"goim/logic/dao"
	"goim/public/context"
	"goim/public/model"
	"log"
)

type handlerService struct{}

var HandlerService = new(handlerService)

// HandleSignIn 处理设备登录
func (s *handlerService) HandleSignIn(ctx *context.Context, signIn model.SignIn) *model.SignInACK {
	device, err := dao.DeviceDao.Get(ctx, signIn.DeviceId)
	if err == sql.ErrNoRows {
		return &model.SignInACK{
			Code:    model.CodeSignInFail,
			Message: "fail",
		}
	}

	if err != nil {
		log.Println(err)
	}

	if device.UserId == signIn.UserId && device.Token == signIn.Token {
		return &model.SignInACK{
			Code:    model.CodeSignInSuccess,
			Message: "success",
		}
	}

	return &model.SignInACK{
		Code:    model.CodeSignInFail,
		Message: "fail",
	}
	return nil
}

// HandleSyncTrigger 处理消息同步触发
func (s *handlerService) HandleSyncTrigger(ctx *context.Context, trigger model.SyncTrigger) error {
	return nil
}

// HandleMessageSend 处理消息发送
func (s *handlerService) HandleMessageSend(ctx *context.Context, send model.MessageSend) error {

}

// HandleMessageACK 处理消息回执
func (s *handlerService) HandleMessageACK(ctx *context.Context, message model.MessageACK) error {

}

// HandleOffLine 处理设备离线
func (s *handlerService) HandleOffLine(ctx *context.Context, deviceId int64) error {

}
