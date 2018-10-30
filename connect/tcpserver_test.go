package connect

import (
	"goim/public/logger"
	"net"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	conf := Conf{
		Address:      "localhost:50002",
		MaxConnCount: 100,
		AcceptCount:  1,
	}
	server := NewTCPServer(conf)
	server.Start()
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	codec := NewCodec(conn)

	codec.Eecode(Package{4, []byte("hello world")}, 2*time.Second)
	codec.Eecode(Package{1, []byte("ok fuck")}, 2*time.Second)
	conn.Close()

}

func TestClientTimeOut(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:50002")
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	codec := NewCodec(conn)

	codec.Eecode(Package{4, []byte("hello world")}, 2*time.Second)
	codec.Eecode(Package{1, []byte("ok fuck")}, 2*time.Second)

	_, err = codec.Read()
	if err != nil {
		logger.Sugar.Error(err)
	}

}
