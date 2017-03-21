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
	"testing"

	"github.com/katydid/katydid/relapse/parser"
)

func TestParse(t *testing.T) {
	patternDecls := []string{
		`->hasQuotes("string_lit")`,
		`a->hasQuotes("string_lit")`,
		`@ref1
		#ref1 = a->eq($int, 123)
		`,
		`@main`,
		`A [
			a->eq($string, "123"), 
			b->eq($string, "456")
		]`,
		`A [
			a->eq($string, "aa"),
			b->eq($int, 123)
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
			!(*) |
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
		`_->greaterThanOne(1) `,
		`[
			(_ -> eq($int, 1) )*,
			bla -> any() 
		]`,
		`( a|b ) -> eq($int, 1) `,
		`( a|_|!(b) ) -> eq($int, 1) `,
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
		`[(*)*, == "a", > 1, :: $string]`,
		`[(*)*, ->eq($string, "a"), > 1, :: $string]`,
		`[*,_ ->eq($string, "a"),_ > 1,_ :: $string]`,
		`[0 ->eq($string, "a"),1 > 1,2 :: $string]`,
		`(0|1) == "a"`,
		`(*)?`,
		`{a:* ; b:*}`,
		`{
			a:* ; 
			b:*;
		}`,
		`.A.1 == "a"`,
		`.1.B == "b"`,
	}
	p := parser.NewParser()
	for i, patternDecl := range patternDecls {
		t.Logf("parsing %d", i)
		st, err := p.ParseGrammar(patternDecl)
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

func TestParseExpr(t *testing.T) {
	positives := []string{
		`-> "String"`,
		"-> false",
		`== "bla"`,
		`-> eq($string, "bla")`,
		`-> 1`,
		`-> double(1.0)`,
		`-> double(-1.0)`,
	}
	negatives := []string{
		"a",
		`= "bla"`,
	}
	for _, in := range positives {
		_, err := parser.NewParser().ParseExpr(in)
		if err != nil {
			t.Errorf("%s results in error: %s", in, err)
		}
	}
	for _, in := range negatives {
		_, err := parser.NewParser().ParseExpr(in)
		if err == nil {
			t.Errorf("%s results in no error", in)
		}
	}
}
