// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
	MessageItem
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

// 消息投递
type Message struct {
	Messages []*MessageItem `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return messageFileDescriptor, []int{0} }

func (m *Message) GetMessages() []*MessageItem {
	if m != nil {
		return m.Messages
	}
	return nil
}

// 单条消息投递
type MessageItem struct {
	SenderType   int32  `protobuf:"varint,1,opt,name=sender_type,json=senderType" json:"sender_type,omitempty"`
	SenderId     int64  `protobuf:"varint,2,opt,name=sender_id,json=senderId" json:"sender_id,omitempty"`
	DeviceId     int64  `protobuf:"varint,3,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	ReceiverType int32  `protobuf:"varint,5,opt,name=receiver_type,json=receiverType" json:"receiver_type,omitempty"`
	ReceiverId   int64  `protobuf:"varint,6,opt,name=receiver_id,json=receiverId" json:"receiver_id,omitempty"`
	Type         int32  `protobuf:"varint,7,opt,name=type" json:"type,omitempty"`
	Content      string `protobuf:"bytes,8,opt,name=content" json:"content,omitempty"`
	Sequence     int64  `protobuf:"varint,9,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *MessageItem) Reset()                    { *m = MessageItem{} }
func (m *MessageItem) String() string            { return proto.CompactTextString(m) }
func (*MessageItem) ProtoMessage()               {}
func (*MessageItem) Descriptor() ([]byte, []int) { return messageFileDescriptor, []int{1} }

func (m *MessageItem) GetSenderType() int32 {
	if m != nil {
		return m.SenderType
	}
	return 0
}

func (m *MessageItem) GetSenderId() int64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *MessageItem) GetDeviceId() int64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *MessageItem) GetReceiverType() int32 {
	if m != nil {
		return m.ReceiverType
	}
	return 0
}

func (m *MessageItem) GetReceiverId() int64 {
	if m != nil {
		return m.ReceiverId
	}
	return 0
}

func (m *MessageItem) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageItem) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MessageItem) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*MessageItem)(nil), "pb.MessageItem")
}

func init() { proto.RegisterFile("message.proto", messageFileDescriptor) }

var messageFileDescriptor = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xc9, 0xd6, 0xdd, 0xb6, 0x53, 0x17, 0x61, 0x4e, 0x41, 0x0f, 0x5b, 0xd6, 0x4b, 0x40,
	0xe8, 0x41, 0xc1, 0x77, 0xc8, 0xc1, 0x4b, 0xf0, 0x2e, 0x36, 0x19, 0x24, 0x87, 0x4d, 0x63, 0x13,
	0x17, 0xf6, 0xcd, 0x3d, 0xca, 0x26, 0x4d, 0xf1, 0x36, 0xf3, 0xfd, 0xff, 0x7c, 0x84, 0xc0, 0xfe,
	0x44, 0x21, 0x7c, 0x7e, 0xd1, 0xe0, 0xe7, 0x29, 0x4e, 0xb8, 0xf1, 0xe3, 0xf1, 0x15, 0xea, 0xb7,
	0x0c, 0xf1, 0x09, 0x9a, 0x65, 0x0c, 0x9c, 0xf5, 0x95, 0xe8, 0x9e, 0xef, 0x06, 0x3f, 0x0e, 0x0b,
	0x93, 0x91, 0x4e, 0x6a, 0x2d, 0x1c, 0x7f, 0x19, 0x74, 0xff, 0x12, 0x3c, 0x40, 0x17, 0xc8, 0x19,
	0x9a, 0x3f, 0xe2, 0xc5, 0x13, 0x67, 0x3d, 0x13, 0x5b, 0x05, 0x19, 0xbd, 0x5f, 0x3c, 0xe1, 0x03,
	0xb4, 0x4b, 0xc1, 0x1a, 0xbe, 0xe9, 0x99, 0xa8, 0x54, 0x93, 0x81, 0x34, 0xd7, 0xd0, 0xd0, 0xd9,
	0x6a, 0xba, 0x86, 0x55, 0x0e, 0x33, 0x90, 0x06, 0x1f, 0x61, 0x3f, 0x93, 0x26, 0x7b, 0x2e, 0xf2,
	0x6d, 0x92, 0xdf, 0x16, 0x98, 0xf4, 0x07, 0xe8, 0xd6, 0x92, 0x35, 0x7c, 0x97, 0x1c, 0x50, 0x90,
	0x34, 0x88, 0x70, 0x93, 0x8e, 0xeb, 0x74, 0x9c, 0x66, 0xe4, 0x50, 0xeb, 0xc9, 0x45, 0x72, 0x91,
	0x37, 0x3d, 0x13, 0xad, 0x2a, 0x2b, 0xde, 0x43, 0x13, 0xe8, 0xfb, 0x87, 0x9c, 0x26, 0xde, 0x96,
	0xc7, 0xe6, 0x7d, 0xdc, 0xa5, 0xdf, 0x7b, 0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x04, 0xed, 0xed,
	0x6a, 0x4e, 0x01, 0x00, 0x00,
}