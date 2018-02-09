// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getmeconfig.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	getmeconfig.proto

It has these top-level messages:
	ConfigId
	Config
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ConfigId struct {
	ConfigId string `protobuf:"bytes,1,opt,name=ConfigId" json:"ConfigId,omitempty"`
}

func (m *ConfigId) Reset()                    { *m = ConfigId{} }
func (m *ConfigId) String() string            { return proto.CompactTextString(m) }
func (*ConfigId) ProtoMessage()               {}
func (*ConfigId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConfigId) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

type Config struct {
	Config []byte `protobuf:"bytes,1,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigId)(nil), "api.ConfigId")
	proto.RegisterType((*Config)(nil), "api.Config")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GetConfig service

type GetConfigClient interface {
	Request(ctx context.Context, in *ConfigId, opts ...grpc.CallOption) (*Config, error)
}

type getConfigClient struct {
	cc *grpc.ClientConn
}

func NewGetConfigClient(cc *grpc.ClientConn) GetConfigClient {
	return &getConfigClient{cc}
}

func (c *getConfigClient) Request(ctx context.Context, in *ConfigId, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := grpc.Invoke(ctx, "/api.GetConfig/Request", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetConfig service

type GetConfigServer interface {
	Request(context.Context, *ConfigId) (*Config, error)
}

func RegisterGetConfigServer(s *grpc.Server, srv GetConfigServer) {
	s.RegisterService(&_GetConfig_serviceDesc, srv)
}

func _GetConfig_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetConfigServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GetConfig/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetConfigServer).Request(ctx, req.(*ConfigId))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GetConfig",
	HandlerType: (*GetConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _GetConfig_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getmeconfig.proto",
}

func init() { proto.RegisterFile("getmeconfig.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4f, 0x2d, 0xc9,
	0x4d, 0x4d, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e,
	0x2c, 0xc8, 0x54, 0x52, 0xe3, 0xe2, 0x70, 0x06, 0x0b, 0x7a, 0xa6, 0x08, 0x49, 0x21, 0xd8, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x92, 0x02, 0x17, 0x1b, 0x84, 0x2d, 0x24, 0x06,
	0x63, 0x81, 0xd5, 0xf0, 0x04, 0x41, 0x79, 0x46, 0x46, 0x5c, 0x9c, 0xee, 0xa9, 0x25, 0x50, 0x45,
	0xaa, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xbc, 0x7a, 0x89, 0x05, 0x99,
	0x7a, 0x30, 0x83, 0xa4, 0xb8, 0x91, 0xb8, 0x49, 0x6c, 0x60, 0x97, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xef, 0x4e, 0x1b, 0xd8, 0x9e, 0x00, 0x00, 0x00,
}
