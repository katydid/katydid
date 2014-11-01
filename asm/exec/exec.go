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
)

type Table interface {
	Trans(src int, input int) (int, error)
	NoEscapeFrom(src int) bool
	StateToName(state int) string
}

type Link interface {
	TokenToStart(token int) (startState int, exists bool)
	GetIf(token int) bool
	IfEval(serialize.Decoder) int
}

type Catcher interface {
	Catch(error)
	GetError() error
	Clear()
}

type Scanner interface {
	Next() error
	IsLeaf() bool
	Id() int
	Up()
	Down()
	serialize.Decoder
}

func NewExec(table Table, link Link, catcher Catcher, startState int, acceptState int) *Exec {
	return &Exec{
		table:       table,
		Link:        link,
		catcher:     catcher,
		startState:  startState,
		acceptState: acceptState,
	}
}

type Exec struct {
	table       Table
	Link        Link
	catcher     Catcher
	startState  int
	acceptState int
	scanner     Scanner
}

func (this *Exec) Eval(scanner Scanner) (bool, error) {
	this.catcher.Clear()
	this.scanner = scanner
	i, err := this.eval(this.startState)
	if err != nil && err != io.EOF {
		return false, err
	}
	return i == this.acceptState, this.catcher.GetError()
}

func (this *Exec) eval(state int) (int, error) {
	if this.table.NoEscapeFrom(state) {
		return state, nil
	}
	err := this.scanner.Next()
	for err == nil {
		if this.scanner.IsLeaf() {
			if this.Link.GetIf(this.scanner.Id()) {
				inputSymbol := this.Link.IfEval(this.scanner)
				state, err = this.table.Trans(state, inputSymbol)
				if err != nil {
					return 0, err
				}
				if this.table.NoEscapeFrom(state) {
					return state, nil
				}
			}
		} else {
			startState, ok := this.Link.TokenToStart(this.scanner.Id())
			if ok {
				this.scanner.Down()
				inputSymbol, err := this.eval(startState)
				if err != nil {
					return 0, err
				}
				this.scanner.Up()
				state, err = this.table.Trans(state, inputSymbol)
				if err != nil {
					return 0, err
				}
				if this.table.NoEscapeFrom(state) {
					return state, nil
				}
			}
		}
		err = this.scanner.Next()
	}
	if err != io.EOF {
		return 0, err
	}
	return state, nil
}
