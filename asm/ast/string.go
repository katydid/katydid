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
	"github.com/katydid/katydid/types"
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
	if this.Transition != nil {
		return this.Transition.String()
	}
	if this.IfExpr != nil {
		return this.IfExpr.String()
	}
	panic("unreachable")
}

func (this *Root) String() string {
	return this.Before.String() + "root" + this.Equal_.String() + this.BeforeQualId.String() + this.Package + "." + this.Message
}

func (this *Init) String() string {
	return this.Before.String() + this.Package + "." + this.Message + this.Equal_.String() + this.BeforeState.String() + this.State
}

func (this *Transition) String() string {
	return this.Before.String() + this.Src + this.BeforeInput.String() + this.Input + this.Equal_.String() + this.BeforeDst.String() + this.Dst
}

func (this *IfExpr) String() string {
	return this.Before.String() + "if" + this.Condition.String() + this.ThenWord.String() + this.ThenClause.String() + this.ElseWord.String() + this.ElseClause.String()
}

func (this *StateExpr) String() string {
	space := this.Before.String()
	if this.State != nil {
		return space + this.GetState()
	}
	return space + "{" + this.GetIfExpr().String() + this.CloseCurly.String()
}

func (this *Expr) String() string {
	space := this.Comma.String()
	if this.Terminal != nil {
		return space + this.GetTerminal().String()
	}
	if this.List != nil {
		return space + this.GetList().String()
	}
	return space + this.GetFunction().String()
}

func (this *List) String() string {
	es := make([]string, len(this.GetElems()))
	for i, v := range this.GetElems() {
		es[i] = v.String()
	}
	return this.Before.String() + listTypeToString(this.Type) + this.OpenCurly.String() + strings.Join(es, "") + this.CloseCurly.String()
}

func listTypeToString(typ types.Type) string {
	switch typ {
	case types.LIST_DOUBLE:
		return "[]double"
	case types.LIST_FLOAT:
		return "[]float"
	case types.LIST_INT64:
		return "[]int64"
	case types.LIST_UINT64:
		return "[]uint64"
	case types.LIST_INT32:
		return "[]int32"
	case types.LIST_BOOL:
		return "[]bool"
	case types.LIST_STRING:
		return "[]string"
	case types.LIST_BYTES:
		return "[][]byte"
	case types.LIST_UINT32:
		return "[]uint32"
	}
	panic("unreachable")
}

func (this *Function) String() string {
	ps := make([]string, len(this.GetParams()))
	for i, v := range this.GetParams() {
		ps[i] = v.String()
	}
	return this.Before.String() + this.GetName() + this.OpenParen.String() + strings.Join(ps, "") + this.CloseParen.String()
}

func (this *Terminal) String() string {
	return this.Before.String() + this.Literal
}

func (this *Variable) String() string {
	return this.Name
}

func (this *Keyword) String() string {
	if this == nil {
		return ""
	}
	return this.Before.String() + this.Value
}

func (this *Space) String() string {
	if this == nil {
		return ""
	}
	return strings.Join(this.Space, "")
}
