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
	"testing"

	"encoding/json"
	protoparser "github.com/gogo/protobuf/parser"
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/asm/compiler"
	"github.com/katydid/katydid/asm/lexer"
	"github.com/katydid/katydid/asm/parser"
	jsonscanner "github.com/katydid/katydid/serialize/json/scanner"
	jsontokens "github.com/katydid/katydid/serialize/json/tokens"
	"github.com/katydid/katydid/serialize/proto/scanner"
	"github.com/katydid/katydid/serialize/proto/tokens"
	reflectscanner "github.com/katydid/katydid/serialize/reflect/scanner"
	"reflect"
)

var (
	gogoprotoImportPath = "../../../../../:../../../../../github.com/gogo/protobuf/protobuf"
)

func test(t *testing.T, protoFilename string, m proto.Message, katydidStr string, positive bool) {
	fileDescriptorSet, err := protoparser.ParseFile(protoFilename, ".", gogoprotoImportPath)
	if err != nil {
		panic(err)
	}
	p := parser.NewParser()
	rules, err := p.ParseRules(lexer.NewLexer([]byte(katydidStr)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	outputStr := rules.String()
	if katydidStr != outputStr {
		t.Logf("input = <<%s>>", katydidStr)
		t.Logf("output = <<%s>>", outputStr)
		t.Fatalf("Parsed string should output same string from ast")
	}

	//Testing Query on Protocol Buffer Marshaled Structures
	protoTokens, err := tokens.NewZipped(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	e, rootToken, err := compiler.Compile(rules, protoTokens)
	if err != nil {
		panic(err)
	}
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := scanner.NewProtoScanner(protoTokens, rootToken)
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	if match, err := e.Eval(s); err != nil {
		t.Errorf("Error: %v", err)
	} else if match != positive {
		t.Errorf("Expected a %v match from \n%v \non \n%v", positive, katydidStr, m)
	}

	//Testing Query on Json Marshaled Structures
	jsonTokens, err := jsontokens.NewZipped(rules, fileDescriptorSet)
	if err != nil {
		panic(err)
	}
	e, rootToken, err = compiler.Compile(rules, jsonTokens)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	jsonS := jsonscanner.NewScanner(jsonTokens, rootToken)
	err = jsonS.Init(jsonData)
	if err != nil {
		panic(err)
	}
	if match, err := e.Eval(jsonS); err != nil {
		t.Errorf("Error: %v", err)
	} else if match != positive {
		t.Errorf("Expected a %v match from \n%v \non \n%v", positive, katydidStr, m)
	}

	//Testing Query on Reflected Structures
	e, rootToken, err = compiler.Compile(rules, jsonTokens)
	if err != nil {
		panic(err)
	}
	if match, err := e.Eval(reflectscanner.NewScanner(jsonTokens, rootToken).Init(reflect.ValueOf(m))); err != nil {
		t.Errorf("Error: %v", err)
	} else if match != positive {
		t.Errorf("Expected a %v match from \n%v \non \n%v", positive, katydidStr, m)
	}
}
