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

package funcs

import (
	"reflect"
	"strings"
)

//Sprint returns the string printout of the function.
func Sprint(i interface{}) string {
	e := reflect.ValueOf(i).Elem()
	uniqName := e.Type().Name()
	name := reverse(uniqName)
	numFields := e.NumField()
	if numFields == 2 {
		p1 := e.Field(0).Interface()
		p2 := e.Field(1).Interface()
		v, ok := isVarConst(p1, p2)
		if ok {
			if name == "eq" {
				return "== " + v
			}
			if name == "ne" {
				return "!= " + v
			}
			if name == "gt" {
				return "> " + v
			}
			if name == "lt" {
				return "< " + v
			}
			if name == "ge" {
				return ">= " + v
			}
			if name == "le" {
				return "<= " + v
			}
		}
	}
	return "->" + sprint(i)
}

func sprint(i interface{}) string {
	e := reflect.ValueOf(i).Elem()
	uniqName := e.Type().Name()
	name := reverse(uniqName)
	if len(name) == 0 {
		strer, ok := i.(stringer)
		if !ok {
			panic("unknown function without String() string method")
		}
		name = strer.String()
	}
	switch i.(type) {
	case Const:
		return name
	case Setter:
		return name
	case ListOf:
		return name
	}
	numFields := e.NumField()
	ss := make([]string, 0, numFields)
	for i := 0; i < numFields; i++ {
		if _, ok := e.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		ss = append(ss, sprint(e.Field(i).Interface()))
	}
	return name + "(" + strings.Join(ss, ",") + ")"
}

func reverse(uniq string) (name string) {
	f, ok := funcsMap.uniqToFunc[uniq]
	if !ok {
		return ""
	}
	return f.name
}

type stringer interface {
	String() string
}

func isVarConst(p1, p2 interface{}) (string, bool) {
	switch p1.(type) {
	case *varBool, *varBytes, *varDouble, *varInt, *varString, *varUint:
		switch p2.(type) {
		case *constBool, *constBytes, *constDouble, *constInt, *constString, *constUint:
			return sprint(p2), true
		}
	}
	switch p2.(type) {
	case *varBool, *varBytes, *varDouble, *varInt, *varString, *varUint:
		switch p1.(type) {
		case *constBool, *constBytes, *constDouble, *constInt, *constString, *constUint:
			return sprint(p1), true
		}
	}
	return "", false
}
