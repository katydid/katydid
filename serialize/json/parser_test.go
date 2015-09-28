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
	"bytes"
	"encoding/json"
	"github.com/katydid/katydid/serialize/debug"
	sjson "github.com/katydid/katydid/serialize/json"
	"testing"
)

func TestJsonParser(t *testing.T) {
	j := map[string][]interface{}{
		"a": {1},
		"b": {
			map[string][]interface{}{
				"ba": {1, 2, 3},
				"bb": {"string"},
			},
		},
	}
	data, err := json.Marshal(j)
	if err != nil {
		t.Fatal(err)
	}
	parser := sjson.NewJsonParser()
	parser.Init(data)
	jout := debug.Walk(parser)
	data2, err := json.Marshal(jout)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(data, data2) {
		t.Error("bytes not equal")
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
	jout := debug.Walk(parser)
	data2, err := json.Marshal(jout)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(data, data2) {
		t.Error("bytes not equal")
	}
}

func TestMultiLineArray(t *testing.T) {
	s := `{
		"A":[1]
	}`
	parser := sjson.NewJsonParser()
	parser.Init([]byte(s))
	jout := debug.Walk(parser)
	t.Logf("%#v", jout)
}
