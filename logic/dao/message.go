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
	_, err := ctx.Session.Exec("insert into t_message(user_id,sender_type,sender,recever_type,recerver,type.content,seq) values(?,?,?,?,?,?,?)",
		message.UserId, message.SenderType, message.Sender, message.ReceiverType,
		message.Receiver, message.Type, message.Content, message.Seq)
	if err != nil {
		log.Println(err)
	}
	return err
}

// List 根据用户id查询大于序号大于seq的消息
func (*messageDao) List(ctx *context.Context, userId int, seq int) ([]*entity.Message, error) {
	rows, err := ctx.Session.Query("select id,user_id,sender_type,sender,recever_type,recerver,type.content,seq,create_time from t_message where user_id = ? and sync_seq > ?")
	if err != nil {
		log.Println(err)
	}

	messages := make([]*entity.Message, 0, 5)
	for rows.Next() {
		message := new(entity.Message)
		err := rows.Scan(&message.Id, &message.UserId, &message.SenderType, &message.Sender, &message.ReceiverType,
			&message.Receiver, &message.Type, &message.Content, &message.Seq, &message.CreateTime)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
