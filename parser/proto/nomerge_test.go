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

package proto_test

import (
	"encoding/binary"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/parser/debug"
	katydidproto "github.com/katydid/katydid/parser/proto"
	"github.com/katydid/katydid/parser/proto/prototests"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func noMerge(data []byte, desc *descriptor.FileDescriptorSet, pkgName, msgName string) error {
	parser := katydidproto.NewProtoNumParser(pkgName, msgName, desc)
	if err := parser.Init(data); err != nil {
		panic(err)
	}
	return katydidproto.NoLatentAppendingOrMerging(parser)
}

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func TestNoMergeNoMerge(t *testing.T) {
	m := debug.Input
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	err = noMerge(data, m.Description(), "debug", "Debug")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoMergeMerge(t *testing.T) {
	m := debug.Input
	m.G = proto.Float64(1.1)
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	key := byte(uint32(7)<<3 | uint32(1))
	data = append(data, key, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	err = noMerge(data, m.Description(), "debug", "Debug")
	if err == nil || !strings.Contains(err.Error(), "G requires merging") {
		t.Fatalf("G should require merging")
	}
}

func TestNoMergeLatent(t *testing.T) {
	m := debug.Input
	m.F = []uint32{5, 6, 7}
	m.G = proto.Float64(1.1)
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	key := byte(uint32(6)<<3 | uint32(5))
	data = append(data, key, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	err = noMerge(data, m.Description(), "debug", "Debug")
	if err == nil || !strings.Contains(err.Error(), "F") {
		t.Fatalf("F should have latent appending")
	}
}

func TestNoMergeNestedNoMerge(t *testing.T) {
	bigm := prototests.NewPopulatedBigMsg(r, true)
	data, err := proto.Marshal(bigm)
	if err != nil {
		panic(err)
	}
	err = noMerge(data, bigm.Description(), "prototests", "BigMsg")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoMergeMessageMerge(t *testing.T) {
	bigm := prototests.NewPopulatedBigMsg(r, true)
	bigm.Msg = prototests.NewPopulatedSmallMsg(r, true)
	data, err := proto.Marshal(bigm)
	if err != nil {
		panic(err)
	}
	key := byte(uint32(3)<<3 | uint32(2))
	fieldkey := byte(uint32(12)<<3 | uint32(5))
	data = append(data, key, 5, fieldkey, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	err = noMerge(data, bigm.Description(), "prototests", "BigMsg")
	if err == nil || !strings.Contains(err.Error(), "requires merging") {
		t.Fatal(err)
	}
}

func TestNoMergeNestedMerge(t *testing.T) {
	bigm := prototests.NewPopulatedBigMsg(r, true)
	m := prototests.NewPopulatedSmallMsg(r, true)
	if m.FlightParachute == nil {
		m.FlightParachute = []uint32{1}
	}
	m.MapShark = proto.String("a")
	key := byte(uint32(12)<<3 | uint32(5))
	m.XXX_unrecognized = append(m.XXX_unrecognized, key, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	bigm.Msg = m
	data, err := proto.Marshal(bigm)
	if err != nil {
		panic(err)
	}
	err = noMerge(data, bigm.Description(), "prototests", "BigMsg")
	if err == nil || !strings.Contains(err.Error(), "FlightParachute requires merging") {
		t.Fatalf("FlightParachute should require merging %#v", bigm)
	}
}

func TestNoMergeExtensionNoMerge(t *testing.T) {
	bigm := prototests.AContainer
	data, err := proto.Marshal(bigm)
	if err != nil {
		panic(err)
	}
	err = noMerge(data, bigm.Description(), "prototests", "Container")
	if err != nil {
		panic(err)
	}
}

func TestNoMergeExtensionMerge(t *testing.T) {
	bigm := prototests.AContainer
	m := &prototests.Small{SmallField: proto.Int64(1)}
	data, err := proto.Marshal(bigm)
	if err != nil {
		panic(err)
	}
	mdata, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	key := uint32(101)<<3 | uint32(2)
	datakey := make([]byte, 10)
	n := binary.PutUvarint(datakey, uint64(key))
	datakey = datakey[:n]
	datalen := make([]byte, 10)
	n = binary.PutUvarint(datalen, uint64(len(mdata)))
	datalen = datalen[:n]
	data = append(data, append(datakey, append(datalen, mdata...)...)...)
	err = noMerge(data, bigm.Description(), "prototests", "Container")
	if err == nil || !strings.Contains(err.Error(), "FieldB requires merging") {
		t.Fatalf("FieldB should require merging, but error is %v", err)
	}
	t.Log(err)
}
