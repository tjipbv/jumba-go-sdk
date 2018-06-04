// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pagination.proto

package service

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

type Pagination struct {
	Total                int64    `protobuf:"varint,1,opt,name=Total" json:"Total,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=Limit" json:"Limit,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=Offset" json:"Offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_pagination_34605a46388f33c2, []int{0}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (dst *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(dst, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Pagination) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Pagination) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*Pagination)(nil), "proto.Pagination")
}

func init() { proto.RegisterFile("pagination.proto", fileDescriptor_pagination_34605a46388f33c2) }

var fileDescriptor_pagination_34605a46388f33c2 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0x4c, 0xcf,
	0xcc, 0x4b, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x4a, 0x01, 0x5c, 0x5c, 0x01, 0x70, 0x29, 0x21, 0x11, 0x2e, 0xd6, 0x90, 0xfc, 0x92, 0xc4, 0x1c,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x07, 0x24, 0xea, 0x93, 0x99, 0x9b, 0x59, 0x22,
	0xc1, 0x04, 0x11, 0x05, 0x73, 0x84, 0xc4, 0xb8, 0xd8, 0xfc, 0xd3, 0xd2, 0x8a, 0x53, 0x4b, 0x24,
	0x98, 0xc1, 0xc2, 0x50, 0x5e, 0x12, 0x1b, 0xd8, 0x60, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x92, 0xb6, 0x43, 0x3f, 0x73, 0x00, 0x00, 0x00,
}