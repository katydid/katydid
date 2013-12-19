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

package readable

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"code.google.com/p/go.text/unicode/norm"
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"github.com/awalterschulze/katydid/exp/readable/ast"
)

type Interpreter struct {
	desc     *descriptor.FileDescriptorSet
	rules    *ast.Rules
	transMap map[string]map[string]string
}

func NewInterpreter(desc *descriptor.FileDescriptorSet, rules *ast.Rules) (*Interpreter, error) {
	if rules.Root == nil {
		return nil, errors.New("No Root Rule")
	}
	transMap := make(map[string]map[string]string)
	for _, t := range rules.GetTransition() {
		if _, ok := transMap[t.GetSrc()]; !ok {
			transMap[t.GetSrc()] = make(map[string]string)
		}
		transMap[t.GetSrc()][t.GetInput()] = t.GetDst()
	}
	return &Interpreter{desc, rules, transMap}, nil
}

func (this *Interpreter) Match(buf []byte) (bool, error) {
	currentState, err := this.interpret(this.rules.Root.GetState(), this.rules.Root.GetPackage(), this.rules.Root.GetMessage(), buf)
	fmt.Printf("final state = %v %v\n", currentState, currentState == "accept")
	return (currentState == "accept"), err
}

func (this *Interpreter) transition(src, input string) (dst string) {
	if d, ok := this.transMap[src][input]; ok {
		return d
	} else if d2, ok := this.transMap[src]["_"]; ok {
		return d2
	}
	panic("no transition specified")
}

func (this *Interpreter) interpret(currentState string, packageName string, messageName string, buf []byte) (string, error) {
	fmt.Printf("interpreting %v %v %v\n", currentState, packageName, messageName)
	l := len(buf)
	offset := 0
	message := this.desc.GetMessage(packageName, messageName)
	if message == nil {
		panic("unknown message: " + packageName + "." + messageName)
	}
	for offset < l {
		for _, r := range this.rules.GetInit() {
			var field *descriptor.FieldDescriptorProto
			for _, f := range message.GetField() {
				if !f.IsMessage() {
					continue
				}
				fmt.Printf("TypeName = %v\n", f.GetTypeName())
				if f.GetTypeName() == ("." + r.GetPackage() + "." + r.GetMessage()) {
					field = f
					break
				}
			}
			if field == nil {
				continue
			}
			key := field.GetKey()
			thisKey := true
			for i := range key {
				if buf[offset+i] != key[i] {
					thisKey = false
					break
				}
			}
			if !thisKey {
				continue
			}
			fmt.Printf("applying %v\n", r)
			mPackageName, mName := this.desc.FindMessage(packageName, messageName, field.GetName())
			if len(mPackageName) == 0 {
				panic("unknown message " + packageName + "-" + messageName + "-" + field.GetName())
			}
			length, n, err := decodeVarint(buf[offset+len(key):])
			if err != nil {
				return "", err
			}
			startOff := offset + len(key) + n
			endOff := startOff + int(length)
			newState, err := this.interpret(r.GetState(), mPackageName, mName, buf[startOff:endOff])
			if err != nil {
				return "", err
			}
			fmt.Printf("transitioning from: %v with input %v \n", currentState, newState)
			currentState = this.transition(currentState, newState)
			fmt.Printf("transitioned to: %v \n", currentState)
		}
		for _, ifexpr := range this.rules.GetIfExpr() {
			v, err := ifexpr.GetTerminalVariable()
			if err != nil {
				return "", err
			}
			if v.GetPackage() != packageName || v.GetMessage() != messageName {
				continue
			}
			field := message.GetFieldDescriptor(v.GetField())
			key := field.GetKey()
			thisKey := true
			for i := range key {
				if buf[offset+i] != key[i] {
					thisKey = false
					break
				}
			}
			if !thisKey {
				continue
			}
			fmt.Printf("ifexpr %v\n", ifexpr)
			ii, err := next(buf[offset:])
			if err != nil {
				return "", err
			}
			newState, err := this.ifExpr(ifexpr, buf[offset+len(key):offset+ii+len(key)])
			if err != nil {
				return "", err
			}
			fmt.Printf("transitioning from: %v with input %v \n", currentState, newState)
			currentState = this.transition(currentState, newState)
			fmt.Printf("transitioned to: %v \n", currentState)
		}
		i, err := next(buf[offset:])
		if err != nil {
			return "", err
		}
		offset += i
	}
	if offset > l {
		return "", errors.New("protobuf underflow")
	}
	return currentState, nil
}

