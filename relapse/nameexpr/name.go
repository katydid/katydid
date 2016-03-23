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

package nameexpr

import (
	"fmt"
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/expr/funcs"
	exprparser "github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/serialize"
)

func EvalName(nameExpr *relapse.NameExpr, tree serialize.Decoder) bool {
	f := NameToFunc(nameExpr)
	b, err := compose.NewBoolFunc(f)
	if err != nil {
		panic(err)
	}
	v, err := b.Eval(tree)
	if err != nil {
		panic(err)
	}
	return v
}

func FuncToName(f funcs.Bool) *relapse.NameExpr {
	exprStr := funcs.Sprint(f)
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return exprToName(expr)
}

func exprToName(e *expr.Expr) *relapse.NameExpr {
	if e.GetBuiltIn() != nil {
		if e.GetBuiltIn().GetSymbol().String() == "==" {
			if e.GetBuiltIn().GetExpr().GetTerminal() != nil {
				t := e.GetBuiltIn().GetExpr().GetTerminal()
				if t.DoubleValue != nil {
					return relapse.NewDoubleName(t.GetDoubleValue())
				}
				if t.IntValue != nil {
					return relapse.NewIntName(t.GetIntValue())
				}
				if t.UintValue != nil {
					return relapse.NewUintName(t.GetUintValue())
				}
				if t.BoolValue != nil {
					return relapse.NewBoolName(t.GetBoolValue())
				}
				if t.StringValue != nil {
					return relapse.NewStringName(t.GetStringValue())
				}
				if t.BytesValue != nil {
					return relapse.NewBytesName(t.GetBytesValue())
				}
			} else {
				panic("todo")
			}
		} else {
			panic("todo")
		}
	}
	if e.GetFunction() != nil {
		if e.GetFunction().GetName() == "not" {
			return relapse.NewAnyNameExcept(exprToName(e.GetFunction().GetParams()[0]))
		}
		if e.GetFunction().GetName() == "or" {
			return relapse.NewNameChoice(
				exprToName(e.GetFunction().GetParams()[0]),
				exprToName(e.GetFunction().GetParams()[1]),
			)
		}
		panic("todo")
	}
	if e.GetTerminal() != nil {
		if e.GetTerminal().BoolValue != nil {
			if e.GetTerminal().GetBoolValue() {
				return relapse.NewAnyName()
			} else {
				panic("todo")
			}
		} else {
			panic("todo")
		}
	}
	panic("todo")
}

func NameToFunc(n *relapse.NameExpr) funcs.Bool {
	typ := n.GetValue()
	switch v := typ.(type) {
	case *relapse.Name:
		if v.DoubleValue != nil {
			return funcs.DoubleEq(funcs.DoubleVar(), funcs.DoubleConst(v.GetDoubleValue()))
		} else if v.IntValue != nil {
			return funcs.IntEq(funcs.IntVar(), funcs.IntConst(v.GetIntValue()))
		} else if v.UintValue != nil {
			return funcs.UintEq(funcs.UintVar(), funcs.UintConst(v.GetUintValue()))
		} else if v.BoolValue != nil {
			return funcs.BoolEq(funcs.BoolVar(), funcs.BoolConst(v.GetBoolValue()))
		} else if v.StringValue != nil {
			return funcs.StringEq(funcs.StringVar(), funcs.StringConst(v.GetStringValue()))
		} else if v.BytesValue != nil {
			return funcs.BytesEq(funcs.BytesVar(), funcs.BytesConst(v.GetBytesValue()))
		}
		panic(fmt.Sprintf("unknown name expr name %#v", v))
	case *relapse.AnyName:
		return funcs.BoolConst(true)
	case *relapse.AnyNameExcept:
		return funcs.Not(NameToFunc(v.GetExcept()))
	case *relapse.NameChoice:
		return funcs.Or(NameToFunc(v.GetLeft()), NameToFunc(v.GetRight()))
	}
	panic(fmt.Sprintf("unknown name expr typ %T", typ))
}
