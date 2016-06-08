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

package mem_test

import (
	reflectparser "github.com/katydid/katydid/parser/reflect"
	. "github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/mem"
	"github.com/katydid/katydid/relapse/tests"
	"reflect"
	"testing"
)

func NewInjectable() *injectableInt {
	return &injectableInt{V: funcs.IntConst(0)}
}

type injectableInt struct {
	V funcs.ConstInt
}

func (this *injectableInt) Eval() (int64, error) {
	return this.V.Eval()
}

func (this *injectableInt) SetValue(v int64) {
	this.V = funcs.IntConst(v)
}

func (this *injectableInt) IsVariable() {
	//If this method is not implemented this function will probably be trimmed
}

func init() {
	funcs.Register("inject", new(injectableInt))

	injectPerson = G{
		"main": InPath("Addresses", InAny(
			InPath("Number", Value(funcs.IntEq(funcs.IntVar(), NewInjectable()))),
		)),
	}
}

type InjectableInt interface {
	SetValue(v int64)
}

var injectPerson = G{}

func testInject(t *testing.T, num int64) bool {
	grammar := injectPerson.Grammar()
	m, err := mem.New(grammar)
	if err != nil {
		t.Fatal(err)
	}
	typ := reflect.TypeOf((*InjectableInt)(nil)).Elem()
	instances := mem.Implements(m, typ)
	if len(instances) == 0 {
		t.Fatal("Expected some functions that implement the interface")
	}
	for _, instance := range instances {
		instance.(InjectableInt).SetValue(num)
	}
	parser := reflectparser.NewReflectParser()
	parser.Init(reflect.ValueOf(tests.RobertPerson))
	res, err := m.Validate(parser)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestInjectPositive(t *testing.T) {
	if !testInject(t, 456) {
		t.Fatalf("expected match")
	}
}

func TestInjectNegative(t *testing.T) {
	if testInject(t, 123) {
		t.Fatalf("expected non match")
	}
}
