package service

import (
	"goim/logic/dao"
	"goim/logic/entity"
	"goim/public/context"
)

type messageService struct{}

var MessageService = new(messageService)

// Add 添加消息
func (*messageService) Add(ctx *context.Context, message entity.Message) error {
	return dao.MessageDao.Add(ctx, message)
}

// ListByUserIdAndSequence 查询消息
func (*messageService) ListByUserIdAndSequence(ctx *context.Context, userId int64, sequence int) ([]*entity.Message, error) {
	return dao.MessageDao.ListByUserIdAndSequence(ctx, userId, sequence)
}
