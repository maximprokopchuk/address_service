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
	Parent               int32    `protobuf:"varint,3,opt,name=parent,proto3" json:"parent,omitempty"`
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

func (m *CreateAddressRequest) GetParent() int32 {
	if m != nil {
		return m.Parent
	}
	return 0
}

type CreateAddressResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Parent               int32    `protobuf:"varint,4,opt,name=parent,proto3" json:"parent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAddressResponse) Reset()         { *m = CreateAddressResponse{} }
func (m *CreateAddressResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAddressResponse) ProtoMessage()    {}
func (*CreateAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{1}
}

func (m *CreateAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAddressResponse.Unmarshal(m, b)
}
func (m *CreateAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAddressResponse.Marshal(b, m, deterministic)
}
func (m *CreateAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAddressResponse.Merge(m, src)
}
func (m *CreateAddressResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAddressResponse.Size(m)
}
func (m *CreateAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAddressResponse proto.InternalMessageInfo

func (m *CreateAddressResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateAddressResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAddressResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateAddressResponse) GetParent() int32 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateAddressRequest)(nil), "api.CreateAddressRequest")
	proto.RegisterType((*CreateAddressResponse)(nil), "api.CreateAddressResponse")
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e) }

var fileDescriptor_982c640dad8fe78e = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x0a,
	0xe3, 0x12, 0x71, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x84, 0xc8, 0x05, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0xd9, 0x20, 0xb1, 0x92, 0xca, 0x82, 0x54, 0x09, 0x26, 0x88, 0x18, 0x88, 0x2d, 0x24,
	0xc6, 0xc5, 0x56, 0x90, 0x58, 0x94, 0x9a, 0x57, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1a, 0x04,
	0xe5, 0x29, 0xa5, 0x73, 0x89, 0xa2, 0x99, 0x5b, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xc4, 0xc7,
	0xc5, 0x94, 0x99, 0x02, 0x36, 0x96, 0x35, 0x88, 0x29, 0x33, 0x05, 0x6e, 0x11, 0x13, 0x16, 0x8b,
	0x98, 0xb1, 0x5a, 0xc4, 0x82, 0x6c, 0x91, 0x51, 0x14, 0x17, 0x1f, 0xd4, 0x8a, 0xe0, 0xd4, 0xa2,
	0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x0f, 0x2e, 0x5e, 0x14, 0xab, 0x85, 0x24, 0xf5, 0x12, 0x0b, 0x32,
	0xf5, 0xb0, 0x79, 0x53, 0x4a, 0x0a, 0x9b, 0x14, 0xc4, 0xa5, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0x80,
	0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x75, 0x76, 0x2b, 0xf0, 0x39, 0x01, 0x00, 0x00,
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
	CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*CreateAddressResponse, error)
}

type addressServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddressServiceClient(cc *grpc.ClientConn) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*CreateAddressResponse, error) {
	out := new(CreateAddressResponse)
	err := c.cc.Invoke(ctx, "/api.AddressService/CreateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
type AddressServiceServer interface {
	CreateAddress(context.Context, *CreateAddressRequest) (*CreateAddressResponse, error)
}

// UnimplementedAddressServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (*UnimplementedAddressServiceServer) CreateAddress(ctx context.Context, req *CreateAddressRequest) (*CreateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
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

var _AddressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAddress",
			Handler:    _AddressService_CreateAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}
