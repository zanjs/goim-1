package dao

import (
	"goim/logic/entity"
	"goim/public/context"
	"log"
)

type messageDao struct{}

var MessageDao = new(messageDao)

// Add 插入一条消息
func (*messageDao) Add(ctx *context.Context, message entity.Message) error {
	_, err := ctx.Session.Exec("insert into t_message(user_id,sender_type,sender_id,device_id,receiver_type,receiver_id,type,content,sequence) values(?,?,?,?,?,?,?,?,?)",
		message.UserId, message.SenderType, message.SenderId, message.DeviceId, message.ReceiverType,
		message.ReceiverId, message.Type, message.Content, message.Sequence)
	if err != nil {
		log.Println(err)
	}
	return err
}

// ListByUserIdAndSequence 根据用户id查询大于序号大于sequence的消息
func (*messageDao) ListByUserIdAndSequence(ctx *context.Context, userId int64, sequence int) ([]*entity.Message, error) {
	rows, err := ctx.Session.Query("select id,user_id,sender_type,sender_id,device_id,receiver_type,receiver_id,type,content,sequence,create_time from t_message where user_id = ? and sequence >= ?",
		userId, sequence)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	messages := make([]*entity.Message, 0, 5)
	for rows.Next() {
		message := new(entity.Message)
		err := rows.Scan(&message.Id, &message.UserId, &message.SenderType, &message.SenderId, &message.DeviceId, &message.ReceiverType,
			&message.ReceiverId, &message.Type, &message.Content, &message.Sequence, &message.CreateTime)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
