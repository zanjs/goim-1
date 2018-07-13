package service

import (
	"goim/logic/entity"
	"goim/logic/lib/session"
)

type UserService struct {
	session.Sessioner
}

func NewUserService(sessioner ...session.Sessioner) *UserService {
	service := new(UserService)
	SetSession(&service.Sessioner, sessioner...)
	return service
}

func Regist() {

}

func (s UserService) Get(id int) (*entity.User, error) {
	return userDao(s).Get(id)
}
