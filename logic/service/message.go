package service

import (
	"goim/connect"
	"goim/logic/dao"
	"goim/logic/model"
	"goim/public/context"
	"goim/public/transfer"
	"log"
)

const (
	ReceiverUser  = 1 // 接收者类型为用户
	ReceiverGroup = 2 // 接收者类型为群组
)

const (
	SendTypeUser = 1 // 用户发送
	SendTypOther = 2 // 其他发送，业务推送
)

type messageService struct{}

var MessageService = new(messageService)

// Add 添加消息
func (*messageService) Add(ctx *context.Context, message model.Message) error {
	return dao.MessageDao.Add(ctx, message)
}

// ListByUserIdAndSequence 查询消息
func (*messageService) ListByUserIdAndSequence(ctx *context.Context, userId int64, sequence int64) ([]*model.Message, error) {
	return dao.MessageDao.ListByUserIdAndSequence(ctx, userId, sequence)
}

// SendToUser 消息发送至用户
func (*messageService) SendToFriend(ctx *context.Context, send transfer.MessageSend) error {
	selfSequence, err := UserRequenceService.GetNext(ctx, send.UserId)
	if err != nil {
		log.Println(err)
		return err
	}
	selfMessage := model.Message{
		UserId:       send.UserId,
		SenderType:   SendTypeUser,
		SenderId:     send.UserId,
		DeviceId:     send.DeviceId,
		ReceiverType: int(send.ReceiverType),
		ReceiverId:   send.ReceiverId,
		Type:         int(send.Type),
		Content:      send.Content,
		Sequence:     selfSequence,
	}

	// 发给发送者
	err = MessageService.SendToUser(ctx, send.UserId, &selfMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	friendSequence, err := UserRequenceService.GetNext(ctx, send.ReceiverId)
	if err != nil {
		log.Println(err)
		return err
	}
	friendMessage := model.Message{
		UserId:       send.ReceiverId,
		SenderType:   SendTypeUser,
		SenderId:     send.UserId,
		DeviceId:     send.DeviceId,
		ReceiverType: int(send.ReceiverType),
		ReceiverId:   send.ReceiverId,
		Type:         int(send.Type),
		Content:      send.Content,
		Sequence:     friendSequence,
	}
	// 发给接收者
	err = MessageService.SendToUser(ctx, send.ReceiverId, &friendMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// SendToGroup 消息发送至群组
func (*messageService) SendToGroup(ctx *context.Context, send transfer.MessageSend) error {
	group, err := GroupService.Get(ctx, send.ReceiverId)
	if err != nil {
		log.Println(err)
		return err
	}

	// 持久化到数据库
	for _, user := range group.GroupUser {
		sequence, err := UserRequenceService.GetNext(ctx, send.UserId)
		if err != nil {
			log.Println(err)
			return err
		}
		message := model.Message{
			UserId:       user.UserId,
			SenderType:   SendTypeUser,
			SenderId:     send.UserId,
			DeviceId:     send.DeviceId,
			ReceiverType: int(send.ReceiverType),
			ReceiverId:   send.ReceiverId,
			Type:         int(send.Type),
			Content:      send.Content,
			Sequence:     sequence,
		}

		err = MessageService.SendToUser(ctx, user.UserId, &message)
		if err != nil {
			log.Println(err)
			return err
		}

	}
	return nil
}

// SendToUser 消息发送至用户
func (*messageService) SendToUser(ctx *context.Context, userId int64, message *model.Message) error {
	err := MessageService.Add(ctx, *message)
	if err != nil {
		log.Println(err)
		return err
	}

	selfItem := transfer.MessageItem{
		SenderType:   message.SenderType,
		SenderId:     message.SenderId,
		DeviceId:     message.DeviceId,
		ReceiverType: message.ReceiverType,
		ReceiverId:   message.ReceiverId,
		Type:         message.Type,
		Content:      message.Content,
		Sequence:     message.Sequence,
	}

	// 查询用户在线设备
	devices, err := dao.DeviceDao.ListOnlineByUserId(ctx, userId)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, v := range devices {
		message := transfer.Message{DeviceId: v.Id, Messages: []transfer.MessageItem{selfItem}}
		connect.HandleMessage(message)
	}
	return nil
}
