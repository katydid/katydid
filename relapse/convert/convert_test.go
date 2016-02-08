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
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/serialize"
	//"github.com/katydid/katydid/serialize/debug"
	"github.com/katydid/katydid/tests"
	//"github.com/katydid/katydid/viper/eval"
	//"strings"
	"testing"
	//"time"
)

type vNot struct {
	has bool
}

func (this *vNot) Visit(node interface{}) relapse.Visitor {
	n, ok := node.(*relapse.Not)
	if !ok {
		return this
	}
	if n.GetPattern().ZAny != nil {
		return this
	}
	this.has = true
	return this
}

func hasNot(g *relapse.Grammar) bool {
	vNot := &vNot{}
	g.Walk(vNot)
	return vNot.has
}

func test(t *testing.T, g *relapse.Grammar, parser serialize.Parser, expected bool, desc string) {
	//t.Skipf("TODO")
	if interp.HasLeftRecursion(g) {
		t.Skipf("convert was not designed to handle recursion")
	}
	// if strings.Contains(desc, "#") && strings.Contains(desc, "Right") {
	// 	t.Fatalf("broken")
	// }
	// if hasNot(g) {
	// 	t.Skipf("TODO: currently skipping tests with the not operator")
	// }
	//parser = debug.NewLogger(parser, debug.NewDelayLogger(time.Millisecond*50))
	refs := relapse.NewRefsLookup(g)
	// if len(refs) > 1 {
	// 	t.Skipf("not today")
	// }
	// auto := convert.Convert(refs, refs["main"])
	// match := convert.Interp(auto, parser)
	match := convert.DerivEval(refs, parser)
	if match != expected {
		t.Fail()
		t.Fatalf("Expected %v given %v\n on \n%s", expected, g, desc)
	}
}

func TestNotNil(t *testing.T) {
	v := tests.Validators["NilNameJohn"]["reflect"]
	refs := relapse.NewRefsLookup(v.Grammar)
	test(t, relapse.NewGrammar(refs), v.Parser(), v.Expected, v.Description)
}
