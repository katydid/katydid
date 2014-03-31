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

type constFloat64 struct {
	v float64
}

func NewFloat64(v float64) Float64 {
	return &constFloat64{v}
}

func (this *constFloat64) Eval(buf []byte) float64 {
	return this.v
}

type constFloat32 struct {
	v float32
}

func NewFloat32(v float32) Float32 {
	return &constFloat32{v}
}

func (this *constFloat32) Eval(buf []byte) float32 {
	return this.v
}

type constInt64 struct {
	v int64
}

func NewInt64(v int64) Int64 {
	return &constInt64{v}
}

func (this *constInt64) Eval(buf []byte) int64 {
	return this.v
}

type constUint64 struct {
	v uint64
}

func NewUint64(v uint64) Uint64 {
	return &constUint64{v}
}

func (this *constUint64) Eval(buf []byte) uint64 {
	return this.v
}

type constInt32 struct {
	v int32
}

func NewInt32(v int32) Int32 {
	return &constInt32{v}
}

func (this *constInt32) Eval(buf []byte) int32 {
	return this.v
}

type constUint32 struct {
	v uint32
}

func NewUint32(v uint32) Uint32 {
	return &constUint32{v}
}

func (this *constUint32) Eval(buf []byte) uint32 {
	return this.v
}

type constBool struct {
	v bool
}

func NewBool(v bool) Bool {
	return &constBool{v}
}

func (this *constBool) Eval(buf []byte) bool {
	return this.v
}

type constString struct {
	v string
}

func NewString(v string) String {
	return &constString{v}
}

func (this *constString) Eval(buf []byte) string {
	return this.v
}

type constBytes struct {
	v []byte
}

func NewBytes(v []byte) Bytes {
	return &constBytes{v}
}

func (this *constBytes) Eval(buf []byte) []byte {
	return this.v
}
