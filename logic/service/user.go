package service

import (
	"database/sql"
	"errors"
	"goim/logic/dao"
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"
)

type UserService struct {
	baseService
}

func NewUserService(session ...*session.Session) *UserService {
	service := new(UserService)
	service.setSession(session...)
	return service
}

var ErrNumberExist = errors.New("user number exist")

func (s *UserService) Regist(user entity.User) (int, error) {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer s.session.Rollback()

	id, err := dao.NewUserDao(s.session).Add(user)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if id == 0 {
		return 0, ErrNumberExist
	}

	err = dao.NewUserSeqDao(s.session).Add(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	s.session.Commit()
	return id, nil
}

var (
	ErrDeviceNotFound = errors.New("device not found")
	ErrToken          = errors.New("error token")
	ErrUserNotFound   = errors.New("user not found")
	ErrPassword       = errors.New("error password")
)

// SignIn 登录
func (s *UserService) SignIn(signIn entity.SignIn) error {
	token, err := dao.NewDeviceDao(s.session).GetToken(signIn.DeviceId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrDeviceNotFound
		}
		log.Println(err)
		return err
	}

	if signIn.Token != token {
		return ErrToken
	}

	password, err := dao.NewUserDao(s.session).GetPassword(signIn.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		log.Println(err)
		return err
	}

	if signIn.Password != password {
		return ErrPassword
	}

	err = dao.NewDeviceDao(s.session).UpdateUserId(signIn.DeviceId, signIn.UserId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
