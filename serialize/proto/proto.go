//  Copyright 2013 Walter Schulze
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

package proto

import (
	"encoding/binary"
	"fmt"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/serialize"
	"io"
	"reflect"
	"unsafe"
)

type errVarint struct {
	buf []byte
}

func newErrVarint(buf []byte) error {
	buffer := make([]byte, 10)
	nn := copy(buffer, buf)
	return &errVarint{buffer[:nn]}
}

func (this *errVarint) Error() string {
	return fmt.Sprintf("error decoding varint from %#v", this.buf)
}

var errBufferOverlow = fmt.Errorf("buffer overflow")

var errUnknownWireType = fmt.Errorf("unknown wire type")

type errUnknown struct {
	msg string
}

func (this *errUnknown) Error() string {
	return "Could not find " + this.msg
}

func uvarint(buf []byte) (uint64, int, error) {
	var uv uint64
	var shift uint
	for i, b := range buf {
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return 0, 0, newErrVarint(buf)
			}
			return uv | uint64(b)<<shift, i + 1, nil
		}
		uv |= (uint64(b) & 0x7F) << shift
		shift += 7
	}
	return 0, 0, newErrVarint(buf)
}

func length(wireType int, buf []byte) (prefix int, l int, err error) {
	switch wireType {
	case 0:
		_, n2, err := uvarint(buf)
		return 0, n2, err
	case 1:
		return 0, 8, nil
	case 2:
		l, n2, err := uvarint(buf)
		return n2, int(l), err
	case 5:
		return 0, 4, nil
	}
	return 0, 0, errUnknownWireType
}

type descMap struct {
	desc       *descriptor.FileDescriptorSet
	root       *descriptor.DescriptorProto
	fieldToMsg map[*descriptor.FieldDescriptorProto]*descriptor.DescriptorProto
	msgToField map[*descriptor.DescriptorProto]map[uint64]*descriptor.FieldDescriptorProto
}

func newDescriptorMap(pkgName, msgName string, desc *descriptor.FileDescriptorSet) (*descMap, error) {
	root := desc.GetMessage(pkgName, msgName)
	d := &descMap{
		desc:       desc,
		root:       root,
		fieldToMsg: make(map[*descriptor.FieldDescriptorProto]*descriptor.DescriptorProto),
		msgToField: make(map[*descriptor.DescriptorProto]map[uint64]*descriptor.FieldDescriptorProto),
	}
	err := d.visit(pkgName, root)
	return d, err
}

func (this *descMap) visit(pkgName string, msg *descriptor.DescriptorProto) error {
	if _, ok := this.msgToField[msg]; ok {
		return nil
	}
	for i, f := range msg.GetField() {
		if _, ok := this.msgToField[msg]; !ok {
			this.msgToField[msg] = make(map[uint64]*descriptor.FieldDescriptorProto)
		}
		this.msgToField[msg][f.GetKeyUint64()] = msg.Field[i]
		if f.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			newPkgName, newMsgName := this.desc.FindMessage(pkgName, msg.GetName(), f.GetName())
			if len(newMsgName) == 0 {
				return &errUnknown{pkgName + "." + msg.GetName() + "." + f.GetName()}
			}
			newMsg := this.desc.GetMessage(newPkgName, newMsgName)
			if newMsg == nil {
				return &errUnknown{newPkgName + "." + newMsgName}
			}
			this.fieldToMsg[msg.Field[i]] = newMsg
			if _, ok := this.msgToField[newMsg]; ok {
				continue
			}
			if err := this.visit(newPkgName, newMsg); err != nil {
				return err
			}
		}
	}
	for i := range msg.GetNestedType() {
		this.visit(pkgName, msg.GetNestedType()[i])
	}
	return nil
}

func (this *descMap) getRoot() *descriptor.DescriptorProto {
	return this.root
}

func (this *descMap) lookupMessage(field *descriptor.FieldDescriptorProto) *descriptor.DescriptorProto {
	return this.fieldToMsg[field]
}

func (this *descMap) lookupField(msg *descriptor.DescriptorProto, key uint64) (*descriptor.FieldDescriptorProto, bool) {
	fields, ok := this.msgToField[msg]
	if !ok {
		return nil, false
	}
	f, ok := fields[key]
	return f, ok
}

