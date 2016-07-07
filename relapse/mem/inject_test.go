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
	return &injectableInt{}
}

type injectableInt struct {
	context *funcs.Context
}

func (this *injectableInt) Eval() (int64, error) {
	return this.context.Value.(int64), nil
}

func (this *injectableInt) SetContext(context *funcs.Context) {
	this.context = context
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

var injectPerson = G{}

func testInject(t *testing.T, m *mem.Mem) bool {
	parser := reflectparser.NewReflectParser()
	parser.Init(reflect.ValueOf(tests.RobertPerson))
	res, err := m.Validate(parser)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func TestInject(t *testing.T) {
	grammar := injectPerson.Grammar()
	m, err := mem.New(grammar)
	if err != nil {
		t.Fatal(err)
	}
	c := &funcs.Context{Value: int64(0)}
	m.SetContext(c)
	c.Value = int64(456)
	if !testInject(t, m) {
		t.Fatalf("expected match")
	}
	c.Value = int64(123)
	if testInject(t, m) {
		t.Fatalf("expected non match")
	}
}
