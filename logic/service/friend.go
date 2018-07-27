package service

import (
	"fmt"
	"goim/logic/dao"
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type FriendService struct {
	baseService
}

func NewFriendService(session ...*session.Session) *FriendService {
	service := new(FriendService)
	service.setSession(session...)
	return service
}

func (s *FriendService) Add(add entity.FriendAdd) error {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer s.session.Rollback()

	friend1 := entity.Friend{
		UserId: add.UserId,
		Friend: add.Friend,
		Label:  add.UserLable,
	}
	err = dao.NewFriendDao(s.session).Add(friend1)
	if err != nil {
		fmt.Println(err)
		return err
	}

	friend2 := entity.Friend{
		UserId: add.Friend,
		Friend: add.UserId,
		Label:  add.FriendLable,
	}
	err = dao.NewFriendDao(s.session).Add(friend2)
	if err != nil {
		fmt.Println(err)
		return err
	}

	s.session.Commit()
	return nil
}
