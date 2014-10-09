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

package ifexpr

import (
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/asm/compose"
	"github.com/awalterschulze/katydid/serialize"
)

type NameToState interface {
	NameToState(name string) (state int, err error)
}

type StateExpr interface {
	Eval(serialize.Decoder) int
}

type IfExpr struct {
	Cond    compose.Bool
	Succ    StateExpr
	Fail    StateExpr
	Catcher Catcher
}

func (this *IfExpr) Eval(dec serialize.Decoder) int {
	cond, err := this.Cond.Eval(dec)
	if err != nil {
		this.Catcher.Catch(err)
	}
	if cond {
		return this.Succ.Eval(dec)
	}
	return this.Fail.Eval(dec)
}

type Catcher interface {
	Catch(error)
	GetError() error
	Clear()
}

func Compile(ifexpr *ast.IfExpr, nameToState NameToState, c Catcher) (StateExpr, error) {
	cond, err := compose.NewBool(ifexpr.Condition)
	if err != nil {
		return nil, err
	}
	succ, err := compile(ifexpr.GetThenClause(), nameToState, c)
	if err != nil {
		return nil, err
	}
	fail, err := compile(ifexpr.GetElseClause(), nameToState, c)
	if err != nil {
		return nil, err
	}
	return &IfExpr{cond, succ, fail, c}, nil
}

type constState struct {
	state int
}

func (this *constState) Eval(serialize.Decoder) int {
	return this.state
}

func compile(stateExpr *ast.StateExpr, nameToState NameToState, c Catcher) (StateExpr, error) {
	if stateExpr.State != nil {
		state, err := nameToState.NameToState(stateExpr.GetState())
		if err != nil {
			return nil, err
		}
		return &constState{state}, nil
	}
	return Compile(stateExpr.GetIfExpr(), nameToState, c)
}
