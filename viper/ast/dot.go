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
	"strconv"
	"strings"
)

func quote(s string) string {
	return strconv.Quote(s)
}

func (this *Rules) Dot() string {
	rs := make([]string, len(this.Rules)+2)
	rs[0] = "digraph viper {"
	for i, r := range this.Rules {
		rs[i+1] = r.Dot()
	}
	rs[len(rs)-1] = "}"
	return strings.Join(rs, "\n")
}

func (this *Rule) Dot() string {
	return this.GetValue().(interface {
		Dot() string
	}).Dot()
}

func (this *Start) Dot() string {
	return "start -> " + this.State.Dot() + ";"
}

func (this *Final) Dot() string {
	return this.State.Dot() + " -> " + "final;"
}

func (this *Internal) Dot() string {
	if this.Dst.Dot() == quote(quote("EmptySet")) {
		return ""
	}
	e := quote("internal(" + this.Expr.String() + ")")
	return this.Src.Dot() + " -> " + this.Dst.Dot() + " [label=" + e + "];"
}

func (this *Call) Dot() string {
	if this.ParentDst.Dot() == quote(quote("EmptySet")) {
		return ""
	}
	if this.ChildDst.Dot() == quote(quote("EmptySet")) {
		return ""
	}
	e := this.Expr.String()
	e1 := quote("callParent(" + e + ")")
	e2 := quote("callChild(" + e + ")")
	t1 := this.Src.Dot() + " -> " + this.ParentDst.Dot() + " [label=" + e1 + "];"
	t2 := this.Src.Dot() + " -> " + this.ChildDst.Dot() + " [label=" + e2 + "];"
	return t1 + "\n" + t2
}

func (this *Return) Dot() string {
	if this.Dst.Dot() == quote(quote("EmptySet")) {
		return ""
	}
	e := this.Expr.String()
	e1 := quote("returnParent(" + e + ")")
	e2 := quote("returnChild(" + e + ")")
	t1 := this.ParentSrc.Dot() + " -> " + this.Dst.Dot() + " [label=" + e1 + "];"
	t2 := this.ChildSrc.Dot() + " -> " + this.Dst.Dot() + " [label=" + e2 + "];"
	return t1 + "\n" + t2
}

func (this *State) Dot() string {
	return quote(this.Name)
}
