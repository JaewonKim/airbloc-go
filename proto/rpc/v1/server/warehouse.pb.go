// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/server/warehouse.proto

package server

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type RawDataRequest struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	OwnerId              string   `protobuf:"bytes,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawDataRequest) Reset()         { *m = RawDataRequest{} }
func (m *RawDataRequest) String() string { return proto.CompactTextString(m) }
func (*RawDataRequest) ProtoMessage()    {}
func (*RawDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ee220902bb5d25, []int{0}
}

func (m *RawDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawDataRequest.Unmarshal(m, b)
}
func (m *RawDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawDataRequest.Marshal(b, m, deterministic)
}
func (m *RawDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawDataRequest.Merge(m, src)
}
func (m *RawDataRequest) XXX_Size() int {
	return xxx_messageInfo_RawDataRequest.Size(m)
}
func (m *RawDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RawDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RawDataRequest proto.InternalMessageInfo

func (m *RawDataRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *RawDataRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *RawDataRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type EncryptedDataRequest struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	OwnerId              string   `protobuf:"bytes,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	EncryptedPayload     []byte   `protobuf:"bytes,3,opt,name=encryptedPayload,proto3" json:"encryptedPayload,omitempty"`
	Capsule              []byte   `protobuf:"bytes,4,opt,name=capsule,proto3" json:"capsule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptedDataRequest) Reset()         { *m = EncryptedDataRequest{} }
func (m *EncryptedDataRequest) String() string { return proto.CompactTextString(m) }
func (*EncryptedDataRequest) ProtoMessage()    {}
func (*EncryptedDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ee220902bb5d25, []int{1}
}

func (m *EncryptedDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedDataRequest.Unmarshal(m, b)
}
func (m *EncryptedDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedDataRequest.Marshal(b, m, deterministic)
}
func (m *EncryptedDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedDataRequest.Merge(m, src)
}
func (m *EncryptedDataRequest) XXX_Size() int {
	return xxx_messageInfo_EncryptedDataRequest.Size(m)
}
func (m *EncryptedDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedDataRequest proto.InternalMessageInfo

func (m *EncryptedDataRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *EncryptedDataRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *EncryptedDataRequest) GetEncryptedPayload() []byte {
	if m != nil {
		return m.EncryptedPayload
	}
	return nil
}

func (m *EncryptedDataRequest) GetCapsule() []byte {
	if m != nil {
		return m.Capsule
	}
	return nil
}

type StoreResult struct {
	BundleId             string   `protobuf:"bytes,1,opt,name=bundleId,proto3" json:"bundleId,omitempty"`
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	DataCount            uint64   `protobuf:"varint,3,opt,name=dataCount,proto3" json:"dataCount,omitempty"`
	GasUsed              uint64   `protobuf:"varint,4,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResult) Reset()         { *m = StoreResult{} }
func (m *StoreResult) String() string { return proto.CompactTextString(m) }
func (*StoreResult) ProtoMessage()    {}
func (*StoreResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ee220902bb5d25, []int{2}
}

func (m *StoreResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResult.Unmarshal(m, b)
}
func (m *StoreResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResult.Marshal(b, m, deterministic)
}
func (m *StoreResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResult.Merge(m, src)
}
func (m *StoreResult) XXX_Size() int {
	return xxx_messageInfo_StoreResult.Size(m)
}
func (m *StoreResult) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResult.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResult proto.InternalMessageInfo

func (m *StoreResult) GetBundleId() string {
	if m != nil {
		return m.BundleId
	}
	return ""
}

func (m *StoreResult) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *StoreResult) GetDataCount() uint64 {
	if m != nil {
		return m.DataCount
	}
	return 0
}

func (m *StoreResult) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

type DeleteBundleRequest struct {
	Uri                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBundleRequest) Reset()         { *m = DeleteBundleRequest{} }
func (m *DeleteBundleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBundleRequest) ProtoMessage()    {}
func (*DeleteBundleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ee220902bb5d25, []int{3}
}

