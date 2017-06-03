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

package auto

import (
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
)

type exprToFunc struct {
	m   map[*ast.Expr]funcs.Bool
	err error
}

func (this *exprToFunc) Visit(node interface{}) interface{} {
	if this.err != nil {
		return this
	}
	leaf, ok := node.(*ast.LeafNode)
	if ok {
		f, err := compose.NewBool(leaf.Expr)
		if err != nil {
			this.err = err
			return this
		}
		this.m[leaf.Expr] = f
	}
	return this
}
