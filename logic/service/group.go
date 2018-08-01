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

// ListByUserId 获取用户群组
func (s *GroupService) ListByUserId(userId int) ([]*entity.Group, error) {
	ids, err := dao.NewGroupUserDao(s.session).ListbyUserId(userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	groups := make([]*entity.Group, 0, 5)
	for i := range ids {
		group, err := NewGroupService(s.session).Get(ids[i])
		if err != nil {
			log.Println(err)
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// ListGroupUser 获取群组的用户信息
func (s *GroupService) Get(id int) (*entity.Group, error) {
	group, err := dao.NewGroupUserDao(s.session).Get(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	group.GroupUser, err = dao.NewGroupUserDao(s.session).ListGroupUser(id)
	if err != nil {
		log.Println(err)
	}
	return group, err
}

// CreateAndAddUser 创建群组并且添加群成员
func (s *GroupService) CreateAndAddUser(add entity.GroupAdd) (int, error) {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer s.session.Rollback()

	id, err := dao.NewGroupDao(s.session).Add(add.Name)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	for i := range add.UserIds {
		err := dao.NewGroupUserDao(s.session).Add(id, add.UserIds[i])
		if err != nil {
			log.Println(err)
			return 0, err
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

func (s *GroupService) UpdateLabel(groupId int, userId int, label string) error {
	return dao.NewGroupUserDao(s.session).UpdateLabel(groupId, userId, label)
}
