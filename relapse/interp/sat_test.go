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
	. "github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/interp"
	"testing"
)

func check(g G, t *testing.T, expected bool) {
	grammar := g.Grammar()
	out := interp.Satisfiable(grammar)
	if out != expected {
		t.Fatalf("expected %v but got %v for %s", expected, out, grammar)
	}
}

func TestSatisfiability(t *testing.T) {
	check(G{
		"main": Any(),
	}, t, true)
	check(G{
		"main": None(),
	}, t, false)
	check(G{
		"main": MatchInAny(
			None(),
		),
	}, t, false)
	check(G{
		"main": MatchIn("A",
			OrMatch(
				MatchIn("B", Any()),
				None(),
			),
		),
	}, t, true)
}
