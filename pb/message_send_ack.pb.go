// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_send_ack.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	message_send_ack.proto

It has these top-level messages:
	MessageSendACK
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var messageSendACK MessageSendACK

// 消息发送回执 type:4
type MessageSendACK struct {
	DeviceId uint64 `protobuf:"varint,2,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Sign     uint32 `protobuf:"varint,3,opt,name=sign" json:"sign,omitempty"`
}

func (m *MessageSendACK) Reset()                    { *m = MessageSendACK{} }
func (m *MessageSendACK) String() string            { return proto.CompactTextString(m) }
func (*MessageSendACK) ProtoMessage()               {}
func (*MessageSendACK) Descriptor() ([]byte, []int) { return messageSendACK.fileDescriptor(), []int{0} }

func (m *MessageSendACK) GetDeviceId() uint64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *MessageSendACK) GetSign() uint32 {
	if m != nil {
		return m.Sign
	}
	return 0
}

func init() {
	proto.RegisterType((*MessageSendACK)(nil), "pb.MessageSendACK")
}

func init() { proto.RegisterFile("message_send_ack.proto", messageSendACK.fileDescriptor()) }

func (MessageSendACK) fileDescriptor() []byte {
	return []byte{
		// 113 bytes of a gzipped FileDescriptorProto
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4d, 0x2d, 0x2e,
		0x4e, 0x4c, 0x4f, 0x8d, 0x2f, 0x4e, 0xcd, 0x4b, 0x89, 0x4f, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca,
		0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe4, 0xe2, 0xf3, 0x85, 0xc8, 0x06, 0xa7, 0xe6,
		0xa5, 0x38, 0x3a, 0x7b, 0x0b, 0x49, 0x73, 0x71, 0xa6, 0xa4, 0x96, 0x65, 0x26, 0xa7, 0xc6, 0x67,
		0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0x71, 0x40, 0x04, 0x3c, 0x53, 0x84, 0x84, 0xb8,
		0x58, 0x8a, 0x33, 0xd3, 0xf3, 0x24, 0x98, 0x15, 0x18, 0x35, 0x78, 0x83, 0xc0, 0xec, 0x24, 0x36,
		0xb0, 0x69, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xfb, 0xb0, 0xf6, 0x67, 0x00, 0x00,
		0x00,
	}
}
