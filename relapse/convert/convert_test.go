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

package convert_test

import (
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/convert"
	"github.com/katydid/katydid/serialize"
	"github.com/katydid/katydid/serialize/debug"
	"github.com/katydid/katydid/tests"
	"github.com/katydid/katydid/viper/eval"
	"testing"
)

func test(t *testing.T, g *relapse.Grammar, scanner serialize.Scanner, expected bool, desc string) {
	scanner = debug.NewLogger(scanner, debug.NewLineLogger())
	rules := convert.ToRules(g)
	//println("Rules:" + rules.String())
	match := eval.Eval(rules, scanner)
	if match != expected {
		t.Fail()
		t.Fatalf("Expected %v given %v converted to \n%s\n on \n%s", expected, g, rules, desc)
	}
}

func TestNotNil(t *testing.T) {
	v := tests.Validators["EmptyOrNilJohn"]["reflect"]
	refs := relapse.NewRefsLookup(v.Grammar)
	delete(refs, "main")
	refs["main"] = refs["nil"]
	delete(refs, "nil")
	delete(refs, "empty")
	test(t, relapse.NewGrammar(refs), v.Scanner(), v.Expected, v.Description)
}
