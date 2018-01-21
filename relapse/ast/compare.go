//  Copyright 2016 Walter Schulze
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
	"bytes"
)

func (this *Grammar) Equal(that *Grammar) bool {
	return deriveEqualGrammar(this, that)
}

func (this *Pattern) Compare(that *Pattern) int {
	return deriveComparePattern(this, that)
}

func (this *Pattern) Equal(that *Pattern) bool {
	return deriveEqualPattern(this, that)
}

func (p *Pattern) Hash() uint64 {
	return deriveHash(p)
}

func (this *NameExpr) Equal(that *NameExpr) bool {
	return deriveEqualNameExpr(this, that)
}

func (this *NameExpr) Compare(that *NameExpr) int {
	return deriveCompareNameExpr(this, that)
}

func (this *Terminal) Equal(that *Terminal) bool {
	return (this == nil && that == nil) ||
		this != nil && that != nil &&
			deriveEqual_(this.Before, that.Before) &&
			// this.Literal == that.Literal &&
			((this.DoubleValue == nil && that.DoubleValue == nil) || (this.DoubleValue != nil && that.DoubleValue != nil && *(this.DoubleValue) == *(that.DoubleValue))) &&
			((this.IntValue == nil && that.IntValue == nil) || (this.IntValue != nil && that.IntValue != nil && *(this.IntValue) == *(that.IntValue))) &&
			((this.UintValue == nil && that.UintValue == nil) || (this.UintValue != nil && that.UintValue != nil && *(this.UintValue) == *(that.UintValue))) &&
			((this.BoolValue == nil && that.BoolValue == nil) || (this.BoolValue != nil && that.BoolValue != nil && *(this.BoolValue) == *(that.BoolValue))) &&
			((this.StringValue == nil && that.StringValue == nil) || (this.StringValue != nil && that.StringValue != nil && *(this.StringValue) == *(that.StringValue))) &&
			bytes.Equal(this.BytesValue, that.BytesValue) &&
			deriveEqual_26(this.Variable, that.Variable)
}

func (this *Terminal) Compare(that *Terminal) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	if c := deriveCompare_18(this.Before, that.Before); c != 0 {
		return c
	}
	// if c := strings.Compare(this.Literal, that.Literal); c != 0 {
	// 	return c
	// }
	if c := deriveCompare_19(this.DoubleValue, that.DoubleValue); c != 0 {
		return c
	}
	if c := deriveCompare_20(this.IntValue, that.IntValue); c != 0 {
		return c
	}
	if c := deriveCompare_21(this.UintValue, that.UintValue); c != 0 {
		return c
	}
	if c := deriveCompare_22(this.BoolValue, that.BoolValue); c != 0 {
		return c
	}
	if c := deriveCompare_23(this.StringValue, that.StringValue); c != 0 {
		return c
	}
	if c := bytes.Compare(this.BytesValue, that.BytesValue); c != 0 {
		return c
	}
	if c := deriveCompare_29(this.Variable, that.Variable); c != 0 {
		return c
	}
	return 0
}
