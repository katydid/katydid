//  Copyright 2017 Walter Schulze
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

package mem

import (
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/sets"
)

type ifExprs struct {
	parentPatterns int
	ifs            []*ifExpr
}

func newIfExprs(parentPatterns int, ifs []*ifExpr) *ifExprs {
	return &ifExprs{parentPatterns, ifs}
}

func (this *Mem) eval(ifs *ifExprs, label parser.Value) (int, int, error) {
	ps := make([]*ast.Pattern, len(ifs.ifs))
	for i := range ifs.ifs {
		p, err := ifs.ifs[i].eval(label)
		if err != nil {
			return 0, 0, err
		}
		ps[i] = p
	}
	zippedPatterns, zipper := sets.Zip(ps)
	zipperIndex := this.zis.Add(zipper)
	stackElement := sets.Pair{
		First:  ifs.parentPatterns,
		Second: zipperIndex,
	}
	stackIndex := this.stackElms.Add(stackElement)
	zippedPatternIndex := this.patterns.Add(zippedPatterns)
	return zippedPatternIndex, stackIndex, nil
}

type ifExpr struct {
	cond funcs.Bool
	then *ast.Pattern
	els  *ast.Pattern
}

func (this *ifExpr) eval(label parser.Value) (*ast.Pattern, error) {
	if this.els == nil {
		return this.then, nil
	}
	f, err := compose.NewBoolFunc(this.cond)
	if err != nil {
		return nil, err
	}
	cond, err := f.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.then, nil
	}
	return this.els, nil
}
