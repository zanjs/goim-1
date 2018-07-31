package service

import (
	"goim/logic/dao"
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type GroupService struct {
	baseService
}

func NewGroupService(session ...*session.Session) *GroupService {
	service := new(GroupService)
	service.setSession(session...)
	return service
}

// ListGroupUser 获取群组的用户信息
func (s *GroupService) ListGroupUser(id int) ([]entity.User, error) {
	users, err := dao.NewGroupUserDao(s.session).ListGroupUser(id)
	if err != nil {
		log.Println(err)
	}
	return users, err
}

// CreateAndAddUser 创建群组并且添加群成员
func (s *GroupService) CreateAndAddUser(add entity.GroupAdd) (int, error) {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	defer s.session.Rollback()

	id, err := dao.NewGroupDao(s.session).Add(add.Name)
	if err != nil {
		log.Println(err)
		return 0, nil
	}

	for i := range add.UserIds {
		err := dao.NewGroupUserDao(s.session).Add(id, add.UserIds[i])
		if err != nil {
			log.Println(err)
			return 0, nil
		}
	}
	s.session.Commit()
	return id, nil
}

// AddUser 给群组添加用户
func (s *GroupService) AddUser(update entity.GroupUserUpdate) error {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer s.session.Rollback()

	for i := range update.UserIds {
		err := dao.NewGroupUserDao(s.session).Add(update.GroupId, update.UserIds[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	s.session.Commit()
	return nil
}

// DeleteUser 从群组移除用户
func (s *GroupService) DeleteUser(update entity.GroupUserUpdate) error {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer s.session.Rollback()

	for i := range update.UserIds {
		err := dao.NewGroupUserDao(s.session).Delete(update.GroupId, update.UserIds[i])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	s.session.Commit()
	return nil
}
