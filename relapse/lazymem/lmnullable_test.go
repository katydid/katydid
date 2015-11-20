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

package lazymem

import (
	. "github.com/katydid/katydid/relapse/combinator"
	"testing"
)

func TestNullableLeftRecursionTrue(t *testing.T) {
	g := G{"main": AnyOf(Eval("main"), Any())}
	p := ConvertGrammar(g.Grammar())
	if !Nullable(p) {
		t.Fatalf("expected nullable for %v", g.Grammar())
	}
}

func TestNullableLeftRecursionFalse(t *testing.T) {
	g := G{"main": AnyOf(Eval("main"), In("a", Any()))}
	p := ConvertGrammar(g.Grammar())
	if Nullable(p) {
		t.Fatalf("expected not nullable for %v", g.Grammar())
	}
}

func TestNullableHiddenLeftRecursionTrue(t *testing.T) {
	g := G{"main": AnyOf(Eval("main1"), Any()), "main1": Eval("main")}
	p := ConvertGrammar(g.Grammar())
	if !Nullable(p) {
		t.Fatalf("expected nullable for %v", g.Grammar())
	}
}

func TestNullableHiddenLeftRecursionFalse(t *testing.T) {
	g := G{"main": AnyOf(Eval("main1"), In("a", Any())), "main1": Eval("main")}
	p := ConvertGrammar(g.Grammar())
	if Nullable(p) {
		t.Fatalf("expected not nullable for %v", g.Grammar())
	}
}

/*
TODO are these tests valid
func TestNullableConcatLeftRecursionTrue(t *testing.T) {
	g := G{"main": InOrder(Eval("main"), Any())}
	p := ConvertGrammar(g.Grammar())
	if !Nullable(p) {
		t.Fatalf("expected nullable for %v", g.Grammar())
	}
}

func TestNullableConcatLeftRecursionFalse(t *testing.T) {
	g := G{"main": InOrder(Eval("main"), In("a", Any()))}
	p := ConvertGrammar(g.Grammar())
	if Nullable(p) {
		t.Fatalf("expected not nullable for %v", g.Grammar())
	}
}
*/
