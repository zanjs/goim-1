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
	SenderType     int32  `protobuf:"varint,1,opt,name=sender_type,json=senderType" json:"sender_type,omitempty"`
	SenderId       int64  `protobuf:"varint,2,opt,name=sender_id,json=senderId" json:"sender_id,omitempty"`
	SenderDeviceId int64  `protobuf:"varint,3,opt,name=sender_device_id,json=senderDeviceId" json:"sender_device_id,omitempty"`
	ReceiverType   int32  `protobuf:"varint,5,opt,name=receiver_type,json=receiverType" json:"receiver_type,omitempty"`
	ReceiverId     int64  `protobuf:"varint,6,opt,name=receiver_id,json=receiverId" json:"receiver_id,omitempty"`
	Type           int32  `protobuf:"varint,7,opt,name=type" json:"type,omitempty"`
	Content        string `protobuf:"bytes,8,opt,name=content" json:"content,omitempty"`
	Sequence       int64  `protobuf:"varint,9,opt,name=sequence" json:"sequence,omitempty"`
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

func (m *MessageItem) GetSenderDeviceId() int64 {
	if m != nil {
		return m.SenderDeviceId
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
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0xeb, 0x6e, 0xdb, 0xa9, 0xab, 0x32, 0xa7, 0xa0, 0x87, 0x2d, 0xeb, 0x25, 0x20,
	0xf4, 0xa0, 0xe0, 0x13, 0x78, 0xe9, 0xc1, 0x4b, 0xf1, 0x2e, 0xb6, 0x19, 0xa4, 0x87, 0x4d, 0x63,
	0x13, 0x17, 0xf6, 0x01, 0x7c, 0x6f, 0x71, 0x92, 0x14, 0x6f, 0x33, 0xdf, 0x7c, 0x93, 0x3f, 0x0c,
	0xec, 0x8e, 0xe4, 0xdc, 0xc7, 0x27, 0xb5, 0x76, 0x99, 0xfd, 0x8c, 0x99, 0x1d, 0x0e, 0xcf, 0x50,
	0xbc, 0x06, 0x88, 0x0f, 0x50, 0xc6, 0xd2, 0x49, 0xd1, 0xe4, 0xaa, 0x7e, 0xbc, 0x6e, 0xed, 0xd0,
	0x46, 0xd6, 0x79, 0x3a, 0xf6, 0xab, 0x70, 0xf8, 0xc9, 0xa0, 0xfe, 0x37, 0xc1, 0x3d, 0xd4, 0x8e,
	0x8c, 0xa6, 0xe5, 0xdd, 0x9f, 0x2d, 0x49, 0xd1, 0x08, 0xb5, 0xe9, 0x21, 0xa0, 0xb7, 0xb3, 0x25,
	0xbc, 0x83, 0x2a, 0x0a, 0x93, 0x96, 0x59, 0x23, 0x54, 0xde, 0x97, 0x01, 0x74, 0x1a, 0x15, 0xdc,
	0xc4, 0xa1, 0xa6, 0xd3, 0x34, 0xd2, 0x9f, 0x93, 0xb3, 0x73, 0x15, 0xf8, 0x0b, 0xe3, 0x4e, 0xe3,
	0x3d, 0xec, 0x16, 0x1a, 0x69, 0x3a, 0xa5, 0xa4, 0x0d, 0x27, 0x5d, 0x26, 0xc8, 0x59, 0x7b, 0xa8,
	0x57, 0x69, 0xd2, 0x72, 0xcb, 0x2f, 0x41, 0x42, 0x9d, 0x46, 0x84, 0x0b, 0x5e, 0x2e, 0x78, 0x99,
	0x6b, 0x94, 0x50, 0x8c, 0xb3, 0xf1, 0x64, 0xbc, 0x2c, 0x1b, 0xa1, 0xaa, 0x3e, 0xb5, 0x78, 0x0b,
	0xa5, 0xa3, 0xaf, 0x6f, 0x32, 0x23, 0xc9, 0x2a, 0xfd, 0x3c, 0xf4, 0xc3, 0x96, 0x4f, 0xf9, 0xf4,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x95, 0x98, 0xfa, 0x5b, 0x01, 0x00, 0x00,
}
