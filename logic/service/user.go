package service

import (
	"goim/logic/lib/session"
)

type UserService struct {
	baseService
}

func NewUserService(session ...session.Session) *UserService {
	service := new(UserService)
	service.setSession(session...)
	return service
}

func Regist() {

}
