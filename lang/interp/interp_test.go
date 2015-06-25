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
	exprparser "github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/lang/ast"
	"github.com/katydid/katydid/lang/interp"
	"github.com/katydid/katydid/serialize/json/scanner"
	"github.com/katydid/katydid/tests"
	"math/rand"
	"testing"
	"time"
)

type G map[string]*lang.Pattern

func test(t *testing.T, m interface{}, ps G, positive bool) {
	g := lang.NewGrammar(ps)
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
		t.Errorf("Expected a %v match from \n%v \non \n%v", positive, g.String(), string(data))
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func any() *lang.Pattern {
	return lang.NewNot(lang.NewEmptySet())
}

func field(nameStr string, exprStr string) *lang.Pattern {
	name, err := exprparser.NewParser().ParseExpr(nameStr)
	if err != nil {
		panic(err)
	}
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return lang.NewTreeNode(name, lang.NewLeafNode(expr))
}

func TestInterp(t *testing.T) {
	var v int64 = 1
	m := &tests.FinanceJudo{
		RumourSpirit: &v,
	}
	test(t, m, G{"main": any()}, true)
	test(t, m, G{"main": lang.NewEmptySet()}, false)
	someSpirit := lang.NewConcat(any(), lang.NewConcat(lang.NewReference("spirit"), any()))
	test(t, m, G{"main": someSpirit, "spirit": field(`"RumourSpirit"`, "eq($int, int(1))")}, true)
	test(t, m, G{"main": someSpirit, "spirit": field(`"RumourSpirit"`, "eq($int, int(2))")}, false)
}
