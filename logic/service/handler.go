package service

import "goim/public/model"

type handlerService struct{}

var HandlerService = new(handlerService)

// HandleSignIn 处理设备登录
func (s *handlerService) HandleSignIn(signIn model.SignIn) (*model.SignInACK, error) {

	return nil, nil
}

// HandleSyncTrigger 处理消息同步触发
func (s *handlerService) HandleSyncTrigger(trigger model.SyncTrigger) error {
	return nil
}

// HandleMessageSend 处理消息发送
func (s *handlerService) HandleMessageSend(send model.MessageSend) {

}

// HandleMessageACK 处理消息回执
func (s *handlerService) HandleMessageACK(message model.MessageACK) {

}
