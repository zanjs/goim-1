package service

import (
	"goim/logic/dao"
	"goim/logic/lib/session"
	"goim/pb"
)

type HandlerService struct {
	baseService
}

func NewHandleService(session ...*session.Session) *HandlerService {
	service := new(HandlerService)
	service.setSession(session...)
	return service
}

// HandleOnline 处理设备上线
func (s *HandlerService) HandleOnline(online pb.OnLine) error {
	dao.NewDeviceDao(s.session).GetToken(online.DeviceId)
}

// HandleOnline 处理设备上线
func (s *HandlerService) HandleMessageSend(messageSend pb.MessageSend) {

}

// HandleOnline 处理设备上线
func (s *HandlerService) HandleMessageAck(messageSend pb.MessageACK) {

}
