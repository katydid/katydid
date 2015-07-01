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

func (this *Rules) String() string {
	lines := make([]string, len(this.Rules))
	for _, r := range this.Rules {
		lines = append(lines, r.String())
	}
	return strings.Join(lines, "") + this.Final.String()
}

func (this *Rule) String() string {
	if this.Root != nil {
		return this.Root.String()
	}
	if this.Init != nil {
		return this.Init.String()
	}
	if this.Final != nil {
		return this.Final.String()
	}
	if this.Transition != nil {
		return this.Transition.String()
	}
	if this.FunctionDecl != nil {
		return this.FunctionDecl.String()
	}
	panic("unreachable")
}

func (this *Root) String() string {
	return this.Before.String() + "root" + this.Eq.String() + this.BeforeQualId.String() + this.Package + "." + this.Message
}

func (this *Init) String() string {
	return this.Before.String() + "init" + this.Eq.String() + this.BeforeState.String() + this.State
}

func (this *Final) String() string {
	return this.Before.String() + "final" + this.Eq.String() + this.BeforeState.String() + this.GetState()
}

func (this *Transition) String() string {
	return this.Before.String() + this.Src + this.BeforeInput.String() +
		this.Input + this.Eq.String() + this.Dst.String()
}

func (this *Destination) String() string {
	return this.OpenParen.String() + this.BeforeChild.String() +
		this.Child + this.CommaOne.String() + this.BeforeSuccess.String() +
		this.Success + this.CommaTwo.String() + this.BeforeFailure.String() +
		this.Failure + this.CloseParen.String()
}

func (this *FunctionDecl) String() string {
	return this.Before.String() + "func" + this.BeforeName.String() +
		this.Name + this.Eq.String() + this.BeforeFunc.String() +
		this.Function.String()
}
