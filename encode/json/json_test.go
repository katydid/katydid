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

package json

import (
	"encoding/json"
	jsonparser "github.com/katydid/katydid/parser/json"
	reflectparser "github.com/katydid/katydid/parser/reflect"
	"reflect"
	"testing"
)

func testTranscode(t *testing.T, input interface{}) {
	inputData, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	t.Logf("testing transcoding: %s", string(inputData))
	p := jsonparser.NewJsonParser()
	if err := p.Init(inputData); err != nil {
		t.Fatal(err)
	}
	outputData, err := Encode(p)
	if err != nil {
		t.Fatal(err)
	}
	output := reflect.New(reflect.TypeOf(input).Elem()).Interface()
	if err := json.Unmarshal(outputData, output); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("expected equal, but got %s", string(outputData))
	}
}

func testMarshal(t *testing.T, input interface{}) {
	t.Logf("testing marshaling: %#v", input)
	p := reflectparser.NewReflectParser()
	p.Init(reflect.ValueOf(input))
	outputData, err := Encode(p)
	if err != nil {
		t.Fatal(err)
	}
	output := reflect.New(reflect.TypeOf(input).Elem()).Interface()
	if err := json.Unmarshal(outputData, output); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("expected equal, but got %s", string(outputData))
	}
}

type A struct {
	Name  string
	Value int64
}

func TestTranscodeA(t *testing.T) {
	testTranscode(t, &A{Name: "a", Value: 1})
	testTranscode(t, &A{Name: "a"})
	testTranscode(t, &A{Value: 1})
	testTranscode(t, &A{})
}

func TestMarshalA(t *testing.T) {
	testMarshal(t, &A{Name: "a", Value: 1})
	testMarshal(t, &A{Name: "a"})
	testMarshal(t, &A{Value: 1})
	testMarshal(t, &A{})
}

type B struct {
	Name string
	A    *A
}

func TestTranscodeB(t *testing.T) {
	testTranscode(t, &B{Name: "b", A: &A{Name: "a", Value: 1}})
	testTranscode(t, &B{Name: "b"})
}

func TestMarshalB(t *testing.T) {
	testMarshal(t, &B{Name: "b", A: &A{Name: "a", Value: 1}})
	testMarshal(t, &B{Name: "b"})
}

type C struct {
	Name string
	List []int64
}

func TestTranscodeC(t *testing.T) {
	testTranscode(t, &C{Name: "c", List: []int64{1, 2, 3, 4}})
	testTranscode(t, &C{Name: "c"})
}

func TestMarshalC(t *testing.T) {
	testMarshal(t, &C{Name: "c", List: []int64{1, 2, 3, 4}})
	testMarshal(t, &C{Name: "c"})
}

type D struct {
	Name string
	List []*A
}

func TestTranscodeD(t *testing.T) {
	testTranscode(t, &D{Name: "d", List: []*A{{Name: "a", Value: 1}, {}, nil}})
}

func TestMarshalD(t *testing.T) {
	testMarshal(t, &D{Name: "d", List: []*A{{Name: "a", Value: 1}, {}, nil}})
}