type protoParser struct {
	descMap *descMap
	state
	stack []state
}

func (this *protoParser) Copy() serialize.Parser {
	s := &protoParser{
		descMap: this.descMap,
		state:   this.state,
		stack:   make([]state, len(this.stack)),
	}
	for i := range this.stack {
		s.stack[i] = this.stack[i]
	}
	return s
}

type state struct {
	buf           []byte
	parent        *descriptor.DescriptorProto
	offset        int
	length        int
	field         *descriptor.FieldDescriptorProto
	isLeaf        bool
	hadLeaf       bool
	isRepeated    bool
	inRepeated    bool
	indexRepeated int
}

type BytesParser interface {
	serialize.Parser
	Init(buf []byte) error
}

func NewProtoParser(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet) BytesParser {
	descMap, err := newDescriptorMap(srcPackage, srcMessage, desc)
	if err != nil {
		panic(err)
	}
	return &protoParser{
		descMap: descMap,
		state: state{
			parent: descMap.getRoot(),
		},
		stack: make([]state, 0, 10),
	}
}

func (s *protoParser) Init(buf []byte) error {
	s.buf = buf
	s.offset = 0
	s.length = 0
	s.field = nil
	s.stack = s.stack[:0]
	return nil
}

func (s *protoParser) Next() error {
	if s.isLeaf {
		if s.hadLeaf {
			return io.EOF
		}
		s.hadLeaf = true
		return nil
	}
	// This field is repeated, but Next is called,
	//   => all elements in the repeated field should be skipped
	//   and the Next field should be retrieved.
	if s.isRepeated && !s.inRepeated {
		if err := s.skipRepeated(); err != nil {
			return err
		}
		s.isRepeated = false
		s.length = 0
	}

	s.offset += s.length
	if s.offset >= len(s.buf) {
		if s.offset == len(s.buf) {
			s.length = 0
			return io.EOF
		}
		return io.ErrShortBuffer
	}

	if s.inRepeated {
		v, n, err := uvarint(s.buf[s.offset:])
		if err != nil {
			return err
		}
		field, ok := s.descMap.lookupField(s.parent, v)
		if !ok {
			s.length = 0
			s.isRepeated = false
			return io.EOF
		}
		if field.GetNumber() != s.field.GetNumber() {
			s.length = 0
			s.isRepeated = false
			return io.EOF
		}
		s.offset += n
		wireType := int(v & 0x7)
		n, l, err := length(wireType, s.buf[s.offset:])
		if err != nil {
			return err
		}
		s.offset += n
		s.length = l
		s.indexRepeated++
		return nil
	}

	v, n, err := uvarint(s.buf[s.offset:])
	if err != nil {
		return err
	}
	field, ok := s.descMap.lookupField(s.parent, v)
	if !ok || !field.IsRepeated() {
		s.offset += n
		wireType := int(v & 0x7)
		n, l, err := length(wireType, s.buf[s.offset:])
		if err != nil {
			return err
		}
		s.offset += n
		s.length = l
	}
	if ok {
		if field.IsRepeated() {
			s.isRepeated = true
			s.length = 0
		}
		s.field = field
		return nil
	}
	return s.Next()
}

func (s *protoParser) skipRepeated() error {
	s.Down()
	err := s.Next()
	for err == nil {
		err = s.Next()
	}
	if err != io.EOF {
		return err
	}
	s.Up()
	return nil
}

func (s *protoParser) IsLeaf() bool {
	return s.isLeaf
}

func (s *protoParser) Value() []byte {
	return s.buf[s.offset : s.offset+s.length]
}

func (s *protoParser) Double() (float64, error) {
	if !s.isLeaf {
		return 0, serialize.ErrNotDouble
	}
	if s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_DOUBLE &&
		s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_FLOAT {
		return 0, serialize.ErrNotDouble
	}
	buf := s.Value()
	if len(buf) == 8 {
		return *(*float64)(unsafe.Pointer(&buf[0])), nil
	}
	if len(buf) == 4 {
		return float64(*(*float32)(unsafe.Pointer(&buf[0]))), nil
	}
	return 0, fmt.Errorf("Double: wrong size buffer %d should be 4 or 8", len(buf))
}

