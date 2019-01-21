// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/server/warehouse.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type RawDataRequest struct {
	// ID of the collection
	CollectionId string `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	// ID of the user (Data Owner)
	OwnerId string `protobuf:"bytes,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	// row ID - merkle tree leaf
	RowId string `protobuf:"bytes,3,opt,name=rowId,proto3" json:"rowId,omitempty"`
	// JSON payload, which follows schema of the given collection.
	Payload              string   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
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

func (m *RawDataRequest) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *RawDataRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *RawDataRequest) GetRowId() string {
	if m != nil {
		return m.RowId
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
	// ID of the collection
	CollectionId string `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	// ID of the user (Data Owner)
	OwnerId string `protobuf:"bytes,2,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	// row ID - merkle tree leaf
	RowId string `protobuf:"bytes,3,opt,name=rowId,proto3" json:"rowId,omitempty"`
	// pre-encrypted JSON payload, which follows schema of the given collection.
	// the payload must be encrypted through ECIES-SECP256k1 using the key in a capsule.
	EncryptedPayload []byte `protobuf:"bytes,4,opt,name=encryptedPayload,proto3" json:"encryptedPayload,omitempty"`
	// symmetric key of the encryptedPayload, encrypted using the provider's key.
	Capsule              []byte   `protobuf:"bytes,5,opt,name=capsule,proto3" json:"capsule,omitempty"`
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

func (m *EncryptedDataRequest) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *EncryptedDataRequest) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *EncryptedDataRequest) GetRowId() string {
	if m != nil {
		return m.RowId
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
	// the ID of the bundle in "{empty}{bundleid}" format.
	BundleId string `protobuf:"bytes,1,opt,name=bundleId,proto3" json:"bundleId,omitempty"`
	// public-accessible URI of the bundle.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// count of the data that have been successfully ingested.
	// if some of your data is being filtered by some reason (e.g. DAuth, Schema Validation),
	// then this count may lower than the amout you've ingested.
	DataCount uint64 `protobuf:"varint,3,opt,name=dataCount,proto3" json:"dataCount,omitempty"`
	// amount of gas (transaction fee) used for registration to blockchain.
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
	// public-accessible URI of the bundle.
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
	// amount of gas (transaction fee) used for interaction to blockchain.
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

type ListBundleRequest struct {
	// Provider (App) ID
	ProviderId string `protobuf:"bytes,1,opt,name=providerId,proto3" json:"providerId,omitempty"`
	// Optional: filter specific collection ID from results
	CollectionId         string   `protobuf:"bytes,2,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBundleRequest) Reset()         { *m = ListBundleRequest{} }
func (m *ListBundleRequest) String() string { return proto.CompactTextString(m) }
func (*ListBundleRequest) ProtoMessage()    {}
func (*ListBundleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ee220902bb5d25, []int{5}
}

func (m *ListBundleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBundleRequest.Unmarshal(m, b)
}
func (m *ListBundleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBundleRequest.Marshal(b, m, deterministic)
}
func (m *ListBundleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBundleRequest.Merge(m, src)
}
func (m *ListBundleRequest) XXX_Size() int {
	return xxx_messageInfo_ListBundleRequest.Size(m)
}
func (m *ListBundleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBundleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBundleRequest proto.InternalMessageInfo

func (m *ListBundleRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *ListBundleRequest) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

type ListBundleResult struct {
	Bundles              []*ListBundleResult_Bundle `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListBundleResult) Reset()         { *m = ListBundleResult{} }
func (m *ListBundleResult) String() string { return proto.CompactTextString(m) }
func (*ListBundleResult) ProtoMessage()    {}
func (*ListBundleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ee220902bb5d25, []int{6}
}

func (m *ListBundleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBundleResult.Unmarshal(m, b)
}
func (m *ListBundleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBundleResult.Marshal(b, m, deterministic)
}
func (m *ListBundleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBundleResult.Merge(m, src)
}
func (m *ListBundleResult) XXX_Size() int {
	return xxx_messageInfo_ListBundleResult.Size(m)
}
func (m *ListBundleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBundleResult.DiscardUnknown(m)
}

var xxx_messageInfo_ListBundleResult proto.InternalMessageInfo

func (m *ListBundleResult) GetBundles() []*ListBundleResult_Bundle {
	if m != nil {
		return m.Bundles
	}
	return nil
}

type ListBundleResult_Bundle struct {
	CollectionId         string   `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	CreatedAt            uint64   `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	DataCount            uint64   `protobuf:"varint,4,opt,name=dataCount,proto3" json:"dataCount,omitempty"`
	Uri                  string   `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBundleResult_Bundle) Reset()         { *m = ListBundleResult_Bundle{} }
func (m *ListBundleResult_Bundle) String() string { return proto.CompactTextString(m) }
func (*ListBundleResult_Bundle) ProtoMessage()    {}
func (*ListBundleResult_Bundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7ee220902bb5d25, []int{6, 0}
}

func (m *ListBundleResult_Bundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBundleResult_Bundle.Unmarshal(m, b)
}
func (m *ListBundleResult_Bundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBundleResult_Bundle.Marshal(b, m, deterministic)
}
func (m *ListBundleResult_Bundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBundleResult_Bundle.Merge(m, src)
}
func (m *ListBundleResult_Bundle) XXX_Size() int {
	return xxx_messageInfo_ListBundleResult_Bundle.Size(m)
}
func (m *ListBundleResult_Bundle) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBundleResult_Bundle.DiscardUnknown(m)
}

var xxx_messageInfo_ListBundleResult_Bundle proto.InternalMessageInfo

func (m *ListBundleResult_Bundle) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *ListBundleResult_Bundle) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ListBundleResult_Bundle) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *ListBundleResult_Bundle) GetDataCount() uint64 {
	if m != nil {
		return m.DataCount
	}
	return 0
}

func (m *ListBundleResult_Bundle) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func init() {
	proto.RegisterType((*RawDataRequest)(nil), "airbloc.rpc.v1.RawDataRequest")
	proto.RegisterType((*EncryptedDataRequest)(nil), "airbloc.rpc.v1.EncryptedDataRequest")
	proto.RegisterType((*StoreResult)(nil), "airbloc.rpc.v1.StoreResult")
	proto.RegisterType((*DeleteBundleRequest)(nil), "airbloc.rpc.v1.DeleteBundleRequest")
	proto.RegisterType((*DeleteBundleResult)(nil), "airbloc.rpc.v1.DeleteBundleResult")
	proto.RegisterType((*ListBundleRequest)(nil), "airbloc.rpc.v1.ListBundleRequest")
	proto.RegisterType((*ListBundleResult)(nil), "airbloc.rpc.v1.ListBundleResult")
	proto.RegisterType((*ListBundleResult_Bundle)(nil), "airbloc.rpc.v1.ListBundleResult.Bundle")
}

func init() {
	proto.RegisterFile("proto/rpc/v1/server/warehouse.proto", fileDescriptor_e7ee220902bb5d25)
}

var fileDescriptor_e7ee220902bb5d25 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0xb4, 0x25, 0xd3, 0x50, 0x85, 0x25, 0x07, 0x2b, 0xa0, 0x2a, 0xb8, 0x48, 0x8d,
	0x90, 0xb0, 0x95, 0xf6, 0x0b, 0x5a, 0xca, 0xa1, 0x52, 0x0f, 0x60, 0x84, 0x22, 0xc1, 0x69, 0xb3,
	0x3b, 0x4a, 0x2c, 0x19, 0xaf, 0xd9, 0x5d, 0x27, 0xf4, 0x2f, 0x38, 0xf3, 0x79, 0x9c, 0xf8, 0x0c,
	0xe4, 0xf5, 0xda, 0xb5, 0x9d, 0xa8, 0xe1, 0xc0, 0x29, 0x99, 0x37, 0xb3, 0x6f, 0x26, 0xf3, 0xde,
	0x04, 0xce, 0x52, 0x29, 0xb4, 0x08, 0x64, 0xca, 0x82, 0xf5, 0x2c, 0x50, 0x28, 0xd7, 0x28, 0x83,
	0x0d, 0x95, 0xb8, 0x12, 0x99, 0x42, 0xdf, 0x64, 0xc9, 0x09, 0x8d, 0xe4, 0x22, 0x16, 0xcc, 0x97,
	0x29, 0xf3, 0xd7, 0x33, 0x6f, 0x05, 0x27, 0x21, 0xdd, 0xdc, 0x50, 0x4d, 0x43, 0xfc, 0x9e, 0xa1,
	0xd2, 0xc4, 0x83, 0x01, 0x13, 0x71, 0x8c, 0x4c, 0x47, 0x22, 0xb9, 0xe5, 0xae, 0x33, 0x71, 0xa6,
	0xfd, 0xb0, 0x81, 0x11, 0x17, 0x8e, 0xc4, 0x26, 0x41, 0x79, 0xcb, 0xdd, 0x8e, 0x49, 0x97, 0x61,
	0x9e, 0x49, 0xe9, 0x7d, 0x2c, 0x28, 0x77, 0xbb, 0x45, 0xc6, 0x86, 0xde, 0x2f, 0x07, 0x46, 0xef,
	0x13, 0x26, 0xef, 0x53, 0x8d, 0xfc, 0xff, 0x35, 0x7c, 0x03, 0x43, 0x2c, 0x59, 0x3f, 0xd4, 0x3a,
	0x0f, 0xc2, 0x2d, 0x3c, 0x67, 0x61, 0x34, 0x55, 0x59, 0x8c, 0x6e, 0xcf, 0x94, 0x94, 0xa1, 0xa7,
	0xe0, 0xf8, 0x93, 0x16, 0x12, 0x43, 0x54, 0x59, 0xac, 0xc9, 0x18, 0x9e, 0x2c, 0xb2, 0x84, 0xc7,
	0x58, 0x8d, 0x53, 0xc5, 0x64, 0x08, 0xdd, 0x4c, 0x46, 0x76, 0x8c, 0xfc, 0x2b, 0x79, 0x09, 0x7d,
	0x4e, 0x35, 0x7d, 0x27, 0xb2, 0x44, 0x9b, 0xde, 0xbd, 0xf0, 0x01, 0xc8, 0x9b, 0x2e, 0xa9, 0xfa,
	0xac, 0x90, 0x9b, 0xa6, 0xbd, 0xb0, 0x0c, 0xbd, 0x73, 0x78, 0x7e, 0x83, 0x31, 0x6a, 0xbc, 0x36,
	0xdc, 0xe5, 0x3e, 0x6c, 0x03, 0xa7, 0x6a, 0xe0, 0xf9, 0x40, 0x9a, 0x85, 0x66, 0xc8, 0x1a, 0xb1,
	0xd3, 0x24, 0x9e, 0xc3, 0xb3, 0xbb, 0x48, 0xe9, 0x26, 0xed, 0x29, 0x40, 0x2a, 0xc5, 0x3a, 0xe2,
	0x66, 0x8b, 0x05, 0x7b, 0x0d, 0xd9, 0x92, 0xa1, 0xb3, 0x2d, 0x83, 0xf7, 0xc7, 0x81, 0x61, 0x9d,
	0xd9, 0xcc, 0x71, 0x05, 0x47, 0xc5, 0x72, 0x94, 0xeb, 0x4c, 0xba, 0xd3, 0xe3, 0x8b, 0x73, 0xbf,
	0x69, 0x32, 0xbf, 0xfd, 0xc4, 0xb7, 0x41, 0xf9, 0x6e, 0xfc, 0xd3, 0x81, 0xc3, 0x02, 0xfb, 0x27,
	0x37, 0x8c, 0xe0, 0x20, 0x4a, 0x38, 0xfe, 0x30, 0x33, 0x3e, 0x0d, 0x8b, 0x20, 0x97, 0x81, 0x49,
	0xa4, 0x1a, 0xf9, 0x55, 0x25, 0x43, 0x05, 0x34, 0x45, 0xea, 0xb5, 0x45, 0xb2, 0x3b, 0x3f, 0xa8,
	0x76, 0x7e, 0xf1, 0xbb, 0x03, 0xfd, 0x79, 0x79, 0x3c, 0xe4, 0xce, 0xfa, 0xc3, 0x0e, 0x79, 0xda,
	0xfe, 0x85, 0xcd, 0x1b, 0x1a, 0xbf, 0x68, 0xe7, 0x6b, 0xe6, 0x9a, 0x3a, 0xe4, 0x2b, 0x8c, 0x0c,
	0x50, 0x9d, 0x83, 0xa5, 0x7d, 0xdd, 0x7e, 0xb6, 0xeb, 0x5e, 0xf6, 0x91, 0xcf, 0x61, 0x50, 0x37,
	0x0b, 0x39, 0x6b, 0x97, 0xef, 0xf0, 0xdc, 0xd8, 0x7b, 0xbc, 0xc8, 0xe8, 0xfc, 0x11, 0xe0, 0x41,
	0x48, 0xf2, 0xea, 0x31, 0x91, 0x0b, 0xd2, 0xc9, 0x3e, 0x1f, 0x5c, 0x5f, 0x7e, 0x99, 0x2d, 0x23,
	0xbd, 0xca, 0x16, 0x3e, 0x13, 0xdf, 0x02, 0x5b, 0x5d, 0x7e, 0xbe, 0x5d, 0x8a, 0x60, 0xc7, 0x5f,
	0xda, 0xe2, 0xd0, 0x80, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x72, 0x9e, 0x0d, 0x36, 0xf0,
	0x04, 0x00, 0x00,
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
	//*
	// StoreBundle receives a data stream, encrypts and stores them as a bundle,
	// and registers the data bundle onto the blockchain.
	StoreBundle(ctx context.Context, opts ...grpc.CallOption) (Warehouse_StoreBundleClient, error)
	//*
	// StoreBundle receives a stream of the data already encrypted, stores them as a bundle,
	// and registers the data bundle onto the blockchain.
	StoreEncryptedBundle(ctx context.Context, opts ...grpc.CallOption) (Warehouse_StoreEncryptedBundleClient, error)
	//*
	// DeleteBundle removes given bundle from the warehouse and the blockchain.
	DeleteBundle(ctx context.Context, in *DeleteBundleRequest, opts ...grpc.CallOption) (*DeleteBundleResult, error)
	//*
	// ListBundle returns a list of ingested data bundles.
	ListBundle(ctx context.Context, in *ListBundleRequest, opts ...grpc.CallOption) (*ListBundleResult, error)
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

func (c *warehouseClient) ListBundle(ctx context.Context, in *ListBundleRequest, opts ...grpc.CallOption) (*ListBundleResult, error) {
	out := new(ListBundleResult)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.Warehouse/ListBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseServer is the server API for Warehouse service.
type WarehouseServer interface {
	//*
	// StoreBundle receives a data stream, encrypts and stores them as a bundle,
	// and registers the data bundle onto the blockchain.
	StoreBundle(Warehouse_StoreBundleServer) error
	//*
	// StoreBundle receives a stream of the data already encrypted, stores them as a bundle,
	// and registers the data bundle onto the blockchain.
	StoreEncryptedBundle(Warehouse_StoreEncryptedBundleServer) error
	//*
	// DeleteBundle removes given bundle from the warehouse and the blockchain.
	DeleteBundle(context.Context, *DeleteBundleRequest) (*DeleteBundleResult, error)
	//*
	// ListBundle returns a list of ingested data bundles.
	ListBundle(context.Context, *ListBundleRequest) (*ListBundleResult, error)
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

func _Warehouse_ListBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).ListBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.Warehouse/ListBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).ListBundle(ctx, req.(*ListBundleRequest))
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
		{
			MethodName: "ListBundle",
			Handler:    _Warehouse_ListBundle_Handler,
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
