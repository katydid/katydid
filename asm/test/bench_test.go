//  Copyright 2013 Walter Schulze
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

package main_test

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/asm/compiler"
	katyexec "github.com/katydid/katydid/asm/exec"
	"github.com/katydid/katydid/asm/lexer"
	"github.com/katydid/katydid/asm/parser"
	main "github.com/katydid/katydid/asm/test"
	"github.com/katydid/katydid/protoparser"
	"github.com/katydid/katydid/serialize/proto/scanner"
	"github.com/katydid/katydid/serialize/proto/tokens"
	"math/rand"
	"testing"
	"time"
)

type bench struct {
	exec    *katyexec.Exec
	scanner scanner.BytesScanner
}

func newBench(protoFilename string, katydidStr string) bench {
	fileDescriptorSet, err := protoparser.ParseFile(protoFilename, ".", gogoprotoImportPath)
	if err != nil {
		panic(err)
	}
	p := parser.NewParser()
	rules, err := p.ParseRules(lexer.NewLexer([]byte(katydidStr)))
	if err != nil {
		panic(err)
	}
	protoTokens, err := tokens.NewZipped(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	e, rootToken, err := compiler.Compile(rules, protoTokens)
	if err != nil {
		panic(err)
	}
	return bench{
		exec:    e,
		scanner: scanner.NewProtoScanner(protoTokens, rootToken),
	}
}

type randyTest interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func (this bench) bench(b *testing.B, newPop func(r randyTest, easy bool) proto.Message) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := 1000
	datas := make([][]byte, num)
	for i := 0; i < num; i++ {
		m := newPop(r, true)
		data, err := proto.Marshal(m)
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := this.scanner.Init(datas[i%num])
		if err != nil {
			panic(err)
		}
		if _, err := this.exec.Eval(this.scanner); err != nil {
			panic(err)
		}
	}
}

func BenchmarkContextPerson(b *testing.B) {
	newBench("person.proto", contextPerson).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}

func BenchmarkRecursiveSrcTree(b *testing.B) {
	pops := []*main.SrcTree{ioUtil, pathSrcTree, runtime, syscall}
	newBench("srctree.proto", recursiveSrcTree).bench(b, func(r randyTest, easy bool) proto.Message {
		return pops[r.Intn(4)]
	})
}

func BenchmarkListIndexAddress(b *testing.B) {
	newBench("person.proto", listIndexAddress).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}

func BenchmarkNilName(b *testing.B) {
	newBench("person.proto", nilName).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}

func BenchmarkLenName(b *testing.B) {
	newBench("person.proto", lenName).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}

func BenchmarkEmptyOrNil(b *testing.B) {
	newBench("person.proto", emptyOrNil).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}

func BenchmarkIncorrectNotName(b *testing.B) {
	newBench("person.proto", incorrentNotName).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}

func BenchmarkCorrectNotName(b *testing.B) {
	newBench("person.proto", correctNotName).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}

func BenchmarkAndNameTelephone(b *testing.B) {
	newBench("person.proto", andNameTelephone).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}

func BenchmarkOrNameTelephone(b *testing.B) {
	newBench("person.proto", orNameTelephone).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPerson(r, easy)
	})
}
