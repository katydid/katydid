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
