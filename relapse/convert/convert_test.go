//  Copyright 2016 Walter Schulze
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
	"github.com/katydid/katydid/relapse/mem"
	"github.com/katydid/katydid/serialize"
	"testing"
)

func test(t *testing.T, g *relapse.Grammar, parser serialize.Parser, expected bool, desc string) {
	if interp.HasLeftRecursion(g) {
		t.Skipf("convert was not designed to handle recursion")
	}
	mem := mem.Compile(g)
	auto := &convert.Auto{
		Refs:        mem.Refs,
		PatternsMap: mem.PatternsMap,
		Calls:       mem.Calls,
		Returns:     mem.Returns,
		Escapables:  mem.Escapables,
		Zis:         mem.Zis,
		Start:       mem.Start,

		StackElms:       mem.StackElms,
		StateToNullable: mem.StateToNullable,
		Nullables:       mem.Nullables,
		Accept:          mem.Accept,
	}
	match := convert.Interpret(auto, parser)
	if match != expected {
		t.Fatalf("Expected %v on given \n%s\n on \n%s", expected, g.String(), desc)
	}
}
