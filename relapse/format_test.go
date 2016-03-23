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
	"github.com/katydid/katydid/relapse"
	"github.com/katydid/katydid/relapse/parser"
	"testing"
)

func format(t *testing.T, s string) string {
	g, err := parser.ParseGrammar(s)
	if err != nil {
		t.Errorf("err: %v given %s", err, s)
		return ""
	}
	relapse.Format(g)
	return g.String()
}

func testFormat(t *testing.T, in string, expected string) {
	got := format(t, in)
	if expected != got {
		t.Errorf("expected \n<<%s>>\n, but got \n<<%s>>\n", expected, got)
	} else {
		if len(got) > 0 {
			got2 := format(t, got)
			if expected != got2 {
				t.Errorf("expected2 \n<<%s>>\n, but got \n<<%s>>\n", expected, got2)
			}
		}
	}
}

func DisabledTestFormat(t *testing.T) {
	testFormat(t, " @main =  *", "*")
	testFormat(t,
		`//attachedcomment
  *`,
		`//attachedcomment
*`)
	testFormat(t,
		`//unattachedcomment

	*`,
		`//unattachedcomment

*`)
	//2 pattern declarations
	testFormat(t,
		`@main = *
		@other = a->any()`,
		`*
@other = a->any()`)
	//3 pattern declarations
	testFormat(t,
		`*
		@other = a->any()

		@more = (*)*`,
		`*
@other = a->any()
@more = (*)*`)
	//treenode
	testFormat(t,
		`@main = 
			"a":*`,
		`a: *`)
}
