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

package compiler

import (
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/asm/exec"
	"github.com/katydid/katydid/asm/link"
	"github.com/katydid/katydid/asm/table"
	"github.com/katydid/katydid/funcs"
)

type Tokens interface {
	GetTokenId(tokenString string) (int, error)
	Len() int
}

type errNoStartState struct {
	pkg string
	msg string
}

func (this *errNoStartState) Error() string {
	return "did not specify any start state for " + this.pkg + "." + this.msg
}

func findStartState(root *ast.Root, inits []*ast.Init, tab table.Table) (int, error) {
	for _, init := range inits {
		if init.GetMessage() == root.GetMessage() {
			if init.GetPackage() == root.GetPackage() {
				startState, err := tab.NameToState(init.GetState())
				if err != nil {
					return 0, err
				}
				return startState, nil
			}
		}
	}
	return 0, &errNoStartState{root.GetPackage(), root.GetMessage()}
}

func Compile(rules *ast.Rules, tokens Tokens) (eval *exec.Exec, rootToken int, err error) {
	tab := table.New(rules.GetTransitions(), rules.GetIfExprs())
	catcher := funcs.NewCatcher(false)
	link, err := link.NewLink(rules, tokens, tab, catcher)
	if err != nil {
		return nil, 0, err
	}
	rootToken, err = tokens.GetTokenId(rules.GetRoot().GetPackage() + "." + rules.GetRoot().GetMessage())
	if err != nil {
		return nil, 0, err
	}
	startState, err := findStartState(rules.GetRoot(), rules.GetInits(), tab)
	if err != nil {
		return nil, 0, err
	}
	acceptState, err := tab.NameToState("accept")
	if err != nil {
		return nil, 0, err
	}
	eval = exec.NewExec(tab, link, catcher, startState, acceptState)
	return eval, rootToken, nil
}
