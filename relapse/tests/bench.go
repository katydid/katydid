//  Copyright 2016 Walter Schulze
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

package tests

import (
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/protonum"
	"reflect"
	"sort"
	"strings"
)

type BenchValidator struct {
	Name        string
	CodecName   string
	Grammar     *ast.Grammar
	PackageName string
	MessageName string
}

func (this *BenchValidator) Record() bool {
	return !strings.Contains(this.CodecName, "xml")
}

var BenchValidators = make(map[string]map[string]BenchValidator)

func BenchValidateProtoNum(name string, grammar combinator.G, m interface{}) {
	packageName := "tests"
	messageName := reflect.TypeOf(m).Elem().Name()
	g, err := protonum.FieldNamesToNumbers(packageName, messageName, m.(ProtoMessage).Description(), grammar.Grammar())
	if err != nil {
		panic(name + ": " + err.Error())
	}
	BenchValidate("ProtoNum"+name, combinator.G(ast.NewRefLookup(g)), []string{"protoNum"}, packageName, messageName)
}

type benchValidatorList []interface{}

func (this benchValidatorList) Less(i, j int) bool {
	if this[i].(BenchValidator).Name < this[j].(BenchValidator).Name {
		return true
	}
	if this[i].(BenchValidator).Name == this[j].(BenchValidator).Name {
		return this[i].(BenchValidator).CodecName < this[j].(BenchValidator).CodecName
	}
	return false
}

func (this benchValidatorList) Len() int {
	return len(this)
}

func (this benchValidatorList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func BenchValidatorList() []interface{} {
	vs := make(benchValidatorList, 0, len(BenchValidators)*3)
	for name, cs := range BenchValidators {
		for cname, _ := range cs {
			vs = append(vs, BenchValidators[name][cname])
		}
	}
	sort.Sort(vs)
	return vs
}

func BenchValidate(name string, grammar combinator.G, codecNames []string, packageName, messageName string) {
	if _, ok := BenchValidators[name]; !ok {
		BenchValidators[name] = make(map[string]BenchValidator)
	}
	for _, codecName := range codecNames {
		BenchValidators[name][codecName] = BenchValidator{
			Name:        name,
			CodecName:   codecName,
			Grammar:     grammar.Grammar(),
			PackageName: packageName,
			MessageName: messageName,
		}
	}
}
