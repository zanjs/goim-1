package connect

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	conf := Conf{
		Address:      "localhost:50002",
		ReadDeadline: 10 * time.Second,
		MaxConnCount: 100,
		AcceptCount:  1,
	}
	server := NewTCPServer(conf)
	server.Start()
}

type handler struct {
}

// OnConnect 当连接建立的时候调用
func (*handler) OnConnect(ctx *ConnContext) {
	log.Println("connect")
}

// OnMessage 当发送消息的时候调用
func (*handler) OnMessage(ctx *ConnContext, message *Message) {
	log.Println(message.Code, string(message.Content))
}

// OnClose 当连接关闭的时候调用
func (*handler) OnClose(*ConnContext) {
	log.Println("close")
}

// OnActive 监听到客户端活动
func (*handler) OnActive(*ConnContext) {
	log.Println("active")
}

// OnInactive 监听到客户端停止活动
func (*handler) OnInactive(*ConnContext) {
	log.Println("inactive")
}

// OnError 当发生错误的时候调用
func (*handler) OnError(*ConnContext, error) {
	log.Println("inactive")
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		log.Println("Error dialing", err.Error())
		return
	}

	codec := NewCodec(conn)

	codec.Eecode(Message{4, []byte("hello world")}, 2*time.Second)
	codec.Eecode(Message{1, []byte("ok fuck")}, 2*time.Second)
	conn.Close()

}

func TestClientTimeOut(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		log.Println("Error dialing", err.Error())
		return
	}

	codec := NewCodec(conn)

	codec.Eecode(Message{4, []byte("hello world")}, 2*time.Second)
	codec.Eecode(Message{1, []byte("ok fuck")}, 2*time.Second)

	_, err = codec.Read()
	if err != nil {
		log.Println(err)
	}

}
