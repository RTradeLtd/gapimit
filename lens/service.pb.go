// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
	fmt "fmt"
	request "github.com/RTradeLtd/grpc/lens/request"
	response "github.com/RTradeLtd/grpc/lens/response"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x12, 0x2d, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0xd1, 0x87, 0xd2, 0x50, 0x61, 0xf1, 0xa2, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x7d, 0x18, 0x03, 0x22, 0x61, 0xd4, 0xc6, 0xc4, 0xc5, 0xe5, 0x99, 0x97, 0x92, 0x5a, 0x91, 0x5a,
	0xe4, 0x18, 0xe0, 0x29, 0xe4, 0xc6, 0x25, 0x14, 0x5c, 0x9a, 0x94, 0x9b, 0x59, 0x02, 0x16, 0x0b,
	0x82, 0x98, 0x21, 0x24, 0xaa, 0x07, 0x33, 0x0d, 0x59, 0x58, 0x4a, 0x5c, 0x0f, 0x6e, 0x18, 0x54,
	0x1c, 0xc2, 0x53, 0x62, 0x10, 0x0a, 0xe5, 0x92, 0x84, 0x98, 0x13, 0x9c, 0x99, 0x5b, 0x90, 0x93,
	0x1a, 0x9c, 0x9a, 0x58, 0x94, 0x9c, 0x01, 0x33, 0x4e, 0x0c, 0x6e, 0x1c, 0x8a, 0xb8, 0x94, 0x1c,
	0xc2, 0x3c, 0x54, 0x6d, 0x70, 0x63, 0x23, 0xb9, 0xa4, 0x21, 0xc6, 0x3a, 0xa6, 0x94, 0x25, 0xe6,
	0x25, 0xa7, 0xa6, 0x10, 0x67, 0xb0, 0x02, 0xc2, 0x60, 0x74, 0x8d, 0x30, 0xa3, 0x93, 0xd8, 0xc0,
	0xe1, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x99, 0x73, 0x81, 0x61, 0x50, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexerAPIClient is the client API for IndexerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexerAPIClient interface {
	// SubmitIndexRequest is used to submit content to be indexed by the lens system
	SubmitIndexRequest(ctx context.Context, in *request.IndexRequest, opts ...grpc.CallOption) (*response.IndexResponse, error)
	// SubmitSimpleSearchRequest is used to perform a simple search against the lens index
	SubmitSimpleSearchRequest(ctx context.Context, in *request.SearchRequest, opts ...grpc.CallOption) (*response.SimpleSearchResponse, error)
	// SubmitAdvancedSearchRequest is used to perform an advanced search against the lens index
	SubmitAdvancedSearchRequest(ctx context.Context, in *request.SearchRequest, opts ...grpc.CallOption) (*response.AdvancedSearchResponse, error)
}

type indexerAPIClient struct {
	cc *grpc.ClientConn
}

func NewIndexerAPIClient(cc *grpc.ClientConn) IndexerAPIClient {
	return &indexerAPIClient{cc}
}

func (c *indexerAPIClient) SubmitIndexRequest(ctx context.Context, in *request.IndexRequest, opts ...grpc.CallOption) (*response.IndexResponse, error) {
	out := new(response.IndexResponse)
	err := c.cc.Invoke(ctx, "/IndexerAPI/SubmitIndexRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerAPIClient) SubmitSimpleSearchRequest(ctx context.Context, in *request.SearchRequest, opts ...grpc.CallOption) (*response.SimpleSearchResponse, error) {
	out := new(response.SimpleSearchResponse)
	err := c.cc.Invoke(ctx, "/IndexerAPI/SubmitSimpleSearchRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexerAPIClient) SubmitAdvancedSearchRequest(ctx context.Context, in *request.SearchRequest, opts ...grpc.CallOption) (*response.AdvancedSearchResponse, error) {
	out := new(response.AdvancedSearchResponse)
	err := c.cc.Invoke(ctx, "/IndexerAPI/SubmitAdvancedSearchRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexerAPIServer is the server API for IndexerAPI service.
type IndexerAPIServer interface {
	// SubmitIndexRequest is used to submit content to be indexed by the lens system
	SubmitIndexRequest(context.Context, *request.IndexRequest) (*response.IndexResponse, error)
	// SubmitSimpleSearchRequest is used to perform a simple search against the lens index
	SubmitSimpleSearchRequest(context.Context, *request.SearchRequest) (*response.SimpleSearchResponse, error)
	// SubmitAdvancedSearchRequest is used to perform an advanced search against the lens index
	SubmitAdvancedSearchRequest(context.Context, *request.SearchRequest) (*response.AdvancedSearchResponse, error)
}

func RegisterIndexerAPIServer(s *grpc.Server, srv IndexerAPIServer) {
	s.RegisterService(&_IndexerAPI_serviceDesc, srv)
}

func _IndexerAPI_SubmitIndexRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerAPIServer).SubmitIndexRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IndexerAPI/SubmitIndexRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerAPIServer).SubmitIndexRequest(ctx, req.(*request.IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexerAPI_SubmitSimpleSearchRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerAPIServer).SubmitSimpleSearchRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IndexerAPI/SubmitSimpleSearchRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerAPIServer).SubmitSimpleSearchRequest(ctx, req.(*request.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexerAPI_SubmitAdvancedSearchRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerAPIServer).SubmitAdvancedSearchRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IndexerAPI/SubmitAdvancedSearchRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerAPIServer).SubmitAdvancedSearchRequest(ctx, req.(*request.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IndexerAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "IndexerAPI",
	HandlerType: (*IndexerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitIndexRequest",
			Handler:    _IndexerAPI_SubmitIndexRequest_Handler,
		},
		{
			MethodName: "SubmitSimpleSearchRequest",
			Handler:    _IndexerAPI_SubmitSimpleSearchRequest_Handler,
		},
		{
			MethodName: "SubmitAdvancedSearchRequest",
			Handler:    _IndexerAPI_SubmitAdvancedSearchRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
