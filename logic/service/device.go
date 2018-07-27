package service

import (
	"goim/logic/dao"
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"

	"github.com/satori/go.uuid"
)

const (
	DeviceOnline  = 1
	DeviceOffline = 0
)

type DeviceService struct {
	baseService
}

func NewDeviceService(session ...*session.Session) *DeviceService {
	service := new(DeviceService)
	service.setSession(session...)
	return service
}

// Regist 注册设备
func (s *DeviceService) Regist(device entity.Device) (int, string, error) {
	err := s.session.Begin()
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	defer s.session.Rollback()

	UUID, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	device.Token = UUID.String()
	id, err := dao.NewDeviceDao(s.session).Add(device)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	err = dao.NewDeviceSeqDao(s.session).Add(id)
	if err != nil {
		log.Println(err)

	}

	return id, device.Token, nil
}
