// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

package api_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddressIndexOptions struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressIndexOptions) Reset()         { *m = AddressIndexOptions{} }
func (m *AddressIndexOptions) String() string { return proto.CompactTextString(m) }
func (*AddressIndexOptions) ProtoMessage()    {}
func (*AddressIndexOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{0}
}
func (m *AddressIndexOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressIndexOptions.Unmarshal(m, b)
}
func (m *AddressIndexOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressIndexOptions.Marshal(b, m, deterministic)
}
func (dst *AddressIndexOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressIndexOptions.Merge(dst, src)
}
func (m *AddressIndexOptions) XXX_Size() int {
	return xxx_messageInfo_AddressIndexOptions.Size(m)
}
func (m *AddressIndexOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressIndexOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressIndexOptions proto.InternalMessageInfo

func (m *AddressIndexOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddressIndexResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressIndexResponse) Reset()         { *m = AddressIndexResponse{} }
func (m *AddressIndexResponse) String() string { return proto.CompactTextString(m) }
func (*AddressIndexResponse) ProtoMessage()    {}
func (*AddressIndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{1}
}
func (m *AddressIndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressIndexResponse.Unmarshal(m, b)
}
func (m *AddressIndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressIndexResponse.Marshal(b, m, deterministic)
}
func (dst *AddressIndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressIndexResponse.Merge(dst, src)
}
func (m *AddressIndexResponse) XXX_Size() int {
	return xxx_messageInfo_AddressIndexResponse.Size(m)
}
func (m *AddressIndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressIndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddressIndexResponse proto.InternalMessageInfo

func (m *AddressIndexResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddressSuggestOptions struct {
	Suggest              string   `protobuf:"bytes,1,opt,name=suggest" json:"suggest,omitempty"`
	Forsale              bool     `protobuf:"varint,2,opt,name=forsale" json:"forsale,omitempty"`
	Type                 int64    `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressSuggestOptions) Reset()         { *m = AddressSuggestOptions{} }
func (m *AddressSuggestOptions) String() string { return proto.CompactTextString(m) }
func (*AddressSuggestOptions) ProtoMessage()    {}
func (*AddressSuggestOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{2}
}
func (m *AddressSuggestOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressSuggestOptions.Unmarshal(m, b)
}
func (m *AddressSuggestOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressSuggestOptions.Marshal(b, m, deterministic)
}
func (dst *AddressSuggestOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressSuggestOptions.Merge(dst, src)
}
func (m *AddressSuggestOptions) XXX_Size() int {
	return xxx_messageInfo_AddressSuggestOptions.Size(m)
}
func (m *AddressSuggestOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressSuggestOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressSuggestOptions proto.InternalMessageInfo

func (m *AddressSuggestOptions) GetSuggest() string {
	if m != nil {
		return m.Suggest
	}
	return ""
}

func (m *AddressSuggestOptions) GetForsale() bool {
	if m != nil {
		return m.Forsale
	}
	return false
}

func (m *AddressSuggestOptions) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

type AddressGetOptions struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Bagid                string   `protobuf:"bytes,3,opt,name=bagid" json:"bagid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressGetOptions) Reset()         { *m = AddressGetOptions{} }
func (m *AddressGetOptions) String() string { return proto.CompactTextString(m) }
func (*AddressGetOptions) ProtoMessage()    {}
func (*AddressGetOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{3}
}
func (m *AddressGetOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressGetOptions.Unmarshal(m, b)
}
func (m *AddressGetOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressGetOptions.Marshal(b, m, deterministic)
}
func (dst *AddressGetOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressGetOptions.Merge(dst, src)
}
func (m *AddressGetOptions) XXX_Size() int {
	return xxx_messageInfo_AddressGetOptions.Size(m)
}
func (m *AddressGetOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressGetOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressGetOptions proto.InternalMessageInfo

func (m *AddressGetOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddressGetOptions) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *AddressGetOptions) GetBagid() string {
	if m != nil {
		return m.Bagid
	}
	return ""
}

type AddressGetByPathOptions struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressGetByPathOptions) Reset()         { *m = AddressGetByPathOptions{} }
func (m *AddressGetByPathOptions) String() string { return proto.CompactTextString(m) }
func (*AddressGetByPathOptions) ProtoMessage()    {}
func (*AddressGetByPathOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{4}
}
func (m *AddressGetByPathOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressGetByPathOptions.Unmarshal(m, b)
}
func (m *AddressGetByPathOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressGetByPathOptions.Marshal(b, m, deterministic)
}
func (dst *AddressGetByPathOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressGetByPathOptions.Merge(dst, src)
}
func (m *AddressGetByPathOptions) XXX_Size() int {
	return xxx_messageInfo_AddressGetByPathOptions.Size(m)
}
func (m *AddressGetByPathOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressGetByPathOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressGetByPathOptions proto.InternalMessageInfo

func (m *AddressGetByPathOptions) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type AddressGetMultiOptions struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressGetMultiOptions) Reset()         { *m = AddressGetMultiOptions{} }
func (m *AddressGetMultiOptions) String() string { return proto.CompactTextString(m) }
func (*AddressGetMultiOptions) ProtoMessage()    {}
func (*AddressGetMultiOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{5}
}
func (m *AddressGetMultiOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressGetMultiOptions.Unmarshal(m, b)
}
func (m *AddressGetMultiOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressGetMultiOptions.Marshal(b, m, deterministic)
}
func (dst *AddressGetMultiOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressGetMultiOptions.Merge(dst, src)
}
func (m *AddressGetMultiOptions) XXX_Size() int {
	return xxx_messageInfo_AddressGetMultiOptions.Size(m)
}
func (m *AddressGetMultiOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressGetMultiOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressGetMultiOptions proto.InternalMessageInfo

func (m *AddressGetMultiOptions) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type AddressListOptions struct {
	Pagination           *Pagination `protobuf:"bytes,1,opt,name=pagination" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddressListOptions) Reset()         { *m = AddressListOptions{} }
func (m *AddressListOptions) String() string { return proto.CompactTextString(m) }
func (*AddressListOptions) ProtoMessage()    {}
func (*AddressListOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{6}
}
func (m *AddressListOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressListOptions.Unmarshal(m, b)
}
func (m *AddressListOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressListOptions.Marshal(b, m, deterministic)
}
func (dst *AddressListOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressListOptions.Merge(dst, src)
}
func (m *AddressListOptions) XXX_Size() int {
	return xxx_messageInfo_AddressListOptions.Size(m)
}
func (m *AddressListOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressListOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressListOptions proto.InternalMessageInfo

func (m *AddressListOptions) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type Address struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Bagid                string   `protobuf:"bytes,3,opt,name=bagid" json:"bagid,omitempty"`
	Legacy               *Legacy  `protobuf:"bytes,4,opt,name=legacy" json:"legacy,omitempty"`
	Suggest              []string `protobuf:"bytes,5,rep,name=suggest" json:"suggest,omitempty"`
	Forsale              bool     `protobuf:"varint,6,opt,name=forsale" json:"forsale,omitempty"`
	Label                string   `protobuf:"bytes,7,opt,name=label" json:"label,omitempty"`
	Type                 int64    `protobuf:"varint,8,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{7}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Address) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Address) GetBagid() string {
	if m != nil {
		return m.Bagid
	}
	return ""
}

func (m *Address) GetLegacy() *Legacy {
	if m != nil {
		return m.Legacy
	}
	return nil
}

func (m *Address) GetSuggest() []string {
	if m != nil {
		return m.Suggest
	}
	return nil
}

func (m *Address) GetForsale() bool {
	if m != nil {
		return m.Forsale
	}
	return false
}

func (m *Address) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Address) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

type Addresses struct {
	Addresses            []*Address  `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Pagination           *Pagination `protobuf:"bytes,2,opt,name=pagination" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Addresses) Reset()         { *m = Addresses{} }
func (m *Addresses) String() string { return proto.CompactTextString(m) }
func (*Addresses) ProtoMessage()    {}
func (*Addresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_3400dfb088bdd8ba, []int{8}
}
func (m *Addresses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addresses.Unmarshal(m, b)
}
func (m *Addresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addresses.Marshal(b, m, deterministic)
}
func (dst *Addresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addresses.Merge(dst, src)
}
func (m *Addresses) XXX_Size() int {
	return xxx_messageInfo_Addresses.Size(m)
}
func (m *Addresses) XXX_DiscardUnknown() {
	xxx_messageInfo_Addresses.DiscardUnknown(m)
}

var xxx_messageInfo_Addresses proto.InternalMessageInfo

func (m *Addresses) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Addresses) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressIndexOptions)(nil), "api_v1.AddressIndexOptions")
	proto.RegisterType((*AddressIndexResponse)(nil), "api_v1.AddressIndexResponse")
	proto.RegisterType((*AddressSuggestOptions)(nil), "api_v1.AddressSuggestOptions")
	proto.RegisterType((*AddressGetOptions)(nil), "api_v1.AddressGetOptions")
	proto.RegisterType((*AddressGetByPathOptions)(nil), "api_v1.AddressGetByPathOptions")
	proto.RegisterType((*AddressGetMultiOptions)(nil), "api_v1.AddressGetMultiOptions")
	proto.RegisterType((*AddressListOptions)(nil), "api_v1.AddressListOptions")
	proto.RegisterType((*Address)(nil), "api_v1.Address")
	proto.RegisterType((*Addresses)(nil), "api_v1.Addresses")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddressServiceClient interface {
	Get(ctx context.Context, in *AddressGetOptions, opts ...grpc.CallOption) (*Address, error)
	GetMulti(ctx context.Context, in *AddressGetMultiOptions, opts ...grpc.CallOption) (*Addresses, error)
	List(ctx context.Context, in *AddressListOptions, opts ...grpc.CallOption) (*Addresses, error)
	Suggest(ctx context.Context, in *AddressSuggestOptions, opts ...grpc.CallOption) (*Addresses, error)
	Update(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
	Index(ctx context.Context, in *AddressIndexOptions, opts ...grpc.CallOption) (*AddressIndexResponse, error)
}

type addressServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddressServiceClient(cc *grpc.ClientConn) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) Get(ctx context.Context, in *AddressGetOptions, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/api_v1.AddressService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) GetMulti(ctx context.Context, in *AddressGetMultiOptions, opts ...grpc.CallOption) (*Addresses, error) {
	out := new(Addresses)
	err := c.cc.Invoke(ctx, "/api_v1.AddressService/GetMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) List(ctx context.Context, in *AddressListOptions, opts ...grpc.CallOption) (*Addresses, error) {
	out := new(Addresses)
	err := c.cc.Invoke(ctx, "/api_v1.AddressService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Suggest(ctx context.Context, in *AddressSuggestOptions, opts ...grpc.CallOption) (*Addresses, error) {
	out := new(Addresses)
	err := c.cc.Invoke(ctx, "/api_v1.AddressService/Suggest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Update(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/api_v1.AddressService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) Index(ctx context.Context, in *AddressIndexOptions, opts ...grpc.CallOption) (*AddressIndexResponse, error) {
	out := new(AddressIndexResponse)
	err := c.cc.Invoke(ctx, "/api_v1.AddressService/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
type AddressServiceServer interface {
	Get(context.Context, *AddressGetOptions) (*Address, error)
	GetMulti(context.Context, *AddressGetMultiOptions) (*Addresses, error)
	List(context.Context, *AddressListOptions) (*Addresses, error)
	Suggest(context.Context, *AddressSuggestOptions) (*Addresses, error)
	Update(context.Context, *Address) (*Address, error)
	Index(context.Context, *AddressIndexOptions) (*AddressIndexResponse, error)
}

func RegisterAddressServiceServer(s *grpc.Server, srv AddressServiceServer) {
	s.RegisterService(&_AddressService_serviceDesc, srv)
}

func _AddressService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressGetOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_v1.AddressService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Get(ctx, req.(*AddressGetOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_GetMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressGetMultiOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GetMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_v1.AddressService/GetMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GetMulti(ctx, req.(*AddressGetMultiOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressListOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_v1.AddressService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).List(ctx, req.(*AddressListOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Suggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressSuggestOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Suggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_v1.AddressService/Suggest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Suggest(ctx, req.(*AddressSuggestOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_v1.AddressService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Update(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressIndexOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_v1.AddressService/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Index(ctx, req.(*AddressIndexOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api_v1.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AddressService_Get_Handler,
		},
		{
			MethodName: "GetMulti",
			Handler:    _AddressService_GetMulti_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AddressService_List_Handler,
		},
		{
			MethodName: "Suggest",
			Handler:    _AddressService_Suggest_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AddressService_Update_Handler,
		},
		{
			MethodName: "Index",
			Handler:    _AddressService_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_address_3400dfb088bdd8ba) }

var fileDescriptor_address_3400dfb088bdd8ba = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0xa6, 0x5f, 0x39, 0x83, 0xae, 0x3d, 0x94, 0x2d, 0x84, 0x82, 0x2a, 0x4b, 0xa0,
	0x0a, 0xb1, 0x96, 0x95, 0x6b, 0x2e, 0x40, 0x48, 0x03, 0xb4, 0x69, 0x93, 0x11, 0x57, 0x5c, 0x4c,
	0xee, 0x62, 0x32, 0x4b, 0x59, 0x12, 0xd5, 0x5e, 0x45, 0x6f, 0x79, 0x05, 0x9e, 0x87, 0x07, 0xe0,
	0x9a, 0x57, 0xe0, 0x41, 0x50, 0x1c, 0x7b, 0xc9, 0xb2, 0x22, 0x21, 0xee, 0x7c, 0x7c, 0xfe, 0xe7,
	0x77, 0xec, 0x73, 0x8e, 0x0d, 0x77, 0x59, 0x18, 0x2e, 0xb9, 0x94, 0xd3, 0x6c, 0x99, 0xaa, 0x14,
	0xdb, 0x2c, 0x13, 0x67, 0xab, 0x83, 0xe0, 0x4e, 0xcc, 0x23, 0x76, 0xbe, 0x2e, 0x76, 0x83, 0x7e,
	0xc6, 0x22, 0x91, 0x30, 0x25, 0xd2, 0xc4, 0xec, 0x8c, 0xa2, 0x34, 0x8d, 0x62, 0x3e, 0x63, 0x99,
	0x98, 0xb1, 0x24, 0x49, 0x95, 0x76, 0x1a, 0x0a, 0x79, 0x02, 0xf7, 0x5e, 0x17, 0xd8, 0xf7, 0x49,
	0xc8, 0xbf, 0x9e, 0x64, 0xda, 0x89, 0x3d, 0x68, 0x88, 0xd0, 0x77, 0xc6, 0xce, 0xc4, 0xa3, 0x0d,
	0x11, 0x92, 0x17, 0x30, 0xac, 0xca, 0x28, 0x97, 0x59, 0x9a, 0x48, 0x8e, 0x3e, 0x74, 0x2e, 0xb9,
	0x94, 0x2c, 0xe2, 0x46, 0x6c, 0x4d, 0x72, 0x06, 0xf7, 0x4d, 0xc4, 0xc7, 0xab, 0x28, 0xe2, 0x52,
	0x59, 0xb4, 0x0f, 0x1d, 0x59, 0xec, 0xd8, 0x10, 0x63, 0xe6, 0x9e, 0x2f, 0xe9, 0x52, 0xb2, 0x98,
	0xfb, 0x8d, 0xb1, 0x33, 0xe9, 0x52, 0x6b, 0x22, 0x42, 0x53, 0xad, 0x33, 0xee, 0xbb, 0x63, 0x67,
	0xe2, 0x52, 0xbd, 0x26, 0xc7, 0x30, 0x30, 0x09, 0x0e, 0xb9, 0xfa, 0xcb, 0xb9, 0xf3, 0xc0, 0x8c,
	0xa9, 0x0b, 0xcd, 0xf3, 0xa8, 0x5e, 0xe3, 0x10, 0x5a, 0x0b, 0x16, 0x89, 0x50, 0xd3, 0x3c, 0x5a,
	0x18, 0x64, 0x1f, 0xf6, 0x4a, 0xdc, 0x9b, 0xf5, 0x29, 0x53, 0x17, 0x16, 0x6a, 0x21, 0x4e, 0x09,
	0x21, 0xcf, 0x60, 0xb7, 0x94, 0x1f, 0x5f, 0xc5, 0x4a, 0x58, 0x75, 0x1f, 0x5c, 0x11, 0x4a, 0xdf,
	0x19, 0xbb, 0x13, 0x8f, 0xe6, 0x4b, 0xf2, 0x0e, 0xd0, 0x68, 0x8f, 0x44, 0x59, 0x87, 0x39, 0x40,
	0xd9, 0x2b, 0xcd, 0xde, 0x9e, 0xe3, 0xb4, 0x68, 0xea, 0xf4, 0xf4, 0xda, 0x43, 0x2b, 0x2a, 0xf2,
	0xd3, 0x81, 0x8e, 0x41, 0xfd, 0xff, 0x55, 0xf1, 0x29, 0xb4, 0x8b, 0x99, 0xf1, 0x9b, 0x3a, 0x6b,
	0xcf, 0x66, 0x3d, 0xd2, 0xbb, 0xd4, 0x78, 0xab, 0x9d, 0x6a, 0xe9, 0xdb, 0x6c, 0xea, 0x54, 0xfb,
	0x66, 0xa7, 0x86, 0xd0, 0x8a, 0xd9, 0x82, 0xc7, 0x7e, 0xa7, 0xc8, 0xa8, 0x8d, 0xeb, 0xfe, 0x75,
	0x2b, 0xfd, 0x4b, 0xc0, 0x33, 0x57, 0xe1, 0x12, 0xf7, 0xc1, 0x63, 0xd6, 0xd0, 0xa5, 0xdb, 0x9e,
	0xef, 0xd8, 0x53, 0x19, 0x15, 0x2d, 0x15, 0xb5, 0xda, 0x35, 0xfe, 0xa5, 0x76, 0xf3, 0x1f, 0x2e,
	0xf4, 0xec, 0x44, 0xf2, 0xe5, 0x4a, 0x9c, 0x73, 0xfc, 0x00, 0xee, 0x21, 0x57, 0xf8, 0xa0, 0x96,
	0xa9, 0x9c, 0xa7, 0xa0, 0x7e, 0x08, 0xb2, 0xf7, 0xed, 0xd7, 0xef, 0xef, 0x8d, 0x01, 0xee, 0xcc,
	0x56, 0x07, 0x33, 0x73, 0xa0, 0x59, 0xc4, 0x15, 0x7e, 0x86, 0xae, 0x9d, 0x04, 0x7c, 0x7c, 0x1b,
	0x58, 0x1d, 0x91, 0x60, 0x50, 0xf3, 0x73, 0x49, 0x46, 0x9a, 0xbb, 0x8b, 0xc3, 0x1a, 0xf7, 0x52,
	0x03, 0x4f, 0xa0, 0x99, 0x8f, 0x0e, 0x06, 0xb5, 0xc0, 0xca, 0x3c, 0x6d, 0x82, 0xfa, 0x1a, 0x8a,
	0xd8, 0xaf, 0x42, 0xe3, 0x1c, 0xf4, 0x0a, 0x3a, 0xe6, 0x59, 0xe2, 0xa3, 0x5a, 0xdc, 0xcd, 0xe7,
	0xba, 0x09, 0xbb, 0x85, 0xcf, 0xa1, 0xfd, 0x29, 0x0b, 0x99, 0xe2, 0x58, 0x2f, 0xd0, 0xed, 0x8a,
	0x6d, 0xe1, 0x5b, 0x68, 0xe9, 0x5f, 0x03, 0x1f, 0xd6, 0x7c, 0xd5, 0x2f, 0x27, 0x18, 0x6d, 0x72,
	0xda, 0x8f, 0x86, 0x6c, 0x2d, 0xda, 0xfa, 0xc3, 0x7a, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2b,
	0xb6, 0x92, 0xe1, 0x07, 0x05, 0x00, 0x00,
}
