package connect

import (
	"goim/public/context"
	"goim/public/logger"
	"goim/public/pb"
	"goim/public/transfer"
	"time"

	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

// LogicRPCer 逻辑层接口
type LogicRPCer interface {
	// SignIn 设备登录
	SignIn(ctx *context.Context, signIn transfer.SignIn) *transfer.SignInACK
	// SyncTrigger 消息同步触发
	SyncTrigger(ctx *context.Context, trigger transfer.SyncTrigger) error
	// MessageSend 消息发送
	MessageSend(ctx *context.Context, send transfer.MessageSend) error
	// MessageACK 消息投递回执
	MessageACK(ctx *context.Context, ack transfer.MessageACK) error
	// OffLine 下线
	OffLine(ctx *context.Context, deviceId int64) error
}

var LogicRPC LogicRPCer

type connectRPC struct{}

var ConnectRPC = new(connectRPC)

// SendMessage 处理消息投递
func (*connectRPC) SendMessage(message transfer.Message) error {
	ctx := load(message.DeviceId)
	if ctx == nil {
		logger.Sugaer.Error("ctx id nil")
		return nil
	}

	messages := make([]*pb.MessageItem, 0, len(message.Messages))
	for _, v := range message.Messages {
		item := new(pb.MessageItem)
		item.SenderType = int32(v.SenderType)
		item.SenderId = v.SenderId
		item.SenderDeviceId = v.SenderDeviceId
		item.ReceiverType = int32(v.ReceiverType)
		item.ReceiverId = v.ReceiverId
		item.Type = int32(v.Type)
		item.Content = v.Content
		item.Sequence = v.Sequence
		messages = append(messages, item)
	}

	content, err := proto.Marshal(&pb.Message{Messages: messages})
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}

	err = ctx.Codec.Eecode(Package{Code: CodeMessage, Content: content}, 10*time.Second)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	logger.Logger.Debug("TCP消息投递", zap.Reflect("message", message))
	return nil
}

// SendMessageSendACK 处理消息发送回执
func (*connectRPC) SendMessageSendACK(ack transfer.MessageSendACK) error {
	content, err := proto.Marshal(&pb.MessageSendACK{ack.SendSequence})
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	ctx := load(ack.DeviceId)
	if ctx == nil {
		logger.Sugaer.Error(err)
		return err
	}

	err = ctx.Codec.Eecode(Package{Code: CodeMessageSendACK, Content: content}, 10*time.Second)
	if err != nil {
		logger.Sugaer.Error(err)
		return err
	}
	return nil
}