func (this *Interpreter) ifExpr(ifExpr *ast.IfExpr, buf []byte) (string, error) {
	s, err := evalIf(this.desc, ifExpr, buf)
	return string(s), err
}

type state string

func evalIf(desc *descriptor.FileDescriptorSet, ifExpr *ast.IfExpr, buf []byte) (state, error) {
	if b, err := evalBoolExpr(desc, ifExpr.GetCondition(), buf); err != nil {
		return "", err
	} else if b {
		return evalStateExpr(desc, ifExpr.GetThen(), buf)
	} else {
		return evalStateExpr(desc, ifExpr.GetElse(), buf)
	}
}

func evalStateExpr(desc *descriptor.FileDescriptorSet, expr *ast.StateExpr, buf []byte) (state, error) {
	if expr.State != nil {
		return state(expr.GetState()), nil
	}
	return evalIf(desc, expr.GetIfExpr(), buf)
}

type kind int

var boolKind kind = 2
var int64Kind kind = 3
var uint64Kind kind = 4
var stringKind kind = 5

func getExprKind(desc *descriptor.FileDescriptorSet, expr *ast.Expr) kind {
	if expr.Terminal != nil {
		return getTerminalKind(desc, expr.GetTerminal())
	}
	return getFunctionKind(desc, expr.GetFunction())
}

func getTerminalKind(desc *descriptor.FileDescriptorSet, term *ast.Terminal) kind {
	if term.BoolValue != nil {
		return boolKind
	}
	if term.Int64Value != nil {
		return int64Kind
	}
	if term.Uint64Value != nil {
		return uint64Kind
	}
	if term.StringValue != nil {
		return stringKind
	}
	field := desc.GetField(term.GetVariable().GetPackage(), term.GetVariable().GetMessage(), term.GetVariable().GetField())
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return boolKind
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return int64Kind
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return uint64Kind
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return stringKind
	}
	panic("unknown kind")
}

func getFunctionKind(desc *descriptor.FileDescriptorSet, fn *ast.Function) kind {
	switch fn.GetName() {
	case "eq":
		return boolKind
	case "contains":
		return boolKind
	case "decString":
		return stringKind
	case "decInt64":
		return int64Kind
	case "nfkc":
		return stringKind
	}
	panic("unknown kind")
}

func evalBoolFunction(desc *descriptor.FileDescriptorSet, fn *ast.Function, buf []byte) (bool, error) {
	switch fn.GetName() {
	case "eq":
		if len(fn.GetParams()) != 2 {
			return false, errors.New("wrong number of parameters")
		}
		typ := getExprKind(desc, fn.GetParams()[1])
		switch typ {
		case int64Kind:
			i1, err := evalInt64Expr(desc, fn.GetParams()[0], buf)
			if err != nil {
				return false, err
			}
			i2, err := evalInt64Expr(desc, fn.GetParams()[1], buf)
			if err != nil {
				return false, err
			}
			return i1 == i2, nil
		case stringKind:
			s1, err := evalStringExpr(desc, fn.GetParams()[0], buf)
			if err != nil {
				return false, err
			}
			s2, err := evalStringExpr(desc, fn.GetParams()[1], buf)
			if err != nil {
				return false, err
			}
			return s1 == s2, nil
		}
	case "ge":
		if len(fn.GetParams()) != 2 {
			return false, errors.New("wrong number of parameters")
		}
		typ := getExprKind(desc, fn.GetParams()[1])
		switch typ {
		case int64Kind:
			i1, err := evalInt64Expr(desc, fn.GetParams()[0], buf)
			if err != nil {
				return false, err
			}
			i2, err := evalInt64Expr(desc, fn.GetParams()[1], buf)
			if err != nil {
				return false, err
			}
			return i1 >= i2, nil
		}
	case "contains":
		if len(fn.GetParams()) != 2 {
			return false, errors.New("wrong number of parameters")
		}
		s1, err := evalStringExpr(desc, fn.GetParams()[0], buf)
		if err != nil {
			return false, err
		}
		s2, err := evalStringExpr(desc, fn.GetParams()[1], buf)
		if err != nil {
			return false, err
		}
		fmt.Printf("strings.Contains(%v, %v)\n", s1, s2)
		return strings.Contains(s1, s2), nil
	case "not":
		if len(fn.GetParams()) != 1 {
			return false, errors.New("wrong number of parameters")
		}
		b1, err := evalBoolExpr(desc, fn.GetParams()[0], buf)
		if err != nil {
			return false, err
		}
		return !b1, nil
	}
	return false, errors.New("unknown function " + fn.GetName())
}

