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

package interp

import (
	"github.com/katydid/katydid/relapse/parser"
	"testing"
)

func TestRecursionPositive(t *testing.T) {
	positives := []string{
		`@main`,
		`#main = @main`,
		`#main = (A:* | @main)`,
		`#main = (A:* | @ref)
		 #ref = @main`,
		`#main = (A:* | @ref)	#ref = (B:* | @main | <empty>)`,
		`#main = A:@ref 		#ref = @ref`,
	}
	for _, p := range positives {
		g, err := parser.ParseGrammar(p)
		if err != nil {
			t.Fatal(err)
		}
		if !HasRecursion(g) {
			t.Errorf("expected recursion for %v", g)
		}
	}
}

func TestRecursionNegative(t *testing.T) {
	negatives := []string{
		`A:@main`,
		`#main = A:@main`,
		`#main = (A:* | B:@main)`,
		`#main = (A:* | @ref)
		 #ref = C:@main`,
		`#main = (A:* | @ref)
		 #ref = (B:* | D:@main | <empty>)`,
	}
	for _, n := range negatives {
		g, err := parser.ParseGrammar(n)
		if err != nil {
			t.Fatal(err)
		}
		if HasRecursion(g) {
			t.Errorf("unexpected recursion for %v", g)
		}
	}
}
