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
	"bytes"
	"code.google.com/p/gogoprotobuf/proto"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func getField(m *DecodeFields) []byte {
	buf, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	_, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return buf[n:]
}

var (
	seed = time.Now().UnixNano()
	r    = rand.New(rand.NewSource(seed))
)

func TestDouble(t *testing.T) {
	v1 := r.Float64()
	m := &DecodeFields{
		DoubleField: proto.Float64(v1),
	}
	v2 := (&(decDouble{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestFloat(t *testing.T) {
	v1 := r.Float32()
	m := &DecodeFields{
		FloatField: proto.Float32(v1),
	}
	v2 := (&(decFloat{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestInt64(t *testing.T) {
	v1 := r.Int63()
	m := &DecodeFields{
		Int64Field: proto.Int64(v1),
	}
	v2 := (&(decInt64{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestUint64(t *testing.T) {
	v1 := uint64(r.Uint32())
	m := &DecodeFields{
		Uint64Field: proto.Uint64(v1),
	}
	v2 := (&(decUint64{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestInt32(t *testing.T) {
	v1 := r.Int31()
	m := &DecodeFields{
		Int32Field: proto.Int32(v1),
	}
	v2 := (&(decInt32{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestFixed64(t *testing.T) {
	v1 := uint64(r.Uint32())
	m := &DecodeFields{
		Fixed64Field: proto.Uint64(v1),
	}
	v2 := (&(decFixed64{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestFixed32(t *testing.T) {
	v1 := r.Uint32()
	m := &DecodeFields{
		Fixed32Field: proto.Uint32(v1),
	}
	v2 := (&(decFixed32{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestBool(t *testing.T) {
	v1 := (r.Int31n(2) == 0)
	m := &DecodeFields{
		BoolField: proto.Bool(v1),
	}
	v2 := (&(decBool{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestString(t *testing.T) {
	v1 := strconv.Itoa(r.Int())
	m := &DecodeFields{
		StringField: proto.String(v1),
	}
	v2 := (&(decString{})).Eval(getField(m))
	if bytes.Equal([]byte(v1), []byte(v2)) {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestBytes(t *testing.T) {
	v1 := []byte(strconv.Itoa(r.Int()))
	m := &DecodeFields{
		BytesField: v1,
	}
	v2 := (&(decBytes{})).Eval(getField(m))
	if bytes.Equal(v1, v2) {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestUint32(t *testing.T) {
	v1 := r.Uint32()
	m := &DecodeFields{
		Uint32Field: proto.Uint32(v1),
	}
	v2 := (&(decUint32{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestEnum(t *testing.T) {
	v1 := r.Int31n(2) + 1
	m := &DecodeFields{
		EnumField: DecodeEnum(v1).Enum(),
	}
	v2 := (&(decEnum{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestSFixed32(t *testing.T) {
	v1 := r.Int31()
	m := &DecodeFields{
		SFixed32Field: proto.Int32(v1),
	}
	v2 := (&(decSFixed32{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestSFixed64(t *testing.T) {
	v1 := r.Int63()
	m := &DecodeFields{
		SFixed64Field: proto.Int64(v1),
	}
	v2 := (&(decSFixed64{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestSInt32(t *testing.T) {
	v1 := r.Int31()
	m := &DecodeFields{
		SInt32Field: proto.Int32(v1),
	}
	v2 := (&(decSInt32{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}

func TestSInt64(t *testing.T) {
	v1 := r.Int63()
	m := &DecodeFields{
		SInt64Field: proto.Int64(v1),
	}
	v2 := (&(decSInt64{})).Eval(getField(m))
	if v1 != v2 {
		t.Fatalf("expected %v got %v", v1, v2)
	}
}
