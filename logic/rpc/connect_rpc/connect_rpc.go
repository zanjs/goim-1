package connect_rpc

import (
	"goim/public/transfer"
)

type ConnectRPCer interface {
	SendMessage(message transfer.Message)
	SendMessageSendACK(ack transfer.MessageSendACK)
}

var ConnectRPC ConnectRPCer
