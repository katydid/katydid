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

package expr

import (
	"fmt"
	"github.com/katydid/katydid/types"
	"strconv"
	"strings"
)

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
	if this.DoubleValue != nil {
		return this.Before.String() + "double(" + strconv.FormatFloat(this.GetDoubleValue(), 'f', -1, 10) + ")"
	}
	if this.FloatValue != nil {
		return this.Before.String() + "float(" + strconv.FormatFloat(float64(this.GetFloatValue()), 'f', -1, 10) + ")"
	}
	if this.Int64Value != nil {
		return this.Before.String() + "int64(" + strconv.FormatInt(this.GetInt64Value(), 10) + ")"
	}
	if this.Uint64Value != nil {
		return this.Before.String() + "uint64(" + strconv.FormatUint(this.GetUint64Value(), 10) + ")"
	}
	if this.Int32Value != nil {
		return this.Before.String() + "int32(" + strconv.FormatInt(int64(this.GetInt32Value()), 10) + ")"
	}
	if this.BoolValue != nil {
		return this.Before.String() + "bool(" + strconv.FormatBool(this.GetBoolValue()) + ")"
	}
	if this.StringValue != nil {
		return this.Before.String() + strconv.Quote(this.GetStringValue())
	}
	if this.BytesValue != nil {
		return this.Before.String() + fmt.Sprintf("%#v", this.GetBytesValue())
	}
	if this.Uint32Value != nil {
		return this.Before.String() + "uint32(" + strconv.FormatUint(uint64(this.GetUint32Value()), 10) + ")"
	}
	if this.Variable != nil {
		return this.Before.String() + this.Variable.String()
	}
	panic("unreachable")
}

func (this *Variable) String() string {
	typ := this.GetType()
	if types.IsList(typ) {
		types.ListToSingle(typ)
	}
	switch typ {
	case types.SINGLE_DOUBLE:
		return "$double"
	case types.SINGLE_FLOAT:
		return "$float"
	case types.SINGLE_INT64:
		return "$int64"
	case types.SINGLE_UINT64:
		return "$uint64"
	case types.SINGLE_INT32:
		return "$int32"
	case types.SINGLE_BOOL:
		return "$bool"
	case types.SINGLE_STRING:
		return "$string"
	case types.SINGLE_BYTES:
		return "$[]byte"
	case types.SINGLE_UINT32:
		return "$uint32"
	}
	panic(fmt.Errorf("unknown type %s", this.GetType()))
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
