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
	fnc      interface{}
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
	rfunc := reflect.ValueOf(fnc)
	returnType := rfunc.MethodByName("Eval").Type()
	uniqName := rfunc.Elem().Type().Name()
	lenFields := rfunc.Elem().NumField()
	res := &funk{
		name:     name,
		uniqName: uniqName,
		Out:      goTypeToProto(returnType.Out(0)),
		fnc:      fnc,
	}
	for i := 0; i < lenFields; i++ {
		meth, ok := rfunc.Elem().Field(i).Type().MethodByName("Eval")
		if !ok {
			panic("no Eval method")
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
	newf := reflect.New(reflect.TypeOf(f.fnc).Elem()).Elem()
	for i := 0; i < newf.NumField(); i++ {
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
