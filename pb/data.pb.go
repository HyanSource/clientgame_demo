// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//定義message格式
type Data01 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Money                int32    `protobuf:"varint,2,opt,name=Money,proto3" json:"Money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data01) Reset()         { *m = Data01{} }
func (m *Data01) String() string { return proto.CompactTextString(m) }
func (*Data01) ProtoMessage()    {}
func (*Data01) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *Data01) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data01.Unmarshal(m, b)
}
func (m *Data01) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data01.Marshal(b, m, deterministic)
}
func (m *Data01) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data01.Merge(m, src)
}
func (m *Data01) XXX_Size() int {
	return xxx_messageInfo_Data01.Size(m)
}
func (m *Data01) XXX_DiscardUnknown() {
	xxx_messageInfo_Data01.DiscardUnknown(m)
}

var xxx_messageInfo_Data01 proto.InternalMessageInfo

func (m *Data01) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Data01) GetMoney() int32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func init() {
	proto.RegisterType((*Data01)(nil), "pb.Data01")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 90 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe2, 0x62, 0x73, 0x49,
	0x2c, 0x49, 0x34, 0x30, 0x14, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x7d, 0xf3, 0xf3, 0x52, 0x2b, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x58, 0x83, 0x20, 0x9c, 0x24, 0x36, 0xb0, 0x76, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x29, 0xc8, 0xf5, 0xf5, 0x4c, 0x00, 0x00, 0x00,
}
