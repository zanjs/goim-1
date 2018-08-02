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

// 消息投递 type:7
type Message struct {
	Messages []*MessageItem `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetMessages() []*MessageItem {
	if m != nil {
		return m.Messages
	}
	return nil
}

// 单条消息投递
type MessageItem struct {
	DeviceId uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Sender   uint64 `protobuf:"varint,2,opt,name=sender" json:"sender,omitempty"`
	Receiver uint64 `protobuf:"varint,3,opt,name=receiver" json:"receiver,omitempty"`
	GroupId  uint64 `protobuf:"varint,5,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	Type     uint32 `protobuf:"varint,6,opt,name=type" json:"type,omitempty"`
	Content  []byte `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	Sn       uint64 `protobuf:"varint,8,opt,name=sn" json:"sn,omitempty"`
}

func (m *MessageItem) Reset()                    { *m = MessageItem{} }
func (m *MessageItem) String() string            { return proto.CompactTextString(m) }
func (*MessageItem) ProtoMessage()               {}
func (*MessageItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MessageItem) GetDeviceId() uint64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *MessageItem) GetSender() uint64 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *MessageItem) GetReceiver() uint64 {
	if m != nil {
		return m.Receiver
	}
	return 0
}

func (m *MessageItem) GetGroupId() uint64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *MessageItem) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageItem) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MessageItem) GetSn() uint64 {
	if m != nil {
		return m.Sn
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*MessageItem)(nil), "pb.MessageItem")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0x77, 0x6d, 0xe2, 0xac, 0xab, 0x30, 0x07, 0x19, 0xf5, 0x12, 0xf6, 0x14, 0x10,
	0x7a, 0x50, 0xf0, 0x1d, 0x7a, 0xf0, 0x92, 0x17, 0x10, 0xdb, 0x0c, 0x4b, 0x0f, 0x9b, 0x84, 0x24,
	0x2e, 0xf8, 0x64, 0xbe, 0x9e, 0x98, 0xb6, 0xb2, 0xb7, 0xff, 0xfb, 0xbf, 0xe4, 0x87, 0x81, 0xfd,
	0x89, 0x73, 0xfe, 0x3c, 0x72, 0x17, 0x53, 0x28, 0x01, 0x9b, 0x38, 0x1c, 0xde, 0x40, 0xbe, 0xcf,
	0x25, 0x3e, 0x83, 0x5a, 0x62, 0x26, 0xa1, 0x37, 0x66, 0xf7, 0x72, 0xd7, 0xc5, 0xa1, 0x5b, 0xba,
	0xbe, 0xf0, 0xc9, 0xfe, 0x3f, 0x38, 0xfc, 0x08, 0xd8, 0x5d, 0x18, 0x7c, 0x82, 0x6b, 0xc7, 0xe7,
	0x69, 0xe4, 0x8f, 0xc9, 0x91, 0xd0, 0xc2, 0x6c, 0xad, 0x9a, 0x8b, 0xde, 0xe1, 0x3d, 0xb4, 0x99,
	0xbd, 0xe3, 0x44, 0x4d, 0x35, 0x0b, 0xe1, 0x23, 0xa8, 0xc4, 0x23, 0x4f, 0x67, 0x4e, 0xb4, 0x99,
	0xff, 0xac, 0x8c, 0x0f, 0xa0, 0x8e, 0x29, 0x7c, 0xc5, 0xbf, 0xbd, 0xab, 0xea, 0x64, 0xe5, 0xde,
	0x21, 0xc2, 0xb6, 0x7c, 0x47, 0xa6, 0x56, 0x0b, 0xb3, 0xb7, 0x35, 0x23, 0x81, 0x1c, 0x83, 0x2f,
	0xec, 0x0b, 0x49, 0x2d, 0xcc, 0x8d, 0x5d, 0x11, 0x6f, 0xa1, 0xc9, 0x9e, 0x54, 0x9d, 0x68, 0xb2,
	0x1f, 0xda, 0x7a, 0xfc, 0xeb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xcc, 0xfc, 0xd8, 0x0d,
	0x01, 0x00, 0x00,
}
