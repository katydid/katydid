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
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/serialize"
	"github.com/katydid/katydid/serialize/debug"
	"testing"

	exprparser "github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/tests"
)

func test(t *testing.T, g *relapse.Grammar, scanner serialize.Scanner, expected bool, desc string) {
	if interp.HasLeftRecursion(g) {
		t.Skipf("interp was not designed to handle left recursion")
	}
	scanner = debug.NewLogger(scanner, debug.NewLineLogger())
	match := interp.Interpret(g, scanner)
	if match != expected {
		t.Fatalf("Expected %v on given \n%s\n on \n%s", expected, g.String(), desc)
	}
}

// var RobertPerson = &Person{
// 	Name: proto.String("Robert"),
// 	Addresses: []*Address{
// 		{
// 			Number: proto.Int64(456),
// 			Street: proto.String("TheStreet"),
// 		},
// 	},
// 	Telephone: proto.String("0127897897"),
// }

func TestWithSomeTreeNode(t *testing.T) {
	expr, err := exprparser.NewParser().ParseExpr(funcs.Sprint(funcs.IntVarEq(funcs.IntConst(456))))
	if err != nil {
		panic(err)
	}
	s := tests.Validators["ContextRobert"]["json"].Scanner()
	g := relapse.NewGrammar(map[string]*relapse.Pattern{
		"main": relapse.NewWithSomeTreeNode(relapse.NewName("Addresses"),
			relapse.NewWithSomeLeafNode(relapse.NewName("Number"), expr)),
	})
	test(t, g, s, true, "WithSomeTreeNode")
}
