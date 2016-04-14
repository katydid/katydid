//  Copyright 2016 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//Package proto encodes a parser.Interface into marshaled protocol buffer.
//This can be used for transcoding or dynamic marshaling.
//Dynamic marshaling is when the protocol buffer has not been compiled into the source and only the descriptor struct is available.
//
//TODO more tests
//
//GROUP not supported
package proto

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/parser"
	protoparser "github.com/katydid/katydid/parser/proto"
	"io"
	"math"
)

const maxVarintSize = 10

//Encoder encodes a parser.Interface into a byte slice containing marshaled protocol buffer.
type Encoder interface {
	Encode([]byte, parser.Interface) ([]byte, error)
}

//NewEncoder returns an Encoder that can marshal a parser.Interface into the specified protocol buffer message.
func NewEncoder(desc *descriptor.FileDescriptorSet, pkgName, msgName string) (Encoder, error) {
	descMap, err := protoparser.NewDescriptorMap(pkgName, msgName, desc)
	if err != nil {
		return nil, err
	}
	return &encoder{msg: newMsg(descMap, descMap.GetRoot())}, nil
}

type encoder struct {
	msg *msg
}

func (this *encoder) Encode(buf []byte, p parser.Interface) ([]byte, error) {
	offset := 0
	var err error
	buf, offset, err = this.msg.encode(buf, offset, p)
	if err != nil {
		return nil, err
	}
	return buf[:offset], nil
}

type msg struct {
	fieldEncoders map[string]fieldEncoder
}

type fieldEncoder interface {
	encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error)
}

func newMsg(desc protoparser.DescMap, protomsg *descriptor.DescriptorProto) *msg {
	msg := &msg{make(map[string]fieldEncoder)}
	for _, f := range protomsg.GetField() {
		fieldEnc := newFieldEncoder(desc, f)
		msg.fieldEncoders[f.GetName()] = fieldEnc
	}
	return msg
}

func newFieldEncoder(desc protoparser.DescMap, f *descriptor.FieldDescriptorProto) fieldEncoder {
	if f.IsRepeated() {
		return newRepeatedField(desc, f)
	}
	return newSingleField(desc, f)
}

func newSingleField(desc protoparser.DescMap, f *descriptor.FieldDescriptorProto) fieldEncoder {
	switch f.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return newDoubleField(f)
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return newFloatField(f)
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return newIntField(f)
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return newUintField(f)
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return newIntField(f)
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return newFixed64UintField(f)
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		return newFixed32UintField(f)
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return newBoolField(f)
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return newStringField(f)
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		return newMsgField(desc, f, desc.LookupMessage(f))
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return newBytesField(f)
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		return newUintField(f)
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		return newIntField(f)
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		return newFixed32IntField(f)
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return newFixed64IntField(f)
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		return newSint32Field(f)
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		return newSint64Field(f)
	}
	panic("unsupported field type " + f.GetType().String())
}

func (this *msg) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				return buf, offset, nil
			} else {
				return nil, 0, err
			}
		}
		name, err := p.String()
		if err != nil {
			return nil, 0, err
		}
		fieldEnc, ok := this.fieldEncoders[name]
		if !ok {
			continue //skip field
		}
		p.Down()
		buf, offset, err = fieldEnc.encode(buf, offset, p)
		if err != nil {
			return nil, 0, err
		}
		p.Up()
	}
}

type repeatedField struct {
	key      []byte
	fieldEnc fieldEncoder
}

func newRepeatedField(desc protoparser.DescMap, f *descriptor.FieldDescriptorProto) *repeatedField {
	return &repeatedField{
		key:      f.GetKey(),
		fieldEnc: newSingleField(desc, f),
	}
}

func (this *repeatedField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				return buf, offset, nil
			} else {
				return nil, 0, err
			}
		}
		_, err := p.Int()
		if err != nil {
			return nil, 0, err
		}
		p.Down()
		buf, offset, err = this.fieldEnc.encode(buf, offset, p)
		if err != nil {
			return nil, 0, err
		}
		p.Up()
	}
}

func newMsgField(desc protoparser.DescMap, field *descriptor.FieldDescriptorProto, msg *descriptor.DescriptorProto) *msgField {
	return &msgField{
		key: field.GetKey(),
		msg: newMsg(desc, msg),
	}
}

type msgField struct {
	key []byte
	msg *msg
}

func (this *msgField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	msginitoff := offset + 20
	msgbuf, msgfinaloff, err := this.msg.encode(buf, msginitoff, p)
	if err != nil {
		return nil, 0, err
	}
	msglen := msgfinaloff - msginitoff
	buf, offset = safeCopy(buf, offset, this.key)
	buf, offset = safeVarint(buf, offset, uint64(msglen))
	buf, offset = safeCopy(buf, offset, msgbuf[msginitoff:msgfinaloff])
	return buf, offset, nil
}

func checkSize(buf []byte, offset int, size int) []byte {
	if offset+size > len(buf) {
		newbuf := make([]byte, len(buf)*3+size+1)
		copy(newbuf, buf)
		buf = newbuf
	}
	return buf
}

