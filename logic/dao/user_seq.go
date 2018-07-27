package dao

import (
	"goim/logic/lib/session"
	"log"
)

type UserSeqDao struct {
	base
}

func NewUserSeqDao(session *session.Session) *UserSeqDao {
	return &UserSeqDao{base{session}}
}

func (d *UserSeqDao) Add(userId int) error {
	_, err := d.session.Exec("insert into t_user_seq(user_id) values(?)", userId)
	if err != nil {
		log.Println(err)
	}
	return err
}

// GetMaxSeq 获取用户最大的序列号
func (d *UserSeqDao) GetMaxSeq(userId int) (int, error) {
	row := d.session.QueryRow("select max_seq from t_user_seq where user_id = ?", userId)
	var maxSeq int
	err := row.Scan(&maxSeq)
	if err != nil {
		log.Println(err)

	}
	return maxSeq, err
}

// GetMaxSeq 获取用户最大的序列号
func (d *UserSeqDao) UpdateMaxSeq(userId int, maxSeq int) error {
	_, err := d.session.Exec("update t_user set max_seq = ? where user_id = ?", maxSeq, userId)
	if err != nil {
		log.Println(err)
	}
	return err
}

// GetSyncSeq 获取用户已经同步消息的序列号
func (d *UserSeqDao) GetSyncSeq(userId int) (int, error) {
	row := d.session.QueryRow("select sync_seq from t_user_seq where user_id = ?", userId)
	var maxSeq int
	err := row.Scan(&maxSeq)
	if err != nil {
		log.Println(err)

	}
	return maxSeq, err
}

// UpdateSyncSeq 更新用户已经同步消息的序列号
func (d *UserSeqDao) UpdateSyncSeq(userId int, maxSeq int) error {
	_, err := d.session.Exec("update t_user set sync_seq = ? where user_id = ?", maxSeq, userId)
	if err != nil {
		log.Println(err)
	}
	return err
}
