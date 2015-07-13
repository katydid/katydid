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

func newEqual() *expr.Keyword {
	return &expr.Keyword{Value: "="}
}

func newSemiColon() *expr.Keyword {
	return &expr.Keyword{Value: ";"}
}

func newState(s string) *State {
	return &State{
		Name: s,
	}
}

func NewRules(rules ...*Rule) *Rules {
	return &Rules{Rules: rules}
}

func NewStart(start string) *Rule {
	return &Rule{
		Start: &Start{
			Eq:        newEqual(),
			State:     newState(start),
			SemiColon: newSemiColon(),
		},
	}
}

func NewFinal(start string) *Rule {
	return &Rule{
		Final: &Final{
			Eq:        newEqual(),
			State:     newState(start),
			SemiColon: newSemiColon(),
		},
	}
}

func NewInternal(src string, expr *expr.Expr, dst string) *Rule {
	return &Rule{
		Internal: &Internal{
			Src:       newState(src),
			Expr:      expr,
			Dst:       newState(dst),
			SemiColon: newSemiColon(),
		},
	}
}

func NewCall(src string, expr *expr.Expr, parentDst, childDst string) *Rule {
	return &Rule{
		Call: &Call{
			Src:       newState(src),
			Expr:      expr,
			ParentDst: newState(parentDst),
			ChildDst:  newState(childDst),
			SemiColon: newSemiColon(),
		},
	}
}

func NewReturn(parentSrc, childSrc string, expr *expr.Expr, dst string) *Rule {
	return &Rule{
		Return: &Return{
			ParentSrc: newState(parentSrc),
			ChildSrc:  newState(childSrc),
			Expr:      expr,
			Dst:       newState(dst),
			SemiColon: newSemiColon(),
		},
	}
}
