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

package asm

import (
	"strings"
)

func (this *Rules) Dot() string {
	ss := make([]string, len(this.Rules))
	for i, rule := range this.Rules {
		ss[i] = rule.Dot()
	}
	return "digraph {\n" +
		strings.Join(ss, "\n") + "\n" +
		"}\n"
}

func (this *Rule) Dot() string {
	if this.Root != nil {
		return this.Root.Dot()
	}
	if this.Init != nil {
		return this.Init.Dot()
	}
	if this.Final != nil {
		return this.Final.Dot()
	}
	if this.Transition != nil {
		return this.Transition.Dot()
	}
	if this.FunctionDecl != nil {
		return this.FunctionDecl.Dot()
	}
	panic("unreachable")
}

func (this *Root) Dot() string {
	return "root -> " + `"` + this.GetPackage() + "." + this.GetMessage() + `"`
}

func (this *Init) Dot() string {
	return `"init"` + " -> " + this.GetState()
}

func (this *Final) Dot() string {
	return `"final"` + " -> " + this.GetState()
}

func (this *Transition) Dot() string {
	return this.GetSrc() + " -> " + this.GetDst().GetChild() + ` [label="child_` + this.GetInput() + `"]
	` + this.GetSrc() + " -> " + this.GetDst().GetSuccess() + ` [label="success_` + this.GetInput() + `"]
	` + this.GetSrc() + " -> " + this.GetDst().GetFailure() + ` [label="failure_` + this.GetInput() + `"]`
}

func (this *FunctionDecl) Dot() string {
	return `"` + this.GetName() + " = " + this.GetFunction().String() + `"`
}
