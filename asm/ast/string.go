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
	"fmt"
	"strings"
)

func (this *Rules) String() string {
	return this.Dot()
}

func (this *Init) String() string {
	return this.Dot()
}

func (this *Transition) String() string {
	return this.Dot()
}

func (this *IfExpr) String() string {
	return this.Dot()
}

func (this *StateExpr) String() string {
	if this.State != nil {
		return this.GetState()
	}
	return this.GetIfExpr().String()
}

func (this *Expr) String() string {
	if this.Terminal != nil {
		return this.GetTerminal().String()
	}
	return this.GetFunction().String()
}

func (this *Function) String() string {
	ps := make([]string, len(this.GetParams()))
	for i, v := range this.GetParams() {
		ps[i] = v.String()
	}
	return this.GetName() + "(" + strings.Join(ps, ",") + ")"
}

func (this *Terminal) String() string {
	if this.BoolValue != nil {
		if this.GetBoolValue() {
			return "true"
		}
		return "false"
	}
	if this.Int64Value != nil {
		return fmt.Sprintf("%d", this.GetInt64Value())
	}
	if this.Uint64Value != nil {
		return fmt.Sprintf("%d", this.GetUint64Value())
	}
	if this.StringValue != nil {
		return this.GetStringValue()
	}
	if this.Variable != nil {
		return this.GetVariable().String()
	}
	if this.BytesValue != nil {
		return fmt.Sprintf("%#v", this.GetBytesValue())
	}
	if this.Int32Value != nil {
		return fmt.Sprintf("%d", this.GetInt32Value())
	}
	if this.Uint32Value != nil {
		return fmt.Sprintf("%d", this.GetUint32Value())
	}
	if this.DoubleValue != nil {
		return fmt.Sprintf("%d", this.GetDoubleValue())
	}
	if this.FloatValue != nil {
		return fmt.Sprintf("%d", this.GetFloatValue())
	}
	panic("unreachable")
}

func (this *Variable) String() string {
	return this.GetPackage() + "." + this.GetMessage() + "." + this.GetField()
}
