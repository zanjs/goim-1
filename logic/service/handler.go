package service

import (
	"goim/logic/dao"
	"goim/pb"
)

type handlerService struct{}

var HandlerService = new(handlerService)

// HandleOnline 处理设备上线
func (s *handlerService) HandleOnline(online pb.OnLine) error {
	dao.DeviceDao.GetToken(online.DeviceId)
}

// HandleOnline 处理设备上线
func (s *handlerService) HandleMessageSend(messageSend pb.MessageSend) {

}

// HandleOnline 处理设备上线
func (s *handlerService) HandleMessageAck(messageSend pb.MessageACK) {

}
