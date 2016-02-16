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
	"strings"
)

func (this *Rules) String() string {
	rs := make([]string, len(this.Rules))
	for i, r := range this.Rules {
		rs[i] = r.String()
	}
	return strings.Join(rs, "") + this.After.String()
}

func (this *Rule) String() string {
	return this.GetValue().(interface {
		String() string
	}).String()
}

func (this *Start) String() string {
	return this.Before.String() + "start" +
		this.Eq.String() + this.State.String() +
		this.SemiColon.String()
}

func (this *Final) String() string {
	return this.Before.String() + "final" +
		this.Eq.String() + this.State.String() +
		this.SemiColon.String()
}

func (this *Internal) String() string {
	return this.Before.String() + "internal" +
		this.Src.String() + this.Expr.String() +
		this.Dst.String() + this.SemiColon.String()
}

func (this *Call) String() string {
	return this.Before.String() + "call" +
		this.Src.String() + this.Expr.String() +
		this.ParentDst.String() + this.ChildDst.String() +
		this.SemiColon.String()
}

func (this *Return) String() string {
	return this.Before.String() + "return" +
		this.ParentSrc.String() + this.ChildSrc.String() +
		this.Expr.String() + this.Dst.String() +
		this.SemiColon.String()
}

func (this *State) String() string {
	return this.Before.String() + this.Name
}
