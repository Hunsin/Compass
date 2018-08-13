// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/Hunsin/compass/trade/pb/APIs.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import trade "github.com/Hunsin/compass/trade"

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

type Status int32

const (
	Status_UNKNOWN Status = 0
	// 1xx are query status
	Status_PING     Status = 100
	Status_SECURITY Status = 101
	Status_QUOTE    Status = 102
	// 2xx are result status
	Status_DONE    Status = 200
	Status_PARTIAL Status = 206
	// 4xx are error status from results
	Status_BAD_REQUEST Status = 400
	Status_CLOSED      Status = 401
	Status_HALTED      Status = 402
	Status_UNLISTED    Status = 403
	Status_NOT_FOUND   Status = 404
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	100: "PING",
	101: "SECURITY",
	102: "QUOTE",
	200: "DONE",
	206: "PARTIAL",
	400: "BAD_REQUEST",
	401: "CLOSED",
	402: "HALTED",
	403: "UNLISTED",
	404: "NOT_FOUND",
}
var Status_value = map[string]int32{
	"UNKNOWN":     0,
	"PING":        100,
	"SECURITY":    101,
	"QUOTE":       102,
	"DONE":        200,
	"PARTIAL":     206,
	"BAD_REQUEST": 400,
	"CLOSED":      401,
	"HALTED":      402,
	"UNLISTED":    403,
	"NOT_FOUND":   404,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_APIs_946756c3f51b7857, []int{0}
}

