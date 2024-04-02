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

type Address struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	ParentId             int32    `protobuf:"varint,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Address) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Address) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Address) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type AddressResponse struct {
	Result               *Address `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressResponse) Reset()         { *m = AddressResponse{} }
func (m *AddressResponse) String() string { return proto.CompactTextString(m) }
func (*AddressResponse) ProtoMessage()    {}
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{2}
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

func (m *AddressResponse) GetResult() *Address {
	if m != nil {
		return m.Result
	}
	return nil
}

type ListAddressesByParentIdAndTypeRequest struct {
	ParentId             int32    `protobuf:"varint,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAddressesByParentIdAndTypeRequest) Reset()         { *m = ListAddressesByParentIdAndTypeRequest{} }
func (m *ListAddressesByParentIdAndTypeRequest) String() string { return proto.CompactTextString(m) }
func (*ListAddressesByParentIdAndTypeRequest) ProtoMessage()    {}
func (*ListAddressesByParentIdAndTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{3}
}

func (m *ListAddressesByParentIdAndTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAddressesByParentIdAndTypeRequest.Unmarshal(m, b)
}
func (m *ListAddressesByParentIdAndTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAddressesByParentIdAndTypeRequest.Marshal(b, m, deterministic)
}
func (m *ListAddressesByParentIdAndTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAddressesByParentIdAndTypeRequest.Merge(m, src)
}
func (m *ListAddressesByParentIdAndTypeRequest) XXX_Size() int {
	return xxx_messageInfo_ListAddressesByParentIdAndTypeRequest.Size(m)
}
func (m *ListAddressesByParentIdAndTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAddressesByParentIdAndTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAddressesByParentIdAndTypeRequest proto.InternalMessageInfo

func (m *ListAddressesByParentIdAndTypeRequest) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *ListAddressesByParentIdAndTypeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ListAddressesByParentIdAndTypeResponse struct {
	Result               []*Address `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListAddressesByParentIdAndTypeResponse) Reset() {
	*m = ListAddressesByParentIdAndTypeResponse{}
}
func (m *ListAddressesByParentIdAndTypeResponse) String() string { return proto.CompactTextString(m) }
func (*ListAddressesByParentIdAndTypeResponse) ProtoMessage()    {}
func (*ListAddressesByParentIdAndTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{4}
}

func (m *ListAddressesByParentIdAndTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAddressesByParentIdAndTypeResponse.Unmarshal(m, b)
}
func (m *ListAddressesByParentIdAndTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAddressesByParentIdAndTypeResponse.Marshal(b, m, deterministic)
}
func (m *ListAddressesByParentIdAndTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAddressesByParentIdAndTypeResponse.Merge(m, src)
}
func (m *ListAddressesByParentIdAndTypeResponse) XXX_Size() int {
	return xxx_messageInfo_ListAddressesByParentIdAndTypeResponse.Size(m)
}
func (m *ListAddressesByParentIdAndTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAddressesByParentIdAndTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAddressesByParentIdAndTypeResponse proto.InternalMessageInfo

func (m *ListAddressesByParentIdAndTypeResponse) GetResult() []*Address {
	if m != nil {
		return m.Result
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
	return fileDescriptor_982c640dad8fe78e, []int{5}
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
	return fileDescriptor_982c640dad8fe78e, []int{6}
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
	return fileDescriptor_982c640dad8fe78e, []int{7}
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
	proto.RegisterType((*Address)(nil), "api.Address")
	proto.RegisterType((*AddressResponse)(nil), "api.AddressResponse")
	proto.RegisterType((*ListAddressesByParentIdAndTypeRequest)(nil), "api.ListAddressesByParentIdAndTypeRequest")
	proto.RegisterType((*ListAddressesByParentIdAndTypeResponse)(nil), "api.ListAddressesByParentIdAndTypeResponse")
	proto.RegisterType((*GetAddressByIdRequest)(nil), "api.GetAddressByIdRequest")
	proto.RegisterType((*DeleteAddressRequest)(nil), "api.DeleteAddressRequest")
	proto.RegisterType((*Empty)(nil), "api.Empty")
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e) }

var fileDescriptor_982c640dad8fe78e = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0xdd, 0x5d, 0x7f, 0xe4, 0x2b, 0x0d, 0x06, 0x03, 0xdb, 0x3a, 0xc8, 0x50, 0x26, 0x05,
	0x1e, 0xec, 0x50, 0x57, 0xcd, 0x08, 0x21, 0x22, 0xb6, 0x0e, 0x41, 0x50, 0x6c, 0xcd, 0x3b, 0x0c,
	0xe8, 0x3a, 0xed, 0x8c, 0xc1, 0xfe, 0x95, 0xfd, 0x4b, 0xe1, 0x38, 0x2b, 0x8c, 0xcd, 0x92, 0xb7,
	0xe1, 0xcd, 0xf3, 0xf3, 0xbe, 0xef, 0xe3, 0x2c, 0x34, 0x62, 0xc6, 0x52, 0x94, 0xb2, 0x2f, 0xd2,
	0xb9, 0x9a, 0x93, 0x20, 0x16, 0x9c, 0xbe, 0x42, 0xeb, 0x26, 0xc5, 0x58, 0xe1, 0x70, 0x75, 0x17,
	0xe1, 0xd7, 0x02, 0xa5, 0x22, 0x04, 0xca, 0x49, 0x3c, 0xc3, 0xb6, 0xd7, 0xf1, 0x7a, 0xf5, 0x48,
	0x9f, 0x97, 0x35, 0x95, 0x09, 0x6c, 0xfb, 0xab, 0xda, 0xf2, 0x4c, 0x8e, 0xa0, 0x2e, 0xe2, 0x14,
	0x13, 0xf5, 0xce, 0x59, 0x3b, 0xe8, 0x78, 0xbd, 0x4a, 0xb4, 0xb3, 0x2a, 0x4c, 0x18, 0x7d, 0x83,
	0x9a, 0xc1, 0x92, 0x26, 0xf8, 0x9c, 0x69, 0x5a, 0x25, 0xf2, 0x39, 0x5b, 0xf3, 0x7d, 0x07, 0x3f,
	0x28, 0xe2, 0x97, 0x37, 0xf8, 0x57, 0xb0, 0xbf, 0x8e, 0x2d, 0xc5, 0x3c, 0x91, 0x48, 0x4e, 0xa0,
	0x9a, 0xa2, 0x5c, 0x4c, 0x95, 0x9e, 0xb5, 0x3b, 0xd8, 0xeb, 0xc7, 0x82, 0xf7, 0xf3, 0x2e, 0x73,
	0x47, 0x5f, 0xe0, 0xf4, 0x9e, 0x4b, 0x65, 0xca, 0x28, 0x47, 0xd9, 0xa3, 0x61, 0x0e, 0x13, 0xf6,
	0x9c, 0x09, 0xcc, 0x35, 0x58, 0xe3, 0x3d, 0x7b, 0xbc, 0xcb, 0x07, 0x7d, 0x80, 0xee, 0x7f, 0x64,
	0x47, 0xd2, 0xa0, 0x30, 0xe9, 0x19, 0x1c, 0xdc, 0x61, 0x8e, 0x1b, 0x65, 0x13, 0x96, 0x27, 0xdb,
	0x10, 0x4a, 0xbb, 0xd0, 0x1a, 0xe3, 0x14, 0xff, 0xfc, 0x91, 0x9b, 0x7d, 0x35, 0xa8, 0xdc, 0xce,
	0x84, 0xca, 0x06, 0x3f, 0x3e, 0x34, 0x4d, 0xef, 0x13, 0xa6, 0xdf, 0xfc, 0x13, 0xc9, 0x08, 0x1a,
	0xd6, 0x63, 0x20, 0x87, 0x3a, 0x93, 0xeb, 0x81, 0x84, 0x2d, 0x2b, 0xae, 0x59, 0x8a, 0x96, 0x48,
	0x06, 0xc7, 0x4e, 0x01, 0x66, 0x7d, 0x72, 0xae, 0x7f, 0xb7, 0x95, 0xfd, 0xf0, 0x62, 0xab, 0xde,
	0xf5, 0xe8, 0x31, 0x34, 0x6d, 0x57, 0x24, 0xd4, 0x00, 0xa7, 0xc0, 0xc2, 0x05, 0xae, 0xa1, 0x61,
	0x89, 0x34, 0x12, 0x5c, 0x72, 0x43, 0xd0, 0x57, 0xda, 0x27, 0x2d, 0x7d, 0x54, 0xf5, 0x77, 0x75,
	0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x2a, 0x75, 0xce, 0x68, 0x03, 0x00, 0x00,
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
	ListAddressesByParentAndType(ctx context.Context, in *ListAddressesByParentIdAndTypeRequest, opts ...grpc.CallOption) (*ListAddressesByParentIdAndTypeResponse, error)
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

func (c *addressServiceClient) ListAddressesByParentAndType(ctx context.Context, in *ListAddressesByParentIdAndTypeRequest, opts ...grpc.CallOption) (*ListAddressesByParentIdAndTypeResponse, error) {
	out := new(ListAddressesByParentIdAndTypeResponse)
	err := c.cc.Invoke(ctx, "/api.AddressService/ListAddressesByParentAndType", in, out, opts...)
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
	ListAddressesByParentAndType(context.Context, *ListAddressesByParentIdAndTypeRequest) (*ListAddressesByParentIdAndTypeResponse, error)
	GetAddressById(context.Context, *GetAddressByIdRequest) (*AddressResponse, error)
	DeleteAddress(context.Context, *DeleteAddressRequest) (*Empty, error)
}

// UnimplementedAddressServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (*UnimplementedAddressServiceServer) CreateAddress(ctx context.Context, req *CreateAddressRequest) (*AddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (*UnimplementedAddressServiceServer) ListAddressesByParentAndType(ctx context.Context, req *ListAddressesByParentIdAndTypeRequest) (*ListAddressesByParentIdAndTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddressesByParentAndType not implemented")
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

func _AddressService_ListAddressesByParentAndType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAddressesByParentIdAndTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).ListAddressesByParentAndType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AddressService/ListAddressesByParentAndType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).ListAddressesByParentAndType(ctx, req.(*ListAddressesByParentIdAndTypeRequest))
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
			MethodName: "ListAddressesByParentAndType",
			Handler:    _AddressService_ListAddressesByParentAndType_Handler,
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
