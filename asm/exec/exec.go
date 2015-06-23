//  Copyright 2013 Walter Schulze
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

package exec

import (
	"github.com/katydid/katydid/serialize"
	"io"
	"log"
	"reflect"
)

type Machine interface {
	Transition(src int, name string) (child int, success int, failure int, err error)
	NoEscapeFrom(src int) bool
	SetFunc(id int) bool
	EvalFunc(serialize.Decoder) bool
	Implements(typ reflect.Type) []interface{}
}

type Catcher interface {
	Catch(error)
	GetError() error
	Clear()
}

type Scanner interface {
	Next() error
	IsLeaf() bool
	Name() string
	Up()
	Down()
	serialize.Decoder
}

func NewExec(machine Machine, catcher Catcher, startState int, acceptStates []int) *Exec {
	return &Exec{
		machine:      machine,
		catcher:      catcher,
		startState:   startState,
		acceptStates: acceptStates,
		debug:        true,
	}
}

type Exec struct {
	machine      Machine
	catcher      Catcher
	startState   int
	acceptStates []int
	scanner      Scanner
	debug        bool
}

func elem(i int, js []int) bool {
	for _, j := range js {
		if j == i {
			return true
		}
	}
	return false
}

func (e *Exec) Implements(typ reflect.Type) []interface{} {
	return e.machine.Implements(typ)
}

func (e *Exec) Eval(scanner Scanner) (bool, error) {
	e.catcher.Clear()
	e.scanner = scanner
	i, err := e.eval(e.startState)
	if err != nil && err != io.EOF {
		return false, err
	}
	return elem(i, e.acceptStates), e.catcher.GetError()
}

func (e *Exec) eval(state int) (int, error) {
	if e.machine.NoEscapeFrom(state) {
		return state, nil
	}
	if e.debug {
		log.Printf("eval -> Next")
	}
	err := e.scanner.Next()
	var child, success, failure int
	for err == nil {
		child, success, failure, err = e.machine.Transition(state, e.scanner.Name())
		if err != nil {
			return -1, err
		}
		if e.scanner.IsLeaf() {
			if ok := e.machine.SetFunc(child); ok {
				if e.machine.EvalFunc(e.scanner) {
					if e.debug {
						log.Printf("Leaf -> success")
					}
					state = success
				} else {
					if e.debug {
						log.Printf("Leaf -> failure")
					}
					state = failure
				}
			} else {
				if e.debug {
					log.Printf("Leaf -> NoFunc")
				}
				state = failure
			}
			if e.machine.NoEscapeFrom(state) {
				return state, nil
			}
		} else {
			if e.debug {
				log.Printf("Node -> Down")
			}
			e.scanner.Down()
			final, err := e.eval(child)
			if err != nil {
				return 0, err
			}
			e.scanner.Up()
			if elem(final, e.acceptStates) {
				if e.debug {
					log.Printf("Node -> Up -> success")
				}
				state = success
			} else {
				if e.debug {
					log.Printf("Node -> Up -> failure")
				}
				state = failure
			}
			if e.machine.NoEscapeFrom(state) {
				return state, nil
			}
		}
		if e.debug {
			log.Printf("loop -> Next")
		}
		err = e.scanner.Next()
		if e.debug {
			if err != nil {
				log.Printf("loop err = %v", err)
			}
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return state, nil
}
