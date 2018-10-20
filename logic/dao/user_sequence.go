package dao

import (
	"goim/public/context"
	"log"
)

type userSequenceDao struct{}

var UserSequenceDao = new(userSequenceDao)

// Add 添加
func (*userSequenceDao) Add(ctx *context.Context, userId int64, sequence int64) error {
	_, err := ctx.Session.Exec("insert into t_user_sequence (user_id,sequence) values(?,?)", userId, sequence)
	if err != nil {
		log.Println(err)
	}
	return err
}

// Increase sequence++
func (*userSequenceDao) Increase(ctx *context.Context, userId int64) error {
	_, err := ctx.Session.Exec("update t_user_sequence set sequence = sequence + 1 where user_id = ?", userId)
	if err != nil {
		log.Println(err)
	}
	return err
}

// GetSequence 获取自增序列
func (*userSequenceDao) GetSequence(ctx *context.Context, userId int64) (int64, error) {
	var sequence int64
	err := ctx.Session.QueryRow("select sequence from t_user_sequence where user_id = ?", userId).
		Scan(&sequence)
	if err != nil {
		log.Println(err)

	}
	return sequence, err
}
