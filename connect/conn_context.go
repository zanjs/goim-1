package connect

import (
	"fmt"
	"goim/public/pb"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
)

const ReadDeadline = 10 * time.Minute

// 消息协议
const (
	CodeSignIn         = 1 // 设备登录
	CodeSignInACK      = 2 // 设备登录回执
	CodeSyncTrigger    = 3 // 消息同步触发
	CodeHeadbeat       = 4 // 心跳
	CodeHeadbeatACK    = 5 // 心跳回执
	CodeMessageSend    = 5 // 消息发送
	CodeMessageSendACK = 6 // 消息发送回执
	CodeMessage        = 7 // 消息投递
	CodeMessageACK     = 8 // 消息投递回执
)

// ConnContext 连接上下文
type ConnContext struct {
	Codec *Codec      // 编解码器
	Info  interface{} // 附加信息
}

// Package 消息包
type Package struct {
	Code    int    // 消息类型
	Content []byte // 消息体
}

func NewConnContext(conn *net.TCPConn) *ConnContext {
	codec := NewCodec(conn)
	return &ConnContext{Codec: codec}
}

// DoConn 处理TCP连接
func (c *ConnContext) DoConn() {
	defer RecoverPanic()

	c.HandleConnect()

	for {
		err := c.Codec.Conn.SetReadDeadline(time.Now().Add(ReadDeadline))
		if err != nil {
			log.Println(err)
			return
		}

		_, err = c.Codec.Read()
		if err != nil {
			c.HandleReadErr(err)
			return
		}

		for {
			message, ok := c.Codec.Decode()
			if ok {
				c.HandleMessage(message)
				continue
			}
			break
		}
	}
}

// HandleConnect 建立连接
func (c *ConnContext) HandleConnect() {
	log.Println("connect")
	return
}

// HandleMessage 处理消息
func (c *ConnContext) HandleMessage(pack *Package) {
	log.Println("message", pack.Code, string(pack.Content))
	switch pack.Code {
	case CodeSignIn:
		var signIn pb.SignIn
		err := proto.Unmarshal(pack.Content, &signIn)
		if err != nil {
			fmt.Println(err)
			c.Close()
			return
		}

	case CodeSyncTrigger:
		var trigger pb.SyncTrigger
		err := proto.Unmarshal(pack.Content, &trigger)
		if err != nil {
			fmt.Println(err)
			c.Close()
			return
		}

	case CodeHeadbeat:
		// 心跳逻辑

	case CodeMessageSend:
		var messageSend pb.MessageSend
		err := proto.Unmarshal(pack.Content, &messageSend)
		if err != nil {
			fmt.Println(err)
			c.Close()
			return
		}
	case CodeMessageACK:
		var messageACK pb.MessageACK
		err := proto.Unmarshal(pack.Content, &messageACK)
		if err != nil {
			fmt.Println(err)
			c.Close()
			return
		}
	}
	return
}

// HandleReadErr 读取conn错误
func (c *ConnContext) HandleReadErr(err error) {
	log.Println(err)
	// 客户端主动关闭连接或者异常程序退出
	if err == io.EOF {
		c.Codec.Conn.Close()
		return
	}
	str := err.Error()
	// SetReadDeadline 之后，超时返回的错误
	if strings.HasSuffix(str, "i/o timeout") {
		c.Codec.Conn.Close()
		return
	}
	// 服务器主动关闭连接
	if strings.HasSuffix(str, "use of closed network connection") {
		return
	}
}

// Close 关闭TCP连接
func (c *ConnContext) Close() {
	log.Println("close")
	c.Codec.Conn.Close()
}
