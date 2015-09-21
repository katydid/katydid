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
	"github.com/katydid/katydid/relapse/lexer"
	"github.com/katydid/katydid/relapse/parser"
	"testing"
)

func TestParse(t *testing.T) {
	patternDecls := []string{
		`main = {"string_lit"}`,
		`main = #ref1
		ref1 = {eq($int, int(123))}
		`,
		`main = "MyParent": [
			{eq($string, "123")}, 
			{eq($string, "456")}
		]`,
		`main = "A" : [
			"a": {eq($string, "aa")},
			"b": {eq($int, int(123))}
		]`,
		`main = "A": [
			!(EmptySet),
			[
				"a": {eq($string, "aa")},
				!(EmptySet)
			]
		]`,
		`main = "Desc": [
			!(EmptySet),
			[
				"Src": { contains($string, "1") },
				[
					"Src": { contains($string, "2") },
					(AnyNameExcept("Src"): !(EmptySet))*
				]
			]
		]`,
		`main = ("MyParent": { int(1) } &  "MyParent": { int(2)})`,
		`main = "A": [
			!(EmptySet),
			"a": { eq($string, "aa") },
			!(EmptySet)
		]`,
		`main = "A": (
			!(EmptySet) &
			"a": { eq($string, "aa") } &
			!(EmptySet)
		)`,
		`main = "A": (
			!(EmptySet) |
			"a": { eq($string, "aa") } |
			!(EmptySet) |
			!(EmptySet) |
			!(EmptySet)
		)`,
		`main ="A": [
			!(EmptySet),
			"a": { eq($string, "aa") },
			( 
			  "b": { contains($string, "bb") }
			  | (
			  	"c": { contains($string, "cc") } &
			  	"c": { contains($string, "see") } &
			  	(
			  		"c": { contains($string, "sea") } |
			  		"c": { contains($string, "ocean") }
			  	)
			  )
			  | "d": { contains($string, "dd") }
			)
		]`,
	}
	p := parser.NewParser()
	for i, patternDecl := range patternDecls {
		t.Logf("parsing %d", i)
		scanner := lexer.NewLexer([]byte(patternDecl))
		st, err := p.ParseGrammar(scanner)
		if err != nil {
			t.Errorf("err = %v, input = %s", err, patternDecl)
		}
		s := st.String()
		if s != patternDecl {
			t.Errorf("String function output = %s\n expected output = %s\n", s, patternDecl)
		}
	}
}
