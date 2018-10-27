package logic_rpc

import (
	"database/sql"
	"goim/logic/dao"
	"goim/logic/rpc/connect_rpc"
	"goim/logic/service"
	"goim/public/context"
	"goim/public/lib"
	"goim/public/logger"
	"goim/public/transfer"
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
		logger.Sugaer.Error(err)
		return nil
	}

	var code int
	var message string
	if device.UserId == signIn.UserId && device.Token == signIn.Token {
		code = transfer.CodeSignInSuccess
		message = "success"
	} else {
		code = transfer.CodeSignInFail
		message = "fail"
	}

	logger.Sugaer.Infow("设备登录",
		"device_id:", signIn.DeviceId,
		"user_id", signIn.UserId,
		"token", signIn.Token,
		"result", message)

	return &transfer.SignInACK{
		Code:    code,
		Message: message,
	}
	return nil
}

// SyncTrigger 处理消息同步触发
func (s *logicRPC) SyncTrigger(ctx *context.Context, trigger transfer.SyncTrigger) error {
	logger.Sugaer.Infow("同步触发",
		"device_id:", trigger.DeviceId,
		"user_id", trigger.UserId,
		"sync_sequence", trigger.SyncSequence)

	dbMessages, err := dao.MessageDao.ListByUserIdAndSequence(ctx, trigger.UserId, trigger.SyncSequence)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}

	messages := make([]transfer.MessageItem, 0, len(dbMessages))
	for _, v := range dbMessages {
		item := transfer.MessageItem{}

		item.MessageId = v.MessageId
		item.SenderType = v.SenderType
		item.SenderId = v.SenderId
		item.SenderDeviceId = v.SenderDeviceId
		item.ReceiverType = v.ReceiverType
		item.ReceiverId = v.ReceiverId
		item.Type = v.Type
		item.Content = v.Content
		item.Sequence = v.Sequence
		item.SendTime = v.SendTime

		messages = append(messages, item)
	}

	message := transfer.Message{DeviceId: trigger.DeviceId, Messages: messages}
	connect_rpc.ConnectRPC.SendMessage(message)

	logger.Sugaer.Infow("消息同步",
		"device_id:", trigger.DeviceId,
		"user_id", trigger.UserId,
		"messages", message.GetLog())
	return nil
}

// MessageSend 处理消息发送
func (s *logicRPC) MessageSend(ctx *context.Context, send transfer.MessageSend) error {
	var err error
	send.MessageId = lib.Lid.Get()

	logger.Sugaer.Infow("消息发送",
		"device_id", send.SenderDeviceId,
		"user_id", send.SenderUserId,
		"message_id", send.MessageId,
		"send_sequence", send.SendSequence)

	// 检查消息是否重复发送
	sendSequence, err := dao.DeviceSendSequenceDao.Get(ctx, send.SenderDeviceId)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	if send.SendSequence <= sendSequence {
		return nil
	}
	err = dao.DeviceSendSequenceDao.UpdateSequence(ctx, send.SenderDeviceId, send.SendSequence)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}

	if send.ReceiverType == service.ReceiverUser {
		service.MessageService.SendToFriend(ctx, send)
	}
	if send.ReceiverType == service.ReceiverGroup {
		service.MessageService.SendToGroup(ctx, send)
	}

	ack := transfer.MessageSendACK{
		MessageId:    send.MessageId,
		DeviceId:     send.SenderDeviceId,
		SendSequence: send.SendSequence,
	}
	// 消息发送回执
	err = connect_rpc.ConnectRPC.SendMessageSendACK(ack)
	if err != nil {
		logger.Sugaer.Error(err)
	}

	logger.Sugaer.Infow("消息发送回执",
		"device_id", ack.DeviceId,
		"user_id", send.SenderUserId,
		"message_id", send.MessageId,
		"send_sequence", ack.SendSequence)

	return nil
}

// MessageACK 处理消息回执
func (s *logicRPC) MessageACK(ctx *context.Context, ack transfer.MessageACK) error {
	err := dao.DeviceSyncSequenceDao.UpdateSequence(ctx, ack.DeviceId, ack.SyncSequence)
	if err != nil {
		logger.Sugaer.Error(err)
	}

	logger.Sugaer.Infow("消息投递回执",
		"device_id", ack.DeviceId,
		"user_id", ack.UserId,
		"message_id", ack.MessageId,
		"sync_sequence", ack.SyncSequence)

	return nil
}

// OffLine 处理设备离线
func (s *logicRPC) OffLine(ctx *context.Context, deviceId int64, userId int64) error {
	err := dao.DeviceDao.UpdateStatus(ctx, deviceId, service.DeviceOffline)
	if err != nil {
		logger.Sugaer.Error(err)
	}

	logger.Sugaer.Infow("设备离线", "device_id", deviceId, "user_id", userId)

	return nil
}
