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

package interp_test

import (
	"encoding/json"
	"fmt"
	"github.com/katydid/katydid/lang/interp"
	"github.com/katydid/katydid/lang/lexer"
	"github.com/katydid/katydid/lang/parser"
	"github.com/katydid/katydid/serialize/json/scanner"
	"github.com/katydid/katydid/tests"
	"math/rand"
	"testing"
	"time"
)

func test(t *testing.T, m interface{}, langStr string, positive bool) {
	p := parser.NewParser()
	g, err := p.ParseGrammar(lexer.NewLexer([]byte(langStr)))
	if err != nil {
		t.Fatalf(err.Error())
	}
	outputStr := g.String()
	if langStr != outputStr {
		t.Logf("input = <<%s>>", langStr)
		t.Logf("output = <<%s>>", outputStr)
		t.Fatalf("Parsed string should output same string from ast")
	}

	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := scanner.NewJsonScanner()
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	match := interp.Interpret(g, s)
	if match != positive {
		t.Errorf("Expected a %v match from \n%v \non \n%v", positive, langStr, string(data))
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestInterp(t *testing.T) {
	var v int64 = 1
	m := &tests.FinanceJudo{
		RumourSpirit: &v,
	}
	spirit1 := fmt.Sprintf("\nspirit = TreeNode(\"RumourSpirit\", LeafNode(eq($int, int(1))))\n")
	spirit2 := fmt.Sprintf("\nspirit = TreeNode(\"RumourSpirit\", LeafNode(eq($int, int(2))))\n")
	test(t, m, `main = Not(EmptySet)`, true)
	test(t, m, `main = EmptySet`, false)
	test(t, m, `main = Concat(Not(EmptySet), Concat(Reference(spirit), Not(EmptySet)))`+spirit1, true)
	test(t, m, `main = Concat(Not(EmptySet), Concat(Reference(spirit), Not(EmptySet)))`+spirit2, false)
}
