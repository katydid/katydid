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

package trim

import (
	"github.com/awalterschulze/katydid/funcs"
	"reflect"
)

var (
	varTyp = reflect.TypeOf((*funcs.Variable)(nil)).Elem()
)

func Bool(f funcs.Bool) funcs.Bool {
	return trim(f).(funcs.Bool)
}

func trim(f interface{}) interface{} {
	if trimable(f) {
		return valToFunc(reflect.ValueOf(f).MethodByName("Eval").Call(nil)[0].Interface())
	}
	e := reflect.ValueOf(f).Elem()
	for i := 0; i < e.NumField(); i++ {
		if _, ok := e.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		field := e.Field(i).Interface()
		e.Field(i).Set(reflect.ValueOf(trim(field)))
	}
	return f
}

func valToFunc(value interface{}) interface{} {
	switch v := value.(type) {
	case float64:
		return funcs.NewConstFloat64(v)
	case float32:
		return funcs.NewConstFloat32(v)
	case int64:
		return funcs.NewConstInt64(v)
	case uint64:
		return funcs.NewConstUint64(v)
	case int32:
		return funcs.NewConstInt32(v)
	case bool:
		return funcs.NewConstBool(v)
	case string:
		return funcs.NewConstString(v)
	case []byte:
		return funcs.NewConstBytes(v)
	case uint32:
		return funcs.NewConstUint32(v)
	case []float64:
		return funcs.NewConstFloat64s(v)
	case []float32:
		return funcs.NewConstFloat32s(v)
	case []int64:
		return funcs.NewConstInt64s(v)
	case []uint64:
		return funcs.NewConstUint64s(v)
	case []int32:
		return funcs.NewConstInt32s(v)
	case []bool:
		return funcs.NewConstBools(v)
	case []string:
		return funcs.NewConstStrings(v)
	case [][]byte:
		return funcs.NewConstListOfBytes(v)
	case []uint32:
		return funcs.NewConstUint32s(v)
	}
	panic("unreachable")
}

func trimable(f interface{}) bool {
	if reflect.ValueOf(f).Type().Implements(varTyp) {
		return false
	}
	e := reflect.ValueOf(f).Elem()
	for i := 0; i < e.NumField(); i++ {
		if _, ok := e.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		if !trimable(e.Field(i).Interface()) {
			return false
		}
	}
	return true
}
