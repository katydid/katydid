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

package json

import (
	"bytes"
	"encoding/json"
	"github.com/katydid/katydid/serialize"
	"io"
	"testing"
)

func TestJsonScanner(t *testing.T) {
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
	scanner := NewJsonScanner()
	scanner.Init(data)
	jout := walk(scanner)
	data2, err := json.Marshal(jout)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(data, data2) {
		t.Error("bytes not equal")
	}
}

func getValue(scanner serialize.Scanner) interface{} {
	var v interface{}
	var err error
	v, err = scanner.Bool()
	if err == nil {
		return v
	}
	v, err = scanner.Bytes()
	if err == nil {
		return v
	}
	v, err = scanner.String()
	if err == nil {
		return v
	}
	v, err = scanner.Int64()
	if err == nil {
		return v
	}
	v, err = scanner.Uint64()
	if err == nil {
		return v
	}
	v, err = scanner.Float64()
	if err == nil {
		return v
	}
	panic("not a value")
}

func walk(scanner serialize.Scanner) map[string][]interface{} {
	m := make(map[string][]interface{})
	for {
		if err := scanner.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		name := scanner.Name()
		var v interface{}
		if !scanner.IsLeaf() {
			scanner.Down()
			mm := walk(scanner)
			scanner.Up()
			v = mm
		} else {
			v = getValue(scanner)
		}
		_, ok := m[name]
		if !ok {
			m[name] = make([]interface{}, 0, 1)
		}
		m[name] = append(m[name], v)
	}
	return m
}
