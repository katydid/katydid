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
	"errors"
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/asm/exec"
	"github.com/katydid/katydid/asm/machine"
	"github.com/katydid/katydid/funcs"
)

type Tokens interface {
	GetTokenId(tokenString string) (int, error)
	Len() int
}

func Compile(rules *asm.Rules, tokens Tokens) (eval *exec.Exec, rootToken int, err error) {
	catcher := funcs.NewCatcher(false)
	mach, err := machine.New(rules.GetTransitions(), rules.GetFunctionDecls(), catcher)
	if err != nil {
		return nil, 0, err
	}
	if !rules.HasInit() {
		return nil, 0, errors.New("missing init state")
	}
	if !rules.HasRoot() {
		return nil, 0, errors.New("missing root")
	}
	rootToken, err = tokens.GetTokenId(rules.GetRoot().GetPackage() + "." + rules.GetRoot().GetMessage())
	if err != nil {
		return nil, 0, err
	}
	startState, ok := mach.GetState(rules.GetInit().GetState())
	if !ok {
		return nil, 0, errors.New("init state does not have any transitions")
	}
	acceptStates := make([]int, 0)
	for _, f := range rules.GetFinals() {
		a, ok := mach.GetState(f.GetState())
		if !ok {
			return nil, 0, errors.New("final state is not reachable")
		}
		acceptStates = append(acceptStates, a)
	}
	eval = exec.NewExec(mach, catcher, startState, acceptStates)
	return eval, rootToken, nil
}
