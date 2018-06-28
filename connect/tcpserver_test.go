package connect

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	conf := Conf{
		Address:      "localhost:50002",
		ReadDeadline: 10 * time.Minute,
		MaxConnCount: 100,
		AcceptCount:  1,
		TypeLen:      2,
		LenLen:       2,
		BufferLen:    100,
	}
	server := NewTCPServer(conf, &handler{})
	server.Start()
}

type handler struct {
}

// OnConnect 当连接建立的时候调用
func (*handler) OnConnect(*ConnContext) {

}

// OnMessage 当发送消息的时候调用
func (*handler) OnMessage(*ConnContext, *Message) {

}

// OnClose 当连接关闭的时候调用
func (*handler) OnClose(*ConnContext) {

}

// OnActive 监听到客户端活动
func (*handler) OnActive(*ConnContext) {

}

// OnInactive 监听到客户端停止活动
func (*handler) OnInactive(*ConnContext) {

}

// OnError 当发生错误的时候调用
func (*handler) OnError(*ConnContext, error) {

}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	var headSize int
	var headBytes = make([]byte, 2)
	s := "hello world"
	content := []byte(s)
	headSize = len(content)
	binary.BigEndian.PutUint16(headBytes, uint16(headSize))
	conn.Write(headBytes)
	conn.Write(content)

	s = "hello go"
	content = []byte(s)
	headSize = len(content)
	binary.BigEndian.PutUint16(headBytes, uint16(headSize))
	conn.Write(headBytes)
	conn.Write(content)

	s = "hello tcp"
	content = []byte(s)
	headSize = len(content)
	binary.BigEndian.PutUint16(headBytes, uint16(headSize))
	conn.Write(headBytes)
	conn.Write(content)
}
