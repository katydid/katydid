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
	. "github.com/katydid/katydid/encode/proto/prototests"
	reflectparser "github.com/katydid/katydid/parser/reflect"
	"math/rand"
	"reflect"
	"testing"
)

func BenchmarkMarshalSimple(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*Simple, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedSimple(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkMarshalNested(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]*Nested, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedNested(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkEncodeSimple(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]reflectparser.ReflectParser, 10000)
	for i := 0; i < 10000; i++ {
		msg := NewPopulatedSimple(popr, false)
		p := reflectparser.NewReflectParser()
		p.Init(reflect.ValueOf(msg))
		pops[i] = p
	}
	buf := make([]byte, 1<<20)
	enc, err := NewEncoder((&Simple{}).Description(), "prototests", "Simple")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := enc.Encode(buf, pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkEncodeNested(b *testing.B) {
	popr := rand.New(rand.NewSource(616))
	total := 0
	pops := make([]reflectparser.ReflectParser, 10000)
	for i := 0; i < 10000; i++ {
		msg := NewPopulatedNested(popr, false)
		p := reflectparser.NewReflectParser()
		p.Init(reflect.ValueOf(msg))
		pops[i] = p
	}
	buf := make([]byte, 1<<20)
	enc, err := NewEncoder((&Nested{}).Description(), "prototests", "Nested")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := enc.Encode(buf, pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}
