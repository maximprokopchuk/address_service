// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateAddressRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ParentId             int32    `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAddressRequest) Reset()         { *m = CreateAddressRequest{} }
func (m *CreateAddressRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAddressRequest) ProtoMessage()    {}
func (*CreateAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0}
}

func (m *CreateAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAddressRequest.Unmarshal(m, b)
}
func (m *CreateAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAddressRequest.Marshal(b, m, deterministic)
}
func (m *CreateAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAddressRequest.Merge(m, src)
}
func (m *CreateAddressRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAddressRequest.Size(m)
}
func (m *CreateAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAddressRequest proto.InternalMessageInfo

func (m *CreateAddressRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAddressRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateAddressRequest) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type AddressResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	ParentId             int32    `protobuf:"varint,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressResponse) Reset()         { *m = AddressResponse{} }
func (m *AddressResponse) String() string { return proto.CompactTextString(m) }
func (*AddressResponse) ProtoMessage()    {}
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{1}
}

func (m *AddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressResponse.Unmarshal(m, b)
}
func (m *AddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressResponse.Marshal(b, m, deterministic)
}
func (m *AddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressResponse.Merge(m, src)
}
func (m *AddressResponse) XXX_Size() int {
	return xxx_messageInfo_AddressResponse.Size(m)
}
func (m *AddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddressResponse proto.InternalMessageInfo

func (m *AddressResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AddressResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddressResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AddressResponse) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type GetAddressesByParentIdRequest struct {
	ParentId             int32    `protobuf:"varint,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAddressesByParentIdRequest) Reset()         { *m = GetAddressesByParentIdRequest{} }
func (m *GetAddressesByParentIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetAddressesByParentIdRequest) ProtoMessage()    {}
func (*GetAddressesByParentIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{2}
}

func (m *GetAddressesByParentIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressesByParentIdRequest.Unmarshal(m, b)
}
func (m *GetAddressesByParentIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressesByParentIdRequest.Marshal(b, m, deterministic)
}
func (m *GetAddressesByParentIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressesByParentIdRequest.Merge(m, src)
}
func (m *GetAddressesByParentIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetAddressesByParentIdRequest.Size(m)
}
func (m *GetAddressesByParentIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressesByParentIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressesByParentIdRequest proto.InternalMessageInfo

func (m *GetAddressesByParentIdRequest) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type GetAddressesByParentIdResponse struct {
	Items                []*AddressResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAddressesByParentIdResponse) Reset()         { *m = GetAddressesByParentIdResponse{} }
func (m *GetAddressesByParentIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetAddressesByParentIdResponse) ProtoMessage()    {}
func (*GetAddressesByParentIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{3}
}

func (m *GetAddressesByParentIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressesByParentIdResponse.Unmarshal(m, b)
}
func (m *GetAddressesByParentIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressesByParentIdResponse.Marshal(b, m, deterministic)
}
func (m *GetAddressesByParentIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressesByParentIdResponse.Merge(m, src)
}
func (m *GetAddressesByParentIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetAddressesByParentIdResponse.Size(m)
}
func (m *GetAddressesByParentIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressesByParentIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressesByParentIdResponse proto.InternalMessageInfo

func (m *GetAddressesByParentIdResponse) GetItems() []*AddressResponse {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetAddressByIdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAddressByIdRequest) Reset()         { *m = GetAddressByIdRequest{} }
func (m *GetAddressByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetAddressByIdRequest) ProtoMessage()    {}
func (*GetAddressByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{4}
}

func (m *GetAddressByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressByIdRequest.Unmarshal(m, b)
}
func (m *GetAddressByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetAddressByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressByIdRequest.Merge(m, src)
}
func (m *GetAddressByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetAddressByIdRequest.Size(m)
}
func (m *GetAddressByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressByIdRequest proto.InternalMessageInfo

func (m *GetAddressByIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteAddressRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAddressRequest) Reset()         { *m = DeleteAddressRequest{} }
func (m *DeleteAddressRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAddressRequest) ProtoMessage()    {}
func (*DeleteAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{5}
}

func (m *DeleteAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAddressRequest.Unmarshal(m, b)
}
func (m *DeleteAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAddressRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAddressRequest.Merge(m, src)
}
func (m *DeleteAddressRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAddressRequest.Size(m)
}
func (m *DeleteAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAddressRequest proto.InternalMessageInfo

func (m *DeleteAddressRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateAddressRequest)(nil), "api.CreateAddressRequest")
	proto.RegisterType((*AddressResponse)(nil), "api.AddressResponse")
	proto.RegisterType((*GetAddressesByParentIdRequest)(nil), "api.GetAddressesByParentIdRequest")
	proto.RegisterType((*GetAddressesByParentIdResponse)(nil), "api.GetAddressesByParentIdResponse")
	proto.RegisterType((*GetAddressByIdRequest)(nil), "api.GetAddressByIdRequest")
	proto.RegisterType((*DeleteAddressRequest)(nil), "api.DeleteAddressRequest")
	proto.RegisterType((*Empty)(nil), "api.Empty")
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e) }

var fileDescriptor_982c640dad8fe78e = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0x5f, 0xdb, 0x55, 0xdd, 0x93, 0x4d, 0x08, 0x13, 0x6a, 0x45, 0x19, 0x11, 0x74, 0x78, 0xd8,
	0x61, 0x5e, 0x3c, 0x78, 0xb1, 0x4e, 0x44, 0xd8, 0x41, 0xea, 0xd1, 0x83, 0x44, 0xf3, 0x84, 0x80,
	0x6d, 0x63, 0x13, 0x85, 0x7e, 0x36, 0xbf, 0x9c, 0x2c, 0xad, 0x9d, 0x8d, 0xed, 0x6e, 0xe1, 0xbd,
	0xf7, 0x7b, 0xbf, 0x3f, 0x2f, 0x30, 0x64, 0x9c, 0xe7, 0xa8, 0xd4, 0x4c, 0xe6, 0x99, 0xce, 0x88,
	0xc7, 0xa4, 0xa0, 0x4f, 0x30, 0xbe, 0xc9, 0x91, 0x69, 0xbc, 0x2e, 0x7b, 0x31, 0x7e, 0x7c, 0xa2,
	0xd2, 0x84, 0x40, 0x3f, 0x65, 0x09, 0x06, 0xce, 0xc4, 0x99, 0x0e, 0x62, 0xf3, 0x5e, 0xd5, 0x74,
	0x21, 0x31, 0x70, 0xcb, 0xda, 0xea, 0x4d, 0x0e, 0x61, 0x20, 0x59, 0x8e, 0xa9, 0x7e, 0x16, 0x3c,
	0xf0, 0x26, 0xce, 0xd4, 0x8f, 0x77, 0xca, 0xc2, 0x3d, 0xa7, 0x6f, 0xb0, 0x57, 0xaf, 0x55, 0x32,
	0x4b, 0x15, 0x92, 0x11, 0xb8, 0x82, 0x9b, 0xad, 0x7e, 0xec, 0x0a, 0x5e, 0xf3, 0xb8, 0x2d, 0x3c,
	0x5e, 0x17, 0x4f, 0xdf, 0xe2, 0xb9, 0x82, 0xa3, 0x3b, 0xd4, 0x15, 0x15, 0xaa, 0xa8, 0x78, 0xa8,
	0x3a, 0xbf, 0x6e, 0x1a, 0x68, 0xc7, 0x42, 0x2f, 0xe1, 0xb8, 0x0b, 0x5d, 0x89, 0x3e, 0x07, 0x5f,
	0x68, 0x4c, 0x54, 0xe0, 0x4c, 0xbc, 0xe9, 0xee, 0x7c, 0x3c, 0x63, 0x52, 0xcc, 0x2c, 0x67, 0x71,
	0x39, 0x42, 0xcf, 0x60, 0x7f, 0xbd, 0x2d, 0x2a, 0xd6, 0x1a, 0x2c, 0xe7, 0xf4, 0x14, 0xc6, 0x0b,
	0x7c, 0xc7, 0x7f, 0xc9, 0xdb, 0x73, 0xdb, 0xe0, 0xdf, 0x26, 0x52, 0x17, 0xf3, 0x6f, 0x17, 0x46,
	0xd5, 0xec, 0x23, 0xe6, 0x5f, 0xe2, 0x15, 0x49, 0x04, 0xc3, 0xc6, 0xf5, 0xc8, 0x81, 0x91, 0xd6,
	0x76, 0xd1, 0xb0, 0x55, 0x35, 0xed, 0x11, 0x84, 0x60, 0x29, 0x94, 0x6e, 0x8b, 0x80, 0x50, 0x83,
	0xd9, 0x98, 0x6d, 0x78, 0xb2, 0x71, 0xa6, 0xa6, 0x59, 0xc0, 0xa8, 0x99, 0x0b, 0x09, 0x2d, 0xe0,
	0x9f, 0xb0, 0x3a, 0xc5, 0x5e, 0xc2, 0xb0, 0x11, 0x5a, 0x65, 0xb8, 0x2d, 0xc8, 0x10, 0x4c, 0xcb,
	0x64, 0x47, 0x7b, 0x2f, 0x5b, 0xe6, 0xd3, 0x5f, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xcf,
	0x07, 0x3f, 0x05, 0x03, 0x00, 0x00,
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
	CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*AddressResponse, error)
	ListGetAddressesByParent(ctx context.Context, in *GetAddressesByParentIdRequest, opts ...grpc.CallOption) (*GetAddressesByParentIdResponse, error)
	GetAddressById(ctx context.Context, in *GetAddressByIdRequest, opts ...grpc.CallOption) (*AddressResponse, error)
	DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*Empty, error)
}

type addressServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddressServiceClient(cc *grpc.ClientConn) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*AddressResponse, error) {
	out := new(AddressResponse)
	err := c.cc.Invoke(ctx, "/api.AddressService/CreateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) ListGetAddressesByParent(ctx context.Context, in *GetAddressesByParentIdRequest, opts ...grpc.CallOption) (*GetAddressesByParentIdResponse, error) {
	out := new(GetAddressesByParentIdResponse)
	err := c.cc.Invoke(ctx, "/api.AddressService/ListGetAddressesByParent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) GetAddressById(ctx context.Context, in *GetAddressByIdRequest, opts ...grpc.CallOption) (*AddressResponse, error) {
	out := new(AddressResponse)
	err := c.cc.Invoke(ctx, "/api.AddressService/GetAddressById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.AddressService/DeleteAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
type AddressServiceServer interface {
	CreateAddress(context.Context, *CreateAddressRequest) (*AddressResponse, error)
	ListGetAddressesByParent(context.Context, *GetAddressesByParentIdRequest) (*GetAddressesByParentIdResponse, error)
	GetAddressById(context.Context, *GetAddressByIdRequest) (*AddressResponse, error)
	DeleteAddress(context.Context, *DeleteAddressRequest) (*Empty, error)
}

// UnimplementedAddressServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (*UnimplementedAddressServiceServer) CreateAddress(ctx context.Context, req *CreateAddressRequest) (*AddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (*UnimplementedAddressServiceServer) ListGetAddressesByParent(ctx context.Context, req *GetAddressesByParentIdRequest) (*GetAddressesByParentIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGetAddressesByParent not implemented")
}
func (*UnimplementedAddressServiceServer) GetAddressById(ctx context.Context, req *GetAddressByIdRequest) (*AddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressById not implemented")
}
func (*UnimplementedAddressServiceServer) DeleteAddress(ctx context.Context, req *DeleteAddressRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}

func RegisterAddressServiceServer(s *grpc.Server, srv AddressServiceServer) {
	s.RegisterService(&_AddressService_serviceDesc, srv)
}

func _AddressService_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AddressService/CreateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).CreateAddress(ctx, req.(*CreateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_ListGetAddressesByParent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressesByParentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).ListGetAddressesByParent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AddressService/ListGetAddressesByParent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).ListGetAddressesByParent(ctx, req.(*GetAddressesByParentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_GetAddressById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GetAddressById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AddressService/GetAddressById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GetAddressById(ctx, req.(*GetAddressByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AddressService/DeleteAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).DeleteAddress(ctx, req.(*DeleteAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAddress",
			Handler:    _AddressService_CreateAddress_Handler,
		},
		{
			MethodName: "ListGetAddressesByParent",
			Handler:    _AddressService_ListGetAddressesByParent_Handler,
		},
		{
			MethodName: "GetAddressById",
			Handler:    _AddressService_GetAddressById_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _AddressService_DeleteAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}
