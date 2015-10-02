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
		`a->hasQuotes("string_lit")`,
		`#ref1
		@ref1 = a->eq($int, int(123))
		`,
		`#main`,
		`A [
			a->eq($string, "123"), 
			b->eq($string, "456")
		]`,
		`A [
			a->eq($string, "aa"),
			b->eq($int, int(123))
		]`,
		`A[
			*,
			[
				a->eq($string, "aa"),
				*
			]
		]`,
		`A: (a->eq($string, "aa"))*`,
		`Desc [
			*,
			[
				Src->contains($string, "1") ,
				[
					Src->contains($string, "2") ,
					(!(Src): *)*
				]
			]
		]`,
		`(MyParent->any() &  MyParent->any())`,
		`A :[
			*,
			a->eq($string, "aa"),
			*
		]`,
		`A: (
			* &
			a -> eq($string, "aa") &
			*
		)`,
		`A: (
			* |
			a -> eq($string, "aa") |
			* |
			!(<emptyset>) |
			<empty>
		)`,
		`[
			*,
			a -> eq($string, "aa") ,
			*,
			( 
			  b -> contains($string, "bb") 
			  | (
			  	c -> contains($string, "cc")  &
			  	c -> contains($string, "see")  &
			  	(
			  		c -> contains($string, "sea")  |
			  		c -> contains($string, "ocean") 
			  	)
			  )
			  | d -> contains($string, "dd") 
			  | d -> contains($string, "dd") 
			  | d -> contains($string, "dd") 
			)
		]`,
		`_->greaterThanOne(int(1)) `,
		`[
			(_ -> eq($int, int(1)) )*,
			bla -> any() 
		]`,
		`( a|b ) -> eq($int, int(1)) `,
		`( a|_|!(b) ) -> eq($int, int(1)) `,
		`"\"a" -> any() `,
		`(
			.a->any() |
			.b->any()
		)`,
		`(
			.a->any() &
			.b [ a->any(), b->any() ] &
			.c->any() &
			.d->any()
		)`,
		`a.b->any()`,
		`.a._.D.b123->any()`,
		`"Whats Up" == "E"`,
		`(* & * & *)`,
		`(* | * | * | ( * & * & * ))`,
		`[ * , * , *]`,
		`[ * , * , *,]`,
	}
	p := parser.NewParser()
	for i, patternDecl := range patternDecls {
		t.Logf("parsing %d", i)
		scanner := lexer.NewLexer([]byte(patternDecl))
		st, err := p.ParseGrammar(scanner)
		if err != nil {
			t.Errorf("err = %v, input = %s", err, patternDecl)
		} else {
			s := st.String()
			if s != patternDecl {
				t.Errorf("String function output = %s\n expected output = %s\n", s, patternDecl)
			}
		}
	}
}
