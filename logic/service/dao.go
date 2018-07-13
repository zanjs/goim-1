package service

import (
	"goim/logic/dao"
	"goim/logic/lib/session"
)

func userDao(sessioner session.Sessioner) *dao.UserDao {
	return &dao.UserDao{sessioner}
}
