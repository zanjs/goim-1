package connect

import (
	"fmt"
	"goim/public/pb"
	"goim/public/transfer"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
)

// HandleMessage 处理消息投递
func HandleMessage(message transfer.Message) {
	ctx := load(message.DeviceId)
	if ctx == nil {
		log.Println("ctx id nil")
		return
	}

	messages := make([]*pb.MessageItem, 0, len(message.Messages))
	for _, v := range message.Messages {
		item := new(pb.MessageItem)
		item.SenderType = int32(v.SenderType)
		item.SenderId = v.SenderId
		item.DeviceId = v.DeviceId
		item.ReceiverType = int32(v.ReceiverType)
		item.ReceiverId = v.ReceiverId
		item.Type = int32(v.Type)
		item.Content = v.Content
		item.Sequence = v.Sequence
		messages = append(messages, item)
	}

	content, err := proto.Marshal(&pb.Message{Messages: messages})
	if err != nil {
		log.Println(err)
		return
	}

	err = ctx.Codec.Eecode(Package{Code: CodeMessage, Content: content}, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
}

// HandleMessageSendACK 处理消息发送回执
func HandleMessageSendACK(ack transfer.MessageSendACK) {
	content, err := proto.Marshal(&pb.MessageSendACK{ack.SendSequence})
	if err != nil {
		log.Println(err)
		return
	}
	ctx := load(ack.DeviceId)
	if ctx == nil {
		log.Println("ctx id nil")
		return
	}
	err = ctx.Codec.Eecode(Package{Code: CodeMessageSendACK, Content: content}, 10*time.Second)
	if err != nil {
		log.Println(err)
		return
	}
}
