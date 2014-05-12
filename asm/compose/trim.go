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

package compose

import (
	"github.com/awalterschulze/katydid/funcs"
	"reflect"
)

var (
	varTyp   = reflect.TypeOf((*funcs.Variable)(nil)).Elem()
	constTyp = reflect.TypeOf((*funcs.Const)(nil)).Elem()
)

func TrimBool(f funcs.Bool) funcs.Bool {
	return trim(f).(funcs.Bool)
}

func trim(f interface{}) interface{} {
	if reflect.TypeOf(f).Implements(varTyp) {
		return f
	}
	if reflect.TypeOf(f).Implements(constTyp) {
		return f
	}
	this := reflect.ValueOf(f).Elem()
	trimable := true
	for i := 0; i < this.NumField(); i++ {
		if _, ok := this.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		if this.Field(i).Elem().Type().Implements(varTyp) {
			trimable = false
			continue
		}
		this.Field(i).Set(reflect.ValueOf(trim(this.Field(i).Interface())))
		if !this.Field(i).Elem().Type().Implements(constTyp) {
			trimable = false
		}
	}
	if !trimable {
		return f
	}
	return funcs.NewConst(reflect.ValueOf(f).MethodByName("Eval").Call(nil)[0].Interface())
}
