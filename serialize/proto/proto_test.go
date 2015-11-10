//  Copyright 2015 Walter Schulze
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
	"github.com/katydid/katydid/serialize"
	"github.com/katydid/katydid/serialize/debug"
	"io"
	"testing"
)

func TestDebug(t *testing.T) {
	p := NewProtoParser("debug", "Debug", debug.DebugDescription())
	p.debug = true
	data, err := proto.Marshal(debug.Input)
	if err != nil {
		panic(err)
	}
	p.Init(data)
	parser := debug.NewLogger(p, debug.NewLineLogger())
	m := debug.Walk(parser)
	if !m.Equal(debug.Output) {
		t.Fatalf("expected %s but got %s", debug.Output, m)
	}
}

func TestRandomDebug(t *testing.T) {
	p := NewProtoParser("debug", "Debug", debug.DebugDescription())
	p.debug = true
	data, err := proto.Marshal(debug.Input)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		p.Init(data)
		//l := debug.NewLogger(p, debug.NewLineLogger())
		debug.RandomWalk(p, debug.NewRand(), 10, 3)
		//t.Logf("original %v vs random %v", debug.Output, m)
	}
}

func next(t *testing.T, parser serialize.Parser) {
	if err := parser.Next(); err != nil {
		if err != io.EOF {
			t.Fatal(err)
		}
	}
}

func TestSkipRepeated1(t *testing.T) {
	p := NewProtoParser("debug", "Debug", debug.DebugDescription())
	p.debug = true
	data, err := proto.Marshal(debug.Input)
	if err != nil {
		panic(err)
	}
	p.Init(data)
	parser := debug.NewLogger(p, debug.NewLineLogger())
	next(t, parser)
	next(t, parser)
	parser.Down()
	next(t, parser)
	parser.Down()
	parser.Up()
	next(t, parser)
	parser.Down()
	next(t, parser)
	next(t, parser)
	parser.Up()
	parser.Up()
	next(t, parser)
}

func TestSkipRepeated2(t *testing.T) {
	p := NewProtoParser("debug", "Debug", debug.DebugDescription())
	p.debug = true
	data, err := proto.Marshal(debug.Input)
	if err != nil {
		panic(err)
	}
	p.Init(data)
	parser := debug.NewLogger(p, debug.NewLineLogger())
	next(t, parser)
	parser.String()
	next(t, parser)
	parser.String()
	parser.Down()
	next(t, parser)
	parser.String()
	parser.Up()
	next(t, parser)
}
