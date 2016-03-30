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

package debug

import (
	"fmt"
	"github.com/katydid/katydid/parser"
	"io"
	"math/rand"
	"strings"
	"time"
)

func getValue(p parser.Interface) interface{} {
	var v interface{}
	var err error
	v, err = p.Int()
	if err == nil {
		return v
	}
	v, err = p.Uint()
	if err == nil {
		return v
	}
	v, err = p.Double()
	if err == nil {
		return v
	}
	v, err = p.Bool()
	if err == nil {
		return v
	}
	v, err = p.String()
	if err == nil {
		return v
	}
	v, err = p.Bytes()
	if err == nil {
		return v
	}
	return nil
}

//Node is a type that represents a node in a tree.
//It has a label an children nodes.
type Node struct {
	Label    string
	Children Nodes
}

//String returns a string representation of Node.
func (this Node) String() string {
	if len(this.Children) == 0 {
		return this.Label
	}
	return this.Label + ":" + this.Children.String()
}

//Equal returns whether two Nodes are the same.
func (this Node) Equal(that Node) bool {
	if this.Label != that.Label {
		return false
	}
	if !this.Children.Equal(that.Children) {
		return false
	}
	return true
}

//Nodes is a list of Node.
type Nodes []Node

//String returns a string representation of Nodes.
func (this Nodes) String() string {
	ss := make([]string, len(this))
	for i := range this {
		ss[i] = this[i].String()
	}
	return "{" + strings.Join(ss, ",") + "}"
}

//Equal returns whether two Node lists are equal.
func (this Nodes) Equal(that Nodes) bool {
	if len(this) != len(that) {
		return false
	}
	for i := range this {
		if !this[i].Equal(that[i]) {
			return false
		}
	}
	return true
}

//Walk walks through the whole parser in a top down manner and records the values into a Nodes structute.
func Walk(p parser.Interface) Nodes {
	a := make(Nodes, 0)
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		value := getValue(p)
		if p.IsLeaf() {
			a = append(a, Node{fmt.Sprintf("%v", value), nil})
		} else {
			name := fmt.Sprintf("%v", value)
			p.Down()
			v := Walk(p)
			p.Up()
			a = append(a, Node{name, v})
		}
	}
	return a
}

//NewRand returns a random integer generator, that can be used with RandomWalk.
//Its seed is the current time.
func NewRand() Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

//Rand is a subset of the interface that is implemented by math/rand representing only the methods used by the RandomWalk function.
type Rand interface {
	Intn(n int) int
}

//RandomWalk does a random walk of the parser, given two probabilities.
//  next is passed to r.Intn and when a zero value is returned the Next method on the parser is called.
//  down is passed to r.Intn and when a non zero value is returned the Down method on the parser is called.
//RandomWalk is great for testing that the implemented parser can handle random skipping of parts of the tree.
func RandomWalk(p parser.Interface, r Rand, next, down int) Nodes {
	a := make(Nodes, 0)
	for {
		if r.Intn(next) == 0 {
			break
		}
		if err := p.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		value := getValue(p)
		if p.IsLeaf() {
			a = append(a, Node{fmt.Sprintf("%#v", value), nil})
		} else {
			name := fmt.Sprintf("%#v", value)
			var v Nodes
			if r.Intn(down) != 0 {
				p.Down()
				v = RandomWalk(p, r, next, down)
				p.Up()
			}
			a = append(a, Node{name, v})
		}
	}
	return a
}
