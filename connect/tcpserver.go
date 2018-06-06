package connect

import (
	"fmt"
	"net"
	"log"
	"encoding/binary"
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
	var (
		buffer     = newBuffer(conn, t.Conf.BufferLen) // 接收消息的buffer
		typeBuf    []byte                              // 数据类型存储字节
		lenBuf     []byte                              // 数据库所占字节数存储字节
		valueType  int                                 // 数据类型
		valueLen   int                                 // 数据内容长度
		valueBuf   []byte                              // 数据内容buffer
		typeLenLen = t.Conf.TypeLen + t.Conf.LenLen    // 数据类型+数据长度 长度
	)
	connContext := &ConnContext{Conn: conn}
	t.Handler.OnConnect(connContext)
	for {
		_, err := buffer.readFromReader()
		if err != nil {
			log.Println(err)
			t.Handler.OnError(connContext, err)
			return
		}
		for {
			// 读取数据类型
			typeBuf, err = buffer.seek(0, t.Conf.TypeLen)
			if err != nil {
				fmt.Println(err)
				break
			}

			// 读取数据长度
			lenBuf, err = buffer.seek(t.Conf.TypeLen, t.Conf.LenLen)
			if err != nil {
				fmt.Println(err)
				break
			}

			// 读取数据内容
			valueType = int(binary.BigEndian.Uint16(typeBuf))
			valueLen = int(binary.BigEndian.Uint16(lenBuf))

			valueBuf,err = buffer.read(typeLenLen, valueLen)
			if err!=nil{
				fmt.Println(valueType)
				fmt.Println(valueLen)
				fmt.Println(string(valueBuf))
				t.Handler.OnMessage(connContext,Message{Code:valueType,Content:valueBuf})
				continue
			}
			break
		}
	}
}