func safeVarint(buf []byte, offset int, v uint64) ([]byte, int) {
	buf = checkSize(buf, offset, maxVarintSize)
	for v >= 1<<7 {
		buf[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	buf[offset] = uint8(v)
	return buf, offset + 1
}

func safeFixed64(buf []byte, offset int, x uint64) ([]byte, int) {
	buf = checkSize(buf, offset, 8)
	buf[offset] = uint8(x)
	buf[offset+1] = uint8(x >> 8)
	buf[offset+2] = uint8(x >> 16)
	buf[offset+3] = uint8(x >> 24)
	buf[offset+4] = uint8(x >> 32)
	buf[offset+5] = uint8(x >> 40)
	buf[offset+6] = uint8(x >> 48)
	buf[offset+7] = uint8(x >> 56)
	return buf, offset + 8
}

func safeFixed32(buf []byte, offset int, x uint32) ([]byte, int) {
	buf = checkSize(buf, offset, 4)
	buf[offset] = uint8(x)
	buf[offset+1] = uint8(x >> 8)
	buf[offset+2] = uint8(x >> 16)
	buf[offset+3] = uint8(x >> 24)
	return buf, offset + 4
}

func safeCopy(dst []byte, offset int, src []byte) ([]byte, int) {
	dst = checkSize(dst, offset, len(src))
	offset += copy(dst[offset:], src)
	return dst, offset
}

func newDoubleField(f *descriptor.FieldDescriptorProto) *doubleField {
	return &doubleField{f.GetKey()}
}

type doubleField struct {
	key []byte
}

func (this *doubleField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Double()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeFixed64(buf, offset, math.Float64bits(v))
	return buf, offset, nil
}

func newFloatField(f *descriptor.FieldDescriptorProto) *floatField {
	return &floatField{f.GetKey()}
}

type floatField struct {
	key []byte
}

func (this *floatField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Double()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeFixed32(buf, offset, math.Float32bits(float32(v)))
	return buf, offset, nil
}

func newIntField(f *descriptor.FieldDescriptorProto) *intField {
	return &intField{f.GetKey()}
}

type intField struct {
	key []byte
}

func (this *intField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Int()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeVarint(buf, offset, uint64(v))
	return buf, offset, nil
}

func newUintField(f *descriptor.FieldDescriptorProto) *uintField {
	return &uintField{f.GetKey()}
}

type uintField struct {
	key []byte
}

func (this *uintField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Uint()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeVarint(buf, offset, uint64(v))
	return buf, offset, nil
}

func newFixed64UintField(f *descriptor.FieldDescriptorProto) *fixed64UintField {
	return &fixed64UintField{f.GetKey()}
}

type fixed64UintField struct {
	key []byte
}

func (this *fixed64UintField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Uint()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeFixed64(buf, offset, uint64(v))
	return buf, offset, nil
}

func newFixed32UintField(f *descriptor.FieldDescriptorProto) *fixed32UintField {
	return &fixed32UintField{f.GetKey()}
}

type fixed32UintField struct {
	key []byte
}

func (this *fixed32UintField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Uint()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeFixed32(buf, offset, uint32(v))
	return buf, offset, nil
}

func newBoolField(f *descriptor.FieldDescriptorProto) *boolField {
	return &boolField{f.GetKey()}
}

type boolField struct {
	key []byte
}

func (this *boolField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	b, err := p.Bool()
	if err != nil {
		return nil, 0, err
	}
	v := uint64(0)
	if b {
		v = 1
	}
	buf, offset = safeVarint(buf, offset, v)
	return buf, offset, nil
}

func newStringField(f *descriptor.FieldDescriptorProto) *stringField {
	return &stringField{f.GetKey()}
}

type stringField struct {
	key []byte
}

func (this *stringField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.String()
	if err != nil {
		return nil, 0, err
	}
	vb := []byte(v)
	buf, offset = safeVarint(buf, offset, uint64(len(vb)))
	buf, offset = safeCopy(buf, offset, vb)
	return buf, offset, nil
}

func newBytesField(f *descriptor.FieldDescriptorProto) *bytesField {
	return &bytesField{f.GetKey()}
}

type bytesField struct {
	key []byte
}

func (this *bytesField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Bytes()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeVarint(buf, offset, uint64(len(v)))
	buf, offset = safeCopy(buf, offset, v)
	return buf, offset, nil
}

func newFixed64IntField(f *descriptor.FieldDescriptorProto) *fixed64IntField {
	return &fixed64IntField{f.GetKey()}
}

type fixed64IntField struct {
	key []byte
}

func (this *fixed64IntField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Int()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeFixed64(buf, offset, uint64(v))
	return buf, offset, nil
}

func newFixed32IntField(f *descriptor.FieldDescriptorProto) *fixed32IntField {
	return &fixed32IntField{f.GetKey()}
}

type fixed32IntField struct {
	key []byte
}

func (this *fixed32IntField) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Int()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeFixed32(buf, offset, uint32(v))
	return buf, offset, nil
}

func newSint32Field(f *descriptor.FieldDescriptorProto) *sint32Field {
	return &sint32Field{f.GetKey()}
}

type sint32Field struct {
	key []byte
}

func (this *sint32Field) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Int()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeVarint(buf, offset, uint64((uint32(v)<<1)^uint32((int32(v)>>31))))
	return buf, offset, nil
}

func newSint64Field(f *descriptor.FieldDescriptorProto) *sint64Field {
	return &sint64Field{f.GetKey()}
}

type sint64Field struct {
	key []byte
}

func (this *sint64Field) encode(buf []byte, offset int, p parser.Interface) ([]byte, int, error) {
	if err := p.Next(); err != nil {
		return buf, offset, nil
	}
	if !p.IsLeaf() {
		return buf, offset, nil
	}
	buf, offset = safeCopy(buf, offset, this.key)
	v, err := p.Int()
	if err != nil {
		return nil, 0, err
	}
	buf, offset = safeVarint(buf, offset, uint64((uint64(v)<<1)^uint64((v>>63))))
	return buf, offset, nil
}
