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

package mem

import (
	"github.com/katydid/katydid/relapse/parser"
	"testing"
)

func TestExplosionAndSameTree(t *testing.T) {
	var input = `.A:.B:
(
	.DeepLevel:.DeeperLevel:.DeepestLevel:->contains($string,"el") &
	(
		.Rs:._:->eq("~a",$string) &
		(
			.NumI32:->contains($int,[]int{28,1,52}) &
			(
				.NumI64:>= 1 &
				(
					.NumU32:<= uint(4) &
					(
						.NumU64:== uint(4) &
						(
							.YesNo:== true &
							(
								.BS:== []byte{0x3, 0x2, 0x1, 0x0} &
								.Uuid: == []byte{0x3, 0x2, 0x1, 0x0}
							)
						)
					)
				)
			)
		)
	)
)
`
	// This one causes a state explosion of over 14000 states.
	// Since we now field names can't repeat the simplifacation above can be made for JSON and proto like serialization formats, but not for XML.
	// var input = (
	// 	.A:.B:.DeepLevel:.DeeperLevel:.DeepestLevel:->contains($string,"el") &
	// 	(
	// 		.A:.B:.Rs:._:->eq("~a",$string) &
	// 		(
	// 			.A:.B:.NumI32:->contains($int,[]int{28,1,52}) &
	// 			(
	// 				.A:.B:.NumI64:>= 1 &
	// 				(
	// 					.A:.B:.NumU32:<= uint(4) &
	// 					(
	// 						.A:.B:.NumU64:== uint(4) &
	// 						(
	// 							.A:.B:.YesNo:== true &
	// 							(
	// 								.A:.B:.BS:== []byte{0x3, 0x2, 0x1, 0x0} &
	// 								.A:.B:.Uuid: == []byte{0x3, 0x2, 0x1, 0x0}
	// 							)
	// 						)
	// 					)
	// 				)
	// 			)
	// 		)
	// 	)
	// )
	g, err := parser.ParseGrammar(input)
	if err != nil {
		panic(err)
	}
	mem := Compile(g)
	t.Logf("number of states %d", len(mem.PatternsMap))
	if len(mem.PatternsMap) > 1000 {
		t.Fatal("number of states exploded")
	}
}
