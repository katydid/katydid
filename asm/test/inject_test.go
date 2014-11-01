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
	"github.com/katydid/katydid/asm/compiler"
	"github.com/katydid/katydid/asm/inject"
	"github.com/katydid/katydid/asm/lexer"
	"github.com/katydid/katydid/asm/parser"
	"github.com/katydid/katydid/funcs"
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
	main.Person = start
	start accept = accept
	start _ = start
	accept _ = accept

	main.Address = start

	if eq($int64(main.Address.Number), inject()) then accept else meh
	`

func TestInject(t *testing.T) {
	fileDescriptorSet, err := protoparser.ParseFile("person.proto", ".", gogoprotoImportPath)
	if err != nil {
		panic(err)
	}
	p := parser.NewParser()
	rules, err := p.ParseRules(lexer.NewLexer([]byte(injectPerson)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	protoTokens, err := tokens.NewZipped(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	e, rootToken, err := compiler.Compile(rules, protoTokens)
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
	s := scanner.NewProtoScanner(protoTokens, rootToken)
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	match, err := e.Eval(s)
	if err != nil {
		panic(err)
	}
	if !match {
		t.Fatalf("expected match")
	}
}
