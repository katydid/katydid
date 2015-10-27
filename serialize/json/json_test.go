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
	"github.com/katydid/katydid/serialize/debug"
	sjson "github.com/katydid/katydid/serialize/json"
	"strconv"
	"testing"
)

func TestDebug(t *testing.T) {
	p := sjson.NewJsonParser()
	data, err := json.Marshal(debug.Input)
	if err != nil {
		t.Fatal(err)
	}
	p.Init(data)
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
		p.Init(data)
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
	parser.Init(data)
	m := debug.Walk(parser)
	name, err := strconv.Unquote(m[0].Label)
	if err != nil {
		t.Fatal(err)
	}
	if name != `a\"` {
		t.Fatalf("wrong escaped name %s", name)
	}
}

func TestMultiLineArray(t *testing.T) {
	s := `{
		"A":[1]
	}`
	parser := sjson.NewJsonParser()
	parser.Init([]byte(s))
	jout := debug.Walk(parser)
	t.Logf("%v", jout)
}
