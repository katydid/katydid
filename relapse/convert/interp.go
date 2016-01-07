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
)

func Interp(auto *auto, tree serialize.Parser) bool {
	f := interpret(auto, auto.states[auto.start], tree)
	return f.final
}

func interpret(auto *auto, current *state, tree serialize.Parser) *state {
	for {
		fmt.Printf("Next current = %v\n", auto.patterns[current.current].String())
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		for _, t := range current.trans {
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
			down := auto.states[t.down]
			up := down
			if tree.IsLeaf() {
				fmt.Printf("leaf\n")
			} else {
				fmt.Printf("Down\n")
				tree.Down()
				up = interpret(auto, down, tree)
				fmt.Printf("Up\n")
				tree.Up()
			}
			for _, u := range t.ups {
				if u.bot == up.current {
					current = auto.states[u.top]
				}
			}
			fmt.Printf("up current = %v\n", auto.patterns[current.current].String())
			break
		}
	}
	return current
}
