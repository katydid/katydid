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

package proto

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/encode/proto/prototests"
	reflectparser "github.com/katydid/katydid/parser/reflect"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestSimple(t *testing.T) {
	msg := prototests.NewPopulatedSimple(r, true)
	p := reflectparser.NewReflectParser()
	p.Init(reflect.ValueOf(msg))
	enc, err := NewEncoder(msg.Description(), "prototests", "Simple")
	if err != nil {
		t.Fatal(err)
	}
	buf, err := enc.Encode(make([]byte, 1024), p)
	if err != nil {
		t.Fatal(err)
	}
	msg2 := &prototests.Simple{}
	if err := proto.Unmarshal(buf, msg2); err != nil {
		t.Fatal(err)
	}
	if !msg.Equal(msg2) {
		t.Fatalf("%#v not equal %#v", msg, msg2)
	}
}

func TestNested(t *testing.T) {
	msg := prototests.NewPopulatedNested(r, true)
	p := reflectparser.NewReflectParser()
	p.Init(reflect.ValueOf(msg))
	enc, err := NewEncoder(msg.Description(), "prototests", "Nested")
	if err != nil {
		t.Fatal(err)
	}
	buf, err := enc.Encode(make([]byte, 1024), p)
	if err != nil {
		t.Fatal(err)
	}
	msg2 := &prototests.Nested{}
	if err := proto.Unmarshal(buf, msg2); err != nil {
		t.Fatal(err)
	}
	if !msg.Equal(msg2) {
		t.Fatalf("%#v not equal %#v", msg, msg2)
	}
}
