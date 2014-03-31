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
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"fmt"
	"reflect"
	"strings"
)

type errUnknownFunction struct {
	f   string
	ins []string
}

func newErrUnknownFunction(name string, ins []descriptor.FieldDescriptorProto_Type) error {
	inss := make([]string, len(ins))
	for i, in := range ins {
		inss[i] = in.String()
	}
	return &errUnknownFunction{name, inss}
}

func (this *errUnknownFunction) Error() string {
	return "unknown function: " + this.f + "(" + strings.Join(this.ins, ", ") + ")"
}

var funcsMap = newFunksMap()

type funksMap struct {
	nameToUniq map[string][]string
	uniqToFunc map[string]*funk
}

func newFunksMap() *funksMap {
	return &funksMap{
		nameToUniq: make(map[string][]string),
		uniqToFunc: make(map[string]*funk),
	}
}

func (this *funksMap) register(f *funk) {
	if _, ok := this.uniqToFunc[f.uniqName]; ok {
		panic("func already registered " + f.uniqName)
	}
	if this.nameToUniq[f.name] == nil {
		this.nameToUniq[f.name] = make([]string, 0, 1)
	}
	this.nameToUniq[f.name] = append(this.nameToUniq[f.name], f.uniqName)
	this.uniqToFunc[f.uniqName] = f
}

func (this *funksMap) which(name string, ins ...descriptor.FieldDescriptorProto_Type) (string, error) {
	uniqs, ok := this.nameToUniq[name]
	if !ok {
		return "", newErrUnknownFunction(name, ins)
	}
	for _, uniq := range uniqs {
		u := this.uniqToFunc[uniq]
		if len(u.In) != len(ins) {
			continue
		}
		eq := true
		for i := range u.In {
			if u.In[i] != ins[i] {
				eq = false
				break
			}
		}
		if !eq {
			continue
		}
		return u.uniqName, nil
	}
	panic(newErrUnknownFunction(name, ins))
	return "", newErrUnknownFunction(name, ins)
}

func (this *funksMap) String() string {
	funcs := make([]string, len(this.uniqToFunc))
	i := 0
	for _, f := range this.uniqToFunc {
		funcs[i] = f.String()
		i++
	}
	return strings.Join(funcs, "\n")
}

type funk struct {
	name     string
	uniqName string
	In       []descriptor.FieldDescriptorProto_Type
	Out      descriptor.FieldDescriptorProto_Type
	newfnc   func() interface{}
}

func (this *funk) String() string {
	ins := make([]string, len(this.In))
	for i, in := range this.In {
		ins[i] = in.String()
	}
	return fmt.Sprintf("func %v as %v(%v) %v", this.uniqName, this.name, strings.Join(ins, ","), this.Out.String())
}

func goTypeToProto(typ reflect.Type) descriptor.FieldDescriptorProto_Type {
	kind := typ.Kind()
	switch kind {
	case reflect.Float64:
		return descriptor.FieldDescriptorProto_TYPE_DOUBLE
	case reflect.Float32:
		return descriptor.FieldDescriptorProto_TYPE_FLOAT
	case reflect.Int64:
		return descriptor.FieldDescriptorProto_TYPE_INT64
	case reflect.Uint64:
		return descriptor.FieldDescriptorProto_TYPE_UINT64
	case reflect.Int32:
		return descriptor.FieldDescriptorProto_TYPE_INT32
	case reflect.Uint32:
		return descriptor.FieldDescriptorProto_TYPE_UINT32
	case reflect.Bool:
		return descriptor.FieldDescriptorProto_TYPE_BOOL
	case reflect.String:
		return descriptor.FieldDescriptorProto_TYPE_STRING
	case reflect.Slice:
		if typ.Elem().Kind() == reflect.Int8 {
			return descriptor.FieldDescriptorProto_TYPE_BYTES
		}
	}
	panic(fmt.Sprintf("go Type %v unsupported", typ))
}

func Register(name string, fnc interface{}) {
	RegisterFactory(name, func() interface{} {
		return reflect.New(reflect.TypeOf(fnc).Elem()).Interface()
	})
}

func RegisterFactory(name string, newFunc func() interface{}) {
	rfunc := reflect.ValueOf(newFunc())
	returnType := rfunc.MethodByName("Eval").Type()
	uniqName := rfunc.Elem().Type().Name()
	lenFields := rfunc.Elem().NumField()
	res := &funk{
		name:     name,
		uniqName: uniqName,
		Out:      goTypeToProto(returnType.Out(0)),
		newfnc:   newFunc,
	}
	for i := 0; i < lenFields; i++ {
		meth, ok := rfunc.Elem().Field(i).Type().MethodByName("Eval")
		if !ok {
			continue
		}
		res.In = append(res.In, goTypeToProto(meth.Type.Out(0)))
	}
	funcsMap.register(res)
}

func newFunc(uniq string, values ...interface{}) (interface{}, error) {
	f, ok := funcsMap.uniqToFunc[uniq]
	if !ok {
		return nil, &errUnknownFunction{uniq, nil}
	}
	newf := reflect.ValueOf(f.newfnc()).Elem()
	for i := 0; i < newf.NumField(); i++ {
		if _, ok := newf.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		newf.Field(i).Set(reflect.ValueOf(values[i]))
	}
	return newf.Addr().Interface(), nil
}

func NewBoolFunc(uniq string, values ...interface{}) (Bool, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Bool), nil
}

func NewInt64Func(uniq string, values ...interface{}) (Int64, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Int64), nil
}

func NewStringFunc(uniq string, values ...interface{}) (String, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(String), nil
}

func Which(name string, ins ...descriptor.FieldDescriptorProto_Type) (string, error) {
	return funcsMap.which(name, ins...)
}

func Out(uniq string) (descriptor.FieldDescriptorProto_Type, error) {
	u, ok := funcsMap.uniqToFunc[uniq]
	if !ok {
		return 0, &errUnknownFunction{uniq, nil}
	}
	return u.Out, nil
}
