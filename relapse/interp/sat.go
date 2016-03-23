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

package interp

import (
	"github.com/katydid/katydid/relapse"
)

// Returns true if there exists a tree that can satisfy the Grammar.
// This function may return some false positives.
func Satisfiable(g *relapse.Grammar) bool {
	refs := relapse.NewRefsLookup(g)
	p := refs["main"]
	sp := Simplify(refs, p)
	return !isNotZany(sp)
}
