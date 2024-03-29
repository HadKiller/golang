// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Book int32

const (
	Book_UNKNOWN Book = 0
	Book_STUDY   Book = 1
	Book_STORY   Book = 2
	Book_TECH    Book = 3
	Book_LAUGHT  Book = 4
)

var Book_name = map[int32]string{
	0: "UNKNOWN",
	1: "STUDY",
	2: "STORY",
	3: "TECH",
	4: "LAUGHT",
}

var Book_value = map[string]int32{
	"UNKNOWN": 0,
	"STUDY":   1,
	"STORY":   2,
	"TECH":    3,
	"LAUGHT":  4,
}

func (x Book) String() string {
	return proto.EnumName(Book_name, int32(x))
}

func (Book) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c7b31125b9b4a3d9, []int{0}
}

type Food int32

const (
	Food_COOKIE    Food = 0
	Food_HUMBURGER Food = 1
	Food_RICE      Food = 2
	Food_BREAD     Food = 3
)

var Food_name = map[int32]string{
	0: "COOKIE",
	1: "HUMBURGER",
	2: "RICE",
	3: "BREAD",
}

var Food_value = map[string]int32{
	"COOKIE":    0,
	"HUMBURGER": 1,
	"RICE":      2,
	"BREAD":     3,
}

func (x Food) String() string {
	return proto.EnumName(Food_name, int32(x))
}

func (Food) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c7b31125b9b4a3d9, []int{1}
}

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7b31125b9b4a3d9, []int{0}
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

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func init() {
	proto.RegisterEnum("example.Book", Book_name, Book_value)
	proto.RegisterEnum("example.Food", Food_name, Food_value)
	proto.RegisterType((*SearchRequest)(nil), "example.SearchRequest")
}

func init() { proto.RegisterFile("go.proto", fileDescriptor_c7b31125b9b4a3d9) }

var fileDescriptor_c7b31125b9b4a3d9 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0xcf, 0x4b, 0x6b, 0xc2, 0x40,
	0x10, 0x07, 0x70, 0xf3, 0x52, 0x33, 0x12, 0xba, 0x0c, 0x3d, 0xe4, 0x56, 0xe9, 0xa1, 0x88, 0x87,
	0x5e, 0x0a, 0xbd, 0x16, 0x13, 0xb7, 0x46, 0x6c, 0x13, 0x59, 0xb3, 0x14, 0x4f, 0x21, 0xb6, 0x43,
	0x0a, 0x55, 0x37, 0xae, 0x09, 0xb4, 0xdf, 0xbe, 0x6c, 0x72, 0x9b, 0xf9, 0xcd, 0x03, 0xfe, 0x30,
	0xae, 0xd4, 0x63, 0xad, 0x55, 0xa3, 0x70, 0x44, 0xbf, 0xe5, 0xa9, 0x3e, 0xd2, 0xfd, 0x19, 0x82,
	0x1d, 0x95, 0xfa, 0xf3, 0x5b, 0xd0, 0xa5, 0xa5, 0x6b, 0x83, 0xb7, 0xe0, 0x5d, 0x5a, 0xd2, 0x7f,
	0xa1, 0x35, 0xb5, 0x66, 0xbe, 0xe8, 0x1b, 0xbc, 0x83, 0x49, 0x5d, 0x56, 0x54, 0x9c, 0xdb, 0xd3,
	0x81, 0x74, 0x68, 0x4f, 0xad, 0x99, 0x27, 0xc0, 0x50, 0xda, 0x09, 0x3e, 0xc0, 0x8d, 0xa6, 0x6b,
	0x7b, 0x6c, 0x8a, 0x9a, 0x74, 0x61, 0x06, 0xa1, 0xd3, 0x2d, 0x05, 0x3d, 0x6f, 0x49, 0x6f, 0xcb,
	0x8a, 0xe6, 0x2f, 0xe0, 0x46, 0x4a, 0xfd, 0xe0, 0x04, 0x46, 0x32, 0xdd, 0xa4, 0xd9, 0x47, 0xca,
	0x06, 0xe8, 0x83, 0xb7, 0xcb, 0xe5, 0x72, 0xcf, 0xac, 0xbe, 0xcc, 0xc4, 0x9e, 0xd9, 0x38, 0x06,
	0x37, 0xe7, 0x71, 0xc2, 0x1c, 0x04, 0x18, 0xbe, 0x2d, 0xe4, 0x2a, 0xc9, 0x99, 0x3b, 0x7f, 0x06,
	0xf7, 0x55, 0xa9, 0x2f, 0x63, 0x71, 0x96, 0x6d, 0xd6, 0x9c, 0x0d, 0x30, 0x00, 0x3f, 0x91, 0xef,
	0x91, 0x14, 0x2b, 0x2e, 0x98, 0x65, 0x0e, 0xc5, 0x3a, 0xe6, 0xcc, 0x36, 0xdf, 0x22, 0xc1, 0x17,
	0x4b, 0xe6, 0x1c, 0x86, 0x5d, 0xf0, 0xa7, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xba, 0x52,
	0x16, 0x04, 0x01, 0x00, 0x00,
}
