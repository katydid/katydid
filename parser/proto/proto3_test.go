//  Copyright 2017 Walter Schulze
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
	"math"
	"strconv"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/parser/debug"
	"github.com/katydid/katydid/parser/proto/prototests"
)

var proto3Input1 = &prototests.Proto3{
	Field: 97824789,
	Msg: &prototests.SmallMsg3{
		ScarBusStop:     "cde",
		FlightParachute: []uint32{1, 2, 3},
	},
	Ints: []int64{math.MinInt64},
}

var proto3Output1 = debug.Nodes{
	debug.Field(`Field`, `97824789`),
	debug.Nested(`Msg`,
		debug.Field(`ScarBusStop`, `cde`),
		debug.Nested(`FlightParachute`,
			debug.Field(`0`, `1`),
			debug.Field(`1`, `2`),
			debug.Field(`2`, `3`),
		),
	),
	debug.Nested(`Ints`,
		debug.Field(`0`, strconv.Itoa(math.MinInt64)),
	),
}

func TestProto31(t *testing.T) {
	p := NewProtoNameParser("prototests", "Proto3", proto3Input1.Description())
	data, err := proto.Marshal(proto3Input1)
	if err != nil {
		panic(err)
	}
	if err := p.Init(data); err != nil {
		t.Fatal(err)
	}
	parser := debug.NewLogger(p, debug.NewLineLogger())
	m := debug.Walk(parser)
	if !m.Equal(proto3Output1) {
		t.Fatalf("expected %s but got %s", proto3Output1, m)
	}
}

func TestRandomProto31(t *testing.T) {
	p := NewProtoNameParser("prototests", "Proto3", proto3Input1.Description())
	data, err := proto.Marshal(proto3Input1)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		if err := p.Init(data); err != nil {
			t.Fatal(err)
		}
		l := debug.NewLogger(p, debug.NewLineLogger())
		debug.RandomWalk(l, debug.NewRand(), 10, 3)
	}
}
