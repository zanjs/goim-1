package service

import (
	"goim/logic/dao"
	"goim/logic/entity"
	"goim/logic/lib/session"
	"log"

	"github.com/satori/go.uuid"
)

type DeviceService struct {
	baseService
}

func NewDeviceService(session ...session.Session) *DeviceService {
	service := new(DeviceService)
	service.setSession(session...)
	return service
}

func (s *DeviceService) Regist(device entity.Device) (int, string, error) {
	UUID, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	token := UUID.String()
	device.Token = token
	id, err := dao.NewDeviceDao(s.session).Insert(device)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	return id, token, nil
}
