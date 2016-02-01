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
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/serialize"
	"io"
	"strings"
)

func Interp(auto *auto, tree serialize.Parser) bool {
	// f := interpret(auto, auto.states[auto.start], tree)
	// return f.final
	n, f := Record(auto, tree)
	fmt.Printf("%v\n", n)
	return f
}
func interpret(auto *auto, current *state, tree serialize.Parser) *state {
	for {
		fmt.Printf("current = %v\n", toStr(current, auto.patterns))
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
				continue
			}
			if !e {
				continue
			}
			fmt.Printf("evaled %v\n", funcs.Sprint(t.value))
			down := auto.states[t.down]
			fmt.Printf("down = %v\n", auto.patterns[down.current])
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
			fmt.Printf("up = %v\n", auto.patterns[up.current])
			for _, cup := range t.ups {
				if cup.bot == up.current {
					current = auto.states[cup.top]
					break
				}
			}
			break
		}
	}
	return current
}

func Record(auto *auto, tree serialize.Parser) (Node, bool) {
	ns, f := record(auto, auto.states[auto.start], tree)
	n := Node{
		Label:    "Root",
		Down:     auto.patterns[auto.start].String(),
		Up:       auto.patterns[f.current].String(),
		Children: ns,
	}
	return n, f.final
}

func getValue(parser serialize.Parser) interface{} {
	var v interface{}
	var err error
	v, err = parser.Int()
	if err == nil {
		return v
	}
	v, err = parser.String()
	if err == nil {
		return v
	}
	v, err = parser.Uint()
	if err == nil {
		return v
	}
	v, err = parser.Double()
	if err == nil {
		return v
	}
	v, err = parser.Bool()
	if err == nil {
		return v
	}
	v, err = parser.Bytes()
	if err == nil {
		return v
	}
	return nil
}

type Node struct {
	Label    string
	Current  string
	Down     string
	Up       string
	Children []Node
}

func (this Node) String() string {
	children := make([]string, len(this.Children))
	for i := range this.Children {
		ss := strings.Split(this.Children[i].String(), "\n")
		for i, s := range ss {
			ss[i] = "\t\t" + s
		}
		children[i] = strings.Join(ss, "\n")
	}
	schild := fmt.Sprintf("\n\tChildren: {\n%v\n\t}", strings.Join(children, "\n"))
	if len(children) == 0 {
		schild = ""
	}
	return fmt.Sprintf("%v:\n\tCurrent: %v\n\tDown: %v%v\n\tUp: %v", this.Label, this.Current, this.Down, schild, this.Up)
}

func record(auto *auto, current *state, tree serialize.Parser) ([]Node, *state) {
	nodes := make([]Node, 0)
	for {
		fmt.Printf("current = %v\n", toStr(current, auto.patterns))
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		value := getValue(tree)
		allfalse := true
		for _, t := range current.trans {
			f, err := compose.NewBoolFunc(t.value)
			if err != nil {
				panic(err)
			}
			e, err := f.Eval(tree)
			if err != nil {
				continue
			}
			if !e {
				continue
			}
			down := auto.states[t.down]
			up := down
			name := fmt.Sprintf("%v", value)
			var children []Node
			if tree.IsLeaf() {
				//do nothing
			} else {
				tree.Down()
				children, up = record(auto, down, tree)
				tree.Up()
			}
			nodes = append(nodes, Node{
				Label:    name,
				Current:  auto.patterns[current.current].String(),
				Down:     auto.patterns[down.current].String(),
				Up:       auto.patterns[up.current].String(),
				Children: children,
			})
			for _, cup := range t.ups {
				if cup.bot == up.current {
					current = auto.states[cup.top]
					allfalse = false
					break
				}
			}
			if allfalse {
				panic("allfalse " + auto.patterns[up.current].String())
			}
			break
		}
		if allfalse {
			panic("allfalse " + fmt.Sprintf("%d ", len(current.trans)) + auto.patterns[current.current].String())
		}
	}
	return nodes, current
}