type Result struct {
	QueryId              string            `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	Status               Status            `protobuf:"varint,2,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	Market               *trade.Market     `protobuf:"bytes,3,opt,name=market,proto3" json:"market,omitempty"`
	Securities           []*trade.Security `protobuf:"bytes,4,rep,name=securities,proto3" json:"securities,omitempty"`
	Quotes               []*trade.Daily    `protobuf:"bytes,5,rep,name=quotes,proto3" json:"quotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_APIs_946756c3f51b7857, []int{0}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *Result) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNKNOWN
}

func (m *Result) GetMarket() *trade.Market {
	if m != nil {
		return m.Market
	}
	return nil
}

func (m *Result) GetSecurities() []*trade.Security {
	if m != nil {
		return m.Securities
	}
	return nil
}

func (m *Result) GetQuotes() []*trade.Daily {
	if m != nil {
		return m.Quotes
	}
	return nil
}

type Query struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ask                  Status   `protobuf:"varint,2,opt,name=ask,proto3,enum=pb.Status" json:"ask,omitempty"`
	Symbol               []string `protobuf:"bytes,3,rep,name=symbol,proto3" json:"symbol,omitempty"`
	FromDate             []string `protobuf:"bytes,4,rep,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	EndDate              []string `protobuf:"bytes,5,rep,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Type                 []string `protobuf:"bytes,6,rep,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_APIs_946756c3f51b7857, []int{1}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (dst *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(dst, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query) GetAsk() Status {
	if m != nil {
		return m.Ask
	}
	return Status_UNKNOWN
}

func (m *Query) GetSymbol() []string {
	if m != nil {
		return m.Symbol
	}
	return nil
}

func (m *Query) GetFromDate() []string {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *Query) GetEndDate() []string {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *Query) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

type Predator struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr                 string          `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Markets              []*trade.Market `protobuf:"bytes,3,rep,name=markets,proto3" json:"markets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Predator) Reset()         { *m = Predator{} }
func (m *Predator) String() string { return proto.CompactTextString(m) }
func (*Predator) ProtoMessage()    {}
func (*Predator) Descriptor() ([]byte, []int) {
	return fileDescriptor_APIs_946756c3f51b7857, []int{2}
}
func (m *Predator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Predator.Unmarshal(m, b)
}
func (m *Predator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Predator.Marshal(b, m, deterministic)
}
func (dst *Predator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Predator.Merge(dst, src)
}
func (m *Predator) XXX_Size() int {
	return xxx_messageInfo_Predator.Size(m)
}
func (m *Predator) XXX_DiscardUnknown() {
	xxx_messageInfo_Predator.DiscardUnknown(m)
}

var xxx_messageInfo_Predator proto.InternalMessageInfo

func (m *Predator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Predator) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Predator) GetMarkets() []*trade.Market {
	if m != nil {
		return m.Markets
	}
	return nil
}

// A Null is nothing.
type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_APIs_946756c3f51b7857, []int{3}
}
func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (dst *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(dst, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Result)(nil), "pb.Result")
	proto.RegisterType((*Query)(nil), "pb.Query")
	proto.RegisterType((*Predator)(nil), "pb.Predator")
	proto.RegisterType((*Null)(nil), "pb.Null")
	proto.RegisterEnum("pb.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompassClient is the client API for Compass service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompassClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (Compass_RegisterClient, error)
	Predators(ctx context.Context, in *Null, opts ...grpc.CallOption) (Compass_PredatorsClient, error)
}

type compassClient struct {
	cc *grpc.ClientConn
}

func NewCompassClient(cc *grpc.ClientConn) CompassClient {
	return &compassClient{cc}
}

func (c *compassClient) Register(ctx context.Context, opts ...grpc.CallOption) (Compass_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Compass_serviceDesc.Streams[0], "/pb.Compass/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &compassRegisterClient{stream}
	return x, nil
}

type Compass_RegisterClient interface {
	Send(*Result) error
	Recv() (*Query, error)
	grpc.ClientStream
}

type compassRegisterClient struct {
	grpc.ClientStream
}

func (x *compassRegisterClient) Send(m *Result) error {
	return x.ClientStream.SendMsg(m)
}

func (x *compassRegisterClient) Recv() (*Query, error) {
	m := new(Query)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compassClient) Predators(ctx context.Context, in *Null, opts ...grpc.CallOption) (Compass_PredatorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Compass_serviceDesc.Streams[1], "/pb.Compass/Predators", opts...)
	if err != nil {
		return nil, err
	}
	x := &compassPredatorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Compass_PredatorsClient interface {
	Recv() (*Predator, error)
	grpc.ClientStream
}

type compassPredatorsClient struct {
	grpc.ClientStream
}

func (x *compassPredatorsClient) Recv() (*Predator, error) {
	m := new(Predator)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompassServer is the server API for Compass service.
type CompassServer interface {
	Register(Compass_RegisterServer) error
	Predators(*Null, Compass_PredatorsServer) error
}

func RegisterCompassServer(s *grpc.Server, srv CompassServer) {
	s.RegisterService(&_Compass_serviceDesc, srv)
}

func _Compass_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CompassServer).Register(&compassRegisterServer{stream})
}

type Compass_RegisterServer interface {
	Send(*Query) error
	Recv() (*Result, error)
	grpc.ServerStream
}

type compassRegisterServer struct {
	grpc.ServerStream
}

func (x *compassRegisterServer) Send(m *Query) error {
	return x.ServerStream.SendMsg(m)
}

func (x *compassRegisterServer) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Compass_Predators_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Null)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompassServer).Predators(m, &compassPredatorsServer{stream})
}

type Compass_PredatorsServer interface {
	Send(*Predator) error
	grpc.ServerStream
}

type compassPredatorsServer struct {
	grpc.ServerStream
}

func (x *compassPredatorsServer) Send(m *Predator) error {
	return x.ServerStream.SendMsg(m)
}

var _Compass_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Compass",
	HandlerType: (*CompassServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _Compass_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Predators",
			Handler:       _Compass_Predators_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/Hunsin/compass/trade/pb/APIs.proto",
}

func init() {
	proto.RegisterFile("github.com/Hunsin/compass/trade/pb/APIs.proto", fileDescriptor_APIs_946756c3f51b7857)
}

var fileDescriptor_APIs_946756c3f51b7857 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x8e, 0x93, 0x50,
	0x14, 0xc7, 0xbd, 0xa5, 0xa5, 0x70, 0xda, 0x19, 0xc9, 0x59, 0x18, 0x1c, 0x5d, 0x34, 0x8d, 0x13,
	0x1b, 0x8d, 0x65, 0x52, 0x9f, 0xa0, 0x0e, 0xe8, 0x10, 0x2b, 0x74, 0x2e, 0x90, 0x89, 0xab, 0x06,
	0x86, 0x3b, 0x23, 0x99, 0xb6, 0x30, 0xdc, 0xcb, 0xa2, 0x6f, 0xe1, 0xd7, 0xc6, 0xc4, 0xd7, 0x31,
	0x71, 0xe5, 0x33, 0x19, 0x2e, 0xd4, 0x98, 0x8c, 0x89, 0x1b, 0x72, 0xce, 0xff, 0x77, 0x6e, 0xf8,
	0x9f, 0x0f, 0x78, 0x71, 0x9d, 0x89, 0x0f, 0x55, 0x32, 0xbd, 0xcc, 0x37, 0xd6, 0x59, 0xb5, 0xe5,
	0xd9, 0xd6, 0xba, 0xcc, 0x37, 0x45, 0xcc, 0xb9, 0x25, 0xca, 0x38, 0x65, 0x56, 0x91, 0x58, 0xf3,
	0xa5, 0xcb, 0xa7, 0x45, 0x99, 0x8b, 0x1c, 0x3b, 0x45, 0x72, 0xf4, 0xfc, 0x7f, 0x4f, 0xe4, 0xb7,
	0x79, 0x30, 0xfe, 0x41, 0x40, 0xa5, 0x8c, 0x57, 0x6b, 0x81, 0x0f, 0x41, 0xbb, 0xad, 0x58, 0xb9,
	0x5b, 0x65, 0xa9, 0x49, 0x46, 0x64, 0xa2, 0xd3, 0xbe, 0xcc, 0xdd, 0x14, 0xc7, 0xa0, 0x72, 0x11,
	0x8b, 0x8a, 0x9b, 0x9d, 0x11, 0x99, 0x1c, 0xce, 0x60, 0x5a, 0x24, 0xd3, 0x40, 0x2a, 0xb4, 0x25,
	0x78, 0x0c, 0xea, 0x26, 0x2e, 0x6f, 0x98, 0x30, 0x95, 0x11, 0x99, 0x0c, 0x66, 0x07, 0xd3, 0xe6,
	0x3f, 0xef, 0xa4, 0x48, 0x5b, 0x88, 0x16, 0x00, 0x67, 0x97, 0x55, 0x99, 0x89, 0x8c, 0x71, 0xb3,
	0x3b, 0x52, 0x26, 0x83, 0xd9, 0xfd, 0xb6, 0x34, 0x68, 0xc0, 0x8e, 0xfe, 0x55, 0x82, 0x4f, 0x40,
	0xbd, 0xad, 0x72, 0xc1, 0xb8, 0xd9, 0x93, 0xc5, 0xc3, 0xb6, 0xd8, 0x8e, 0xb3, 0xf5, 0x8e, 0xb6,
	0x6c, 0xfc, 0x8d, 0x40, 0xef, 0xbc, 0x76, 0x8b, 0x87, 0xd0, 0xf9, 0xd3, 0x40, 0x27, 0x4b, 0xf1,
	0x31, 0x28, 0x31, 0xbf, 0xf9, 0x87, 0xf1, 0x5a, 0xc6, 0x07, 0xa0, 0xf2, 0xdd, 0x26, 0xc9, 0xd7,
	0xa6, 0x32, 0x52, 0x26, 0x3a, 0x6d, 0x33, 0x7c, 0x04, 0xfa, 0x55, 0x99, 0x6f, 0x56, 0x69, 0x2c,
	0x98, 0x74, 0xa9, 0x53, 0xad, 0x16, 0xec, 0x58, 0xb0, 0x7a, 0x52, 0x6c, 0x9b, 0x36, 0xac, 0x27,
	0x59, 0x9f, 0x6d, 0x53, 0x89, 0x10, 0xba, 0x62, 0x57, 0x30, 0x53, 0x95, 0xb2, 0x8c, 0xc7, 0x17,
	0xa0, 0x2d, 0x4b, 0x96, 0xc6, 0x22, 0x2f, 0xef, 0xb8, 0x43, 0xe8, 0xc6, 0x69, 0x5a, 0x4a, 0x7b,
	0x3a, 0x95, 0x31, 0x3e, 0x85, 0x7e, 0x33, 0x2c, 0x2e, 0x4d, 0xdd, 0x19, 0xe5, 0x9e, 0x8e, 0x55,
	0xe8, 0x7a, 0xd5, 0x7a, 0xfd, 0xec, 0x3b, 0x01, 0xb5, 0x69, 0x0a, 0x07, 0xd0, 0x8f, 0xbc, 0xb7,
	0x9e, 0x7f, 0xe1, 0x19, 0xf7, 0x50, 0x83, 0xee, 0xd2, 0xf5, 0xde, 0x18, 0x29, 0x0e, 0x41, 0x0b,
	0x9c, 0xd3, 0x88, 0xba, 0xe1, 0x7b, 0x83, 0xa1, 0x0e, 0xbd, 0xf3, 0xc8, 0x0f, 0x1d, 0xe3, 0x0a,
	0x75, 0xe8, 0xda, 0xbe, 0xe7, 0x18, 0x3f, 0x09, 0x0e, 0xa1, 0xbf, 0x9c, 0xd3, 0xd0, 0x9d, 0x2f,
	0x8c, 0x5f, 0x04, 0x0d, 0x18, 0xbc, 0x9a, 0xdb, 0x2b, 0xea, 0x9c, 0x47, 0x4e, 0x10, 0x1a, 0x1f,
	0x15, 0x1c, 0x80, 0x7a, 0xba, 0xf0, 0x03, 0xc7, 0x36, 0x3e, 0xc9, 0xe4, 0x6c, 0xbe, 0x08, 0x1d,
	0xdb, 0xf8, 0xac, 0xe0, 0x01, 0x68, 0x91, 0xb7, 0x70, 0x83, 0x3a, 0xfd, 0xa2, 0xe0, 0x21, 0xe8,
	0x9e, 0x1f, 0xae, 0x5e, 0xfb, 0x91, 0x67, 0x1b, 0x5f, 0x95, 0xd9, 0x05, 0xf4, 0x4f, 0x9b, 0x03,
	0xc4, 0x63, 0xd0, 0x28, 0xbb, 0xce, 0xb8, 0x60, 0x25, 0xca, 0x5d, 0x34, 0xb7, 0x77, 0xa4, 0xd7,
	0xb1, 0xdc, 0xdf, 0x84, 0x9c, 0x10, 0x3c, 0x06, 0x7d, 0x3f, 0x31, 0x8e, 0x5a, 0xcd, 0xea, 0x3e,
	0x8f, 0x86, 0x75, 0xb4, 0x07, 0x27, 0x24, 0x51, 0xe5, 0x0d, 0xbf, 0xfc, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x1a, 0x89, 0xf3, 0x25, 0x03, 0x00, 0x00,
}
