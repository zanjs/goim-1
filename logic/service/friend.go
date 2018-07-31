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

func (s *FriendService) ListUserFriend(userId int) ([]entity.FriendUser, error) {
	users, err := dao.NewFriendDao(s.session).ListUserFriend(userId)
	if err != nil {
		log.Println(err)
	}
	return users, err

}

// Add 添加好友关系
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
		Label:  add.UserLabel,
	}
	err = dao.NewFriendDao(s.session).Add(friend1)
	if err != nil {
		fmt.Println(err)
		return err
	}

	friend2 := entity.Friend{
		UserId: add.Friend,
		Friend: add.UserId,
		Label:  add.FriendLabel,
	}
	err = dao.NewFriendDao(s.session).Add(friend2)
	if err != nil {
		fmt.Println(err)
		return err
	}

	s.session.Commit()
	return nil
}

// Delete 删除好友关系
func (s *FriendService) Delete(userId, friend int) error {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
	}
	defer s.session.Rollback()

	err = dao.NewFriendDao(s.session).Delete(userId, friend)
	if err != nil {
		log.Println(err)
		return err
	}

	err = dao.NewFriendDao(s.session).Delete(friend, userId)
	if err != nil {
		log.Println(err)
		return err
	}

	s.session.Commit()
	return nil
}
