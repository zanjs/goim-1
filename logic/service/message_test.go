package service

import (
	"fmt"
	"goim/logic/model"
	"log"
	"step-wx/lib"
	"testing"
)

func TestMessageService_Add(t *testing.T) {
	message := model.Message{
		UserId:         1,
		SenderType:     1,
		SenderId:       1,
		SenderDeviceId: 1,
		ReceiverType:   1,
		ReceiverId:     1,
		Type:           1,
		Content:        "1",
		Sequence:       1,
	}
	err := MessageService.Add(ctx, message)
	fmt.Println(err)
}

func TestMessageService_ListByUserIdAndSequence(t *testing.T) {
	messages, err := MessageService.ListByUserIdAndSequence(ctx, 1, 0)
	if err != nil {
		log.Println(err)
		return
	}
	for _, message := range messages {
		fmt.Printf("%#v\n", message)
		fmt.Println(lib.FormatTime(message.CreateTime))
	}
}