func (m *DeleteBundleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBundleRequest.Unmarshal(m, b)
}
func (m *DeleteBundleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBundleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBundleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBundleRequest.Merge(m, src)
}
func (m *DeleteBundleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBundleRequest.Size(m)
}
func (m *DeleteBundleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBundleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBundleRequest proto.InternalMessageInfo

func (m *DeleteBundleRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type DeleteBundleResult struct {
	GasUsed              uint64   `protobuf:"varint,1,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBundleResult) Reset()         { *m = DeleteBundleResult{} }
func (m *DeleteBundleResult) String() string { return proto.CompactTextString(m) }
func (*DeleteBundleResult) ProtoMessage()    {}
func (*DeleteBundleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ee220902bb5d25, []int{4}
}

func (m *DeleteBundleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBundleResult.Unmarshal(m, b)
}
func (m *DeleteBundleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBundleResult.Marshal(b, m, deterministic)
}
func (m *DeleteBundleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBundleResult.Merge(m, src)
}
func (m *DeleteBundleResult) XXX_Size() int {
	return xxx_messageInfo_DeleteBundleResult.Size(m)
}
func (m *DeleteBundleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBundleResult.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBundleResult proto.InternalMessageInfo

func (m *DeleteBundleResult) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func init() {
	proto.RegisterType((*RawDataRequest)(nil), "airbloc.rpc.v1.RawDataRequest")
	proto.RegisterType((*EncryptedDataRequest)(nil), "airbloc.rpc.v1.EncryptedDataRequest")
	proto.RegisterType((*StoreResult)(nil), "airbloc.rpc.v1.StoreResult")
	proto.RegisterType((*DeleteBundleRequest)(nil), "airbloc.rpc.v1.DeleteBundleRequest")
	proto.RegisterType((*DeleteBundleResult)(nil), "airbloc.rpc.v1.DeleteBundleResult")
}

func init() {
	proto.RegisterFile("proto/rpc/v1/server/warehouse.proto", fileDescriptor_e7ee220902bb5d25)
}

var fileDescriptor_e7ee220902bb5d25 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x5d, 0x8b, 0xda, 0x40,
	0x14, 0x25, 0x2a, 0x6d, 0xbd, 0x15, 0x91, 0xa9, 0x0f, 0xc1, 0x16, 0x29, 0xb1, 0x50, 0x29, 0x34,
	0xc1, 0xfa, 0x0f, 0xac, 0x7d, 0x10, 0xfa, 0x50, 0x52, 0x16, 0x61, 0xf7, 0x69, 0x32, 0x73, 0xd1,
	0xc0, 0x6c, 0x26, 0x3b, 0x1f, 0x8a, 0xbf, 0x62, 0xdf, 0xf6, 0xf7, 0x2e, 0x19, 0x13, 0x37, 0x71,
	0x65, 0xf7, 0x65, 0x9f, 0x92, 0x7b, 0xcf, 0x99, 0x73, 0xee, 0xdc, 0xc3, 0xc0, 0x24, 0x57, 0xd2,
	0xc8, 0x48, 0xe5, 0x2c, 0xda, 0xcd, 0x22, 0x8d, 0x6a, 0x87, 0x2a, 0xda, 0x53, 0x85, 0x5b, 0x69,
	0x35, 0x86, 0x0e, 0x25, 0x7d, 0x9a, 0xaa, 0x44, 0x48, 0x16, 0xaa, 0x9c, 0x85, 0xbb, 0x59, 0xc0,
	0xa1, 0x1f, 0xd3, 0xfd, 0x92, 0x1a, 0x1a, 0xe3, 0x9d, 0x45, 0x6d, 0xc8, 0x18, 0x80, 0x49, 0x21,
	0x90, 0x99, 0x54, 0x66, 0xbe, 0xf7, 0xd5, 0x9b, 0x76, 0xe3, 0x5a, 0x87, 0xf8, 0xf0, 0x5e, 0xee,
	0x33, 0x54, 0x2b, 0xee, 0xb7, 0x1c, 0x58, 0x95, 0x05, 0x92, 0xd3, 0x83, 0x90, 0x94, 0xfb, 0xed,
	0x23, 0x52, 0x96, 0xc1, 0x83, 0x07, 0xc3, 0x3f, 0x19, 0x53, 0x87, 0xdc, 0x20, 0x7f, 0x1b, 0xb3,
	0x1f, 0x30, 0xc0, 0x4a, 0xf1, 0x5f, 0xcd, 0xb5, 0x17, 0x3f, 0xeb, 0x17, 0x2a, 0x8c, 0xe6, 0xda,
	0x0a, 0xf4, 0x3b, 0x8e, 0x52, 0x95, 0x81, 0x86, 0x8f, 0xff, 0x8d, 0x54, 0x18, 0xa3, 0xb6, 0xc2,
	0x90, 0x11, 0x7c, 0x48, 0x6c, 0xc6, 0x05, 0xae, 0x78, 0x39, 0xcc, 0xa9, 0x26, 0x03, 0x68, 0x5b,
	0x95, 0x96, 0x63, 0x14, 0xbf, 0xe4, 0x0b, 0x74, 0x39, 0x35, 0xf4, 0xb7, 0xb4, 0x99, 0x71, 0xde,
	0x9d, 0xf8, 0xa9, 0x51, 0x98, 0x6e, 0xa8, 0xbe, 0xd2, 0xc8, 0x9d, 0x69, 0x27, 0xae, 0xca, 0xe0,
	0x3b, 0x7c, 0x5a, 0xa2, 0x40, 0x83, 0x0b, 0xa7, 0x5d, 0xed, 0xa2, 0x34, 0xf0, 0x4e, 0x06, 0x41,
	0x08, 0xa4, 0x49, 0x74, 0x43, 0xd6, 0x84, 0xbd, 0x86, 0xf0, 0xaf, 0xfb, 0x16, 0x74, 0xd7, 0x55,
	0xe0, 0xe4, 0x6f, 0x79, 0xb7, 0xe3, 0x61, 0x32, 0x0e, 0x9b, 0xd1, 0x87, 0xcd, 0xdc, 0x47, 0x9f,
	0xcf, 0xf1, 0xda, 0x62, 0xa6, 0x1e, 0xb9, 0x81, 0xa1, 0x6b, 0x9c, 0x62, 0x2c, 0x65, 0xbf, 0x9d,
	0x1f, 0xbb, 0x94, 0xf3, 0x6b, 0xe2, 0x6b, 0xe8, 0xd5, 0x2f, 0x4a, 0x26, 0xe7, 0xf4, 0x0b, 0xfb,
	0x1a, 0x05, 0x2f, 0x93, 0x0a, 0xe9, 0xc5, 0xfc, 0x7a, 0xb6, 0x49, 0xcd, 0xd6, 0x26, 0x21, 0x93,
	0xb7, 0x51, 0xc9, 0xaf, 0xbe, 0x3f, 0x37, 0x32, 0xba, 0xf0, 0x66, 0x92, 0x77, 0xae, 0x39, 0x7f,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x86, 0xe3, 0xf9, 0x51, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WarehouseClient is the client API for Warehouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WarehouseClient interface {
	StoreBundle(ctx context.Context, opts ...grpc.CallOption) (Warehouse_StoreBundleClient, error)
	StoreEncryptedBundle(ctx context.Context, opts ...grpc.CallOption) (Warehouse_StoreEncryptedBundleClient, error)
	DeleteBundle(ctx context.Context, in *DeleteBundleRequest, opts ...grpc.CallOption) (*DeleteBundleResult, error)
}

type warehouseClient struct {
	cc *grpc.ClientConn
}

func NewWarehouseClient(cc *grpc.ClientConn) WarehouseClient {
	return &warehouseClient{cc}
}

func (c *warehouseClient) StoreBundle(ctx context.Context, opts ...grpc.CallOption) (Warehouse_StoreBundleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Warehouse_serviceDesc.Streams[0], "/airbloc.rpc.v1.Warehouse/StoreBundle", opts...)
	if err != nil {
		return nil, err
	}
	x := &warehouseStoreBundleClient{stream}
	return x, nil
}

type Warehouse_StoreBundleClient interface {
	Send(*RawDataRequest) error
	CloseAndRecv() (*StoreResult, error)
	grpc.ClientStream
}

type warehouseStoreBundleClient struct {
	grpc.ClientStream
}

func (x *warehouseStoreBundleClient) Send(m *RawDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *warehouseStoreBundleClient) CloseAndRecv() (*StoreResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StoreResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *warehouseClient) StoreEncryptedBundle(ctx context.Context, opts ...grpc.CallOption) (Warehouse_StoreEncryptedBundleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Warehouse_serviceDesc.Streams[1], "/airbloc.rpc.v1.Warehouse/StoreEncryptedBundle", opts...)
	if err != nil {
		return nil, err
	}
	x := &warehouseStoreEncryptedBundleClient{stream}
	return x, nil
}

type Warehouse_StoreEncryptedBundleClient interface {
	Send(*EncryptedDataRequest) error
	CloseAndRecv() (*StoreResult, error)
	grpc.ClientStream
}

type warehouseStoreEncryptedBundleClient struct {
	grpc.ClientStream
}

func (x *warehouseStoreEncryptedBundleClient) Send(m *EncryptedDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *warehouseStoreEncryptedBundleClient) CloseAndRecv() (*StoreResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StoreResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *warehouseClient) DeleteBundle(ctx context.Context, in *DeleteBundleRequest, opts ...grpc.CallOption) (*DeleteBundleResult, error) {
	out := new(DeleteBundleResult)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Warehouse/DeleteBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseServer is the server API for Warehouse service.
type WarehouseServer interface {
	StoreBundle(Warehouse_StoreBundleServer) error
	StoreEncryptedBundle(Warehouse_StoreEncryptedBundleServer) error
	DeleteBundle(context.Context, *DeleteBundleRequest) (*DeleteBundleResult, error)
}

func RegisterWarehouseServer(s *grpc.Server, srv WarehouseServer) {
	s.RegisterService(&_Warehouse_serviceDesc, srv)
}

func _Warehouse_StoreBundle_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WarehouseServer).StoreBundle(&warehouseStoreBundleServer{stream})
}

type Warehouse_StoreBundleServer interface {
	SendAndClose(*StoreResult) error
	Recv() (*RawDataRequest, error)
	grpc.ServerStream
}

type warehouseStoreBundleServer struct {
	grpc.ServerStream
}

func (x *warehouseStoreBundleServer) SendAndClose(m *StoreResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *warehouseStoreBundleServer) Recv() (*RawDataRequest, error) {
	m := new(RawDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Warehouse_StoreEncryptedBundle_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WarehouseServer).StoreEncryptedBundle(&warehouseStoreEncryptedBundleServer{stream})
}

type Warehouse_StoreEncryptedBundleServer interface {
	SendAndClose(*StoreResult) error
	Recv() (*EncryptedDataRequest, error)
	grpc.ServerStream
}

type warehouseStoreEncryptedBundleServer struct {
	grpc.ServerStream
}

func (x *warehouseStoreEncryptedBundleServer) SendAndClose(m *StoreResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *warehouseStoreEncryptedBundleServer) Recv() (*EncryptedDataRequest, error) {
	m := new(EncryptedDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Warehouse_DeleteBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).DeleteBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Warehouse/DeleteBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).DeleteBundle(ctx, req.(*DeleteBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Warehouse_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.rpc.v1.Warehouse",
	HandlerType: (*WarehouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteBundle",
			Handler:    _Warehouse_DeleteBundle_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StoreBundle",
			Handler:       _Warehouse_StoreBundle_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StoreEncryptedBundle",
			Handler:       _Warehouse_StoreEncryptedBundle_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/rpc/v1/server/warehouse.proto",
}