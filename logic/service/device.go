package service

import (
	"goim/logic/dao"
	"goim/logic/entity"
	"log"

	"goim/lib/context"

	"github.com/satori/go.uuid"
)

const (
	DeviceOnline  = 1
	DeviceOffline = 0
)

type deviceService struct{}

var DeviceService = new(deviceService)

// Regist 注册设备
func (*deviceService) Regist(ctx *context.Context, device entity.Device) (int64, string, error) {
	err := ctx.Session.Begin()
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	defer ctx.Session.Rollback()

	UUID, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	device.Token = UUID.String()
	id, err := dao.DeviceDao.Add(ctx, device)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}

	err = dao.DeviceSequenceDao.Add(ctx, id)
	if err != nil {
		log.Println(err)

	}
	return id, device.Token, nil
}
