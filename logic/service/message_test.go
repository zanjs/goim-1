package service

import (
	"fmt"
	"goim/logic/model"
	"goim/public/logger"
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
	logger.Sugaer.Error(err)
}

func TestMessageService_ListByUserIdAndSequence(t *testing.T) {
	messages, err := MessageService.ListByUserIdAndSequence(ctx, 1, 0)
	if err != nil {
		logger.Sugaer.Error(err)
		return
	}
	for _, message := range messages {
		fmt.Println(message)
		fmt.Println(lib.FormatTime(message.CreateTime))
	}
}
