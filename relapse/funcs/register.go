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

	"github.com/katydid/katydid/relapse/types"
)

//Register registers a function as function that can composed.
func Register(uniqName string, name string, fnc interface{}) {
	typ := reflect.TypeOf(fnc)
	eval, ok := typ.Out(0).MethodByName("Eval")
	if !ok {
		panic("return type has no eval method")
	}
	returnType := eval.Type
	ins := typ.NumIn()
	res := &funk{
		name:     name,
		uniqName: uniqName,
		Out:      types.FromGo(returnType.Out(0)),
		newfnc:   fnc,
	}
	for i := 0; i < ins; i++ {
		meth, ok := typ.In(i).MethodByName("Eval")
		if !ok {
			continue
		}
		res.InConst = append(res.InConst, IsConst(typ.In(i)))
		res.In = append(res.In, types.FromGo(meth.Type.Out(0)))
	}
	funcsMap.register(res)
}

//IsConst returns whether a reflected type is a function that is actually a constant value.
func IsConst(typ reflect.Type) bool {
	switch typ {
	case typConstDouble:
	case typConstInt:
	case typConstUint:
	case typConstBool:
	case typConstString:
	case typConstBytes:
	case typConstDoubles:
	case typConstInts:
	case typConstUints:
	case typConstBools:
	case typConstStrings:
	case typConstListOfBytes:
	default:
		return false
	}
	return true
}

//Which returns the unique name of the function given the function name and parameter types.
func Which(name string, ins ...types.Type) (string, error) {
	return funcsMap.which(name, ins...)
}

//Out returns the output type given the unique function name.
func Out(uniq string) (types.Type, error) {
	u, ok := funcsMap.uniqToFunc[uniq]
	if !ok {
		return 0, &errUnknownFunction{uniq, nil}
	}
	return u.Out, nil
}

type errUnknownFunction struct {
	f   string
	ins []string
}

func newErrUnknownFunction(name string, ins []types.Type) error {
	inss := make([]string, len(ins))
	for i, in := range ins {
		inss[i] = in.String()
	}
	return &errUnknownFunction{name, inss}
}

func (this *errUnknownFunction) Error() string {
	return "relapse/funcs: unknown function: " + this.f + "(" + strings.Join(this.ins, ", ") + ")"
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

func (this *funksMap) which(name string, ins ...types.Type) (string, error) {
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
	In       []types.Type
	InConst  []bool
	Out      types.Type
	newfnc   interface{}
}

func (this *funk) String() string {
	ins := make([]string, len(this.In))
	for i, in := range this.In {
		ins[i] = in.String()
	}
	return fmt.Sprintf("func %v as %v(%v) %v", this.uniqName, this.name, strings.Join(ins, ","), this.Out.String())
}

func newFunc(uniq string, values ...interface{}) (interface{}, error) {
	f, ok := funcsMap.uniqToFunc[uniq]
	if !ok {
		return nil, &errUnknownFunction{uniq, nil}
	}
	newf := reflect.ValueOf(f.newfnc)
	rvalues := make([]reflect.Value, len(values))
	for i := range rvalues {
		rvalues[i] = reflect.ValueOf(values[i])
	}
	res := newf.Call(rvalues)
	return res[0].Interface(), nil
}
