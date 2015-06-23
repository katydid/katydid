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
	"github.com/katydid/katydid/asm/lexer"
	"github.com/katydid/katydid/asm/parser"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/protoparser"
	"github.com/katydid/katydid/serialize/proto/scanner"
	"github.com/katydid/katydid/serialize/proto/tokens"
	"reflect"
	"testing"
)

type injectableInt64 struct {
	v int64
}

func (this *injectableInt64) Eval() int64 {
	return this.v
}

func (this *injectableInt64) SetValue(v int64) {
	this.v = v
}

func (this *injectableInt64) IsVariable() {
	//If this method is not implemented this function will probably be trimmed
}

func init() {
	funcs.Register("inject", new(injectableInt64))
}

type InjectableInt64 interface {
	SetValue(v int64)
}

var injectPerson = `root = main.Person
	init = start
	final = accept
	start Addresses = (address, accept, start)
	start _ = (accept, start, start)
	accept _ = (accept, accept, accept)
	reject _ = (reject, reject, reject)
	address Number = (number, accept, reject)
	address _ = (accept, address, reject)
	func number = eq($int64, inject())
`

func testInject(t *testing.T, num int64) bool {
	fileDescriptorSet, err := protoparser.ParseFile("person.proto", ".", gogoprotoImportPath)
	if err != nil {
		panic(err)
	}
	p := parser.NewParser()
	rules, err := p.ParseRules(lexer.NewLexer([]byte(injectPerson)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	protoTokens, err := tokens.New(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	e, rootToken, err := compiler.Compile(rules, protoTokens)
	if err != nil {
		panic(err)
	}
	typ := reflect.TypeOf((*InjectableInt64)(nil)).Elem()
	instances := e.Implements(typ)
	for _, instance := range instances {
		instance.(InjectableInt64).SetValue(num)
	}
	m := robert
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := scanner.NewProtoScanner(protoTokens, rootToken)
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	match, err := e.Eval(s)
	if err != nil {
		panic(err)
	}
	return match
}

func TestInjectPositive(t *testing.T) {
	if !testInject(t, 456) {
		t.Fatalf("expected match")
	}
}

func TestInjectNegative(t *testing.T) {
	if testInject(t, 123) {
		t.Fatalf("expected non match")
	}
}
