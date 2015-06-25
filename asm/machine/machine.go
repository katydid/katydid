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

package machine

import (
	"fmt"
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/serialize"
	"log"
	"reflect"
	"time"
)

type Catcher interface {
	Catch(error)
	GetError() error
	Clear()
}

type Machine interface {
	Transition(current int, name string) (child int, success int, failure int, err error)
	NoEscapeFrom(current int) bool
	GetState(name string) (int, bool)
	SetFunc(current int) bool
	EvalFunc(serialize.Decoder) bool
	Implements(typ reflect.Type) []interface{}
}

type dst struct {
	child, success, failure int
}

type machine struct {
	table       map[int]map[string]dst
	nameToState map[string]int
	stateToName map[int]string
	maxState    int
	funcs       map[int]compose.Bool
	fnc         compose.Bool
	catcher     Catcher
	debug       bool
}

func New(transitions []*asm.Transition, funcDecls []*asm.FunctionDecl, catcher Catcher) (Machine, error) {
	m := &machine{
		table:       make(map[int]map[string]dst),
		nameToState: make(map[string]int),
		stateToName: make(map[int]string),
		maxState:    0,
		funcs:       make(map[int]compose.Bool),
		catcher:     catcher,
		debug:       true,
	}
	m.getState("_")
	for _, t := range transitions {
		src := m.getState(t.GetSrc())
		name := t.GetInput()
		child := m.getState(t.GetDst().GetChild())
		success := m.getState(t.GetDst().GetSuccess())
		failure := m.getState(t.GetDst().GetFailure())
		if _, ok := m.table[src]; !ok {
			m.table[src] = make(map[string]dst)
		}
		m.table[src][name] = dst{child, success, failure}
	}
	for _, f := range funcDecls {
		fState := m.getState(f.GetName())
		fnc, err := compose.NewBool(&expr.Expr{Function: f.GetFunction()})
		if err != nil {
			return nil, err
		}
		m.funcs[fState] = fnc
	}
	return m, nil
}

//Returns all the function structs instances that implements the given interface
func (m *machine) Implements(typ reflect.Type) []interface{} {
	var is []interface{}
	for _, f := range m.funcs {
		is = append(is, compose.FuncImplements(f, typ)...)
	}
	return is
}

func (m *machine) getState(s string) int {
	if s, ok := m.nameToState[s]; ok {
		return s
	}
	newState := m.maxState
	m.nameToState[s] = newState
	m.stateToName[newState] = s
	m.maxState++
	return newState
}

func (m *machine) GetState(name string) (int, bool) {
	s, ok := m.nameToState[name]
	return s, ok
}

func (m *machine) SetFunc(current int) bool {
	f, ok := m.funcs[current]
	if !ok {
		return false
	}
	m.fnc = f
	return true
}

func (m *machine) EvalFunc(value serialize.Decoder) bool {
	e, err := m.fnc.Eval(value)
	if err != nil {
		m.catcher.Catch(err)
	}
	return e
}

func (m *machine) NoEscapeFrom(current int) bool {
	c, ok := m.table[current]
	if !ok {
		if m.debug {
			log.Printf("NoEscapeFrom %s", m.stateToName[current])
		}
		return true
	}
	for _, d := range c {
		if d.success != current {
			return false
		}
		if d.failure != current {
			return false
		}
	}
	if m.debug {
		log.Printf("NoEscapeFrom %s", m.stateToName[current])
	}
	return true
}

func (m *machine) Transition(current int, name string) (int, int, int, error) {
	c, ok := m.table[current]
	if !ok {
		return -1, -1, -1, fmt.Errorf("unknown transition from %s {%#v}", m.stateToName[current], m)
	}
	d, ok := c[name]
	if !ok {
		d2, ok2 := c["_"]
		if !ok2 {
			return -1, -1, -1, fmt.Errorf("unknown transition from %s given %s {%#v}", m.stateToName[current], name, m)
		}
		if m.debug {
			log.Printf("Transition(%s, _{%s}) -> (%s, %s, %s)", m.stateToName[current], name, m.stateToName[d2.child], m.stateToName[d2.success], m.stateToName[d2.failure])
			time.Sleep(1e8)
		}
		return d2.child, d2.success, d2.failure, nil
	}
	if m.debug {
		log.Printf("Transition(%s, %s) -> (%s, %s, %s)", m.stateToName[current], name, m.stateToName[d.child], m.stateToName[d.success], m.stateToName[d.failure])
		time.Sleep(1e8)
	}
	return d.child, d.success, d.failure, nil
}
