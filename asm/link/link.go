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

package link

import (
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/asm/ifexpr"
	"github.com/katydid/katydid/asm/table"
	"github.com/katydid/katydid/serialize"
)

type Link interface {
	TokenToStart(token int) (startState int, exists bool)
	GetIf(token int) bool
	IfEval(serialize.Decoder) int
	GetIfs() []ifexpr.StateExpr
}

type Tokens interface {
	GetTokenId(tokenString string) (int, error)
	Len() int
}

type link struct {
	tokenToStart []int
	tokenToIf    []ifexpr.StateExpr
	currentIf    ifexpr.StateExpr
}

func NewLink(rules *ast.Rules, tokens Tokens, tab table.Table, c ifexpr.Catcher) (Link, error) {
	l := tokens.Len()
	link := &link{
		tokenToStart: make([]int, l),
		tokenToIf:    make([]ifexpr.StateExpr, l),
	}
	for _, rule := range rules.Rules {
		if init := rule.Init; init != nil {
			pstate, err := tokens.GetTokenId(init.GetPackage() + "." + init.GetMessage())
			if err != nil {
				return nil, err
			}
			tstate, err := tab.NameToState(init.GetState())
			if err != nil {
				return nil, err
			}
			link.tokenToStart[pstate] = tstate
		}
	}
	for _, rule := range rules.Rules {
		if ifExpr := rule.IfExpr; ifExpr != nil {
			var1, err := ifExpr.GetVariable()
			if err != nil {
				return nil, err
			}
			pstate, err := tokens.GetTokenId(var1.Name)
			if err != nil {
				return nil, err
			}
			stateExpr, err := ifexpr.Compile(ifExpr, tab, c)
			if err != nil {
				return nil, err
			}
			link.tokenToIf[pstate] = stateExpr
		}
	}
	return link, nil
}

func (this *link) TokenToStart(token int) (startState int, exists bool) {
	if token < 0 || token >= len(this.tokenToStart) {
		return 0, false
	}
	s := this.tokenToStart[token]
	return s, (s != 0)
}

func (this *link) GetIf(token int) bool {
	if token < 0 || token >= len(this.tokenToIf) {
		return false
	}
	ifExpr := this.tokenToIf[token]
	if ifExpr == nil {
		return false
	}
	this.currentIf = ifExpr
	return true
}

func (this *link) IfEval(decoder serialize.Decoder) int {
	return this.currentIf.Eval(decoder)
}

func (this *link) GetIfs() []ifexpr.StateExpr {
	return this.tokenToIf
}
