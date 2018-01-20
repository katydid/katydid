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

var errTyp = reflect.TypeOf((*error)(nil)).Elem()
var funcTyp = reflect.TypeOf((*Func)(nil)).Elem()

//Register registers a function as function that can composed.
func Register(name string, fnc interface{}) {
	typ := reflect.TypeOf(fnc)
	if typ.Kind() != reflect.Func {
		panic(fmt.Sprintf("expecting constructor function for %s, but got %T", name, fnc))
	}
	if typ.NumOut() == 2 {
		if !typ.Out(1).Implements(errTyp) {
			panic(fmt.Sprintf("the second return type of the constructor for %s is not an error", name))
		}
	}
	if typ.NumOut() > 2 {
		panic(fmt.Sprintf("the constructor for %s has more than 2 return values", name))
	}
	if typ.NumOut() == 0 {
		panic(fmt.Sprintf("the constructor for %s has no return values", name))
	}
	eval, ok := typ.Out(0).MethodByName("Eval")
	if !ok {
		panic(fmt.Sprintf("the constructor for %s returns a type without an Eval method", name))
	}
	if !typ.Out(0).Implements(funcTyp) {
		panic(fmt.Sprintf("the constructor for %s returns a type that does not implement funcs.Func", name))
	}
	returnType := eval.Type
	ins := typ.NumIn()
	res := &Funk{
		Name:   name,
		Out:    types.FromGo(returnType.Out(0)),
		newfnc: fnc,
	}
	for i := 0; i < ins; i++ {
		meth, ok := typ.In(i).MethodByName("Eval")
		if !ok {
			continue
		}
		if !typ.In(i).Implements(funcTyp) {
			panic(fmt.Sprintf("the constructor for %s has an input parameter (number %d) that does not implement funcs.Func", name, i))
		}
		res.InConst = append(res.InConst, IsConst(typ.In(i)))
		inType := types.FromGo(meth.Type.Out(0))
		res.In = append(res.In, inType)
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

//Which returns the Funk (function creator) of the function given the function name and parameter types.
func Which(name string, ins ...types.Type) (*Funk, error) {
	return funcsMap.which(name, ins...)
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

type funksMap map[string][]*Funk

func newFunksMap() funksMap {
	return make(map[string][]*Funk)
}

func (this funksMap) register(f *Funk) {
	if _, ok := this[f.Name]; !ok {
		this[f.Name] = []*Funk{}
	}
	this[f.Name] = append(this[f.Name], f)
}

func (this funksMap) which(name string, ins ...types.Type) (*Funk, error) {
	funks, ok := this[name]
	if !ok {
		return nil, newErrUnknownFunction(name, ins)
	}
	for _, f := range funks {
		if len(f.In) != len(ins) {
			continue
		}
		eq := true
		for i := range f.In {
			if f.In[i] != ins[i] {
				eq = false
				break
			}
		}
		if !eq {
			continue
		}
		return f, nil
	}
	return nil, newErrUnknownFunction(name, ins)
}

type Funk struct {
	Name    string
	In      []types.Type
	InConst []bool
	Out     types.Type
	newfnc  interface{}
}

func (this *Funk) String() string {
	ins := make([]string, len(this.In))
	for i, in := range this.In {
		ins[i] = in.String()
	}
	return fmt.Sprintf("func %v(%v) %v", this.Name, strings.Join(ins, ","), this.Out.String())
}

func (f *Funk) New(values ...interface{}) (interface{}, error) {
	newf := reflect.ValueOf(f.newfnc)
	rvalues := make([]reflect.Value, len(values))
	for i := range rvalues {
		rvalues[i] = reflect.ValueOf(values[i])
	}
	res := newf.Call(rvalues)
	if len(res) == 2 {
		if !res[1].IsNil() {
			return res[0].Interface(), res[1].Interface().(error)
		}
	}
	return res[0].Interface(), nil
}

//IsFalse returns whether a function is a false constant.
func IsFalse(fn Bool) bool {
	v, ok := fn.(*constBool)
	if !ok {
		return false
	}
	return v.v == false
}

//IsTrue returns whether a function is a true constant.
func IsTrue(fn Bool) bool {
	v, ok := fn.(*constBool)
	if !ok {
		return false
	}
	return v.v == true
}

//Equal returns whether two functions are equal.
func Equal(l, r Comparable) bool {
	hl := l.Hash()
	hr := r.Hash()
	if hl != hr {
		return false
	}
	return l.Compare(r) == 0
}

func isVarConst(a, b interface{}) (string, bool) {
	if c, aok := a.(aConst); aok {
		if _, bok := b.(aVariable); bok {
			return c.String(), true
		}
	} else if c, bok := b.(aConst); bok {
		if _, aok := a.(aVariable); aok {
			return c.String(), true
		}
	}
	return "", false
}
