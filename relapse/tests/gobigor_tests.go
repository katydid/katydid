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
	"fmt"

	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/funcs"
)

func init() {
	num := 70
	ors := make([]*ast.Pattern, num)
	for i := 0; i < num; i++ {
		ors[i] = ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), combinator.Value(funcs.IntEq(funcs.IntVar(), funcs.IntConst(int64(i))))))
	}
	bigor := ast.NewOr(ors...)
	for i := 5; i < 11; i++ {
		Validate(fmt.Sprintf("GoBigOr%d", i), combinator.G{"main": bigor}, JsonString(fmt.Sprintf("{\"A\": %d}", i)), true)
	}
}
