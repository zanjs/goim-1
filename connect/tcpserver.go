package connect

import (
	"fmt"
	"net"
	"log"
)

// Conf server配置文件
type Conf struct {
	Port          string // 端口
	HeadbeatTime  int    // 心跳时间
	HeadbeatCount int    // 心跳次数
	MaxConnCount  int    // 最大连接数
	AcceptCount   int    // 接收建立连接的groutine数量
	TypeLen       int    // 消息type字节数组的长度
	LenLen        int    // 消息length字节数组的长度
	BufferLen     int    // buffer大小,建议不小于最大消息体的字节长度
}

// TCPServer TCP服务器
type TCPServer struct {
	Handler Handler // 回调处理接口
	Conf    Conf    // 配置
}

// NewTCPServer 创建TCP服务器
func NewTCPServer() *TCPServer {
	return new(TCPServer)
}

// Start 启动服务器
func (t *TCPServer) Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("error listening", err.Error())
		return
	}
	for i := 0; i < t.Conf.AcceptCount; i++ {
		go t.Accept(listener)
	}
}

// Accept 接收客户端的TCP长连接的建立
func (t *TCPServer) Accept(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go t.DoConn(conn)
	}
}

// DoConn 处理连接请求
func (t *TCPServer) DoConn(conn net.Conn) {
	codec := NewCodec(conn, t.Conf.BufferLen, t.Conf.TypeLen, t.Conf.LenLen)

	connContext := &ConnContext{Conn: conn}
	t.Handler.OnConnect(connContext)
	for {
		_, err := codec.Read()
		if err != nil {
			return
		}
		for {
			message, ok := codec.Decode()
			if ok {
				fmt.Println(message)
				continue
			}
			break
		}
	}
}
