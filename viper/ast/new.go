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

package viper

import (
	"github.com/katydid/katydid/expr/ast"
)

func newSpace() *expr.Space {
	return &expr.Space{Space: []string{" "}}
}

func newEqual() *expr.Keyword {
	return &expr.Keyword{Before: newSpace(), Value: "="}
}

func newSemiColon() *expr.Keyword {
	return &expr.Keyword{Value: ";"}
}

func newNewLine() *expr.Space {
	return &expr.Space{Space: []string{"\n"}}
}

func newState(s string) *State {
	return &State{
		Before: newSpace(),
		Name:   s,
	}
}

func newExpr(e *expr.Expr) *expr.Expr {
	// if e.Terminal != nil {
	// 	e.Terminal.Before = newSpace()
	// }
	// if e.List != nil {
	// 	e.List.Before = newSpace()
	// }
	// if e.Function != nil {
	// 	e.Function.Before = newSpace()
	// }
	return e
}

func NewRules(rules ...*Rule) *Rules {
	return &Rules{Rules: rules}
}

func NewStart(start string) *Rule {
	return &Rule{
		Start: &Start{
			Before:    newNewLine(),
			Eq:        newEqual(),
			State:     newState(start),
			SemiColon: newSemiColon(),
		},
	}
}

func NewFinal(final string) *Rule {
	return &Rule{
		Final: &Final{
			Before:    newNewLine(),
			Eq:        newEqual(),
			State:     newState(final),
			SemiColon: newSemiColon(),
		},
	}
}

func NewInternal(src string, expr *expr.Expr, dst string) *Rule {
	return &Rule{
		Internal: &Internal{
			Before:    newNewLine(),
			Src:       newState(src),
			Expr:      newExpr(expr),
			Dst:       newState(dst),
			SemiColon: newSemiColon(),
		},
	}
}

func NewCall(src string, expr *expr.Expr, parentDst, childDst string) *Rule {
	return &Rule{
		Call: &Call{
			Before:    newNewLine(),
			Src:       newState(src),
			Expr:      newExpr(expr),
			ParentDst: newState(parentDst),
			ChildDst:  newState(childDst),
			SemiColon: newSemiColon(),
		},
	}
}

func NewReturn(parentSrc, childSrc string, expr *expr.Expr, dst string) *Rule {
	return &Rule{
		Return: &Return{
			Before:    newNewLine(),
			ParentSrc: newState(parentSrc),
			ChildSrc:  newState(childSrc),
			Expr:      newExpr(expr),
			Dst:       newState(dst),
			SemiColon: newSemiColon(),
		},
	}
}
