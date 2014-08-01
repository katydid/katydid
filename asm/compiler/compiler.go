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
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/asm/exec"
	"github.com/awalterschulze/katydid/asm/ifexpr"
	"github.com/awalterschulze/katydid/asm/link"
	"github.com/awalterschulze/katydid/asm/protomap"
	"github.com/awalterschulze/katydid/asm/table"
	"github.com/awalterschulze/katydid/funcs"
)

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

func findIncludes(rules *ast.Rules) ([]string, error) {
	s := make([]string, 0, 1+len(rules.Rules))
	s = append(s, rules.GetRoot().GetPackage()+"."+rules.GetRoot().GetMessage())
	for _, rule := range rules.Rules {
		if init := rule.Init; init != nil {
			s = append(s, init.GetPackage()+"."+init.GetMessage())
		}
	}
	for _, rule := range rules.Rules {
		if ifExpr := rule.IfExpr; ifExpr != nil {
			v, err := ifexpr.GetVariable(ifExpr)
			if err != nil {
				return nil, err
			}
			s = append(s, v.GetPackage()+"."+v.GetMessage()+"."+v.GetField())
		}
	}
	return s, nil
}

func Compile(rules *ast.Rules, desc *descriptor.FileDescriptorSet) (*exec.Exec, error) {
	includes, err := findIncludes(rules)
	if err != nil {
		return nil, err
	}
	pmap, err := protomap.NewZipped(rules.GetRoot().GetPackage(), rules.GetRoot().GetMessage(), desc, includes)
	if err != nil {
		return nil, err
	}
	tab := table.New(rules.GetTransitions(), rules.GetIfExprs())
	catcher := funcs.NewCatcher(false)
	link, err := link.NewLink(rules, pmap, tab, catcher)
	if err != nil {
		return nil, err
	}
	rootState, err := pmap.State(rules.GetRoot().GetPackage(), rules.GetRoot().GetMessage(), "")
	if err != nil {
		return nil, err
	}
	startState, err := findStartState(rules.GetRoot(), rules.GetInits(), tab)
	if err != nil {
		return nil, err
	}
	acceptState, err := tab.NameToState("accept")
	if err != nil {
		return nil, err
	}
	return exec.NewExec(pmap, tab, link, catcher, rootState, startState, acceptState), nil
}
