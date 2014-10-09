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
	"github.com/awalterschulze/katydid/asm/compiler"
	"github.com/awalterschulze/katydid/asm/lexer"
	"github.com/awalterschulze/katydid/asm/parser"
	"github.com/awalterschulze/katydid/serialize/proto/scanner"
	"github.com/awalterschulze/katydid/serialize/proto/tokens"
	"testing"
)

var compileErr = `//Is this person's telephone number 0123456789

root = main.Person
main.Person = start
start tel = accept
start _ = start
accept _ = accept

if eq($string(main.Person.Telephone), elem([]string{"0123456789"}, int64(3)))
  then tel
  else notel
`

func TestCompileError(t *testing.T) {
	fileDescriptorSet, err := protoparser.ParseFile("person.proto", ".", gogoprotoImportPath)
	if err != nil {
		panic(err)
	}
	p := parser.NewParser()
	rules, err := p.ParseRules(lexer.NewLexer([]byte(compileErr)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	protoTokens, err := tokens.NewZipped(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	_, _, err = compiler.Compile(rules, protoTokens)
	if err == nil {
		t.Fatalf("expected error")
	}
}

var runtimeErr = `
root = main.Person
main.Person = start
start tel = accept
start _ = start
accept _ = accept

if eq(elem([]string{$string(main.Person.Telephone)}, int64(3)), "0123456789") 
  then tel
  else notel
`

func TestRuntimeError(t *testing.T) {
	fileDescriptorSet, err := protoparser.ParseFile("person.proto", ".", gogoprotoImportPath)
	if err != nil {
		panic(err)
	}
	p := parser.NewParser()
	rules, err := p.ParseRules(lexer.NewLexer([]byte(runtimeErr)))
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
	if err == nil {
		t.Fatalf("expected error")
	}
	if match {
		t.Fatalf("expected match")
	}
}
