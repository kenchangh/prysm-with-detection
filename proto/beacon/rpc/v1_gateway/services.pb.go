// Code generated by protoc-gen-go. DO NOT EDIT.
// proto/beacon/rpc/v1/services.proto is a deprecated file.

package ethereum_beacon_rpc_v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ValidatorRole int32 // Deprecated: Do not use.
const (
	ValidatorRole_UNKNOWN    ValidatorRole = 0
	ValidatorRole_ATTESTER   ValidatorRole = 1
	ValidatorRole_PROPOSER   ValidatorRole = 2
	ValidatorRole_AGGREGATOR ValidatorRole = 3
)

var ValidatorRole_name = map[int32]string{
	0: "UNKNOWN",
	1: "ATTESTER",
	2: "PROPOSER",
	3: "AGGREGATOR",
}

var ValidatorRole_value = map[string]int32{
	"UNKNOWN":    0,
	"ATTESTER":   1,
	"PROPOSER":   2,
	"AGGREGATOR": 3,
}

func (x ValidatorRole) String() string {
	return proto.EnumName(ValidatorRole_name, int32(x))
}

func (ValidatorRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9eb4e94b85965285, []int{0}
}

// Deprecated: Do not use.
type AggregationRequest struct {
	Slot                 uint64   `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	CommitteeIndex       uint64   `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	SlotSignature        []byte   `protobuf:"bytes,4,opt,name=slot_signature,json=slotSignature,proto3" json:"slot_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregationRequest) Reset()         { *m = AggregationRequest{} }
func (m *AggregationRequest) String() string { return proto.CompactTextString(m) }
func (*AggregationRequest) ProtoMessage()    {}
func (*AggregationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb4e94b85965285, []int{0}
}

func (m *AggregationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregationRequest.Unmarshal(m, b)
}
func (m *AggregationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregationRequest.Marshal(b, m, deterministic)
}
func (m *AggregationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregationRequest.Merge(m, src)
}
func (m *AggregationRequest) XXX_Size() int {
	return xxx_messageInfo_AggregationRequest.Size(m)
}
func (m *AggregationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AggregationRequest proto.InternalMessageInfo

func (m *AggregationRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *AggregationRequest) GetCommitteeIndex() uint64 {
	if m != nil {
		return m.CommitteeIndex
	}
	return 0
}

func (m *AggregationRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *AggregationRequest) GetSlotSignature() []byte {
	if m != nil {
		return m.SlotSignature
	}
	return nil
}

// Deprecated: Do not use.
type AggregationResponse struct {
	Root                 []byte   `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregationResponse) Reset()         { *m = AggregationResponse{} }
func (m *AggregationResponse) String() string { return proto.CompactTextString(m) }
func (*AggregationResponse) ProtoMessage()    {}
func (*AggregationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb4e94b85965285, []int{1}
}

