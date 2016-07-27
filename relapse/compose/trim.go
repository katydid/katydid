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
	"fmt"
	"github.com/katydid/katydid/relapse/funcs"
	"reflect"
)

type errNotConst struct {
	f     interface{}
	field reflect.Value
}

func (this *errNotConst) Error() string {
	return fmt.Sprintf("relapse/compose: %s has constant %s which has a variable parameter", reflect.ValueOf(this.f).Elem().Type().Name(), this.field.Elem().Type())
}

//TrimBool returns a simplified version of the function work that can be done at compile time has been completed.
//This means that any part of the function, that does not contain any variables, has been evaluated and replaced with a constant.
func TrimBool(f funcs.Bool) (funcs.Bool, error) {
	trimmed, err := trim(f)
	if err != nil {
		return nil, err
	}
	return trimmed.(funcs.Bool), nil
}

func trimList(f interface{}) (interface{}, error) {
	this := reflect.ValueOf(f).Elem()
	list := this.Field(0)
	newList := reflect.MakeSlice(list.Type(), list.Len(), list.Len())
	trimable := true
	for i := 0; i < list.Len(); i++ {
		trimmed, err := trim(list.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		trimmedValue := reflect.ValueOf(trimmed)
		newList.Index(i).Set(trimmedValue)
		if !trimmedValue.Type().Implements(constTyp) {
			trimable = false
		}
	}
	this.Field(0).Set(newList)
	if !trimable {
		return f, nil
	}
	return funcs.NewConst(reflect.ValueOf(f).MethodByName("Eval").Call(nil)[0].Interface()), nil
}

func trim(f interface{}) (interface{}, error) {
	if reflect.TypeOf(f).Implements(varTyp) {
		return f, nil
	}
	if reflect.TypeOf(f).Implements(constTyp) {
		return f, nil
	}
	this := reflect.ValueOf(f).Elem()
	trimable := true
	for i := 0; i < this.NumField(); i++ {
		var trimmed interface{}
		var err error
		if this.Field(i).Type().Kind() == reflect.Slice {
			if _, ok := this.Field(i).Type().Elem().MethodByName("Eval"); ok {
				trimable = false
			}
		}
		if _, ok := this.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		if this.Field(i).Elem().Type().Implements(varTyp) {
			if funcs.IsConst(this.Field(i).Type()) {
				return nil, &errNotConst{f, this.Field(i)}
			}
			trimable = false
			continue
		}
		if this.Field(i).Elem().Type().Implements(listOfTyp) {
			trimmed, err = trimList(this.Field(i).Interface())
		} else {
			trimmed, err = trim(this.Field(i).Interface())
		}
		if err != nil {
			return nil, err
		}
		this.Field(i).Set(reflect.ValueOf(trimmed))
		if !this.Field(i).Elem().Type().Implements(constTyp) {
			if funcs.IsConst(this.Field(i).Type()) {
				return nil, &errNotConst{f, this.Field(i)}
			}
			trimable = false
		}
	}
	if !trimable {
		return f, nil
	}
	if inits, ok := f.(funcs.Init); ok {
		err := inits.Init()
		if err != nil {
			return nil, err
		}
	}
	return funcs.NewConst(reflect.ValueOf(f).MethodByName("Eval").Call(nil)[0].Interface()), nil
}
