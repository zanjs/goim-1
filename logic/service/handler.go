package service

import (
	"database/sql"
	"goim/connect"
	"goim/logic/dao"
	"goim/public/context"
	"goim/public/transfer"
	"log"
)

type handlerService struct{}

var HandlerService = new(handlerService)

// HandleSignIn 处理设备登录
func (s *handlerService) HandleSignIn(ctx *context.Context, signIn transfer.SignIn) *transfer.SignInACK {
	device, err := dao.DeviceDao.Get(ctx, signIn.DeviceId)
	if err == sql.ErrNoRows {
		return &transfer.SignInACK{
			Code:    transfer.CodeSignInFail,
			Message: "fail",
		}
	}

	if err != nil {
		log.Println(err)
	}

	if device.UserId == signIn.UserId && device.Token == signIn.Token {
		return &transfer.SignInACK{
			Code:    transfer.CodeSignInSuccess,
			Message: "success",
		}
	}

	return &transfer.SignInACK{
		Code:    transfer.CodeSignInFail,
		Message: "fail",
	}
	return nil
}

// HandleSyncTrigger 处理消息同步触发
func (s *handlerService) HandleSyncTrigger(ctx *context.Context, trigger transfer.SyncTrigger) error {
	dbMessages, err := dao.MessageDao.ListByUserIdAndSequence(ctx, trigger.UserId, trigger.Sequence)
	if err != nil {
		log.Println(err)
		return err
	}

	messages := make([]transfer.MessageItem, 0, len(dbMessages))
	for _, v := range dbMessages {
		item := transfer.MessageItem{}
		item.SenderType = v.SenderType
		item.SenderId = v.SenderId
		item.DeviceId = v.DeviceId
		item.ReceiverType = v.ReceiverType
		item.ReceiverId = v.ReceiverId
		item.Type = v.Type
		item.Content = v.Content
		item.Sequence = v.Sequence
		messages = append(messages, item)
	}

	message := transfer.Message{DeviceId: trigger.DeviceId, Messages: messages}
	connect.HandleMessage(message)
	return nil
}

// HandleMessageSend 处理消息发送
func (s *handlerService) HandleMessageSend(ctx *context.Context, send transfer.MessageSend) error {
	// 检查消息是否重复发送
	sendSequence, err := dao.DeviceSendSequenceDao.Get(ctx, send.DeviceId)
	if err != nil {
		log.Println(err)
		return err
	}
	if send.SendSequence <= sendSequence {
		return nil
	}
	err = dao.DeviceSendSequenceDao.UpdateSequence(ctx, send.DeviceId, send.SendSequence)
	if err != nil {
		log.Println(err)
		return err
	}

	if send.ReceiverType == ReceiverUser {
		MessageService.SendToUser(ctx, send)
		return nil
	}
	if send.ReceiverType == ReceiverGroup {
		MessageService.SendToGroup(ctx, send)
		return nil
	}

	return nil
}

// HandleMessageACK 处理消息回执
func (s *handlerService) HandleMessageACK(ctx *context.Context, ack transfer.MessageACK) error {
	err := dao.DeviceSyncSequenceDao.UpdateSequence(ctx, ack.DeviceId, ack.Sequence)
	if err != nil {
		log.Println(err)
	}
	return nil
}

// HandleOffLine 处理设备离线
func (s *handlerService) HandleOffLine(ctx *context.Context, deviceId int64) error {
	err := dao.DeviceDao.UpdateStatus(ctx, deviceId, DeviceOffline)
	if err != nil {
		log.Println(err)
	}
	return nil
}
