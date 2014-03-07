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
	"github.com/awalterschulze/katydid/exp/asm/ast"
	"github.com/awalterschulze/katydid/exp/asm/compose"
	"github.com/awalterschulze/katydid/exp/funcs"
)

type errUnmatchedVariables struct {
	var1 string
	var2 string
}

func newErrUnmatchedVariables(var1 *ast.Variable, var2 *ast.Variable) error {
	return &errUnmatchedVariables{
		var1: var1.GetPackage() + "." + var1.GetMessage() + "." + var1.GetField(),
		var2: var2.GetPackage() + "." + var2.GetMessage() + "." + var2.GetField(),
	}
}

func (this *errUnmatchedVariables) Error() string {
	return "Variables in if expr do not match " + this.var1 + " != " + this.var2
}

type NameToState interface {
	NameToState(name string) (state int, err error)
}

type StateExpr interface {
	Eval(buf []byte) int
}

type ifExpr struct {
	cond funcs.Bool
	succ StateExpr
	fail StateExpr
}

func (this *ifExpr) Eval(buf []byte) int {
	if this.cond.Eval(buf) {
		return this.succ.Eval(buf)
	}
	return this.fail.Eval(buf)
}

func Compile(ifexpr *ast.IfExpr, nameToState NameToState) (StateExpr, error) {
	cond, err := compose.ComposeBool(ifexpr.Condition)
	if err != nil {
		return nil, err
	}
	succ, err := compile(ifexpr.GetThen(), nameToState)
	if err != nil {
		return nil, err
	}
	fail, err := compile(ifexpr.GetElse(), nameToState)
	if err != nil {
		return nil, err
	}
	return &ifExpr{cond, succ, fail}, nil
}

func GetVariable(ifExpr *ast.IfExpr) (*ast.Variable, error) {
	var1, err := getVariable(ifExpr.GetCondition())
	if err != nil {
		return nil, err
	}
	if ifExpr.GetThen().IfExpr != nil {
		var2, err := GetVariable(ifExpr.GetThen().GetIfExpr())
		if err != nil {
			return nil, err
		}
		if var1 == nil {
			var1 = var2
		} else if var2 != nil {
			if var1.GetPackage() != var2.GetPackage() ||
				var1.GetMessage() != var2.GetMessage() ||
				var1.GetField() != var2.GetField() {
				return nil, newErrUnmatchedVariables(var1, var2)
			}
		}
	}
	if ifExpr.GetElse().IfExpr != nil {
		var2, err := GetVariable(ifExpr.GetElse().GetIfExpr())
		if err != nil {
			return nil, err
		}
		if var1 == nil {
			var1 = var2
		} else if var2 != nil {
			if var1.GetPackage() != var2.GetPackage() ||
				var1.GetMessage() != var2.GetMessage() ||
				var1.GetField() != var2.GetField() {
				return nil, newErrUnmatchedVariables(var1, var2)
			}
		}
	}
	return var1, nil
}

func getVariable(expr *ast.Expr) (*ast.Variable, error) {
	if expr.Terminal != nil {
		return expr.GetTerminal().GetVariable(), nil
	}
	var var1 *ast.Variable
	for _, param := range expr.GetFunction().GetParams() {
		var2, err := getVariable(param)
		if err != nil {
			return nil, err
		}
		if var1 == nil {
			var1 = var2
		} else if var2 != nil {
			if var1.GetPackage() != var2.GetPackage() ||
				var1.GetMessage() != var2.GetMessage() ||
				var1.GetField() != var2.GetField() {
				return nil, newErrUnmatchedVariables(var1, var2)
			}
		}
	}
	return var1, nil
}

type constState struct {
	state int
}

func (this *constState) Eval(buf []byte) int {
	return this.state
}

func compile(stateExpr *ast.StateExpr, nameToState NameToState) (StateExpr, error) {
	if stateExpr.State != nil {
		state, err := nameToState.NameToState(stateExpr.GetState())
		if err != nil {
			return nil, err
		}
		return &constState{state}, nil
	}
	return Compile(stateExpr.GetIfExpr(), nameToState)
}
