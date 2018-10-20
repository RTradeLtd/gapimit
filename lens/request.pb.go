// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package lens

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type IndexRequest struct {
	// dataType is the "type" of data, such as IPLD
	DataType string `protobuf:"bytes,1,opt,name=dataType,proto3" json:"dataType,omitempty"`
	// objectIdentifier is the identifier of the object, such as an IPFS content hash
	ObjectIdentifier     string   `protobuf:"bytes,2,opt,name=objectIdentifier,proto3" json:"objectIdentifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexRequest) Reset()         { *m = IndexRequest{} }
func (m *IndexRequest) String() string { return proto.CompactTextString(m) }
func (*IndexRequest) ProtoMessage()    {}
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *IndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexRequest.Unmarshal(m, b)
}
func (m *IndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexRequest.Marshal(b, m, deterministic)
}
func (m *IndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRequest.Merge(m, src)
}
func (m *IndexRequest) XXX_Size() int {
	return xxx_messageInfo_IndexRequest.Size(m)
}
func (m *IndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRequest proto.InternalMessageInfo

func (m *IndexRequest) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *IndexRequest) GetObjectIdentifier() string {
	if m != nil {
		return m.ObjectIdentifier
	}
	return ""
}

type SearchRequest struct {
	// repeated means there can be 0 or more "keywords"
	Keywords             []string `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{1}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexRequest)(nil), "request.IndexRequest")
	proto.RegisterType((*SearchRequest)(nil), "request.SearchRequest")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xc2, 0xb8,
	0x78, 0x3c, 0xf3, 0x52, 0x52, 0x2b, 0x82, 0x20, 0x7c, 0x21, 0x29, 0x2e, 0x8e, 0x94, 0xc4, 0x92,
	0xc4, 0x90, 0xca, 0x82, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x48, 0x8b,
	0x4b, 0x20, 0x3f, 0x29, 0x2b, 0x35, 0xb9, 0xc4, 0x33, 0x25, 0x35, 0xaf, 0x24, 0x33, 0x2d, 0x33,
	0xb5, 0x48, 0x82, 0x09, 0xac, 0x06, 0x43, 0x5c, 0x49, 0x9b, 0x8b, 0x37, 0x38, 0x35, 0xb1, 0x28,
	0x39, 0x03, 0xc9, 0xe0, 0xec, 0xd4, 0xca, 0xf2, 0xfc, 0xa2, 0x94, 0x62, 0x09, 0x46, 0x05, 0x66,
	0x90, 0xc1, 0x30, 0x7e, 0x12, 0x1b, 0xd8, 0x51, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08,
	0x1d, 0x44, 0xc0, 0xa5, 0x00, 0x00, 0x00,
}
