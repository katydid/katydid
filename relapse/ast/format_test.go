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

package relapse_test

import (
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/parser"
	"testing"
)

func format(t *testing.T, s string) string {
	g, err := parser.ParseGrammar(s)
	if err != nil {
		t.Fatal(err)
	}
	relapse.Format(g)
	return g.String()
}

func testFormat(t *testing.T, in string, expected string) {
	got := format(t, in)
	if expected != got {
		t.Fatalf("expected \n<<%s>>\n, but got \n<<%s>>\n", expected, got)
	}
	got2 := format(t, got)
	if expected != got2 {
		t.Fatalf("expected2 \n<<%s>>\n, but got \n<<%s>>\n", expected, got2)
	}
}

func TestFormat(t *testing.T) {
	testFormat(t, " main =  Empty", "main = Empty")
	testFormat(t,
		`//attachedcomment
    main = EmptySet`,
		`//attachedcomment
main = EmptySet`)
	testFormat(t,
		`//unattachedcomment

    main = EmptySet`,
		`//unattachedcomment

main = EmptySet`)
	//2 pattern declarations
	testFormat(t,
		`main = Empty
		other = EmptySet`,
		`main = Empty
other = EmptySet`)
	//3 pattern declarations
	testFormat(t,
		`main = Empty
		other = EmptySet

		more = ZeroOrMore(Empty)`,
		`main = Empty
other = EmptySet
more = ZeroOrMore(Empty)`)
	//treenode
	testFormat(t,
		`main = 
			Name("a"):Empty`,
		`main = Name("a"): Empty`)
}
