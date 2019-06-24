// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/synchronizer.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
	pb "open-match.dev/open-match/pkg/pb"
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

type RegisterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dbd55595bca25da, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

type RegisterResponse struct {
	// Identifier for this request valid for the current synchronization cycle.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dbd55595bca25da, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EvaluateProposalsRequest struct {
	// List of proposals to evaluate in the current synchronization cycle.
	Match []*pb.Match `protobuf:"bytes,1,rep,name=match,proto3" json:"match,omitempty"`
	// Identifier for this request issued during request registration.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvaluateProposalsRequest) Reset()         { *m = EvaluateProposalsRequest{} }
func (m *EvaluateProposalsRequest) String() string { return proto.CompactTextString(m) }
func (*EvaluateProposalsRequest) ProtoMessage()    {}
func (*EvaluateProposalsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dbd55595bca25da, []int{2}
}

func (m *EvaluateProposalsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateProposalsRequest.Unmarshal(m, b)
}
func (m *EvaluateProposalsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateProposalsRequest.Marshal(b, m, deterministic)
}
func (m *EvaluateProposalsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateProposalsRequest.Merge(m, src)
}
func (m *EvaluateProposalsRequest) XXX_Size() int {
	return xxx_messageInfo_EvaluateProposalsRequest.Size(m)
}
func (m *EvaluateProposalsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateProposalsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateProposalsRequest proto.InternalMessageInfo

func (m *EvaluateProposalsRequest) GetMatch() []*pb.Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *EvaluateProposalsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EvaluateProposalsResponse struct {
	// Results from evaluating proposals for this request.
	Match                []*pb.Match `protobuf:"bytes,1,rep,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EvaluateProposalsResponse) Reset()         { *m = EvaluateProposalsResponse{} }
func (m *EvaluateProposalsResponse) String() string { return proto.CompactTextString(m) }
func (*EvaluateProposalsResponse) ProtoMessage()    {}
func (*EvaluateProposalsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9dbd55595bca25da, []int{3}
}

func (m *EvaluateProposalsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateProposalsResponse.Unmarshal(m, b)
}
func (m *EvaluateProposalsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateProposalsResponse.Marshal(b, m, deterministic)
}
func (m *EvaluateProposalsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateProposalsResponse.Merge(m, src)
}
func (m *EvaluateProposalsResponse) XXX_Size() int {
	return xxx_messageInfo_EvaluateProposalsResponse.Size(m)
}
func (m *EvaluateProposalsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateProposalsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateProposalsResponse proto.InternalMessageInfo

func (m *EvaluateProposalsResponse) GetMatch() []*pb.Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "api.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "api.RegisterResponse")
	proto.RegisterType((*EvaluateProposalsRequest)(nil), "api.EvaluateProposalsRequest")
	proto.RegisterType((*EvaluateProposalsResponse)(nil), "api.EvaluateProposalsResponse")
}

func init() { proto.RegisterFile("api/synchronizer.proto", fileDescriptor_9dbd55595bca25da) }

var fileDescriptor_9dbd55595bca25da = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x1d, 0x28, 0xb0, 0x45, 0xd0, 0xae, 0xa0, 0x4a, 0xc3, 0x87, 0xb6, 0xee, 0xa5, 0x8a,
	0x88, 0xb7, 0x0d, 0x3d, 0x05, 0x55, 0x6a, 0x81, 0x1e, 0x2a, 0x85, 0x0f, 0xb9, 0x12, 0x87, 0xde,
	0x36, 0xf6, 0x60, 0x2f, 0xb2, 0x77, 0x97, 0x9d, 0x75, 0x0a, 0x1c, 0xfb, 0x13, 0xe0, 0xc6, 0x0f,
	0xe1, 0x8f, 0x70, 0xe6, 0xc6, 0x19, 0xfe, 0x02, 0xb2, 0x9d, 0x90, 0x92, 0xb4, 0xe2, 0x64, 0x79,
	0xde, 0x9b, 0xf7, 0x3c, 0x6f, 0xc6, 0x64, 0x4d, 0x18, 0xc9, 0xf1, 0xa3, 0x8a, 0x33, 0xab, 0x95,
	0xfc, 0x04, 0x36, 0x34, 0x56, 0x3b, 0x4d, 0x5b, 0xc2, 0xc8, 0x0e, 0xad, 0xc0, 0x02, 0x10, 0x45,
	0x0a, 0xd8, 0x00, 0x9d, 0xfb, 0xa9, 0xd6, 0x69, 0x0e, 0xbc, 0x82, 0x84, 0x52, 0xda, 0x09, 0x27,
	0xb5, 0x9a, 0xa2, 0x8f, 0xea, 0x47, 0xdc, 0x4b, 0x41, 0xf5, 0xf0, 0x54, 0xa4, 0x29, 0x58, 0xae,
	0x4d, 0xcd, 0x58, 0x64, 0x07, 0xab, 0xe4, 0x76, 0x04, 0xa9, 0x44, 0x07, 0x36, 0x82, 0xf7, 0x25,
	0xa0, 0x0b, 0x02, 0xb2, 0x32, 0x2b, 0xa1, 0xd1, 0x0a, 0x81, 0xde, 0x22, 0xbe, 0x4c, 0xda, 0x1e,
	0xf3, 0xb6, 0x6e, 0x44, 0xbe, 0x4c, 0x82, 0x21, 0x69, 0x1f, 0x8e, 0x45, 0x5e, 0x0a, 0x07, 0xaf,
	0xad, 0x36, 0x1a, 0x45, 0x8e, 0x93, 0x7e, 0xca, 0xc8, 0xd5, 0x42, 0xb8, 0x38, 0x6b, 0x7b, 0xac,
	0xb5, 0xb5, 0xdc, 0x27, 0xa1, 0x30, 0x32, 0x7c, 0x51, 0x55, 0xa2, 0x06, 0x98, 0xa8, 0xf9, 0x7f,
	0xd5, 0xf6, 0xc8, 0xfa, 0x05, 0x6a, 0x13, 0xeb, 0xff, 0xca, 0xf5, 0x7f, 0x7b, 0xe4, 0xe6, 0xf1,
	0xb9, 0xfc, 0xe8, 0x09, 0xb9, 0x3e, 0x9d, 0x80, 0xde, 0xa9, 0xf9, 0x73, 0x33, 0x76, 0xee, 0xce,
	0x55, 0x1b, 0xaf, 0x60, 0xe3, 0xec, 0xfb, 0xcf, 0x2f, 0xfe, 0x3d, 0xba, 0xce, 0xc7, 0x3b, 0xff,
	0xac, 0x84, 0xdb, 0xa9, 0xde, 0x99, 0x47, 0x56, 0x17, 0x3e, 0x96, 0x3e, 0xa8, 0xf5, 0x2e, 0x8b,
	0xa4, 0xf3, 0xf0, 0x32, 0x78, 0xe2, 0x1b, 0xd6, 0xbe, 0x5b, 0xc1, 0xe6, 0x82, 0xaf, 0x99, 0x72,
	0x07, 0x30, 0xe9, 0x1e, 0x78, 0xdd, 0xa7, 0xbf, 0xfc, 0xcf, 0x07, 0x3f, 0x7c, 0xfa, 0x6d, 0x6e,
	0xf0, 0xe0, 0x88, 0x90, 0x57, 0x06, 0x14, 0xab, 0xe3, 0xa1, 0x6b, 0x99, 0x73, 0x06, 0x07, 0x9c,
	0x6b, 0x03, 0xaa, 0x57, 0x67, 0x15, 0x26, 0x30, 0xee, 0x6c, 0xce, 0xde, 0x7b, 0x89, 0xc4, 0xb8,
	0x44, 0xdc, 0x6f, 0x4e, 0x2a, 0xb5, 0xba, 0x34, 0x18, 0xc6, 0xba, 0xe8, 0xbe, 0x21, 0xf4, 0xc0,
	0x88, 0x38, 0x03, 0xd6, 0x0f, 0xb7, 0xd9, 0x50, 0xc6, 0x50, 0x2d, 0x63, 0x7f, 0x2a, 0x99, 0x4a,
	0x97, 0x95, 0xa3, 0x8a, 0xc9, 0x9b, 0xd6, 0xb7, 0xda, 0xa6, 0xa2, 0x00, 0x3c, 0x67, 0xc6, 0x47,
	0xb9, 0x1e, 0xf1, 0x42, 0x54, 0xb9, 0xf1, 0xe1, 0xd1, 0xb3, 0xc3, 0x97, 0xc7, 0x87, 0xfd, 0xd6,
	0x4e, 0xb8, 0xdd, 0xf5, 0x3d, 0xbf, 0xbf, 0x22, 0x8c, 0xc9, 0x65, 0x5c, 0x5f, 0x23, 0x7f, 0x87,
	0x5a, 0x0d, 0x16, 0x2a, 0xd1, 0x13, 0xd2, 0xda, 0xdd, 0xde, 0xa5, 0xbb, 0xa4, 0x1b, 0x81, 0x2b,
	0xad, 0x82, 0x84, 0x9d, 0x66, 0xa0, 0x98, 0xcb, 0x80, 0x59, 0x40, 0x5d, 0xda, 0x18, 0x58, 0xa2,
	0x01, 0x99, 0xd2, 0x8e, 0xc1, 0x07, 0x89, 0x2e, 0xa4, 0x4b, 0xe4, 0xca, 0x57, 0xdf, 0xbb, 0x66,
	0xf7, 0x48, 0x7b, 0x16, 0x06, 0x7b, 0xae, 0xe3, 0xb2, 0x00, 0xd5, 0x5c, 0x3f, 0xdd, 0xb8, 0x38,
	0x1a, 0x8e, 0xd2, 0x01, 0x4f, 0x74, 0x8c, 0xfc, 0x64, 0x59, 0x2a, 0x07, 0x56, 0x89, 0x9c, 0x9b,
	0xd1, 0x68, 0xa9, 0xfe, 0x5b, 0x1e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xcf, 0x1f, 0xc5,
	0xac, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SynchronizerClient is the client API for Synchronizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SynchronizerClient interface {
	// Register associates this request with the current synchronization cycle and
	// returns an identifier for this registration. The caller returns this
	// identifier back in the evaluation
	// request. This enables identify stale evaluation requests belonging to a
	// prior window when synchronizing evaluation requests for a window.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// EvaluateProposals accepts a list of proposals and a registration identifier
	// for this request. If the synchronization cycle to which the request was
	// registered is completed, this request fails otherwise the proposals are
	// added to the list of proposals to be evaluated in the current cycle. At the
	//  end of the cycle, the user defined evaluation method is triggered and the
	// matches accepted by it are returned as results.
	EvaluateProposals(ctx context.Context, in *EvaluateProposalsRequest, opts ...grpc.CallOption) (*EvaluateProposalsResponse, error)
}

type synchronizerClient struct {
	cc *grpc.ClientConn
}

func NewSynchronizerClient(cc *grpc.ClientConn) SynchronizerClient {
	return &synchronizerClient{cc}
}

func (c *synchronizerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/api.Synchronizer/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *synchronizerClient) EvaluateProposals(ctx context.Context, in *EvaluateProposalsRequest, opts ...grpc.CallOption) (*EvaluateProposalsResponse, error) {
	out := new(EvaluateProposalsResponse)
	err := c.cc.Invoke(ctx, "/api.Synchronizer/EvaluateProposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SynchronizerServer is the server API for Synchronizer service.
type SynchronizerServer interface {
	// Register associates this request with the current synchronization cycle and
	// returns an identifier for this registration. The caller returns this
	// identifier back in the evaluation
	// request. This enables identify stale evaluation requests belonging to a
	// prior window when synchronizing evaluation requests for a window.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// EvaluateProposals accepts a list of proposals and a registration identifier
	// for this request. If the synchronization cycle to which the request was
	// registered is completed, this request fails otherwise the proposals are
	// added to the list of proposals to be evaluated in the current cycle. At the
	//  end of the cycle, the user defined evaluation method is triggered and the
	// matches accepted by it are returned as results.
	EvaluateProposals(context.Context, *EvaluateProposalsRequest) (*EvaluateProposalsResponse, error)
}

func RegisterSynchronizerServer(s *grpc.Server, srv SynchronizerServer) {
	s.RegisterService(&_Synchronizer_serviceDesc, srv)
}

func _Synchronizer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synchronizer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Synchronizer_EvaluateProposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateProposalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SynchronizerServer).EvaluateProposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Synchronizer/EvaluateProposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SynchronizerServer).EvaluateProposals(ctx, req.(*EvaluateProposalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Synchronizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Synchronizer_Register_Handler,
		},
		{
			MethodName: "EvaluateProposals",
			Handler:    _Synchronizer_EvaluateProposals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/synchronizer.proto",
}