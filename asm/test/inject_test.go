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
	protoparser "code.google.com/p/gogoprotobuf/parser"
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/asm/compiler"
	"github.com/awalterschulze/katydid/asm/inject"
	"github.com/awalterschulze/katydid/asm/lexer"
	"github.com/awalterschulze/katydid/asm/parser"
	"github.com/awalterschulze/katydid/funcs"
	"reflect"
	"testing"
)

type injectableInt64 struct {
	v int64
}

func (this *injectableInt64) Eval(buf []byte) int64 {
	return this.v
}

func (this *injectableInt64) SetValue(v int64) {
	this.v = v
}

func init() {
	funcs.Register("inject", new(injectableInt64))
}

type InjectableInt64 interface {
	SetValue(v int64)
}

var injectPerson = `root = main.Person
	main.Person = start
	start accept = accept
	start _ = start
	accept _ = accept

	main.Address = start

	if eq(decInt64(main.Address.Number), inject()) then accept else meh
	`

func TestInject(t *testing.T) {
	fileDescriptorSet, err := protoparser.ParseFile("person.proto", ".", gogoprotoImportPath)
	if err != nil {
		panic(err)
	}
	p := parser.NewParser()
	r, err := p.Parse(lexer.NewLexer([]byte(injectPerson)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	rules := r.(*ast.Rules)
	e, err := compiler.Compile(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	typ := reflect.TypeOf((*InjectableInt64)(nil)).Elem()
	instances := inject.Implements(e, typ)
	for _, instance := range instances {
		instance.(InjectableInt64).SetValue(456)
	}
	m := robert
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	match, err := e.Eval(data)
	if err != nil {
		panic(err)
	}
	if !match {
		t.Fatalf("expected match")
	}
}
