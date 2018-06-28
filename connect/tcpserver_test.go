package connect

import (
	"fmt"
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
	server := NewTCPServer(conf, &handler{})
	server.Start()
}

type handler struct {
}

// OnConnect 当连接建立的时候调用
func (*handler) OnConnect(ctx *ConnContext) {
	fmt.Println("connect")
}

// OnMessage 当发送消息的时候调用
func (*handler) OnMessage(ctx *ConnContext, message *Message) {
	fmt.Println(message.Code, string(message.Content))
}

// OnClose 当连接关闭的时候调用
func (*handler) OnClose(*ConnContext) {
	fmt.Println("close")
}

// OnActive 监听到客户端活动
func (*handler) OnActive(*ConnContext) {
	fmt.Println("active")
}

// OnInactive 监听到客户端停止活动
func (*handler) OnInactive(*ConnContext) {
	fmt.Println("inactive")
}

// OnError 当发生错误的时候调用
func (*handler) OnError(*ConnContext, error) {
	fmt.Println("inactive")
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	codec := NewCodec(conn)

	codec.Eecode(Message{4, []byte("hello world")})
	codec.Eecode(Message{1, []byte("ok fuck")})
	conn.Close()

}

func TestClientTimeOut(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	codec := NewCodec(conn)

	codec.Eecode(Message{4, []byte("hello world")})
	codec.Eecode(Message{1, []byte("ok fuck")})

	_, err = codec.Read()
	if err != nil {
		log.Println(err)
	}

}
