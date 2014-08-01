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

package funcs

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

type decDouble struct {
	D Bytes
}

func (this *decDouble) Eval() float64 {
	buf := this.D.Eval()
	if len(buf) < 8 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*float64)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decDouble", new(decDouble))
}

type decFloat struct {
	D Bytes
}

func (this *decFloat) Eval() float32 {
	buf := this.D.Eval()
	if len(buf) < 4 {
		panic(fmt.Sprintf("decodeFloat: buffer too short"))
	}
	return *(*float32)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decFloat", new(decFloat))
}

type decInt64 struct {
	D Bytes
}

func (this *decInt64) Eval() int64 {
	buf := this.D.Eval()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int64(v)
}

func init() {
	Register("decInt64", new(decInt64))
}

type decUint64 struct {
	D Bytes
}

func (this *decUint64) Eval() uint64 {
	buf := this.D.Eval()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return v
}

func init() {
	Register("decUint64", new(decUint64))
}

type decInt32 struct {
	D Bytes
}

func (this *decInt32) Eval() int32 {
	buf := this.D.Eval()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int32(v)
}

func init() {
	Register("decInt32", new(decInt32))
}

type decFixed64 struct {
	D Bytes
}

func (this *decFixed64) Eval() uint64 {
	buf := this.D.Eval()
	if len(buf) < 8 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*uint64)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decFixed64", new(decFixed64))
}

type decFixed32 struct {
	D Bytes
}

func (this *decFixed32) Eval() uint32 {
	buf := this.D.Eval()
	if len(buf) < 4 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*uint32)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decFixed32", new(decFixed32))
}

type decBool struct {
	D Bytes
}

func (this *decBool) Eval() bool {
	buf := this.D.Eval()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return v != 0
}

func init() {
	Register("decBool", new(decBool))
}

type decString struct {
	D Bytes
}

func (this *decString) Eval() string {
	buf := this.D.Eval()
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	strHeader := reflect.StringHeader{Data: header.Data, Len: header.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}

func init() {
	Register("decString", new(decString))
}

type decBytes struct {
	D Bytes
}

func (this *decBytes) Eval() []byte {
	buf := this.D.Eval()
	return buf
}

func init() {
	Register("decBytes", new(decBytes))
}

type decUint32 struct {
	D Bytes
}

func (this *decUint32) Eval() uint32 {
	buf := this.D.Eval()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return uint32(v)
}

func init() {
	Register("decUint32", new(decUint32))
}

type decEnum struct {
	D Bytes
}

func (this *decEnum) Eval() int32 {
	buf := this.D.Eval()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int32(v)
}

func init() {
	Register("decEnum", new(decEnum))
}

type decSFixed32 struct {
	D Bytes
}

func (this *decSFixed32) Eval() int32 {
	buf := this.D.Eval()
	if len(buf) < 4 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*int32)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decSFixed32", new(decSFixed32))
}

type decSFixed64 struct {
	D Bytes
}

func (this *decSFixed64) Eval() int64 {
	buf := this.D.Eval()
	if len(buf) < 8 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*int64)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decSFixed64", new(decSFixed64))
}

type decSInt32 struct {
	D Bytes
}

func (this *decSInt32) Eval() int32 {
	buf := this.D.Eval()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
}

func init() {
	Register("decSInt32", new(decSInt32))
}

type decSInt64 struct {
	D Bytes
}

func (this *decSInt64) Eval() int64 {
	buf := this.D.Eval()
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int64((v >> 1) ^ uint64((int64(v&1)<<63)>>63))
}

func init() {
	Register("decSInt64", new(decSInt64))
}
