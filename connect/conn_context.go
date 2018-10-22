package connect

import (
	"goim/public/pb"
	"goim/public/transfer"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"fmt"

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
	Codec    *Codec // 编解码器
	DeviceId int64  // 设备id
	UserId   int64  // 用户id
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
				c.HandlePackage(message)
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

// HandlePackage 处理消息包
func (c *ConnContext) HandlePackage(pack *Package) {
	log.Println("package:code", pack.Code)
	switch pack.Code {
	case CodeSignIn:
		c.HandlePackageSignIn(pack)
	case CodeSyncTrigger:
		c.HandlePackageSyncTrigger(pack)
	case CodeHeadbeat:
		c.HandlePackageHeadbeat()
	case CodeMessageSend:
		c.HandlePackageMessageSend(pack)
	case CodeMessageACK:
		c.HandlePackageMessageACK(pack)
	}
	return
}

// HandlePackageSignIn 处理登录消息包
func (c *ConnContext) HandlePackageSignIn(pack *Package) {
	var signIn pb.SignIn
	err := proto.Unmarshal(pack.Content, &signIn)
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}

	// 处理设备登录逻辑
	ack := LogicRPC.SignIn(Context(), transfer.SignIn{
		DeviceId: signIn.DeviceId,
		UserId:   signIn.UserId,
		Token:    signIn.Token,
	})

	content, err := proto.Marshal(&pb.SignInACK{Code: int32(ack.Code), Message: ack.Message})
	if err != nil {
		log.Println(err)
		return
	}

	err = c.Codec.Eecode(Package{Code: CodeSignInACK, Content: content}, 10*time.Second)
	if err != nil {
		log.Println(err)
	}

	if ack.Code == 1 {
		c.DeviceId = signIn.DeviceId
		c.UserId = signIn.UserId
		store(c.DeviceId, c)
	}

}

// HandlePackageSyncTrigger 处理同步触发消息包
func (c *ConnContext) HandlePackageSyncTrigger(pack *Package) {
	var trigger pb.SyncTrigger
	err := proto.Unmarshal(pack.Content, &trigger)
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}
	LogicRPC.SyncTrigger(Context(), transfer.SyncTrigger{DeviceId: c.DeviceId, UserId: c.UserId, SyncSequence: trigger.SyncSequence})
}

// HandlePackageHeadbeat 处理心跳包
func (c *ConnContext) HandlePackageHeadbeat() {
	log.Println("收到心跳")
	err := c.Codec.Eecode(Package{Code: CodeHeadbeatACK, Content: []byte{}}, 10*time.Second)
	if err != nil {
		log.Println(err)
	}
}

// HandlePackageMessageSend 处理消息发送包
func (c *ConnContext) HandlePackageMessageSend(pack *Package) {
	var send pb.MessageSend
	err := proto.Unmarshal(pack.Content, &send)
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}
	fmt.Printf("消息发送:%#v\n", send)
	err = LogicRPC.MessageSend(Context(), transfer.MessageSend{
		SenderDeviceId: c.DeviceId,
		SenderUserId:   c.UserId,
		ReceiverType:   send.ReceiverType,
		ReceiverId:     send.ReceiverId,
		Type:           send.Type,
		Content:        send.Content,
		SendSequence:   send.SendSequence,
	})
	if err != nil {
		log.Println(err)
	}
}

// HandlePackageMessageACK 处理消息回执消息包
func (c *ConnContext) HandlePackageMessageACK(pack *Package) {
	var messageACK pb.MessageACK
	err := proto.Unmarshal(pack.Content, &messageACK)
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}
	LogicRPC.MessageACK(Context(), transfer.MessageACK{
		DeviceId:     c.DeviceId,
		UserId:       c.UserId,
		SyncSequence: messageACK.SyncSequence,
	})
}

// HandleReadErr 读取conn错误
func (c *ConnContext) HandleReadErr(err error) {
	log.Println(err)
	delete(c.DeviceId)
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
