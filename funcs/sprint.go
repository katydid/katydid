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
	"fmt"
	"reflect"
	"strings"
)

func reverse(uniq string) (name string) {
	for n, us := range funcsMap.nameToUniq {
		for _, u := range us {
			if uniq == u {
				return n
			}
		}
	}
	return ""
}

type stringer interface {
	String() string
}

func isVarConst(p1, p2 interface{}) (string, bool) {
	switch p1.(type) {
	case *varBool, *varBytes, *varDouble, *varInt, *varString, *varUint:
		switch p2.(type) {
		case *constBool, *constBytes, *constDouble, *constInt, *constString, *constUint:
			return Sprint(p2), true
		}
	}
	switch p2.(type) {
	case *varBool, *varBytes, *varDouble, *varInt, *varString, *varUint:
		switch p1.(type) {
		case *constBool, *constBytes, *constDouble, *constInt, *constString, *constUint:
			return Sprint(p1), true
		}
	}
	return "", false
}

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
	return sprint(i)
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
	case Decoder:
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

func Equal(l, r interface{}) bool {
	le := reflect.ValueOf(l).Elem()
	lUniqName := le.Type().Name()
	re := reflect.ValueOf(r).Elem()
	rUniqName := re.Type().Name()
	if lUniqName != rUniqName {
		return false
	}
	if len(reverse(lUniqName)) == 0 {
		//TODO maybe this could be done better or we could just always convert functions to strings to compare
		return Sprint(l) == Sprint(r)
	}
	numFields := le.NumField()
	for i := 0; i < numFields; i++ {
		if _, ok := le.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		if !Equal(le.Field(i).Interface(), re.Field(i).Interface()) {
			return false
		}
	}
	if Sprint(l) != Sprint(r) {
		fmt.Printf("wtf %v == %v\n", Sprint(l), Sprint(r))
		panic("two non equal functions are equal")
	}
	return true
}

func Simplify(f Bool) Bool {
	if ff, ok := f.(*and); ok {
		v1 := Simplify(ff.V1)
		v2 := Simplify(ff.V2)
		if l, ok := v1.(*constBool); ok {
			if l.v == false {
				return BoolConst(false)
			} else {
				return v2
			}
		}
		if r, ok := v2.(*constBool); ok {
			if r.v == false {
				return BoolConst(false)
			} else {
				return v1
			}
		}
		if Equal(v1, v2) {
			return v1
		}
		if l, ok := v1.(*not); ok {
			if Equal(l.V1, v2) {
				return BoolConst(false)
			}
		}
		if r, ok := v2.(*not); ok {
			if Equal(v1, r.V1) {
				return BoolConst(false)
			}
		}
		switch vv1 := v1.(type) {
		case *stringEq:
			if vv2, ok := v2.(*stringEq); ok {
				if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
					if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
						if vvv1 != vvv2 {
							return BoolConst(false)
						}
					}
				}
			}
			if vv2, ok := v2.(*stringNe); ok {
				if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
					if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
						if vvv1 != vvv2 {
							return v1
						} else {
							return BoolConst(false)
						}
					}
				}
			}
		case *stringNe:
			if vv2, ok := v2.(*stringEq); ok {
				if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
					if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
						if vvv1 != vvv2 {
							return v2
						} else {
							return BoolConst(false)
						}
					}
				}
			}
		case *intEq:
			if vv2, ok := v2.(*intEq); ok {
				if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
					if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
						if vvv1 != vvv2 {
							return BoolConst(false)
						}
					}
				}
			}
			if vv2, ok := v2.(*intNe); ok {
				if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
					if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
						if vvv1 != vvv2 {
							return v1
						} else {
							return BoolConst(false)
						}
					}
				}
			}
		case *intNe:
			if vv2, ok := v2.(*intEq); ok {
				if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
					if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
						if vvv1 != vvv2 {
							return v2
						} else {
							return BoolConst(false)
						}
					}
				}
			}
		}
		return And(v1, v2)
	}
	if ff, ok := f.(*or); ok {
		v1 := Simplify(ff.V1)
		v2 := Simplify(ff.V2)
		if l, ok := v1.(*constBool); ok {
			if l.v == true {
				return BoolConst(true)
			} else {
				return v2
			}
		}
		if r, ok := v2.(*constBool); ok {
			if r.v == true {
				return BoolConst(true)
			} else {
				return v1
			}
		}
		if Equal(v1, v2) {
			return v1
		}
		if l, ok := v1.(*not); ok {
			if Equal(l.V1, v2) {
				return BoolConst(true)
			}
		}
		if r, ok := v2.(*not); ok {
			if Equal(v1, r.V1) {
				return BoolConst(true)
			}
		}
		return Or(v1, v2)
	}
	if ff, ok := f.(*not); ok {
		v1 := Simplify(ff.V1)
		if vv, ok := v1.(*not); ok {
			return vv.V1
		}
		if vv, ok := v1.(*constBool); ok {
			return BoolConst(!vv.v)
		}
		if vv, ok := v1.(*and); ok {
			return Simplify(Or(Not(vv.V1), Not(vv.V2)))
		}
		if vv, ok := v1.(*or); ok {
			return Simplify(And(Not(vv.V1), Not(vv.V2)))
		}
		switch vv := v1.(type) {
		case *boolEq:
			return BoolNe(vv.V1, vv.V2)
		case *bytesEq:
			return BytesNe(vv.V1, vv.V2)
		case *doubleEq:
			return DoubleNe(vv.V1, vv.V2)
		case *intEq:
			return IntNe(vv.V1, vv.V2)
		case *stringEq:
			return StringNe(vv.V1, vv.V2)
		case *uintEq:
			return UintNe(vv.V1, vv.V2)
		case *boolNe:
			return BoolEq(vv.V1, vv.V2)
		case *bytesNe:
			return BytesEq(vv.V1, vv.V2)
		case *doubleNe:
			return DoubleEq(vv.V1, vv.V2)
		case *intNe:
			return IntEq(vv.V1, vv.V2)
		case *stringNe:
			return StringEq(vv.V1, vv.V2)
		case *uintNe:
			return UintEq(vv.V1, vv.V2)
		}
		return Not(v1)
	}
	return f
}
