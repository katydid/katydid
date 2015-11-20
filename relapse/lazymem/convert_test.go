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

var RecursiveTurn = G{"main": AnyOf(
	InAny(Eval("main")),
	In("Turn", Any()),
)}

func TestConvert(t *testing.T) {
	g := RecursiveTurn.Grammar()
	t.Logf("%v", g)
	p := ConvertGrammar(g)
	gg := ConvertPattern(p)
	t.Logf("%v", gg)
	if !g.Equal(gg) {
		t.Fatalf("expected equal grammars")
	}
}
