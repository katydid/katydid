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

type decDouble struct{}

func (this *decDouble) Eval(buf []byte) float64 {
	if len(buf) < 8 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*float64)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decDouble", new(decDouble))
}

type decFloat struct{}

func (this *decFloat) Eval(buf []byte) float32 {
	if len(buf) < 4 {
		panic(fmt.Sprintf("decodeFloat: buffer too short"))
	}
	return *(*float32)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decFloat", new(decFloat))
}

type decInt64 struct{}

func (this *decInt64) Eval(buf []byte) int64 {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int64(v)
}

func init() {
	Register("decInt64", new(decInt64))
}

type decUint64 struct{}

func (this *decUint64) Eval(buf []byte) uint64 {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return v
}

func init() {
	Register("decUint64", new(decUint64))
}

type decInt32 struct{}

func (this *decInt32) Eval(buf []byte) int32 {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int32(v)
}

func init() {
	Register("decInt32", new(decInt32))
}

type decFixed64 struct{}

func (this *decFixed64) Eval(buf []byte) uint64 {
	if len(buf) < 8 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*uint64)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decFixed64", new(decFixed64))
}

type decFixed32 struct{}

func (this *decFixed32) Eval(buf []byte) uint32 {
	if len(buf) < 4 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*uint32)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decFixed32", new(decFixed32))
}

type decBool struct{}

func (this *decBool) Eval(buf []byte) bool {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return v != 0
}

func init() {
	Register("decBool", new(decBool))
}

type decString struct{}

func (this *decString) Eval(buf []byte) string {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	strHeader := reflect.StringHeader{header.Data, header.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}

func init() {
	Register("decString", new(decString))
}

type decBytes struct{}

func (this *decBytes) Eval(buf []byte) []byte {
	return buf
}

func init() {
	Register("decBytes", new(decBytes))
}

type decUint32 struct{}

func (this *decUint32) Eval(buf []byte) uint32 {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return uint32(v)
}

func init() {
	Register("decUint32", new(decUint32))
}

type decEnum struct{}

func (this *decEnum) Eval(buf []byte) int32 {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int32(v)
}

func init() {
	Register("decEnum", new(decEnum))
}

type decSFixed32 struct{}

func (this *decSFixed32) Eval(buf []byte) int32 {
	if len(buf) < 4 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*int32)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decSFixed32", new(decSFixed32))
}

type decSFixed64 struct{}

func (this *decSFixed64) Eval(buf []byte) int64 {
	if len(buf) < 8 {
		panic(fmt.Sprintf("decodeDouble: buffer too short"))
	}
	return *(*int64)(unsafe.Pointer(&buf[0]))
}

func init() {
	Register("decSFixed64", new(decSFixed64))
}

type decSInt32 struct{}

func (this *decSInt32) Eval(buf []byte) int32 {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
}

func init() {
	Register("decSInt32", new(decSInt32))
}

type decSInt64 struct{}

func (this *decSInt64) Eval(buf []byte) int64 {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int64((v >> 1) ^ uint64((int64(v&1)<<63)>>63))
}

func init() {
	Register("decSInt64", new(decSInt64))
}
