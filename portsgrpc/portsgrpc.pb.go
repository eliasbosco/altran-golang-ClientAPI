// Code generated by protoc-gen-go. DO NOT EDIT.
// source: portsgrpc.proto

package portsgrpc

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

type Ports struct {
	PortsBody            []*PortsBody `protobuf:"bytes,1,rep,name=ports_body,json=portsBody,proto3" json:"ports_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Ports) Reset()         { *m = Ports{} }
func (m *Ports) String() string { return proto.CompactTextString(m) }
func (*Ports) ProtoMessage()    {}
func (*Ports) Descriptor() ([]byte, []int) {
	return fileDescriptor_09e7f44acbd5ff64, []int{0}
}

func (m *Ports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ports.Unmarshal(m, b)
}
func (m *Ports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ports.Marshal(b, m, deterministic)
}
func (m *Ports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ports.Merge(m, src)
}
func (m *Ports) XXX_Size() int {
	return xxx_messageInfo_Ports.Size(m)
}
func (m *Ports) XXX_DiscardUnknown() {
	xxx_messageInfo_Ports.DiscardUnknown(m)
}

var xxx_messageInfo_Ports proto.InternalMessageInfo

func (m *Ports) GetPortsBody() []*PortsBody {
	if m != nil {
		return m.PortsBody
	}
	return nil
}

type PortsBody struct {
	PortId               string    `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string    `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string    `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string  `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string  `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []float32 `protobuf:"fixed32,7,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string    `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string    `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string  `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 string    `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PortsBody) Reset()         { *m = PortsBody{} }
func (m *PortsBody) String() string { return proto.CompactTextString(m) }
func (*PortsBody) ProtoMessage()    {}
func (*PortsBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_09e7f44acbd5ff64, []int{1}
}

func (m *PortsBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortsBody.Unmarshal(m, b)
}
func (m *PortsBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortsBody.Marshal(b, m, deterministic)
}
func (m *PortsBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortsBody.Merge(m, src)
}
func (m *PortsBody) XXX_Size() int {
	return xxx_messageInfo_PortsBody.Size(m)
}
func (m *PortsBody) XXX_DiscardUnknown() {
	xxx_messageInfo_PortsBody.DiscardUnknown(m)
}

var xxx_messageInfo_PortsBody proto.InternalMessageInfo

func (m *PortsBody) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *PortsBody) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PortsBody) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *PortsBody) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *PortsBody) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *PortsBody) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *PortsBody) GetCoordinates() []float32 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *PortsBody) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *PortsBody) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *PortsBody) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *PortsBody) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Response struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_09e7f44acbd5ff64, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Request struct {
	Skip                 int32    `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	PortId               string   `protobuf:"bytes,3,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_09e7f44acbd5ff64, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSkip() int32 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *Request) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Request) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func init() {
	proto.RegisterType((*Ports)(nil), "portsgrpc.Ports")
	proto.RegisterType((*PortsBody)(nil), "portsgrpc.PortsBody")
	proto.RegisterType((*Response)(nil), "portsgrpc.Response")
	proto.RegisterType((*Request)(nil), "portsgrpc.Request")
}

func init() { proto.RegisterFile("portsgrpc.proto", fileDescriptor_09e7f44acbd5ff64) }

var fileDescriptor_09e7f44acbd5ff64 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x41, 0x8b, 0xdb, 0x30,
	0x10, 0x85, 0xe3, 0x38, 0xb6, 0xe3, 0xc9, 0xa1, 0x45, 0x0d, 0xad, 0xc8, 0xc9, 0xf8, 0xe4, 0x53,
	0xa0, 0x49, 0x0f, 0x3d, 0xf4, 0x54, 0x0a, 0xa5, 0xd0, 0xc3, 0x22, 0xd8, 0x73, 0x70, 0x6c, 0x11,
	0xc4, 0xc6, 0x1a, 0xaf, 0xa4, 0x2c, 0x78, 0x7f, 0xcb, 0xfe, 0xd8, 0x45, 0xe3, 0xd8, 0xf1, 0xee,
	0xde, 0xe6, 0x7b, 0x6f, 0x9e, 0x35, 0x1a, 0x19, 0x3e, 0xb5, 0x68, 0x9c, 0x3d, 0x99, 0xb6, 0xda,
	0xb6, 0x06, 0x1d, 0xb2, 0x74, 0x14, 0xf2, 0x5f, 0x10, 0xdd, 0x79, 0x60, 0x7b, 0x00, 0x52, 0x0f,
	0x47, 0xac, 0x3b, 0x1e, 0x64, 0x61, 0xb1, 0xda, 0xad, 0xb7, 0xb7, 0x24, 0x75, 0xfd, 0xc6, 0xba,
	0x13, 0x7d, 0xda, 0x97, 0xf9, 0xcb, 0x1c, 0xd2, 0xd1, 0x60, 0xdf, 0x20, 0xf1, 0xd6, 0x41, 0xd5,
	0x3c, 0xc8, 0x82, 0x22, 0x15, 0xb1, 0xc7, 0x7f, 0x35, 0x63, 0xb0, 0xd0, 0x65, 0x23, 0xf9, 0x9c,
	0x54, 0xaa, 0xbd, 0x56, 0x29, 0xd7, 0xf1, 0xb0, 0xd7, 0x7c, 0xcd, 0x38, 0x24, 0x15, 0x5e, 0xb4,
	0x33, 0x1d, 0x5f, 0x90, 0x3c, 0x20, 0x5b, 0x43, 0x54, 0x9e, 0x55, 0x69, 0x79, 0x94, 0x85, 0x45,
	0x2a, 0x7a, 0xf0, 0xfd, 0x46, 0x9e, 0x14, 0x6a, 0xcb, 0x63, 0xd2, 0x07, 0x64, 0x19, 0xac, 0x2a,
	0x44, 0x53, 0x2b, 0x5d, 0x3a, 0x69, 0x79, 0x92, 0x85, 0xc5, 0x5c, 0x4c, 0x25, 0xb6, 0x81, 0x65,
	0x6b, 0xf0, 0x49, 0xe9, 0x4a, 0xf2, 0x25, 0x1d, 0x36, 0xb2, 0xf7, 0x9c, 0x6a, 0xe4, 0x33, 0x6a,
	0xc9, 0xd3, 0xde, 0x1b, 0x98, 0x7d, 0x85, 0xf8, 0xa2, 0xcf, 0x58, 0x59, 0x0e, 0x74, 0xe4, 0x95,
	0xe8, 0x3e, 0x58, 0x4b, 0xbe, 0xba, 0xde, 0x07, 0x6b, 0x99, 0xff, 0x84, 0xa5, 0x90, 0xb6, 0x45,
	0x6d, 0xe5, 0xe8, 0x07, 0x37, 0xdf, 0xcf, 0xdf, 0x48, 0x6b, 0xcb, 0xd3, 0xb0, 0x9a, 0x01, 0xf3,
	0xff, 0x90, 0x08, 0xf9, 0x78, 0x91, 0xd6, 0xf9, 0xa0, 0x7d, 0x50, 0x2d, 0x05, 0x23, 0x41, 0xb5,
	0x5f, 0xc7, 0x59, 0x35, 0xca, 0x51, 0x2c, 0x12, 0x3d, 0x4c, 0xf7, 0x1f, 0x4e, 0xf7, 0xbf, 0x33,
	0x90, 0xd0, 0x2b, 0xfd, 0x39, 0xb2, 0xef, 0x10, 0xdf, 0xb7, 0x56, 0x1a, 0xc7, 0x3e, 0xbf, 0x7f,
	0xdc, 0xcd, 0x97, 0x89, 0x32, 0xcc, 0x9d, 0xcf, 0xd8, 0x0f, 0x80, 0xbf, 0xd2, 0x0d, 0x1f, 0x60,
	0x6f, 0x9a, 0x68, 0xc4, 0xcd, 0x87, 0x4f, 0xe5, 0xb3, 0x63, 0x4c, 0xbf, 0xda, 0xfe, 0x35, 0x00,
	0x00, 0xff, 0xff, 0x36, 0xfa, 0x27, 0x8f, 0x7d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PortsDbClient is the client API for PortsDb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortsDbClient interface {
	Upsert(ctx context.Context, in *Ports, opts ...grpc.CallOption) (*Response, error)
	GetPortsDb(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Ports, error)
}

type portsDbClient struct {
	cc *grpc.ClientConn
}

func NewPortsDbClient(cc *grpc.ClientConn) PortsDbClient {
	return &portsDbClient{cc}
}

func (c *portsDbClient) Upsert(ctx context.Context, in *Ports, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/portsgrpc.PortsDb/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portsDbClient) GetPortsDb(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Ports, error) {
	out := new(Ports)
	err := c.cc.Invoke(ctx, "/portsgrpc.PortsDb/GetPortsDb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortsDbServer is the server API for PortsDb service.
type PortsDbServer interface {
	Upsert(context.Context, *Ports) (*Response, error)
	GetPortsDb(context.Context, *Request) (*Ports, error)
}

// UnimplementedPortsDbServer can be embedded to have forward compatible implementations.
type UnimplementedPortsDbServer struct {
}

func (*UnimplementedPortsDbServer) Upsert(ctx context.Context, req *Ports) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (*UnimplementedPortsDbServer) GetPortsDb(ctx context.Context, req *Request) (*Ports, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortsDb not implemented")
}

func RegisterPortsDbServer(s *grpc.Server, srv PortsDbServer) {
	s.RegisterService(&_PortsDb_serviceDesc, srv)
}

func _PortsDb_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ports)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsDbServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portsgrpc.PortsDb/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsDbServer).Upsert(ctx, req.(*Ports))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortsDb_GetPortsDb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsDbServer).GetPortsDb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portsgrpc.PortsDb/GetPortsDb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsDbServer).GetPortsDb(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortsDb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "portsgrpc.PortsDb",
	HandlerType: (*PortsDbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upsert",
			Handler:    _PortsDb_Upsert_Handler,
		},
		{
			MethodName: "GetPortsDb",
			Handler:    _PortsDb_GetPortsDb_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portsgrpc.proto",
}
