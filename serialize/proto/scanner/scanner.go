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

package scanner

import (
	"encoding/binary"
	"fmt"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/serialize"
	"github.com/katydid/katydid/serialize/proto/tokens"
	"io"
	//"log"
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

type ProtoTokens interface {
	LookupKey(src int, key uint64) (int, string, bool)
	IsLeaf(int) bool
	LookupType(src int) descriptor.FieldDescriptorProto_Type
	LookupName(src int) string
}

type protoScanner struct {
	tokens ProtoTokens
	state
	stack []state
}

func (this *protoScanner) Copy() serialize.Scanner {
	s := &protoScanner{
		tokens: this.tokens,
		state:  this.state,
		stack:  make([]state, len(this.stack)),
	}
	for i := range this.stack {
		s.stack[i] = this.stack[i]
	}
	return s
}

type state struct {
	buf         []byte
	parentToken int
	offset      int
	length      int
	tokenId     int
	name        string
}

type BytesScanner interface {
	serialize.Scanner
	Init(buf []byte) error
	Value() []byte
}

func NewProtoScanner(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet) BytesScanner {
	toks, err := tokens.New(srcPackage, srcMessage, desc)
	if err != nil {
		panic(err)
	}
	rootToken, err := toks.GetTokenId(srcPackage + "." + srcMessage)
	if err != nil {
		panic(err)
	}
	return &protoScanner{
		tokens: toks,
		state: state{
			parentToken: rootToken,
		},
		stack: make([]state, 0, 10),
	}
}

func (s *protoScanner) Init(buf []byte) error {
	s.buf = buf
	s.offset = 0
	s.length = 0
	s.tokenId = 0
	s.stack = s.stack[:0]
	return nil
}

func (s *protoScanner) Next() error {
	//log.Printf("Next %d/%d", s.offset, len(s.buf))
	s.offset += s.length
	if s.offset >= len(s.buf) {
		if s.offset == len(s.buf) {
			return io.EOF
		}
		return io.ErrShortBuffer
	}
	v, n, err := uvarint(s.buf[s.offset:])
	if err != nil {
		return err
	}
	s.offset += n
	wireType := int(v & 0x7)
	n, l, err := length(wireType, s.buf[s.offset:])
	if err != nil {
		return err
	}
	s.offset += n
	s.length = l
	tokenId, name, ok := s.tokens.LookupKey(s.parentToken, v)
	if ok {
		//log.Printf("Return %d", s.offset)
		s.tokenId = tokenId
		s.name = name
		return nil
	}
	return s.Next()
}

func (s *protoScanner) Id() int {
	return s.tokenId
}

func (s *protoScanner) IsLeaf() bool {
	return s.tokens.IsLeaf(s.tokenId)
}

func (s *protoScanner) Value() []byte {
	return s.buf[s.offset : s.offset+s.length]
}

func (s *protoScanner) Name() string {
	return s.name
}

func (s *protoScanner) Float64() (float64, error) {
	buf := s.Value()
	if len(buf) == 8 {
		return *(*float64)(unsafe.Pointer(&buf[0])), nil
	}
	if len(buf) == 4 {
		return float64(*(*float32)(unsafe.Pointer(&buf[0]))), nil
	}
	return 0, fmt.Errorf("Double: wrong size buffer %d should be 4 or 8", len(buf))
}

func (s *protoScanner) Int64() (int64, error) {
	typ := s.tokens.LookupType(s.tokenId)
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
	return 0, serialize.ErrNotInt64
}

func (s *protoScanner) Uint64() (uint64, error) {
	typ := s.tokens.LookupType(s.tokenId)
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
	return 0, serialize.ErrNotUint64
}

func (s *protoScanner) Bool() (bool, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return false, fmt.Errorf("decodeVarint n = %d", n)
	}
	return v != 0, nil
}

func (s *protoScanner) String() (string, error) {
	buf := s.Value()
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	strHeader := reflect.StringHeader{Data: header.Data, Len: header.Len}
	return *(*string)(unsafe.Pointer(&strHeader)), nil
}

func (s *protoScanner) Bytes() ([]byte, error) {
	return s.Value(), nil
}

func (s *protoScanner) decodeInt64() (int64, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int64(v), nil
}

func (s *protoScanner) decodeUint64() (uint64, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return v, nil
}

func (s *protoScanner) decodeInt32() (int32, error) {
	v, n := binary.Uvarint(s.Value())
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int32(v), nil
}

func (s *protoScanner) decodeFixed64() (uint64, error) {
	buf := s.Value()
	if len(buf) < 8 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*uint64)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoScanner) decodeFixed32() (uint32, error) {
	buf := s.Value()
	if len(buf) < 4 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*uint32)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoScanner) decodeUint32() (uint32, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return uint32(v), nil
}

func (s *protoScanner) decodeSfixed32() (int32, error) {
	buf := s.Value()
	if len(buf) < 4 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*int32)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoScanner) decodeSfixed64() (int64, error) {
	buf := s.Value()
	if len(buf) < 8 {
		return 0, fmt.Errorf("decodeDouble: buffer too short")
	}
	return *(*int64)(unsafe.Pointer(&buf[0])), nil
}

func (s *protoScanner) decodeSint32() (int32, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31)), nil
}

func (s *protoScanner) decodeSint64() (int64, error) {
	buf := s.Value()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		return 0, fmt.Errorf("decodeVarint n = %d", n)
	}
	return int64((v >> 1) ^ uint64((int64(v&1)<<63)>>63)), nil
}

func (s *protoScanner) Up() {
	top := len(s.stack) - 1
	s.state = s.stack[top]
	s.stack = s.stack[:top]
}

func (s *protoScanner) Down() {
	s.stack = append(s.stack, s.state)
	s.buf = s.buf[s.offset : s.offset+s.length]
	s.parentToken = s.tokenId
	s.offset = 0
	s.length = 0
	s.tokenId = 0
}
