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

package parser_test

import (
	"github.com/katydid/katydid/lang/lexer"
	"github.com/katydid/katydid/lang/parser"
	"testing"
)

func TestParse(t *testing.T) {
	patternDecls := []string{
		`main = LeafNode("string_lit")`,
		`main = Reference(ref1)
		ref1 = LeafNode(eq($int, int(123)))
		`,
		`main = TreeNode("MyParent", Concat(
			LeafNode(eq($string, "123")), 
			LeafNode(eq($string, "456"))
		))`,
		`main = TreeNode("A", Concat(
			TreeNode("a", LeafNode(eq($string, "aa"))),
			TreeNode("b", LeafNode(eq($int, int(123))))
		))`,
		`main = TreeNode("A", Concat(
			Not(EmptySet),
			Concat(
				TreeNode("a", LeafNode(eq($string, "aa"))),
				Not(EmptySet)
			)
		))`,
		`main = TreeNode("Desc", Concat(
			Not(EmptySet),
			Concat(
				TreeNode("Src", LeafNode(contains($string, "1"))),
				Concat(
					TreeNode("Src", LeafNode(contains($string, "2"))),
					ZeroOrMore(TreeNode(not("Src"), Not(EmptySet)))
				)
			)
		))`,
		`main = And(TreeNode("MyParent", LeafNode(int(1))), TreeNode(  "MyParent", LeafNode(int(2))))`,
	}
	p := parser.NewParser()
	for _, patternDecl := range patternDecls {
		scanner := lexer.NewLexer([]byte(patternDecl))
		st, err := p.ParseGrammar(scanner)
		if err != nil {
			t.Fatalf("err = %v, input = %s", err, patternDecl)
		}
		s := st.String()
		if s != patternDecl {
			t.Fatalf("String function output = %s\n expected output = %s\n", s, patternDecl)
		}
	}
}
