package connect

import (
	"io"
	"log"
	"net"
	"strings"
	"time"
)

const (
	CodeEOF          = 0 // 客户端主动关闭连接或者异常程序退出
	CodeTimeOut      = 1 // SetReadDeadline 之后，超时返回的错误
	CodeServerClosed = 2 // 服务器主动关闭连接
)

// Conf server配置文件
type Conf struct {
	Address      string        // 端口
	ReadDeadline time.Duration // 读取超时时间，单位为秒
	MaxConnCount int           // 最大连接数
	AcceptCount  int           // 接收建立连接的groutine数量
	TypeLen      int           // 消息type字节数组的长度
	LenLen       int           // 消息length字节数组的长度
	BufferLen    int           // buffer大小,建议不小于最大消息体的字节长度
}

// TCPServer TCP服务器
type TCPServer struct {
	Conf    Conf          // 配置
	Handler Handler       // 回调处理接口
	Factory *CodecFactory // 生成解码器的工厂
}

// NewTCPServer 创建TCP服务器
func NewTCPServer(conf Conf, handler Handler) *TCPServer {
	return &TCPServer{
		Conf:    conf,
		Handler: handler,
		Factory: NewCodecFactory(conf.TypeLen, conf.LenLen, conf.BufferLen),
	}
}

// Start 启动服务器
func (t *TCPServer) Start() {
	addr, err := net.ResolveTCPAddr("tcp", t.Conf.Address)
	if err != nil {
		log.Println(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Println("error listening", err.Error())
		return
	}
	for i := 0; i < t.Conf.AcceptCount; i++ {
		go t.Accept(listener)
	}
}

// Accept 接收客户端的TCP长连接的建立
func (t *TCPServer) Accept(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		go t.DoConn(conn)
	}
}

// DoConn 处理连接请求
func (t *TCPServer) DoConn(conn *net.TCPConn) {

	conn.SetKeepAlive(true)
	codec := t.Factory.NewCodec(conn)

	ctx := &ConnContext{Conn: conn}
	t.Handler.OnConnect(ctx)
	for {
		conn.SetReadDeadline(time.Now().Add(t.Conf.ReadDeadline))
		_, err := codec.Read()
		if err != nil {
			code := ErrCode(err)
			if code == CodeEOF {
				t.Handler.OnClose(ctx)
				conn.Close()
				return
			}
			if code == CodeTimeOut {
				t.Handler.OnInactive(ctx)
				conn.Close()
				return
			}
			if code == CodeServerClosed {
				// 当服务器主动关闭连接的时候，结束掉协程
				return
			}
			t.Handler.OnError(ctx, err)
			break
		}
		for {
			message, ok := codec.Decode()
			if ok {
				t.Handler.OnMessage(ctx, message)
				continue
			}
			break
		}
	}
}

func ErrCode(err error) int {
	// 客户端主动关闭连接或者异常断开
	if err == io.EOF {
		return CodeEOF
	}
	str := err.Error()
	// SetReadDeadline 之后，超时返回的错误
	if strings.HasSuffix(str, "i/o timeout") {
		return CodeTimeOut
	}
	// 服务器主动关闭连接
	if strings.HasSuffix(str, "use of closed network connection") {
		return CodeServerClosed
	}
	return 0
}