func (s *protoParser) Int() (int64, error) {
	if !s.isLeaf {
		if s.inRepeated {
			return int64(s.indexRepeated - 1), nil
		}
		return 0, serialize.ErrNotInt
	}
	typ := s.field.GetType()
	switch typ {
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return s.decodeInt64()
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return s.decodeSfixed64()
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		return s.decodeSint64()
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		i, err := s.decodeInt32()
		return int64(i), err
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		i, err := s.decodeSfixed32()
		return int64(i), err
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		i, err := s.decodeSint32()
		return int64(i), err
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		i, err := s.decodeInt32()
		return int64(i), err
	}
	return 0, serialize.ErrNotInt
}

func (s *protoParser) Uint() (uint64, error) {
	if !s.isLeaf {
		return 0, serialize.ErrNotUint
	}
	typ := s.field.GetType()
	switch typ {
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return s.decodeUint64()
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return s.decodeFixed64()
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		u, err := s.decodeUint32()
		return uint64(u), err
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		u, err := s.decodeFixed32()
		return uint64(u), err
	}
	return 0, serialize.ErrNotUint
}

func (s *protoParser) Bool() (bool, error) {
	if !s.isLeaf {
		return false, serialize.ErrNotBool
	}
	if s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_BOOL {
		return false, serialize.ErrNotBool
	}
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return false, fmt.Errorf("decodeVarint n = %d", n)
	}
	return v != 0, nil
}

func (s *protoParser) String() (string, error) {
	if !s.isLeaf {
		return s.field.GetName(), nil
	}
	if s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_STRING {
		return "", serialize.ErrNotString
	}
	buf := s.Value()
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	strHeader := reflect.StringHeader{Data: header.Data, Len: header.Len}
	return *(*string)(unsafe.Pointer(&strHeader)), nil
}

func (s *protoParser) Bytes() ([]byte, error) {
	if !s.isLeaf {
		return nil, serialize.ErrNotBytes
	}
	if s.field.GetType() != descriptor.FieldDescriptorProto_TYPE_BYTES {
		return nil, serialize.ErrNotBytes
	}
	return s.Value(), nil
}

func (s *protoParser) decodeInt64() (int64, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int64(v), nil
}

func (s *protoParser) decodeUint64() (uint64, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return v, nil
}

func (s *protoParser) decodeInt32() (int32, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int32(v), nil
}

func (s *protoParser) decodeFixed64() (uint64, error) {
	buf := s.Value()
	if len(buf) < 8 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*uint64)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoParser) decodeFixed32() (uint32, error) {
	buf := s.Value()
	if len(buf) < 4 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*uint32)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoParser) decodeUint32() (uint32, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return uint32(v), nil
}

func (s *protoParser) decodeSfixed32() (int32, error) {
	buf := s.Value()
	if len(buf) < 4 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*int32)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoParser) decodeSfixed64() (int64, error) {
	buf := s.Value()
	if len(buf) < 8 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*int64)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoParser) decodeSint32() (int32, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31)), nil
}

func (s *protoParser) decodeSint64() (int64, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int64((v >> 1) ^ uint64((int64(v&1)<<63)>>63)), nil
}

func (s *protoParser) Up() {
	top := len(s.stack) - 1
	offset := s.offset
	length := s.length
	inRepeated := s.inRepeated
	s.state = s.stack[top]
	s.stack = s.stack[:top]
	if inRepeated {
		s.offset = offset + length
	}
}

func (s *protoParser) Down() {
	if s.isRepeated && !s.inRepeated {
		s.stack = append(s.stack, s.state)
		s.inRepeated = true
		s.indexRepeated = 0
		s.length = 0
	} else if s.field.IsMessage() {
		s.stack = append(s.stack, s.state)
		s.buf = s.buf[s.offset : s.offset+s.length]
		s.parent = s.descMap.lookupMessage(s.field)
		s.offset = 0
		s.length = 0
		s.field = nil
		s.inRepeated = false
		s.isRepeated = false
	} else {
		s.stack = append(s.stack, s.state)
		s.isLeaf = true
		s.inRepeated = false
		s.isRepeated = false
	}
}
