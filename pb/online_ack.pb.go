// Code generated by protoc-gen-go. DO NOT EDIT.
// source: online_ack.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	online_ack.proto

It has these top-level messages:
	OnLineACK
*/
package pb

import "github.com/golang/protobuf/proto"
import "fmt"
import "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var onlineACK OnLineACK

// 设备上线回执 type:2
type OnLineACK struct {
	Ok      bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *OnLineACK) Reset()                    { *m = OnLineACK{} }
func (m *OnLineACK) String() string            { return proto.CompactTextString(m) }
func (*OnLineACK) ProtoMessage()               {}
func (*OnLineACK) Descriptor() ([]byte, []int) { return onlineACK.fileDescriptor(), []int{0} }

func (m *OnLineACK) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *OnLineACK) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*OnLineACK)(nil), "pb.OnLineACK")
}

func init() { proto.RegisterFile("online_ack.proto", onlineACK.fileDescriptor()) }

func (OnLineACK) fileDescriptor() []byte {
	return []byte{
		// 100 bytes of a gzipped FileDescriptorProto
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcf, 0xcb, 0xc9,
		0xcc, 0x4b, 0x8d, 0x4f, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
		0x52, 0x32, 0xe5, 0xe2, 0xf4, 0xcf, 0xf3, 0xc9, 0xcc, 0x4b, 0x75, 0x74, 0xf6, 0x16, 0xe2, 0xe3,
		0x62, 0xca, 0xcf, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x62, 0xca, 0xcf, 0x16, 0x92, 0xe0,
		0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
		0x71, 0x93, 0xd8, 0xc0, 0x26, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xd6, 0x0e, 0xb2,
		0x55, 0x00, 0x00, 0x00,
	}
}
