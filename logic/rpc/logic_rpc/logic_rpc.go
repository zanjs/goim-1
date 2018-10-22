package logic_rpc

import (
	"database/sql"
	"goim/logic/dao"
	"goim/logic/rpc/connect_rpc"
	"goim/logic/service"
	"goim/public/context"
	"goim/public/transfer"
	"log"
)

type logicRPC struct{}

var LogicRPC = new(logicRPC)

// SignIn 处理设备登录
func (s *logicRPC) SignIn(ctx *context.Context, signIn transfer.SignIn) *transfer.SignInACK {
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

// SyncTrigger 处理消息同步触发
func (s *logicRPC) SyncTrigger(ctx *context.Context, trigger transfer.SyncTrigger) error {
	dbMessages, err := dao.MessageDao.ListByUserIdAndSequence(ctx, trigger.UserId, trigger.SyncSequence)
	if err != nil {
		log.Println(err)
		return err
	}

	messages := make([]transfer.MessageItem, 0, len(dbMessages))
	for _, v := range dbMessages {
		item := transfer.MessageItem{}
		item.SenderType = v.SenderType
		item.SenderId = v.SenderId
		item.SenderDeviceId = v.SenderDeviceId
		item.ReceiverType = v.ReceiverType
		item.ReceiverId = v.ReceiverId
		item.Type = v.Type
		item.Content = v.Content
		item.Sequence = v.Sequence
		messages = append(messages, item)
	}

	message := transfer.Message{DeviceId: trigger.DeviceId, Messages: messages}
	connect_rpc.ConnectRPC.SendMessage(message)
	return nil
}

// MessageSend 处理消息发送
func (s *logicRPC) MessageSend(ctx *context.Context, send transfer.MessageSend) error {
	// 检查消息是否重复发送
	sendSequence, err := dao.DeviceSendSequenceDao.Get(ctx, send.SenderDeviceId)
	if err != nil {
		log.Println(err)
		return err
	}
	if send.SendSequence <= sendSequence {
		return nil
	}
	err = dao.DeviceSendSequenceDao.UpdateSequence(ctx, send.SenderDeviceId, send.SendSequence)
	if err != nil {
		log.Println(err)
		return err
	}

	if send.ReceiverType == service.ReceiverUser {
		service.MessageService.SendToFriend(ctx, send)
		return nil
	}
	if send.ReceiverType == service.ReceiverGroup {
		service.MessageService.SendToGroup(ctx, send)
		return nil
	}

	return nil
}

// MessageACK 处理消息回执
func (s *logicRPC) MessageACK(ctx *context.Context, ack transfer.MessageACK) error {
	err := dao.DeviceSyncSequenceDao.UpdateSequence(ctx, ack.DeviceId, ack.SyncSequence)
	if err != nil {
		log.Println(err)
	}
	return nil
}

// OffLine 处理设备离线
func (s *logicRPC) OffLine(ctx *context.Context, deviceId int64) error {
	err := dao.DeviceDao.UpdateStatus(ctx, deviceId, service.DeviceOffline)
	if err != nil {
		log.Println(err)
	}
	return nil
}
