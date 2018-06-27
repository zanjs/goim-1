package connect

import (
	"net"
)

// ConnContext 连接上下文
type ConnContext struct {
	Conn net.Conn
	Info interface{} // 附加信息
}

// Message 消息
type Message struct {
	Code    int    // 消息类型
	Content []byte // 消息体
}

// Handler 事件处理
type Handler interface {
	// OnConnect 当连接建立的时候调用
	OnConnect(*ConnContext)

	// OnMessage 当发送消息的时候调用
	OnMessage(*ConnContext, *Message)

	// OnClose 当连接关闭的时候调用
	OnClose(*ConnContext)

	// OnActive 监听到客户端活动
	OnActive(*ConnContext)

	// OnInactive 监听到客户端停止活动
	OnInactive(*ConnContext)

	// OnError 当发生错误的时候调用
	OnError(*ConnContext, error)
}
