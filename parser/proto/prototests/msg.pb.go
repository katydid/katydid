// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

package prototests

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// BigMsg contains a field and a message field.
type BigMsg struct {
	Field            *int64    `protobuf:"varint,1,opt,name=Field" json:"Field,omitempty"`
	Msg              *SmallMsg `protobuf:"bytes,3,opt,name=Msg" json:"Msg,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BigMsg) Reset()         { *m = BigMsg{} }
func (m *BigMsg) String() string { return proto.CompactTextString(m) }
func (*BigMsg) ProtoMessage()    {}
func (*BigMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_07dd33c081c61e70, []int{0}
}
func (m *BigMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigMsg.Unmarshal(m, b)
}
func (m *BigMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigMsg.Marshal(b, m, deterministic)
}
func (dst *BigMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigMsg.Merge(dst, src)
}
func (m *BigMsg) XXX_Size() int {
	return xxx_messageInfo_BigMsg.Size(m)
}
func (m *BigMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BigMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BigMsg proto.InternalMessageInfo

func (m *BigMsg) GetField() int64 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *BigMsg) GetMsg() *SmallMsg {
	if m != nil {
		return m.Msg
	}
	return nil
}

// SmallMsg only contains some native fields.
type SmallMsg struct {
	ScarBusStop      *string  `protobuf:"bytes,1,opt,name=ScarBusStop" json:"ScarBusStop,omitempty"`
	FlightParachute  []uint32 `protobuf:"fixed32,12,rep,name=FlightParachute" json:"FlightParachute,omitempty"`
	MapShark         *string  `protobuf:"bytes,18,opt,name=MapShark" json:"MapShark,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmallMsg) Reset()         { *m = SmallMsg{} }
func (m *SmallMsg) String() string { return proto.CompactTextString(m) }
func (*SmallMsg) ProtoMessage()    {}
func (*SmallMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_07dd33c081c61e70, []int{1}
}
func (m *SmallMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmallMsg.Unmarshal(m, b)
}
func (m *SmallMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmallMsg.Marshal(b, m, deterministic)
}
func (dst *SmallMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmallMsg.Merge(dst, src)
}
func (m *SmallMsg) XXX_Size() int {
	return xxx_messageInfo_SmallMsg.Size(m)
}
func (m *SmallMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SmallMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SmallMsg proto.InternalMessageInfo

func (m *SmallMsg) GetScarBusStop() string {
	if m != nil && m.ScarBusStop != nil {
		return *m.ScarBusStop
	}
	return ""
}

func (m *SmallMsg) GetFlightParachute() []uint32 {
	if m != nil {
		return m.FlightParachute
	}
	return nil
}

func (m *SmallMsg) GetMapShark() string {
	if m != nil && m.MapShark != nil {
		return *m.MapShark
	}
	return ""
}

// Packed contains some repeated packed fields.
type Packed struct {
	Ints             []int64   `protobuf:"varint,4,rep,packed,name=Ints" json:"Ints,omitempty"`
	Floats           []float64 `protobuf:"fixed64,5,rep,packed,name=Floats" json:"Floats,omitempty"`
	Uints            []uint32  `protobuf:"varint,6,rep,packed,name=Uints" json:"Uints,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Packed) Reset()         { *m = Packed{} }
func (m *Packed) String() string { return proto.CompactTextString(m) }
func (*Packed) ProtoMessage()    {}
func (*Packed) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_07dd33c081c61e70, []int{2}
}
func (m *Packed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packed.Unmarshal(m, b)
}
func (m *Packed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packed.Marshal(b, m, deterministic)
}
func (dst *Packed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packed.Merge(dst, src)
}
func (m *Packed) XXX_Size() int {
	return xxx_messageInfo_Packed.Size(m)
}
func (m *Packed) XXX_DiscardUnknown() {
	xxx_messageInfo_Packed.DiscardUnknown(m)
}

var xxx_messageInfo_Packed proto.InternalMessageInfo

func (m *Packed) GetInts() []int64 {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *Packed) GetFloats() []float64 {
	if m != nil {
		return m.Floats
	}
	return nil
}

func (m *Packed) GetUints() []uint32 {
	if m != nil {
		return m.Uints
	}
	return nil
}

func init() {
	proto.RegisterType((*BigMsg)(nil), "prototests.BigMsg")
	proto.RegisterType((*SmallMsg)(nil), "prototests.SmallMsg")
	proto.RegisterType((*Packed)(nil), "prototests.Packed")
}
func (this *BigMsg) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *SmallMsg) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *Packed) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func MsgDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3929 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x4b, 0x70, 0x23, 0xd7,
		0x75, 0x1d, 0x7c, 0x09, 0x5c, 0x80, 0x60, 0xf3, 0x91, 0xe2, 0x60, 0x28, 0x6b, 0x86, 0x03, 0xfd,
		0xa8, 0x51, 0xcc, 0x71, 0x8d, 0x34, 0x23, 0x0d, 0x26, 0xb6, 0x02, 0x92, 0x20, 0x0d, 0x85, 0x1f,
		0xb8, 0x41, 0xea, 0xe7, 0x4a, 0x75, 0x3d, 0x36, 0x1e, 0x80, 0x9e, 0x69, 0x74, 0xb7, 0xbb, 0x1b,
		0x33, 0xe2, 0x54, 0x16, 0x4a, 0x29, 0x9f, 0x72, 0xa5, 0xf2, 0x4f, 0x55, 0x6c, 0x45, 0x56, 0x62,
		0xa7, 0x12, 0x25, 0xce, 0xcf, 0x8e, 0x13, 0x27, 0x76, 0x36, 0xd9, 0x38, 0xc9, 0x2a, 0x15, 0xef,
		0xb2, 0xc8, 0x22, 0x8e, 0x94, 0xca, 0x4f, 0x89, 0x95, 0x44, 0x0b, 0x57, 0x69, 0x93, 0x7a, 0xbf,
		0x46, 0xa3, 0x01, 0x4e, 0x83, 0xae, 0x92, 0xbc, 0x22, 0xfb, 0xbe, 0x7b, 0x4e, 0xdf, 0x77, 0xdf,
		0x7d, 0xf7, 0xde, 0xf7, 0x1a, 0xf0, 0xdd, 0xeb, 0xb0, 0xd2, 0xb5, 0xed, 0xae, 0x49, 0x2e, 0x3b,
		0xae, 0xed, 0xdb, 0x47, 0x83, 0xce, 0xe5, 0x36, 0xf1, 0x74, 0xd7, 0x70, 0x7c, 0xdb, 0x5d, 0x63,
		0x32, 0x34, 0xc7, 0x35, 0xd6, 0xa4, 0x46, 0x65, 0x17, 0xe6, 0xb7, 0x0c, 0x93, 0x6c, 0x06, 0x8a,
		0x2d, 0xe2, 0xa3, 0xa7, 0x21, 0xdd, 0x31, 0x4c, 0x52, 0x4e, 0xac, 0xa4, 0x56, 0x0b, 0x57, 0x1e,
		0x5a, 0x8b, 0x80, 0xd6, 0x46, 0x11, 0x4d, 0x2a, 0x56, 0x19, 0xa2, 0xf2, 0x76, 0x1a, 0x16, 0x26,
		0x8c, 0x22, 0x04, 0x69, 0x0b, 0xf7, 0x29, 0x63, 0x62, 0x35, 0xaf, 0xb2, 0xff, 0x51, 0x19, 0x66,
		0x1c, 0xac, 0xdf, 0xc2, 0x5d, 0x52, 0x4e, 0x32, 0xb1, 0x7c, 0x44, 0xe7, 0x01, 0xda, 0xc4, 0x21,
		0x56, 0x9b, 0x58, 0xfa, 0x71, 0x39, 0xb5, 0x92, 0x5a, 0xcd, 0xab, 0x21, 0x09, 0x7a, 0x1c, 0xe6,
		0x9d, 0xc1, 0x91, 0x69, 0xe8, 0x5a, 0x48, 0x0d, 0x56, 0x52, 0xab, 0x19, 0x55, 0xe1, 0x03, 0x9b,
		0x43, 0xe5, 0x47, 0x61, 0xee, 0x0e, 0xc1, 0xb7, 0xc2, 0xaa, 0x05, 0xa6, 0x5a, 0xa2, 0xe2, 0x90,
		0xe2, 0x06, 0x14, 0xfb, 0xc4, 0xf3, 0x70, 0x97, 0x68, 0xfe, 0xb1, 0x43, 0xca, 0x69, 0x36, 0xfb,
		0x95, 0xb1, 0xd9, 0x47, 0x67, 0x5e, 0x10, 0xa8, 0x83, 0x63, 0x87, 0xa0, 0x1a, 0xe4, 0x89, 0x35,
		0xe8, 0x73, 0x86, 0xcc, 0x09, 0xfe, 0xab, 0x5b, 0x83, 0x7e, 0x94, 0x25, 0x47, 0x61, 0x82, 0x62,
		0xc6, 0x23, 0xee, 0x6d, 0x43, 0x27, 0xe5, 0x2c, 0x23, 0x78, 0x74, 0x8c, 0xa0, 0xc5, 0xc7, 0xa3,
		0x1c, 0x12, 0x87, 0x36, 0x20, 0x4f, 0x5e, 0xf6, 0x89, 0xe5, 0x19, 0xb6, 0x55, 0x9e, 0x61, 0x24,
		0x0f, 0x4f, 0x58, 0x45, 0x62, 0xb6, 0xa3, 0x14, 0x43, 0x1c, 0xba, 0x06, 0x33, 0xb6, 0xe3, 0x1b,
		0xb6, 0xe5, 0x95, 0x73, 0x2b, 0x89, 0xd5, 0xc2, 0x95, 0x8f, 0x4c, 0x0c, 0x84, 0x7d, 0xae, 0xa3,
		0x4a, 0x65, 0xd4, 0x00, 0xc5, 0xb3, 0x07, 0xae, 0x4e, 0x34, 0xdd, 0x6e, 0x13, 0xcd, 0xb0, 0x3a,
		0x76, 0x39, 0xcf, 0x08, 0x2e, 0x8c, 0x4f, 0x84, 0x29, 0x6e, 0xd8, 0x6d, 0xd2, 0xb0, 0x3a, 0xb6,
		0x5a, 0xf2, 0x46, 0x9e, 0xd1, 0x12, 0x64, 0xbd, 0x63, 0xcb, 0xc7, 0x2f, 0x97, 0x8b, 0x2c, 0x42,
		0xc4, 0x53, 0xe5, 0x1b, 0x59, 0x98, 0x9b, 0x26, 0xc4, 0x6e, 0x40, 0xa6, 0x43, 0x67, 0x59, 0x4e,
		0x9e, 0xc6, 0x07, 0x1c, 0x33, 0xea, 0xc4, 0xec, 0xf7, 0xe9, 0xc4, 0x1a, 0x14, 0x2c, 0xe2, 0xf9,
		0xa4, 0xcd, 0x23, 0x22, 0x35, 0x65, 0x4c, 0x01, 0x07, 0x8d, 0x87, 0x54, 0xfa, 0xfb, 0x0a, 0xa9,
		0x17, 0x60, 0x2e, 0x30, 0x49, 0x73, 0xb1, 0xd5, 0x95, 0xb1, 0x79, 0x39, 0xce, 0x92, 0xb5, 0xba,
		0xc4, 0xa9, 0x14, 0xa6, 0x96, 0xc8, 0xc8, 0x33, 0xda, 0x04, 0xb0, 0x2d, 0x62, 0x77, 0xb4, 0x36,
		0xd1, 0xcd, 0x72, 0xee, 0x04, 0x2f, 0xed, 0x53, 0x95, 0x31, 0x2f, 0xd9, 0x5c, 0xaa, 0x9b, 0xe8,
		0xfa, 0x30, 0xd4, 0x66, 0x4e, 0x88, 0x94, 0x5d, 0xbe, 0xc9, 0xc6, 0xa2, 0xed, 0x10, 0x4a, 0x2e,
		0xa1, 0x71, 0x4f, 0xda, 0x62, 0x66, 0x79, 0x66, 0xc4, 0x5a, 0xec, 0xcc, 0x54, 0x01, 0xe3, 0x13,
		0x9b, 0x75, 0xc3, 0x8f, 0xe8, 0x41, 0x08, 0x04, 0x1a, 0x0b, 0x2b, 0x60, 0x59, 0xa8, 0x28, 0x85,
		0x7b, 0xb8, 0x4f, 0x96, 0xef, 0x42, 0x69, 0xd4, 0x3d, 0x68, 0x11, 0x32, 0x9e, 0x8f, 0x5d, 0x9f,
		0x45, 0x61, 0x46, 0xe5, 0x0f, 0x48, 0x81, 0x14, 0xb1, 0xda, 0x2c, 0xcb, 0x65, 0x54, 0xfa, 0x2f,
		0xfa, 0x91, 0xe1, 0x84, 0x53, 0x6c, 0xc2, 0x8f, 0x8c, 0xaf, 0xe8, 0x08, 0x73, 0x74, 0xde, 0xcb,
		0x4f, 0xc1, 0xec, 0xc8, 0x04, 0xa6, 0x7d, 0x75, 0xe5, 0xc7, 0xe1, 0xbe, 0x89, 0xd4, 0xe8, 0x05,
		0x58, 0x1c, 0x58, 0x86, 0xe5, 0x13, 0xd7, 0x71, 0x09, 0x8d, 0x58, 0xfe, 0xaa, 0xf2, 0xbf, 0xce,
		0x9c, 0x10, 0x73, 0x87, 0x61, 0x6d, 0xce, 0xa2, 0x2e, 0x0c, 0xc6, 0x85, 0x97, 0xf2, 0xb9, 0x7f,
		0x9b, 0x51, 0x5e, 0x79, 0xe5, 0x95, 0x57, 0x92, 0x95, 0xcf, 0x65, 0x61, 0x71, 0xd2, 0x9e, 0x99,
		0xb8, 0x7d, 0x97, 0x20, 0x6b, 0x0d, 0xfa, 0x47, 0xc4, 0x65, 0x4e, 0xca, 0xa8, 0xe2, 0x09, 0xd5,
		0x20, 0x63, 0xe2, 0x23, 0x62, 0x96, 0xd3, 0x2b, 0x89, 0xd5, 0xd2, 0x95, 0xc7, 0xa7, 0xda, 0x95,
		0x6b, 0x3b, 0x14, 0xa2, 0x72, 0x24, 0xfa, 0x04, 0xa4, 0x45, 0x8a, 0xa6, 0x0c, 0x97, 0xa6, 0x63,
		0xa0, 0x7b, 0x49, 0x65, 0x38, 0x74, 0x3f, 0xe4, 0xe9, 0x5f, 0x1e, 0x1b, 0x59, 0x66, 0x73, 0x8e,
		0x0a, 0x68, 0x5c, 0xa0, 0x65, 0xc8, 0xb1, 0x6d, 0xd2, 0x26, 0xb2, 0xb4, 0x05, 0xcf, 0x34, 0xb0,
		0xda, 0xa4, 0x83, 0x07, 0xa6, 0xaf, 0xdd, 0xc6, 0xe6, 0x80, 0xb0, 0x80, 0xcf, 0xab, 0x45, 0x21,
		0x7c, 0x8e, 0xca, 0xd0, 0x05, 0x28, 0xf0, 0x5d, 0x65, 0x58, 0x6d, 0xf2, 0x32, 0xcb, 0x9e, 0x19,
		0x95, 0x6f, 0xb4, 0x06, 0x95, 0xd0, 0xd7, 0xdf, 0xf4, 0x6c, 0x4b, 0x86, 0x26, 0x7b, 0x05, 0x15,
		0xb0, 0xd7, 0x3f, 0x15, 0x4d, 0xdc, 0x0f, 0x4c, 0x9e, 0x5e, 0x34, 0xa6, 0x2a, 0x5f, 0x4f, 0x42,
		0x9a, 0xe5, 0x8b, 0x39, 0x28, 0x1c, 0xbc, 0xd8, 0xac, 0x6b, 0x9b, 0xfb, 0x87, 0xeb, 0x3b, 0x75,
		0x25, 0x81, 0x4a, 0x00, 0x4c, 0xb0, 0xb5, 0xb3, 0x5f, 0x3b, 0x50, 0x92, 0xc1, 0x73, 0x63, 0xef,
		0xe0, 0xda, 0x93, 0x4a, 0x2a, 0x00, 0x1c, 0x72, 0x41, 0x3a, 0xac, 0xf0, 0xc4, 0x15, 0x25, 0x83,
		0x14, 0x28, 0x72, 0x82, 0xc6, 0x0b, 0xf5, 0xcd, 0x6b, 0x4f, 0x2a, 0xd9, 0x51, 0xc9, 0x13, 0x57,
		0x94, 0x19, 0x34, 0x0b, 0x79, 0x26, 0x59, 0xdf, 0xdf, 0xdf, 0x51, 0x72, 0x01, 0x67, 0xeb, 0x40,
		0x6d, 0xec, 0x6d, 0x2b, 0xf9, 0x80, 0x73, 0x5b, 0xdd, 0x3f, 0x6c, 0x2a, 0x10, 0x30, 0xec, 0xd6,
		0x5b, 0xad, 0xda, 0x76, 0x5d, 0x29, 0x04, 0x1a, 0xeb, 0x2f, 0x1e, 0xd4, 0x5b, 0x4a, 0x71, 0xc4,
		0xac, 0x27, 0xae, 0x28, 0xb3, 0xc1, 0x2b, 0xea, 0x7b, 0x87, 0xbb, 0x4a, 0x09, 0xcd, 0xc3, 0x2c,
		0x7f, 0x85, 0x34, 0x62, 0x2e, 0x22, 0xba, 0xf6, 0xa4, 0xa2, 0x0c, 0x0d, 0xe1, 0x2c, 0xf3, 0x23,
		0x82, 0x6b, 0x4f, 0x2a, 0xa8, 0xb2, 0x01, 0x19, 0x16, 0x5d, 0x08, 0x41, 0x69, 0xa7, 0xb6, 0x5e,
		0xdf, 0xd1, 0xf6, 0x9b, 0x07, 0x8d, 0xfd, 0xbd, 0xda, 0x8e, 0x92, 0x18, 0xca, 0xd4, 0xfa, 0xa7,
		0x0e, 0x1b, 0x6a, 0x7d, 0x53, 0x49, 0x86, 0x65, 0xcd, 0x7a, 0xed, 0xa0, 0xbe, 0xa9, 0xa4, 0x2a,
		0x3a, 0x2c, 0x4e, 0xca, 0x93, 0x13, 0x77, 0x46, 0x68, 0x89, 0x93, 0x27, 0x2c, 0x31, 0xe3, 0x1a,
		0x5b, 0xe2, 0xb7, 0x92, 0xb0, 0x30, 0xa1, 0x56, 0x4c, 0x7c, 0xc9, 0x33, 0x90, 0xe1, 0x21, 0xca,
		0xab, 0xe7, 0x63, 0x13, 0x8b, 0x0e, 0x0b, 0xd8, 0xb1, 0x0a, 0xca, 0x70, 0xe1, 0x0e, 0x22, 0x75,
		0x42, 0x07, 0x41, 0x29, 0xc6, 0x72, 0xfa, 0x8f, 0x8d, 0xe5, 0x74, 0x5e, 0xf6, 0xae, 0x4d, 0x53,
		0xf6, 0x98, 0xec, 0x74, 0xb9, 0x3d, 0x33, 0x21, 0xb7, 0xdf, 0x80, 0xf9, 0x31, 0xa2, 0xa9, 0x73,
		0xec, 0xab, 0x09, 0x28, 0x9f, 0xe4, 0x9c, 0x98, 0x4c, 0x97, 0x1c, 0xc9, 0x74, 0x37, 0xa2, 0x1e,
		0xbc, 0x78, 0xf2, 0x22, 0x8c, 0xad, 0xf5, 0x9b, 0x09, 0x58, 0x9a, 0xdc, 0x29, 0x4e, 0xb4, 0xe1,
		0x13, 0x90, 0xed, 0x13, 0xbf, 0x67, 0xcb, 0x6e, 0xe9, 0x91, 0x09, 0x35, 0x98, 0x0e, 0x47, 0x17,
		0x5b, 0xa0, 0xc2, 0x45, 0x3c, 0x75, 0x52, 0xbb, 0xc7, 0xad, 0x19, 0xb3, 0xf4, 0xb3, 0x49, 0xb8,
		0x6f, 0x22, 0xf9, 0x44, 0x43, 0x1f, 0x00, 0x30, 0x2c, 0x67, 0xe0, 0xf3, 0x8e, 0x88, 0x27, 0xd8,
		0x3c, 0x93, 0xb0, 0xe4, 0x45, 0x93, 0xe7, 0xc0, 0x0f, 0xc6, 0x53, 0x6c, 0x1c, 0xb8, 0x88, 0x29,
		0x3c, 0x3d, 0x34, 0x34, 0xcd, 0x0c, 0x3d, 0x7f, 0xc2, 0x4c, 0xc7, 0x02, 0xf3, 0x63, 0xa0, 0xe8,
		0xa6, 0x41, 0x2c, 0x5f, 0xf3, 0x7c, 0x97, 0xe0, 0xbe, 0x61, 0x75, 0x59, 0x05, 0xc9, 0x55, 0x33,
		0x1d, 0x6c, 0x7a, 0x44, 0x9d, 0xe3, 0xc3, 0x2d, 0x39, 0x4a, 0x11, 0x2c, 0x80, 0xdc, 0x10, 0x22,
		0x3b, 0x82, 0xe0, 0xc3, 0x01, 0xa2, 0xf2, 0xb5, 0x1c, 0x14, 0x42, 0x7d, 0x35, 0xba, 0x08, 0xc5,
		0x9b, 0xf8, 0x36, 0xd6, 0xe4, 0x59, 0x89, 0x7b, 0xa2, 0x40, 0x65, 0x4d, 0x71, 0x5e, 0xfa, 0x18,
		0x2c, 0x32, 0x15, 0x7b, 0xe0, 0x13, 0x57, 0xd3, 0x4d, 0xec, 0x79, 0xcc, 0x69, 0x39, 0xa6, 0x8a,
		0xe8, 0xd8, 0x3e, 0x1d, 0xda, 0x90, 0x23, 0xe8, 0x2a, 0x2c, 0x30, 0x44, 0x7f, 0x60, 0xfa, 0x86,
		0x63, 0x12, 0x8d, 0x9e, 0xde, 0x3c, 0x56, 0x49, 0x02, 0xcb, 0xe6, 0xa9, 0xc6, 0xae, 0x50, 0xa0,
		0x16, 0x79, 0x68, 0x13, 0x1e, 0x60, 0xb0, 0x2e, 0xb1, 0x88, 0x8b, 0x7d, 0xa2, 0x91, 0xcf, 0x0c,
		0xb0, 0xe9, 0x69, 0xd8, 0x6a, 0x6b, 0x3d, 0xec, 0xf5, 0xca, 0x8b, 0x94, 0x60, 0x3d, 0x59, 0x4e,
		0xa8, 0xe7, 0xa8, 0xe2, 0xb6, 0xd0, 0xab, 0x33, 0xb5, 0x9a, 0xd5, 0xfe, 0x24, 0xf6, 0x7a, 0xa8,
		0x0a, 0x4b, 0x8c, 0xc5, 0xf3, 0x5d, 0xc3, 0xea, 0x6a, 0x7a, 0x8f, 0xe8, 0xb7, 0xb4, 0x81, 0xdf,
		0x79, 0xba, 0x7c, 0x7f, 0xf8, 0xfd, 0xcc, 0xc2, 0x16, 0xd3, 0xd9, 0xa0, 0x2a, 0x87, 0x7e, 0xe7,
		0x69, 0xd4, 0x82, 0x22, 0x5d, 0x8c, 0xbe, 0x71, 0x97, 0x68, 0x1d, 0xdb, 0x65, 0xa5, 0xb1, 0x34,
		0x21, 0x35, 0x85, 0x3c, 0xb8, 0xb6, 0x2f, 0x00, 0xbb, 0x76, 0x9b, 0x54, 0x33, 0xad, 0x66, 0xbd,
		0xbe, 0xa9, 0x16, 0x24, 0xcb, 0x96, 0xed, 0xd2, 0x80, 0xea, 0xda, 0x81, 0x83, 0x0b, 0x3c, 0xa0,
		0xba, 0xb6, 0x74, 0xef, 0x55, 0x58, 0xd0, 0x75, 0x3e, 0x67, 0x43, 0xd7, 0xc4, 0x19, 0xcb, 0x2b,
		0x2b, 0x23, 0xce, 0xd2, 0xf5, 0x6d, 0xae, 0x20, 0x62, 0xdc, 0x43, 0xd7, 0xe1, 0xbe, 0xa1, 0xb3,
		0xc2, 0xc0, 0xf9, 0xb1, 0x59, 0x46, 0xa1, 0x57, 0x61, 0xc1, 0x39, 0x1e, 0x07, 0xa2, 0x91, 0x37,
		0x3a, 0xc7, 0x51, 0xd8, 0x53, 0xb0, 0xe8, 0xf4, 0x9c, 0x71, 0xdc, 0xa5, 0x30, 0x0e, 0x39, 0x3d,
		0x27, 0x0a, 0x7c, 0x98, 0x1d, 0xb8, 0x5d, 0xa2, 0x63, 0x9f, 0xb4, 0xcb, 0x67, 0xc3, 0xea, 0xa1,
		0x01, 0x74, 0x19, 0x14, 0x5d, 0xd7, 0x88, 0x85, 0x8f, 0x4c, 0xa2, 0x61, 0x97, 0x58, 0xd8, 0x2b,
		0x5f, 0x08, 0x2b, 0x97, 0x74, 0xbd, 0xce, 0x46, 0x6b, 0x6c, 0x10, 0x5d, 0x82, 0x79, 0xfb, 0xe8,
		0xa6, 0xce, 0x43, 0x52, 0x73, 0x5c, 0xd2, 0x31, 0x5e, 0x2e, 0x3f, 0xc4, 0xfc, 0x3b, 0x47, 0x07,
		0x58, 0x40, 0x36, 0x99, 0x18, 0x3d, 0x06, 0x8a, 0xee, 0xf5, 0xb0, 0xeb, 0xb0, 0x9c, 0xec, 0x39,
		0x58, 0x27, 0xe5, 0x87, 0xb9, 0x2a, 0x97, 0xef, 0x49, 0x31, 0xdd, 0x12, 0xde, 0x1d, 0xa3, 0xe3,
		0x4b, 0xc6, 0x47, 0xf9, 0x96, 0x60, 0x32, 0xc1, 0xb6, 0x0a, 0x0a, 0x75, 0xc5, 0xc8, 0x8b, 0x57,
		0x99, 0x5a, 0xc9, 0xe9, 0x39, 0xe1, 0xf7, 0x3e, 0x08, 0xb3, 0x54, 0x73, 0xf8, 0xd2, 0xc7, 0x78,
		0x43, 0xe6, 0xf4, 0x42, 0x6f, 0xfc, 0xc0, 0x7a, 0xe3, 0x4a, 0x15, 0x8a, 0xe1, 0xf8, 0x44, 0x79,
		0xe0, 0x11, 0xaa, 0x24, 0x68, 0xb3, 0xb2, 0xb1, 0xbf, 0x49, 0xdb, 0x8c, 0x97, 0xea, 0x4a, 0x92,
		0xb6, 0x3b, 0x3b, 0x8d, 0x83, 0xba, 0xa6, 0x1e, 0xee, 0x1d, 0x34, 0x76, 0xeb, 0x4a, 0x2a, 0xdc,
		0x57, 0x7f, 0x2b, 0x09, 0xa5, 0xd1, 0x23, 0x12, 0xfa, 0x61, 0x38, 0x2b, 0xef, 0x33, 0x3c, 0xe2,
		0x6b, 0x77, 0x0c, 0x97, 0x6d, 0x99, 0x3e, 0xe6, 0xe5, 0x2b, 0x58, 0xb4, 0x45, 0xa1, 0xd5, 0x22,
		0xfe, 0xf3, 0x86, 0x4b, 0x37, 0x44, 0x1f, 0xfb, 0x68, 0x07, 0x2e, 0x58, 0xb6, 0xe6, 0xf9, 0xd8,
		0x6a, 0x63, 0xb7, 0xad, 0x0d, 0x6f, 0x92, 0x34, 0xac, 0xeb, 0xc4, 0xf3, 0x6c, 0x5e, 0xaa, 0x02,
		0x96, 0x8f, 0x58, 0x76, 0x4b, 0x28, 0x0f, 0x73, 0x78, 0x4d, 0xa8, 0x46, 0x02, 0x2c, 0x75, 0x52,
		0x80, 0xdd, 0x0f, 0xf9, 0x3e, 0x76, 0x34, 0x62, 0xf9, 0xee, 0x31, 0x6b, 0x8c, 0x73, 0x6a, 0xae,
		0x8f, 0x9d, 0x3a, 0x7d, 0xfe, 0x70, 0xce, 0x27, 0xff, 0x98, 0x82, 0x62, 0xb8, 0x39, 0xa6, 0x67,
		0x0d, 0x9d, 0xd5, 0x91, 0x04, 0xcb, 0x34, 0x0f, 0xde, 0xb3, 0x95, 0x5e, 0xdb, 0xa0, 0x05, 0xa6,
		0x9a, 0xe5, 0x2d, 0xab, 0xca, 0x91, 0xb4, 0xb8, 0xd3, 0xdc, 0x42, 0x78, 0x8b, 0x90, 0x53, 0xc5,
		0x13, 0xda, 0x86, 0xec, 0x4d, 0x8f, 0x71, 0x67, 0x19, 0xf7, 0x43, 0xf7, 0xe6, 0x7e, 0xb6, 0xc5,
		0xc8, 0xf3, 0xcf, 0xb6, 0xb4, 0xbd, 0x7d, 0x75, 0xb7, 0xb6, 0xa3, 0x0a, 0x38, 0x3a, 0x07, 0x69,
		0x13, 0xdf, 0x3d, 0x1e, 0x2d, 0x45, 0x4c, 0x34, 0xad, 0xe3, 0xcf, 0x41, 0xfa, 0x0e, 0xc1, 0xb7,
		0x46, 0x0b, 0x00, 0x13, 0x7d, 0x80, 0xa1, 0x7f, 0x19, 0x32, 0xcc, 0x5f, 0x08, 0x40, 0x78, 0x4c,
		0x39, 0x83, 0x72, 0x90, 0xde, 0xd8, 0x57, 0x69, 0xf8, 0x2b, 0x50, 0xe4, 0x52, 0xad, 0xd9, 0xa8,
		0x6f, 0xd4, 0x95, 0x64, 0xe5, 0x2a, 0x64, 0xb9, 0x13, 0xe8, 0xd6, 0x08, 0xdc, 0xa0, 0x9c, 0x11,
		0x8f, 0x82, 0x23, 0x21, 0x47, 0x0f, 0x77, 0xd7, 0xeb, 0xaa, 0x92, 0x0c, 0x2f, 0xaf, 0x07, 0xc5,
		0x70, 0x5f, 0xfc, 0xe1, 0xc4, 0xd4, 0x37, 0x13, 0x50, 0x08, 0xf5, 0xb9, 0xb4, 0x41, 0xc1, 0xa6,
		0x69, 0xdf, 0xd1, 0xb0, 0x69, 0x60, 0x4f, 0x04, 0x05, 0x30, 0x51, 0x8d, 0x4a, 0xa6, 0x5d, 0xb4,
		0x0f, 0xc5, 0xf8, 0x37, 0x12, 0xa0, 0x44, 0x5b, 0xcc, 0x88, 0x81, 0x89, 0x1f, 0xa8, 0x81, 0xaf,
		0x27, 0xa0, 0x34, 0xda, 0x57, 0x46, 0xcc, 0xbb, 0xf8, 0x03, 0x35, 0xef, 0x9f, 0x92, 0x30, 0x3b,
		0xd2, 0x4d, 0x4e, 0x6b, 0xdd, 0x67, 0x60, 0xde, 0x68, 0x93, 0xbe, 0x63, 0xfb, 0xc4, 0xd2, 0x8f,
		0x35, 0x93, 0xdc, 0x26, 0x66, 0xb9, 0xc2, 0x12, 0xc5, 0xe5, 0x7b, 0xf7, 0xab, 0x6b, 0x8d, 0x21,
		0x6e, 0x87, 0xc2, 0xaa, 0x0b, 0x8d, 0xcd, 0xfa, 0x6e, 0x73, 0xff, 0xa0, 0xbe, 0xb7, 0xf1, 0xa2,
		0x76, 0xb8, 0xf7, 0xa3, 0x7b, 0xfb, 0xcf, 0xef, 0xa9, 0x8a, 0x11, 0x51, 0xfb, 0x00, 0xb7, 0x7a,
		0x13, 0x94, 0xa8, 0x51, 0xe8, 0x2c, 0x4c, 0x32, 0x4b, 0x39, 0x83, 0x16, 0x60, 0x6e, 0x6f, 0x5f,
		0x6b, 0x35, 0x36, 0xeb, 0x5a, 0x7d, 0x6b, 0xab, 0xbe, 0x71, 0xd0, 0xe2, 0x37, 0x10, 0x81, 0xf6,
		0xc1, 0xe8, 0xa6, 0x7e, 0x2d, 0x05, 0x0b, 0x13, 0x2c, 0x41, 0x35, 0x71, 0x76, 0xe0, 0xc7, 0x99,
		0x8f, 0x4e, 0x63, 0xfd, 0x1a, 0x2d, 0xf9, 0x4d, 0xec, 0xfa, 0xe2, 0xa8, 0xf1, 0x18, 0x50, 0x2f,
		0x59, 0xbe, 0xd1, 0x31, 0x88, 0x2b, 0x2e, 0x6c, 0xf8, 0x81, 0x62, 0x6e, 0x28, 0xe7, 0x77, 0x36,
		0x3f, 0x04, 0xc8, 0xb1, 0x3d, 0xc3, 0x37, 0x6e, 0x13, 0xcd, 0xb0, 0xe4, 0xed, 0x0e, 0x3d, 0x60,
		0xa4, 0x55, 0x45, 0x8e, 0x34, 0x2c, 0x3f, 0xd0, 0xb6, 0x48, 0x17, 0x47, 0xb4, 0x69, 0x02, 0x4f,
		0xa9, 0x8a, 0x1c, 0x09, 0xb4, 0x2f, 0x42, 0xb1, 0x6d, 0x0f, 0x68, 0xd7, 0xc5, 0xf5, 0x68, 0xbd,
		0x48, 0xa8, 0x05, 0x2e, 0x0b, 0x54, 0x44, 0x3f, 0x3d, 0xbc, 0x56, 0x2a, 0xaa, 0x05, 0x2e, 0xe3,
		0x2a, 0x8f, 0xc2, 0x1c, 0xee, 0x76, 0x5d, 0x4a, 0x2e, 0x89, 0xf8, 0x09, 0xa1, 0x14, 0x88, 0x99,
		0xe2, 0xf2, 0xb3, 0x90, 0x93, 0x7e, 0xa0, 0x25, 0x99, 0x7a, 0x42, 0x73, 0xf8, 0xb1, 0x37, 0xb9,
		0x9a, 0x57, 0x73, 0x96, 0x1c, 0xbc, 0x08, 0x45, 0xc3, 0xd3, 0x86, 0xb7, 0xe4, 0xc9, 0x95, 0xe4,
		0x6a, 0x4e, 0x2d, 0x18, 0x5e, 0x70, 0xc3, 0x58, 0x79, 0x33, 0x09, 0xa5, 0xd1, 0x5b, 0x7e, 0xb4,
		0x09, 0x39, 0xd3, 0xd6, 0x31, 0x0b, 0x2d, 0xfe, 0x89, 0x69, 0x35, 0xe6, 0xc3, 0xc0, 0xda, 0x8e,
		0xd0, 0x57, 0x03, 0xe4, 0xf2, 0xdf, 0x25, 0x20, 0x27, 0xc5, 0x68, 0x09, 0xd2, 0x0e, 0xf6, 0x7b,
		0x8c, 0x2e, 0xb3, 0x9e, 0x54, 0x12, 0x2a, 0x7b, 0xa6, 0x72, 0xcf, 0xc1, 0x16, 0x0b, 0x01, 0x21,
		0xa7, 0xcf, 0x74, 0x5d, 0x4d, 0x82, 0xdb, 0xec, 0xf8, 0x61, 0xf7, 0xfb, 0xc4, 0xf2, 0x3d, 0xb9,
		0xae, 0x42, 0xbe, 0x21, 0xc4, 0xe8, 0x71, 0x98, 0xf7, 0x5d, 0x6c, 0x98, 0x23, 0xba, 0x69, 0xa6,
		0xab, 0xc8, 0x81, 0x40, 0xb9, 0x0a, 0xe7, 0x24, 0x6f, 0x9b, 0xf8, 0x58, 0xef, 0x91, 0xf6, 0x10,
		0x94, 0x65, 0xd7, 0x0c, 0x67, 0x85, 0xc2, 0xa6, 0x18, 0x97, 0xd8, 0xca, 0xb7, 0x13, 0x30, 0x2f,
		0x0f, 0x4c, 0xed, 0xc0, 0x59, 0xbb, 0x00, 0xd8, 0xb2, 0x6c, 0x3f, 0xec, 0xae, 0xf1, 0x50, 0x1e,
		0xc3, 0xad, 0xd5, 0x02, 0x90, 0x1a, 0x22, 0x58, 0xee, 0x03, 0x0c, 0x47, 0x4e, 0x74, 0xdb, 0x05,
		0x28, 0x88, 0x4f, 0x38, 0xec, 0x3b, 0x20, 0x3f, 0x62, 0x03, 0x17, 0xd1, 0x93, 0x15, 0x5a, 0x84,
		0xcc, 0x11, 0xe9, 0x1a, 0x96, 0xb8, 0x98, 0xe5, 0x0f, 0xf2, 0x22, 0x24, 0x1d, 0x5c, 0x84, 0xac,
		0x7f, 0x1a, 0x16, 0x74, 0xbb, 0x1f, 0x35, 0x77, 0x5d, 0x89, 0x1c, 0xf3, 0xbd, 0x4f, 0x26, 0x5e,
		0x82, 0x61, 0x8b, 0xf9, 0xbd, 0x44, 0xe2, 0x4b, 0xc9, 0xd4, 0x76, 0x73, 0xfd, 0xcb, 0xc9, 0xe5,
		0x6d, 0x0e, 0x6d, 0xca, 0x99, 0xaa, 0xa4, 0x63, 0x12, 0x9d, 0x5a, 0x0f, 0xff, 0x72, 0x09, 0x3e,
		0xda, 0x35, 0xfc, 0xde, 0xe0, 0x68, 0x4d, 0xb7, 0xfb, 0x97, 0xbb, 0x76, 0xd7, 0x1e, 0x7e, 0xfa,
		0xa4, 0x4f, 0xec, 0x81, 0xfd, 0x27, 0x3e, 0x7f, 0xe6, 0x03, 0xe9, 0x72, 0xec, 0xb7, 0xd2, 0xea,
		0x1e, 0x2c, 0x08, 0x65, 0x8d, 0x7d, 0x7f, 0xe1, 0xa7, 0x08, 0x74, 0xcf, 0x3b, 0xac, 0xf2, 0x57,
		0xdf, 0x66, 0xe5, 0x5a, 0x9d, 0x17, 0x50, 0x3a, 0xc6, 0x0f, 0x1a, 0x55, 0x15, 0xee, 0x1b, 0xe1,
		0xe3, 0x5b, 0x93, 0xb8, 0x31, 0x8c, 0xdf, 0x12, 0x8c, 0x0b, 0x21, 0xc6, 0x96, 0x80, 0x56, 0x37,
		0x60, 0xf6, 0x34, 0x5c, 0x7f, 0x2d, 0xb8, 0x8a, 0x24, 0x4c, 0xb2, 0x0d, 0x73, 0x8c, 0x44, 0x1f,
		0x78, 0xbe, 0xdd, 0x67, 0x79, 0xef, 0xde, 0x34, 0x7f, 0xf3, 0x36, 0xdf, 0x2b, 0x25, 0x0a, 0xdb,
		0x08, 0x50, 0xd5, 0x2a, 0xb0, 0x4f, 0x4e, 0x6d, 0xa2, 0x9b, 0x31, 0x0c, 0x7f, 0x2b, 0x0c, 0x09,
		0xf4, 0xab, 0xcf, 0xc1, 0x22, 0xfd, 0x9f, 0xa5, 0xa5, 0xb0, 0x25, 0xf1, 0x17, 0x5e, 0xe5, 0x6f,
		0xbf, 0xca, 0xb7, 0xe3, 0x42, 0x40, 0x10, 0xb2, 0x29, 0xb4, 0x8a, 0x5d, 0xe2, 0xfb, 0xc4, 0xf5,
		0x34, 0x6c, 0x4e, 0x32, 0x2f, 0x74, 0x63, 0x50, 0xfe, 0xfc, 0x3b, 0xa3, 0xab, 0xb8, 0xcd, 0x91,
		0x35, 0xd3, 0xac, 0x1e, 0xc2, 0xd9, 0x09, 0x51, 0x31, 0x05, 0xe7, 0x6b, 0x82, 0x73, 0x71, 0x2c,
		0x32, 0x28, 0x6d, 0x13, 0xa4, 0x3c, 0x58, 0xcb, 0x29, 0x38, 0x7f, 0x5d, 0x70, 0x22, 0x81, 0x95,
		0x4b, 0x4a, 0x19, 0x9f, 0x85, 0xf9, 0xdb, 0xc4, 0x3d, 0xb2, 0x3d, 0x71, 0x4b, 0x33, 0x05, 0xdd,
		0xeb, 0x82, 0x6e, 0x4e, 0x00, 0xd9, 0xb5, 0x0d, 0xe5, 0xba, 0x0e, 0xb9, 0x0e, 0xd6, 0xc9, 0x14,
		0x14, 0x5f, 0x10, 0x14, 0x33, 0x54, 0x9f, 0x42, 0x6b, 0x50, 0xec, 0xda, 0xa2, 0x32, 0xc5, 0xc3,
		0xdf, 0x10, 0xf0, 0x82, 0xc4, 0x08, 0x0a, 0xc7, 0x76, 0x06, 0x26, 0x2d, 0x5b, 0xf1, 0x14, 0xbf,
		0x21, 0x29, 0x24, 0x46, 0x50, 0x9c, 0xc2, 0xad, 0xbf, 0x29, 0x29, 0xbc, 0x90, 0x3f, 0x9f, 0x81,
		0x82, 0x6d, 0x99, 0xc7, 0xb6, 0x35, 0x8d, 0x11, 0x5f, 0x14, 0x0c, 0x20, 0x20, 0x94, 0xe0, 0x06,
		0xe4, 0xa7, 0x5d, 0x88, 0xdf, 0x7e, 0x47, 0x6e, 0x0f, 0xb9, 0x02, 0xdb, 0x30, 0x27, 0x13, 0x94,
		0x61, 0x5b, 0x53, 0x50, 0xfc, 0x8e, 0xa0, 0x28, 0x85, 0x60, 0x62, 0x1a, 0x3e, 0xf1, 0xfc, 0x2e,
		0x99, 0x86, 0xe4, 0x4d, 0x39, 0x0d, 0x01, 0x11, 0xae, 0x3c, 0x22, 0x96, 0xde, 0x9b, 0x8e, 0xe1,
		0x77, 0xa5, 0x2b, 0x25, 0x86, 0x52, 0x6c, 0xc0, 0x6c, 0x1f, 0xbb, 0x5e, 0x0f, 0x9b, 0x53, 0x2d,
		0xc7, 0xef, 0x09, 0x8e, 0x62, 0x00, 0x12, 0x1e, 0x19, 0x58, 0xa7, 0xa1, 0xf9, 0xb2, 0xf4, 0x48,
		0x08, 0x26, 0xb6, 0x9e, 0xe7, 0xb3, 0x2b, 0xad, 0xd3, 0xb0, 0xfd, 0xbe, 0xdc, 0x7a, 0x1c, 0xbb,
		0x1b, 0x66, 0xbc, 0x01, 0x79, 0xcf, 0xb8, 0x3b, 0x15, 0xcd, 0x1f, 0xc8, 0x95, 0x66, 0x00, 0x0a,
		0x7e, 0x11, 0xce, 0x4d, 0x2c, 0x13, 0x53, 0x90, 0xfd, 0xa1, 0x20, 0x5b, 0x9a, 0x50, 0x2a, 0x44,
		0x4a, 0x38, 0x2d, 0xe5, 0x1f, 0xc9, 0x94, 0x40, 0x22, 0x5c, 0x4d, 0x7a, 0x56, 0xf0, 0x70, 0xe7,
		0x74, 0x5e, 0xfb, 0x63, 0xe9, 0x35, 0x8e, 0x1d, 0xf1, 0xda, 0x01, 0x2c, 0x09, 0xc6, 0xd3, 0xad,
		0xeb, 0x57, 0x64, 0x62, 0xe5, 0xe8, 0xc3, 0xd1, 0xd5, 0xfd, 0x34, 0x2c, 0x07, 0xee, 0x94, 0x4d,
		0xa9, 0xa7, 0xf5, 0xb1, 0x33, 0x05, 0xf3, 0x57, 0x05, 0xb3, 0xcc, 0xf8, 0x41, 0x57, 0xeb, 0xed,
		0x62, 0x87, 0x92, 0xbf, 0x00, 0x65, 0x49, 0x3e, 0xb0, 0x5c, 0xa2, 0xdb, 0x5d, 0xcb, 0xb8, 0x4b,
		0xda, 0x53, 0x50, 0xff, 0x49, 0x64, 0xa9, 0x0e, 0x43, 0x70, 0xca, 0xdc, 0x00, 0x25, 0xe8, 0x55,
		0x34, 0xa3, 0xef, 0xd8, 0xae, 0x1f, 0xc3, 0xf8, 0x35, 0xb9, 0x52, 0x01, 0xae, 0xc1, 0x60, 0xd5,
		0x3a, 0x94, 0xd8, 0xe3, 0xb4, 0x21, 0xf9, 0xa7, 0x82, 0x68, 0x76, 0x88, 0x12, 0x89, 0x43, 0xb7,
		0xfb, 0x0e, 0x76, 0xa7, 0xc9, 0x7f, 0x7f, 0x26, 0x13, 0x87, 0x80, 0x88, 0xc4, 0xe1, 0x1f, 0x3b,
		0x84, 0x56, 0xfb, 0x29, 0x18, 0xbe, 0x2e, 0x13, 0x87, 0xc4, 0x08, 0x0a, 0xd9, 0x30, 0x4c, 0x41,
		0xf1, 0xe7, 0x92, 0x42, 0x62, 0x28, 0xc5, 0xa7, 0x86, 0x85, 0xd6, 0x25, 0x5d, 0xc3, 0xf3, 0x5d,
		0xde, 0x0a, 0xdf, 0x9b, 0xea, 0x2f, 0xde, 0x19, 0x6d, 0xc2, 0xd4, 0x10, 0x94, 0x66, 0x22, 0x71,
		0x85, 0xca, 0x4e, 0x4a, 0xf1, 0x86, 0x7d, 0x43, 0x66, 0xa2, 0x10, 0x8c, 0xda, 0x16, 0xea, 0x10,
		0xa9, 0xdb, 0x75, 0x7a, 0x3e, 0x98, 0x82, 0xee, 0x9b, 0x11, 0xe3, 0x5a, 0x12, 0x4b, 0x39, 0x43,
		0xfd, 0xcf, 0xc0, 0xba, 0x45, 0x8e, 0xa7, 0x8a, 0xce, 0xbf, 0x8c, 0xf4, 0x3f, 0x87, 0x1c, 0xc9,
		0x73, 0xc8, 0x5c, 0xa4, 0x9f, 0x42, 0x71, 0x3f, 0xd6, 0x29, 0xff, 0xc4, 0x7b, 0x62, 0xbe, 0xa3,
		0xed, 0x54, 0x75, 0x87, 0x06, 0xf9, 0x68, 0xd3, 0x13, 0x4f, 0xf6, 0xea, 0x7b, 0x41, 0x9c, 0x8f,
		0xf4, 0x3c, 0xd5, 0x2d, 0x98, 0x1d, 0x69, 0x78, 0xe2, 0xa9, 0x7e, 0x52, 0x50, 0x15, 0xc3, 0xfd,
		0x4e, 0xf5, 0x2a, 0xa4, 0x69, 0xf3, 0x12, 0x0f, 0xff, 0x29, 0x01, 0x67, 0xea, 0xd5, 0x8f, 0x43,
		0x4e, 0x36, 0x2d, 0xf1, 0xd0, 0x9f, 0x16, 0xd0, 0x00, 0x42, 0xe1, 0xb2, 0x61, 0x89, 0x87, 0xff,
		0x8c, 0x84, 0x4b, 0x08, 0x85, 0x4f, 0xef, 0xc2, 0xbf, 0xfa, 0xd9, 0xb4, 0x28, 0x3a, 0xd2, 0x77,
		0x37, 0x60, 0x46, 0x74, 0x2a, 0xf1, 0xe8, 0xcf, 0x8a, 0x97, 0x4b, 0x44, 0xf5, 0x29, 0xc8, 0x4c,
		0xe9, 0xf0, 0x9f, 0x13, 0x50, 0xae, 0x5f, 0xdd, 0x80, 0x42, 0xa8, 0x3b, 0x89, 0x87, 0xff, 0xbc,
		0x80, 0x87, 0x51, 0xd4, 0x74, 0xd1, 0x9d, 0xc4, 0x13, 0xfc, 0x82, 0x34, 0x5d, 0x20, 0xa8, 0xdb,
		0x64, 0x63, 0x12, 0x8f, 0xfe, 0x45, 0xe9, 0x75, 0x09, 0xa9, 0x3e, 0x03, 0xf9, 0xa0, 0xd8, 0xc4,
		0xe3, 0x7f, 0x49, 0xe0, 0x87, 0x18, 0xea, 0x81, 0x50, 0xb1, 0x8b, 0xa7, 0xf8, 0x65, 0xe9, 0x81,
		0x10, 0x8a, 0x6e, 0xa3, 0x68, 0x03, 0x13, 0xcf, 0xf4, 0x2b, 0x72, 0x1b, 0x45, 0xfa, 0x17, 0xba,
		0x9a, 0x2c, 0xe7, 0xc7, 0x53, 0xfc, 0xaa, 0x5c, 0x4d, 0xa6, 0x4f, 0xcd, 0x88, 0x76, 0x04, 0xf1,
		0x1c, 0xbf, 0x26, 0xcd, 0x88, 0x34, 0x04, 0xd5, 0x26, 0xa0, 0xf1, 0x6e, 0x20, 0x9e, 0xef, 0x73,
		0x82, 0x6f, 0x7e, 0xac, 0x19, 0xa8, 0x3e, 0x0f, 0x4b, 0x93, 0x3b, 0x81, 0x78, 0xd6, 0xcf, 0xbf,
		0x17, 0x39, 0xbb, 0x85, 0x1b, 0x81, 0xea, 0xc1, 0xb0, 0xa4, 0x84, 0xbb, 0x80, 0x78, 0xda, 0xd7,
		0xde, 0x1b, 0x4d, 0xdc, 0xe1, 0x26, 0xa0, 0x5a, 0x03, 0x18, 0x16, 0xe0, 0x78, 0xae, 0xd7, 0x05,
		0x57, 0x08, 0x44, 0xb7, 0x86, 0xa8, 0xbf, 0xf1, 0xf8, 0x2f, 0xc8, 0xad, 0x21, 0x10, 0x74, 0x6b,
		0xc8, 0xd2, 0x1b, 0x8f, 0x7e, 0x43, 0x6e, 0x0d, 0x09, 0xa1, 0x91, 0x1d, 0xaa, 0x6e, 0xf1, 0x0c,
		0x5f, 0x94, 0x91, 0x1d, 0x42, 0x55, 0xf7, 0x60, 0x7e, 0xac, 0x20, 0xc6, 0x53, 0x7d, 0x49, 0x50,
		0x29, 0xd1, 0x7a, 0x18, 0x2e, 0x5e, 0xa2, 0x18, 0xc6, 0xb3, 0xfd, 0x56, 0xa4, 0x78, 0x89, 0x5a,
		0x58, 0xbd, 0x01, 0x39, 0x6b, 0x60, 0x9a, 0x74, 0xf3, 0xa0, 0x7b, 0xff, 0xc0, 0xae, 0xfc, 0xef,
		0xef, 0x0b, 0xef, 0x48, 0x40, 0xf5, 0x2a, 0x64, 0x48, 0xff, 0x88, 0xb4, 0xe3, 0x90, 0xff, 0xf1,
		0xbe, 0x4c, 0x98, 0x54, 0xbb, 0xfa, 0x0c, 0x00, 0xbf, 0x1a, 0x61, 0x9f, 0xfd, 0x62, 0xb0, 0xff,
		0xf9, 0xbe, 0xf8, 0xe9, 0xcb, 0x10, 0x32, 0x24, 0xe0, 0x3f, 0xa4, 0xb9, 0x37, 0xc1, 0x3b, 0xa3,
		0x04, 0x6c, 0x45, 0xae, 0xc3, 0xcc, 0x4d, 0xcf, 0xb6, 0x7c, 0xdc, 0x8d, 0x43, 0xff, 0x97, 0x40,
		0x4b, 0x7d, 0xea, 0xb0, 0xbe, 0xed, 0x12, 0x1f, 0x77, 0xbd, 0x38, 0xec, 0x7f, 0x0b, 0x6c, 0x00,
		0xa0, 0x60, 0x1d, 0x7b, 0xfe, 0x34, 0xf3, 0xfe, 0xae, 0x04, 0x4b, 0x00, 0x35, 0x9a, 0xfe, 0x7f,
		0x8b, 0x1c, 0xc7, 0x61, 0xdf, 0x95, 0x46, 0x0b, 0xfd, 0xea, 0xc7, 0x21, 0x4f, 0xff, 0xe5, 0xbf,
		0x67, 0x8b, 0x01, 0xff, 0x8f, 0x00, 0x0f, 0x11, 0xf4, 0xcd, 0x9e, 0xdf, 0xf6, 0x8d, 0x78, 0x67,
		0xff, 0xaf, 0x58, 0x69, 0xa9, 0x5f, 0xad, 0x41, 0xc1, 0xf3, 0xdb, 0xed, 0x81, 0xe8, 0x4f, 0x63,
		0xe0, 0xff, 0xf7, 0x7e, 0x70, 0x65, 0x11, 0x60, 0xd6, 0xeb, 0x93, 0x6f, 0x5f, 0x61, 0xdb, 0xde,
		0xb6, 0xf9, 0xbd, 0xeb, 0x4b, 0x95, 0xf8, 0x0b, 0x54, 0x78, 0x37, 0x09, 0xf9, 0xbe, 0xd7, 0x15,
		0x77, 0xa8, 0x3c, 0xc1, 0xd0, 0x02, 0xea, 0x2d, 0x9f, 0xee, 0xfa, 0xb5, 0xb2, 0x05, 0xd9, 0x75,
		0xa3, 0xbb, 0xeb, 0x75, 0xd1, 0x22, 0x64, 0x98, 0xf9, 0xec, 0xdb, 0x61, 0x4a, 0xe5, 0x0f, 0xe8,
		0x11, 0x48, 0xed, 0x7a, 0x5d, 0xf1, 0xeb, 0xb1, 0xc5, 0xb5, 0xe1, 0x8b, 0xd6, 0x5a, 0x7d, 0x6c,
		0x9a, 0xbb, 0x5e, 0x57, 0xa5, 0x0a, 0x15, 0x17, 0x72, 0x52, 0x80, 0x56, 0xa0, 0xd0, 0xd2, 0xb1,
		0xbb, 0x3e, 0xf0, 0x5a, 0xbe, 0xed, 0xc8, 0x5f, 0x47, 0x85, 0x44, 0x68, 0x15, 0xe6, 0xb6, 0x4c,
		0xa3, 0xdb, 0xf3, 0x9b, 0xd8, 0xc5, 0x7a, 0x6f, 0xe0, 0x93, 0x72, 0x71, 0x25, 0xb5, 0x3a, 0xa3,
		0x46, 0xc5, 0x68, 0x19, 0x72, 0xbb, 0xd8, 0x69, 0xf5, 0xb0, 0x7b, 0x8b, 0xfd, 0xd6, 0x26, 0xaf,
		0x06, 0xcf, 0x95, 0xe7, 0x20, 0xdb, 0xe4, 0x9f, 0xed, 0x97, 0x20, 0xdd, 0xe0, 0xdf, 0x00, 0x52,
		0xab, 0x29, 0x7e, 0x69, 0x4e, 0x9f, 0xd1, 0x32, 0x64, 0xb7, 0x4c, 0x1b, 0xfb, 0x1e, 0xfb, 0x3d,
		0x61, 0x82, 0x8d, 0x08, 0x09, 0x2a, 0x43, 0xe6, 0xd0, 0x90, 0xdf, 0x00, 0x66, 0xd9, 0x10, 0x17,
		0xac, 0x2f, 0xbe, 0xfb, 0x9d, 0xf3, 0x89, 0xef, 0x7d, 0xe7, 0x7c, 0xe2, 0x2b, 0xff, 0x7c, 0x3e,
		0xf1, 0xf7, 0x6f, 0x9d, 0x3f, 0xf3, 0x0f, 0x6f, 0x9d, 0x3f, 0xf3, 0xff, 0x01, 0x00, 0x00, 0xff,
		0xff, 0x74, 0xad, 0xaa, 0xf5, 0xcd, 0x33, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *BigMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&prototests.BigMsg{")
	if this.Field != nil {
		s = append(s, "Field: "+valueToGoStringMsg(this.Field, "int64")+",\n")
	}
	if this.Msg != nil {
		s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SmallMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.SmallMsg{")
	if this.ScarBusStop != nil {
		s = append(s, "ScarBusStop: "+valueToGoStringMsg(this.ScarBusStop, "string")+",\n")
	}
	if this.FlightParachute != nil {
		s = append(s, "FlightParachute: "+fmt.Sprintf("%#v", this.FlightParachute)+",\n")
	}
	if this.MapShark != nil {
		s = append(s, "MapShark: "+valueToGoStringMsg(this.MapShark, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Packed) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.Packed{")
	if this.Ints != nil {
		s = append(s, "Ints: "+fmt.Sprintf("%#v", this.Ints)+",\n")
	}
	if this.Floats != nil {
		s = append(s, "Floats: "+fmt.Sprintf("%#v", this.Floats)+",\n")
	}
	if this.Uints != nil {
		s = append(s, "Uints: "+fmt.Sprintf("%#v", this.Uints)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMsg(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedBigMsg(r randyMsg, easy bool) *BigMsg {
	this := &BigMsg{}
	if r.Intn(10) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Field = &v1
	}
	if r.Intn(10) != 0 {
		this.Msg = NewPopulatedSmallMsg(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 4)
	}
	return this
}

func NewPopulatedSmallMsg(r randyMsg, easy bool) *SmallMsg {
	this := &SmallMsg{}
	if r.Intn(10) != 0 {
		v2 := string(randStringMsg(r))
		this.ScarBusStop = &v2
	}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.FlightParachute = make([]uint32, v3)
		for i := 0; i < v3; i++ {
			this.FlightParachute[i] = uint32(r.Uint32())
		}
	}
	if r.Intn(10) != 0 {
		v4 := string(randStringMsg(r))
		this.MapShark = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 19)
	}
	return this
}

func NewPopulatedPacked(r randyMsg, easy bool) *Packed {
	this := &Packed{}
	if r.Intn(10) != 0 {
		v5 := r.Intn(10)
		this.Ints = make([]int64, v5)
		for i := 0; i < v5; i++ {
			this.Ints[i] = int64(r.Int63())
			if r.Intn(2) == 0 {
				this.Ints[i] *= -1
			}
		}
	}
	if r.Intn(10) != 0 {
		v6 := r.Intn(10)
		this.Floats = make([]float64, v6)
		for i := 0; i < v6; i++ {
			this.Floats[i] = float64(r.Float64())
			if r.Intn(2) == 0 {
				this.Floats[i] *= -1
			}
		}
	}
	if r.Intn(10) != 0 {
		v7 := r.Intn(10)
		this.Uints = make([]uint32, v7)
		for i := 0; i < v7; i++ {
			this.Uints[i] = uint32(r.Uint32())
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 7)
	}
	return this
}

type randyMsg interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMsg(r randyMsg) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMsg(r randyMsg) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneMsg(r)
	}
	return string(tmps)
}
func randUnrecognizedMsg(r randyMsg, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMsg(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMsg(dAtA []byte, r randyMsg, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(v9))
	case 1:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMsg(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_msg_07dd33c081c61e70) }

var fileDescriptor_msg_07dd33c081c61e70 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8d, 0x3f, 0x4e, 0xc3, 0x30,
	0x1c, 0x85, 0x6b, 0xdc, 0x86, 0xf6, 0x57, 0x10, 0xc8, 0x8a, 0x90, 0x95, 0x21, 0xb2, 0x32, 0x20,
	0x2f, 0xa4, 0x12, 0x47, 0xc8, 0x10, 0x89, 0x21, 0x52, 0x95, 0x08, 0x76, 0x37, 0x0d, 0x4e, 0xd4,
	0x04, 0x47, 0xb6, 0x73, 0x2f, 0x8e, 0xc2, 0xc8, 0x5a, 0x7a, 0x01, 0x46, 0x46, 0x54, 0x87, 0x7f,
	0x62, 0xb2, 0xbf, 0xef, 0xf9, 0x3d, 0xc3, 0xa2, 0x33, 0x32, 0xee, 0xb5, 0xb2, 0x8a, 0x80, 0x3b,
	0x6c, 0x65, 0xac, 0x09, 0x6e, 0x64, 0x63, 0xeb, 0x61, 0x13, 0x97, 0xaa, 0x5b, 0x49, 0x25, 0xd5,
	0xca, 0x65, 0x9b, 0xe1, 0xd1, 0x91, 0x03, 0x77, 0x1b, 0xab, 0x51, 0x0a, 0x5e, 0xd2, 0xc8, 0xcc,
	0x48, 0xe2, 0xc3, 0x2c, 0x6d, 0xaa, 0x76, 0x4b, 0x11, 0x43, 0x1c, 0xe7, 0x23, 0x90, 0x6b, 0xc0,
	0x99, 0x91, 0x14, 0x33, 0xc4, 0x97, 0xb7, 0x7e, 0xfc, 0xfb, 0x51, 0x5c, 0x74, 0xa2, 0x6d, 0x33,
	0x23, 0xf3, 0xe3, 0x83, 0x48, 0xc3, 0xfc, 0x5b, 0x10, 0x06, 0xcb, 0xa2, 0x14, 0x3a, 0x19, 0x4c,
	0x61, 0x55, 0xef, 0xf6, 0x16, 0xf9, 0x5f, 0x45, 0x38, 0x5c, 0xa4, 0x6d, 0x23, 0x6b, 0xbb, 0x16,
	0x5a, 0x94, 0xf5, 0x60, 0x2b, 0x7a, 0xc6, 0x30, 0x3f, 0xcd, 0xff, 0x6b, 0x12, 0xc0, 0x3c, 0x13,
	0x7d, 0x51, 0x0b, 0xbd, 0xa3, 0xc4, 0x0d, 0xfd, 0x70, 0xf4, 0x00, 0xde, 0x5a, 0x94, 0xbb, 0x6a,
	0x4b, 0xae, 0x60, 0x7a, 0xf7, 0x64, 0x0d, 0x9d, 0x32, 0xcc, 0x71, 0x72, 0x72, 0x89, 0x72, 0xc7,
	0x24, 0x00, 0x2f, 0x6d, 0x95, 0xb0, 0x86, 0xce, 0x18, 0xe6, 0xc8, 0x25, 0x5f, 0x86, 0x50, 0x98,
	0xdd, 0x37, 0xc7, 0x92, 0xc7, 0x30, 0x3f, 0x77, 0xd1, 0x28, 0x12, 0xff, 0x7d, 0x1f, 0xa2, 0x8f,
	0x7d, 0x88, 0x9e, 0xdf, 0x42, 0xf4, 0x72, 0x08, 0x27, 0xaf, 0x87, 0x70, 0xf2, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x8b, 0xe2, 0xad, 0x70, 0x70, 0x01, 0x00, 0x00,
}
