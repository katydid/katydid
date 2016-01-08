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

package convert

import (
	"fmt"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/serialize"
	"io"
	"strings"
)

func Interp(auto *auto, tree serialize.Parser) bool {
	start := []*stackstate{{auto.states[auto.start], 0, 0}}
	fs := interpret(auto, start, tree)
	for _, f := range fs {
		if f.current.final {
			return true
		}
	}
	return false
}

type stackstate struct {
	current    *state
	stackState int
	stackFunc  int
}

func stacks(auto *auto, currents []*stackstate) string {
	ss := make([]string, len(currents))
	for i, c := range currents {
		ss[i] = auto.patterns[c.current.current].String()
	}
	return strings.Join(ss, ", ")
}

func interpret(auto *auto, currents []*stackstate, tree serialize.Parser) []*stackstate {
	for {
		fmt.Printf("currents = %v\n", stacks(auto, currents))
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		downs := []*stackstate{}
		for currenti, current := range currents {
			for trani, t := range current.current.trans {
				f, err := compose.NewBoolFunc(t.value)
				if err != nil {
					panic(err)
				}
				e, err := f.Eval(tree)
				if err != nil {
					panic(err)
				}
				if !e {
					continue
				}
				downs = append(downs, &stackstate{
					current:    auto.states[t.down],
					stackState: currenti,
					stackFunc:  trani,
				})
			}
		}
		fmt.Printf("downs = %v\n", stacks(auto, downs))
		ups := downs
		if tree.IsLeaf() {
			fmt.Printf("leaf\n")
		} else {
			fmt.Printf("Down\n")
			tree.Down()
			ups = interpret(auto, downs, tree)
			fmt.Printf("Up\n")
			tree.Up()
		}
		fmt.Printf("ups = %v\n", stacks(auto, ups))
		newCurrents := []*stackstate{}
		for _, up := range ups {
			current := currents[up.stackState]
			tran := current.current.trans[up.stackFunc]
			for _, u := range tran.ups {
				if u.bot == up.current.current {
					newCurrents = append(newCurrents, &stackstate{
						current:    auto.states[u.top],
						stackState: current.stackState,
						stackFunc:  current.stackFunc,
					})
				}
			}
		}
		currents = newCurrents
	}
	return currents
}