func (m *AggregationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregationResponse.Unmarshal(m, b)
}
func (m *AggregationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregationResponse.Marshal(b, m, deterministic)
}
func (m *AggregationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregationResponse.Merge(m, src)
}
func (m *AggregationResponse) XXX_Size() int {
	return xxx_messageInfo_AggregationResponse.Size(m)
}
func (m *AggregationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AggregationResponse proto.InternalMessageInfo

func (m *AggregationResponse) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

// Deprecated: Do not use.
type ExitedValidatorsRequest struct {
	PublicKeys           [][]byte `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitedValidatorsRequest) Reset()         { *m = ExitedValidatorsRequest{} }
func (m *ExitedValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*ExitedValidatorsRequest) ProtoMessage()    {}
func (*ExitedValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb4e94b85965285, []int{2}
}

func (m *ExitedValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitedValidatorsRequest.Unmarshal(m, b)
}
func (m *ExitedValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitedValidatorsRequest.Marshal(b, m, deterministic)
}
func (m *ExitedValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitedValidatorsRequest.Merge(m, src)
}
func (m *ExitedValidatorsRequest) XXX_Size() int {
	return xxx_messageInfo_ExitedValidatorsRequest.Size(m)
}
func (m *ExitedValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitedValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExitedValidatorsRequest proto.InternalMessageInfo

func (m *ExitedValidatorsRequest) GetPublicKeys() [][]byte {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

// Deprecated: Do not use.
type ExitedValidatorsResponse struct {
	PublicKeys           [][]byte `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitedValidatorsResponse) Reset()         { *m = ExitedValidatorsResponse{} }
func (m *ExitedValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*ExitedValidatorsResponse) ProtoMessage()    {}
func (*ExitedValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eb4e94b85965285, []int{3}
}

func (m *ExitedValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitedValidatorsResponse.Unmarshal(m, b)
}
func (m *ExitedValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitedValidatorsResponse.Marshal(b, m, deterministic)
}
func (m *ExitedValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitedValidatorsResponse.Merge(m, src)
}
func (m *ExitedValidatorsResponse) XXX_Size() int {
	return xxx_messageInfo_ExitedValidatorsResponse.Size(m)
}
func (m *ExitedValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitedValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExitedValidatorsResponse proto.InternalMessageInfo

func (m *ExitedValidatorsResponse) GetPublicKeys() [][]byte {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func init() {
	proto.RegisterEnum("ethereum.beacon.rpc.v1.ValidatorRole", ValidatorRole_name, ValidatorRole_value)
	proto.RegisterType((*AggregationRequest)(nil), "ethereum.beacon.rpc.v1.AggregationRequest")
	proto.RegisterType((*AggregationResponse)(nil), "ethereum.beacon.rpc.v1.AggregationResponse")
	proto.RegisterType((*ExitedValidatorsRequest)(nil), "ethereum.beacon.rpc.v1.ExitedValidatorsRequest")
	proto.RegisterType((*ExitedValidatorsResponse)(nil), "ethereum.beacon.rpc.v1.ExitedValidatorsResponse")
}

func init() { proto.RegisterFile("proto/beacon/rpc/v1/services.proto", fileDescriptor_9eb4e94b85965285) }

var fileDescriptor_9eb4e94b85965285 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x75, 0x92, 0xe2, 0xc7, 0xdd, 0x6e, 0xad, 0x23, 0xb8, 0x43, 0x45, 0x2c, 0x05, 0xb1, 0xac,
	0x98, 0xb0, 0xfa, 0xe6, 0x83, 0x12, 0x21, 0x14, 0x59, 0x69, 0xcb, 0xa4, 0xea, 0x63, 0xc9, 0xc7,
	0xdd, 0x38, 0x98, 0x64, 0xe2, 0xcc, 0xa4, 0x6c, 0xff, 0x81, 0x0f, 0xfe, 0x03, 0xff, 0x88, 0x3f,
	0x4f, 0x32, 0xd9, 0x54, 0x64, 0x7d, 0xd8, 0xb7, 0x3b, 0xe7, 0x9e, 0x73, 0xe7, 0x1c, 0xee, 0x85,
	0x59, 0xad, 0xa4, 0x91, 0x7e, 0x82, 0x71, 0x2a, 0x2b, 0x5f, 0xd5, 0xa9, 0xbf, 0x3b, 0xf3, 0x35,
	0xaa, 0x9d, 0x48, 0x51, 0x7b, 0xb6, 0x49, 0x1f, 0xa1, 0xf9, 0x8a, 0x0a, 0x9b, 0xd2, 0xeb, 0x68,
	0x9e, 0xaa, 0x53, 0x6f, 0x77, 0x36, 0x79, 0x9c, 0x4b, 0x99, 0x17, 0xe8, 0x5b, 0x56, 0xd2, 0x5c,
	0xf8, 0x58, 0xd6, 0x66, 0xdf, 0x89, 0x66, 0xbf, 0x08, 0xd0, 0x20, 0xcf, 0x15, 0xe6, 0xb1, 0x11,
	0xb2, 0xe2, 0xf8, 0xbd, 0x41, 0x6d, 0x28, 0x85, 0x81, 0x2e, 0xa4, 0x61, 0x64, 0x4a, 0xe6, 0x03,
	0x6e, 0x6b, 0xfa, 0x1c, 0xee, 0xa7, 0xb2, 0x2c, 0x85, 0x31, 0x88, 0x5b, 0x51, 0x65, 0x78, 0xc9,
	0x1c, 0xdb, 0x1e, 0x1d, 0xe0, 0x0f, 0x2d, 0x4a, 0x9f, 0x00, 0xd4, 0x4d, 0x52, 0x88, 0x74, 0xfb,
	0x0d, 0xf7, 0xcc, 0x9d, 0x92, 0xf9, 0x90, 0xdf, 0xeb, 0x90, 0x73, 0xdc, 0xd3, 0x67, 0x30, 0x6a,
	0xe7, 0x6d, 0xb5, 0xc8, 0xab, 0xd8, 0x34, 0x0a, 0xd9, 0xc0, 0x52, 0x8e, 0x5b, 0x34, 0xea, 0xc1,
	0x37, 0x0e, 0x23, 0xb3, 0x97, 0xf0, 0xf0, 0x1f, 0x73, 0xba, 0x96, 0x95, 0xc6, 0xd6, 0x9d, 0x92,
	0x57, 0xee, 0x86, 0xdc, 0xd6, 0x96, 0xfe, 0x16, 0x4e, 0xc2, 0x4b, 0x61, 0x30, 0xfb, 0x1c, 0x17,
	0x22, 0x8b, 0x8d, 0x54, 0xba, 0x0f, 0xf4, 0x14, 0x8e, 0xfe, 0x7a, 0xd2, 0x8c, 0x4c, 0xdd, 0xf9,
	0x90, 0xc3, 0xc1, 0x94, 0xb6, 0xfa, 0x77, 0xc0, 0xae, 0xeb, 0xaf, 0xfe, 0xbc, 0xc9, 0x80, 0xd3,
	0x8f, 0x70, 0x7c, 0x90, 0x72, 0x59, 0x20, 0x3d, 0x82, 0x3b, 0x9f, 0x96, 0xe7, 0xcb, 0xd5, 0x97,
	0xe5, 0xf8, 0x16, 0x1d, 0xc2, 0xdd, 0x60, 0xb3, 0x09, 0xa3, 0x4d, 0xc8, 0xc7, 0xa4, 0x7d, 0xad,
	0xf9, 0x6a, 0xbd, 0x8a, 0x42, 0x3e, 0x76, 0xe8, 0x08, 0x20, 0x58, 0x2c, 0x78, 0xb8, 0x08, 0x36,
	0x2b, 0x3e, 0x76, 0x27, 0x0e, 0x23, 0xaf, 0x7e, 0x12, 0x78, 0xd0, 0xc7, 0x97, 0x2a, 0xea, 0xb6,
	0x4d, 0x15, 0x9c, 0x44, 0x4d, 0x52, 0x0a, 0xd3, 0xb7, 0x30, 0xa8, 0xb2, 0xb5, 0x92, 0xf2, 0x82,
	0x9e, 0x7a, 0xff, 0x3f, 0x01, 0xef, 0xfa, 0x86, 0x27, 0x2f, 0x6e, 0xc4, 0xed, 0xc2, 0x4f, 0xdc,
	0x1f, 0x0e, 0x79, 0xef, 0xfe, 0x26, 0x24, 0xb9, 0x6d, 0xcf, 0xe6, 0xf5, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0x61, 0x38, 0x79, 0x91, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AggregatorServiceClient is the client API for AggregatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type AggregatorServiceClient interface {
	SubmitAggregateAndProof(ctx context.Context, in *AggregationRequest, opts ...grpc.CallOption) (*AggregationResponse, error)
}

type aggregatorServiceClient struct {
	cc *grpc.ClientConn
}

// Deprecated: Do not use.
func NewAggregatorServiceClient(cc *grpc.ClientConn) AggregatorServiceClient {
	return &aggregatorServiceClient{cc}
}

func (c *aggregatorServiceClient) SubmitAggregateAndProof(ctx context.Context, in *AggregationRequest, opts ...grpc.CallOption) (*AggregationResponse, error) {
	out := new(AggregationResponse)
	err := c.cc.Invoke(ctx, "/ethereum.beacon.rpc.v1.AggregatorService/SubmitAggregateAndProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServiceServer is the server API for AggregatorService service.
//
// Deprecated: Do not use.
type AggregatorServiceServer interface {
	SubmitAggregateAndProof(context.Context, *AggregationRequest) (*AggregationResponse, error)
}

// Deprecated: Do not use.
// UnimplementedAggregatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAggregatorServiceServer struct {
}

func (*UnimplementedAggregatorServiceServer) SubmitAggregateAndProof(ctx context.Context, req *AggregationRequest) (*AggregationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAggregateAndProof not implemented")
}

// Deprecated: Do not use.
func RegisterAggregatorServiceServer(s *grpc.Server, srv AggregatorServiceServer) {
	s.RegisterService(&_AggregatorService_serviceDesc, srv)
}

func _AggregatorService_SubmitAggregateAndProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).SubmitAggregateAndProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.beacon.rpc.v1.AggregatorService/SubmitAggregateAndProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).SubmitAggregateAndProof(ctx, req.(*AggregationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AggregatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.beacon.rpc.v1.AggregatorService",
	HandlerType: (*AggregatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitAggregateAndProof",
			Handler:    _AggregatorService_SubmitAggregateAndProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/beacon/rpc/v1/services.proto",
}