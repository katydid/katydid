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

package json_test

import (
	"encoding/json"
	"github.com/katydid/katydid/parser/debug"
	sjson "github.com/katydid/katydid/parser/json"
	"testing"
)

func TestDebug(t *testing.T) {
	p := sjson.NewJsonParser()
	data, err := json.Marshal(debug.Input)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(data); err != nil {
		t.Fatal(err)
	}
	m := debug.Walk(p)
	if !m.Equal(debug.Output) {
		t.Fatalf("expected %s but got %s", debug.Output, m)
	}
}

func TestRandomDebug(t *testing.T) {
	p := sjson.NewJsonParser()
	data, err := json.Marshal(debug.Input)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if err := p.Init(data); err != nil {
			t.Fatal(err)
		}
		//l := debug.NewLogger(p, debug.NewLineLogger())
		debug.RandomWalk(p, debug.NewRand(), 10, 3)
		//t.Logf("original %v vs random %v", debug.Output, m)
	}
}

func TestEscapedChar(t *testing.T) {
	j := map[string][]interface{}{
		`a\"`: {1},
	}
	data, err := json.Marshal(j)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", string(data))
	parser := sjson.NewJsonParser()
	if err := parser.Init(data); err != nil {
		t.Fatal(err)
	}
	m := debug.Walk(parser)
	name := m[0].Label
	if name != `a\"` {
		t.Fatalf("wrong escaped name %s", name)
	}
}

func TestMultiLineArray(t *testing.T) {
	s := `{
		"A":[1]
	}`
	parser := sjson.NewJsonParser()
	if err := parser.Init([]byte(s)); err != nil {
		t.Fatal(err)
	}
	jout := debug.Walk(parser)
	t.Logf("%v", jout)
}

func TestIntWithExponent(t *testing.T) {
	s := `{"A":1e+08}`
	parser := sjson.NewJsonParser()
	if err := parser.Init([]byte(s)); err != nil {
		t.Fatal(err)
	}
	if err := parser.Next(); err != nil {
		t.Fatal(err)
	}
	parser.Down()
	if err := parser.Next(); err != nil {
		t.Fatal(err)
	}
	if !parser.IsLeaf() {
		t.Fatal("incorrect walk, please adjust the path above")
	}
	if i, err := parser.Int(); err != nil {
		t.Fatalf("did not expect error %v", err)
	} else if i != 1e+08 {
		t.Fatalf("got %d", i)
	}
}

func testValue(t *testing.T, input, output string) {
	parser := sjson.NewJsonParser()
	if err := parser.Init([]byte(input)); err != nil {
		t.Fatalf("init error: %v", err)
	}
	jout := debug.Walk(parser)
	if len(jout) != 1 {
		t.Fatalf("expected one node")
	}
	if len(jout[0].Children) != 0 {
		t.Fatalf("did not expected any children")
	}
	if jout[0].Label != output {
		t.Fatalf("expected %s got %s", output, jout[0].Label)
	}
}

func TestValues(t *testing.T) {
	testValue(t, "0", "0")
	testValue(t, "1", "1")
	testValue(t, "-1", "-1")
	testValue(t, "123", "123")
	testValue(t, "1.1", "1.1")
	testValue(t, "1.123", "1.123")
	testValue(t, "1.1e1", "11")
	testValue(t, "1.1e-1", "0.11")
	testValue(t, "1.1e10", "11000000000")
	testValue(t, "1.1e+10", "11000000000")
	testValue(t, `"a"`, "a")
	testValue(t, `"abc"`, "abc")
	testValue(t, `""`, "")
	testValue(t, `"\b"`, "\b")
	testValue(t, `true`, "true")
	testValue(t, `false`, "false")
	testValue(t, `null`, "<nil>")
}

func testArray(t *testing.T, s string) {
	parser := sjson.NewJsonParser()
	if err := parser.Init([]byte(s)); err != nil {
		t.Fatal(err)
	}
	jout := debug.Walk(parser)
	t.Logf("%v", jout)
}

func TestArray(t *testing.T) {
	testArray(t, `[1]`)
	testArray(t, `[1,2.3e5]`)
	testArray(t, `[1,"a"]`)
	testArray(t, `[true, false, null]`)
	testArray(t, `[{"a": true, "b": [1,2]}]`)
}
