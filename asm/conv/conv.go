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

package conv

import (
	"github.com/awalterschulze/katydid/asm/lexer"
	"github.com/awalterschulze/katydid/asm/parser"
	"github.com/awalterschulze/katydid/types"
)

type ErrUnknownType struct {
	typ string
}

func (this *ErrUnknownType) Error() string {
	return "unknown type: " + this.typ
}

type ErrWrongType struct {
	exp string
	got string
}

func (this *ErrWrongType) Error() string {
	return "wrong type: expected " + this.exp + ", but got " + this.got
}

func FromGo(v interface{}) (string, error) {
	switch value := v.(type) {
	case float64:
		return DoubleFromGo(value), nil
	case float32:
		return FloatFromGo(value), nil
	case int64:
		return Int64FromGo(value), nil
	case uint64:
		return Uint64FromGo(value), nil
	case int32:
		return Int32FromGo(value), nil
	case bool:
		return BoolFromGo(value), nil
	case string:
		return StringFromGo(value), nil
	case []byte:
		return BytesFromGo(value), nil
	case uint32:
		return Uint32FromGo(value), nil
	case []float64:
		return DoublesFromGo(value), nil
	case []float32:
		return FloatsFromGo(value), nil
	case []int64:
		return Int64sFromGo(value), nil
	case []uint64:
		return Uint64sFromGo(value), nil
	case []int32:
		return Int32sFromGo(value), nil
	case []bool:
		return BoolsFromGo(value), nil
	case []string:
		return StringsFromGo(value), nil
	case [][]byte:
		return ListOfBytesFromGo(value), nil
	case []uint32:
		return Uint32sFromGo(value), nil
	}
	panic("not implemented")
}

func ToGo(s string) (interface{}, error) {
	expr, err := parser.NewParser().ParseExpr(lexer.NewLexer([]byte(s)))
	if err != nil {
		return nil, err
	}
	if expr.List == nil {
		if expr.Terminal == nil {
			return nil, &ErrUnknownType{expr.String()}
		}
		term := expr.GetTerminal()
		if term.DoubleValue != nil {
			return doubleToGo(term)
		}
		if term.FloatValue != nil {
			return floatToGo(term)
		}
		if term.Int64Value != nil {
			return int64ToGo(term)
		}
		if term.Uint64Value != nil {
			return uint64ToGo(term)
		}
		if term.Int32Value != nil {
			return int32ToGo(term)
		}
		if term.BoolValue != nil {
			return boolToGo(term)
		}
		if term.StringValue != nil {
			return stringToGo(term)
		}
		if term.BytesValue != nil {
			return bytesToGo(term)
		}
		if term.Uint32Value != nil {
			return uint32ToGo(term)
		}
		return nil, &ErrUnknownType{expr.String()}
	}
	list := expr.GetList()
	elems := list.GetElems()
	switch list.GetType() {
	case types.LIST_DOUBLE:
		return doublesToGo(elems)
	case types.LIST_FLOAT:
		return floatsToGo(elems)
	case types.LIST_INT64:
		return int64sToGo(elems)
	case types.LIST_UINT64:
		return uint64sToGo(elems)
	case types.LIST_INT32:
		return int32sToGo(elems)
	case types.LIST_BOOL:
		return boolsToGo(elems)
	case types.LIST_STRING:
		return stringsToGo(elems)
	case types.LIST_BYTES:
		return listOfBytesToGo(elems)
	case types.LIST_UINT32:
		return uint32sToGo(elems)
	}
	return nil, &ErrUnknownType{expr.String()}
}
