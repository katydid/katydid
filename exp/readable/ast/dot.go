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

package ast

import (
	"strings"
)

func (this *Rules) Dot() string {
	is := make([]string, len(this.GetInit()))
	for i, v := range this.GetInit() {
		is[i] = v.Dot()
	}
	ts := make([]string, len(this.GetTransition()))
	for i, v := range this.GetTransition() {
		ts[i] = v.Dot()
	}
	ifs := make([]string, len(this.GetIfExpr()))
	for i, v := range this.GetIfExpr() {
		ifs[i] = v.Dot()
	}
	return "digraph {\n" +
		this.GetRoot().String() + "\n" +
		strings.Join(is, "\n") + "\n" +
		strings.Join(ts, "\n") + "\n" +
		strings.Join(ifs, "\n") + "\n" +
		"}\n"
}

func (this *Init) Dot() string {
	return `"` + this.GetPackage() + "." + this.GetMessage() + `"` + " -> " + this.GetState()
}

func (this *Transition) Dot() string {
	return this.GetSrc() + " -> " + this.GetDst() + ` [label="` + this.GetInput() + `"]`
}

func (this *IfExpr) Dot() string {
	return this.GetCondition().Dot() + " -> {" + this.GetThen().Dot() + `} [label="true"]` + "\n" +
		this.GetCondition().Dot() + " -> {" + this.GetElse().Dot() + `} [label="false"]`
}

func (this *StateExpr) Dot() string {
	if this.State != nil {
		return this.GetState()
	}
	return this.GetIfExpr().Dot()
}

func (this *Expr) Dot() string {
	if this.Terminal != nil {
		return this.GetTerminal().Dot()
	}
	return this.GetFunction().Dot()
}

func (this *Function) Dot() string {
	return `<` + this.String() + `>`
}

func (this *Terminal) Dot() string {
	return `"` + this.String() + `"`
}

func (this *Variable) Dot() string {
	return `"` + this.String() + `"`
}
