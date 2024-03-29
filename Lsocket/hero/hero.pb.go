// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hero.proto

package hero

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

type HeroAttack struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Damage               int32    `protobuf:"varint,2,opt,name=damage,proto3" json:"damage,omitempty"`
	Weapons              []string `protobuf:"bytes,4,rep,name=weapons,proto3" json:"weapons,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeroAttack) Reset()         { *m = HeroAttack{} }
func (m *HeroAttack) String() string { return proto.CompactTextString(m) }
func (*HeroAttack) ProtoMessage()    {}
func (*HeroAttack) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{0}
}

func (m *HeroAttack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroAttack.Unmarshal(m, b)
}
func (m *HeroAttack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroAttack.Marshal(b, m, deterministic)
}
func (m *HeroAttack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroAttack.Merge(m, src)
}
func (m *HeroAttack) XXX_Size() int {
	return xxx_messageInfo_HeroAttack.Size(m)
}
func (m *HeroAttack) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroAttack.DiscardUnknown(m)
}

var xxx_messageInfo_HeroAttack proto.InternalMessageInfo

func (m *HeroAttack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HeroAttack) GetDamage() int32 {
	if m != nil {
		return m.Damage
	}
	return 0
}

func (m *HeroAttack) GetWeapons() []string {
	if m != nil {
		return m.Weapons
	}
	return nil
}

func init() {
	proto.RegisterType((*HeroAttack)(nil), "hero.HeroAttack")
}

func init() { proto.RegisterFile("hero.proto", fileDescriptor_d5c1be3f86b99613) }

var fileDescriptor_d5c1be3f86b99613 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x48, 0x2d, 0xca,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x82, 0xb8, 0xb8, 0x3c, 0x52,
	0x8b, 0xf2, 0x1d, 0x4b, 0x4a, 0x12, 0x93, 0xb3, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x31, 0x2e, 0xb6, 0x94, 0xc4, 0xdc,
	0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x28, 0x4f, 0x48, 0x82, 0x8b, 0xbd,
	0x3c, 0x35, 0xb1, 0x20, 0x3f, 0xaf, 0x58, 0x82, 0x45, 0x81, 0x59, 0x83, 0x33, 0x08, 0xc6, 0x4d,
	0x62, 0x03, 0x5b, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x98, 0xe2, 0x66, 0x03, 0x6e, 0x00,
	0x00, 0x00,
}
