package dao

import (
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type MessageDao struct {
	base
}

func NewMessageDao(session *session.Session) *MessageDao {
	return &MessageDao{base{session}}
}

// Insert 插入一条消息
func (d *MessageDao) Insert(message entity.Message) error {
	_, err := d.session.Exec("insert into t_message(user_id,sender_type,sender,recever_type,recerver,type.content,seq) values(?,?,?,?,?,?,?)",
		message.UserId, message.SenderType, message.Sender, message.ReceiverType,
		message.Receiver, message.Type, message.Content, message.Seq)
	if err != nil {
		log.Println(err)
	}
	return err
}

// List 根据用户id查询大于序号大于seq的消息
func (d *MessageDao) List(userId int, seq int) ([]*entity.Message, error) {
	rows, err := d.session.Query("select id,user_id,sender_type,sender,recever_type,recerver,type.content,seq,create_time from t_message where user_id = ? and sync_seq > ?")
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