func evalInt64Function(desc *descriptor.FileDescriptorSet, fn *ast.Function, buf []byte) (int64, error) {
	switch fn.GetName() {
	case "decInt64":
		if len(fn.GetParams()) != 1 {
			return 0, errors.New("wrong number of parameters")
		}
		buffy, err := evalBytesTerminal(desc, fn.GetParams()[0].GetTerminal(), buf, descriptor.FieldDescriptorProto_TYPE_INT64)
		if err != nil {
			return 0, err
		}
		v, _, err := decodeVarint(buffy)
		fmt.Printf("decoded int64 %v\n", int64(v))
		return int64(v), err
	}
	panic("not implemented")
}

func evalStringFunction(desc *descriptor.FileDescriptorSet, fn *ast.Function, buf []byte) (string, error) {
	switch fn.GetName() {
	case "nfkc":
		if len(fn.GetParams()) != 1 {
			return "", errors.New("wrong number of parameters")
		}
		s1, err := evalStringExpr(desc, fn.GetParams()[0], buf)
		if err != nil {
			return "", err
		}
		s2 := norm.NFKC.String(s1)
		return s2, nil
	case "decString":
		if len(fn.GetParams()) != 1 {
			return "", errors.New("wrong number of parameters")
		}
		buffy, err := evalBytesTerminal(desc, fn.GetParams()[0].GetTerminal(), buf, descriptor.FieldDescriptorProto_TYPE_STRING)
		if err != nil {
			return "", err
		}
		fmt.Printf("decoded string %v\n", string(buffy))
		return string(buffy), nil
	}
	return "", errors.New("unkonwn function " + fn.GetName())
}

func evalInt64Expr(desc *descriptor.FileDescriptorSet, expr *ast.Expr, buf []byte) (int64, error) {
	if expr.Terminal != nil {
		return evalInt64Terminal(desc, expr.GetTerminal(), buf)
	}
	return evalInt64Function(desc, expr.GetFunction(), buf)
}

func evalStringExpr(desc *descriptor.FileDescriptorSet, expr *ast.Expr, buf []byte) (string, error) {
	if expr.Terminal != nil {
		return evalStringTerminal(desc, expr.GetTerminal(), buf)
	}
	return evalStringFunction(desc, expr.GetFunction(), buf)
}

func evalBoolExpr(desc *descriptor.FileDescriptorSet, expr *ast.Expr, buf []byte) (bool, error) {
	if expr.Terminal != nil {
		return evalBoolTerminal(desc, expr.GetTerminal(), buf)
	}
	return evalBoolFunction(desc, expr.GetFunction(), buf)
}

func evalInt64Terminal(desc *descriptor.FileDescriptorSet, expr *ast.Terminal, buf []byte) (int64, error) {
	if expr.Int64Value != nil {
		return expr.GetInt64Value(), nil
	}
	return 0, errors.New("Not an int64 type")
}

func evalBytesTerminal(desc *descriptor.FileDescriptorSet, expr *ast.Terminal, buf []byte, typ descriptor.FieldDescriptorProto_Type) ([]byte, error) {
	if expr.Variable != nil {
		field := desc.GetField(expr.GetVariable().GetPackage(), expr.GetVariable().GetMessage(), expr.GetVariable().GetField())
		if field == nil {
			return nil, errors.New("could not find field")
		}
		switch expr.GetVariable().GetPart() {
		case "value":
			if field.GetType() != typ {
				return nil, errors.New("not a " + typ.String())
			}
			switch typ {
			case descriptor.FieldDescriptorProto_TYPE_INT64:
				n, err := endOfVarint(buf)
				if err != nil {
					return nil, err
				}
				return buf[:n], nil
			case descriptor.FieldDescriptorProto_TYPE_STRING:
				v, n, err := decodeVarint(buf)
				if err != nil {
					return nil, err
				}
				return buf[n : n+int(v)], nil
			}
		case "len":
			n, err := endOfVarint(buf)
			if err != nil {
				return nil, err
			}
			return buf[:n], nil
		default:
			panic("invalid part " + expr.GetVariable().GetPart())
		}
	}
	return nil, errors.New("Not a bytes type")
}

func evalStringTerminal(desc *descriptor.FileDescriptorSet, expr *ast.Terminal, buf []byte) (string, error) {
	if expr.StringValue != nil {
		return strconv.Unquote(expr.GetStringValue())
	}
	return "", errors.New("Not a string type")
}

func evalBoolTerminal(desc *descriptor.FileDescriptorSet, expr *ast.Terminal, buf []byte) (bool, error) {
	if expr.BoolValue != nil {
		return expr.GetBoolValue(), nil
	}
	return false, errors.New("Not a bool type")
}
