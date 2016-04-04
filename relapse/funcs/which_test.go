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
	"github.com/katydid/katydid/relapse/types"
	"testing"
)

type which struct {
	exp string
}

func (this which) test(t *testing.T, name string, params ...types.Type) {
	uniq, err := funcsMap.which(name, params...)
	if err != nil {
		panic(err)
	}
	if uniq != this.exp {
		t.Fatalf("expected %v got %v", this.exp, uniq)
	}
}

func TestWhichStringEq(t *testing.T) {
	which{"stringEq"}.test(t, "eq", types.SINGLE_STRING, types.SINGLE_STRING)
}

func TestWhichInt64Eq(t *testing.T) {
	which{"intEq"}.test(t, "eq", types.SINGLE_INT, types.SINGLE_INT)
}

func TestWhichInt64Ge(t *testing.T) {
	which{"intGE"}.test(t, "ge", types.SINGLE_INT, types.SINGLE_INT)
}

func TestWhichElem(t *testing.T) {
	which{"elemUints"}.test(t, "elem", types.LIST_UINT, types.SINGLE_INT)
}
