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

func Sprint(i interface{}) string {
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
	numFields := e.NumField()
	ss := make([]string, 0, numFields)
	for i := 0; i < numFields; i++ {
		if _, ok := e.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		ss = append(ss, Sprint(e.Field(i).Interface()))
	}
	if len(ss) == 0 {
		return name
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
		return Not(v1)
	}
	return f
}
