//  Copyright 2015 Walter Schulze
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

package interp

import (
	"fmt"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/lang/ast"
	"github.com/katydid/katydid/serialize"
)

type RefLookup map[string]*lang.Pattern

func refs(g *lang.Grammar) RefLookup {
	decls := g.GetPatternDecls()
	refs := make(RefLookup, len(decls))
	for _, d := range decls {
		refs[d.Name] = d.Pattern
	}
	return refs
}

func nullable(refs RefLookup, p *lang.Pattern) bool {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *lang.Empty:
		return true
	case *lang.EmptySet:
		return false
	case *lang.TreeNode:
		return false
	case *lang.LeafNode:
		return false
	case *lang.Concat:
		return nullable(refs, v.GetLeftPattern()) && nullable(refs, v.GetRightPattern())
	case *lang.Or:
		return nullable(refs, v.GetLeftPattern()) || nullable(refs, v.GetRightPattern())
	case *lang.And:
		return nullable(refs, v.GetLeftPattern()) && nullable(refs, v.GetRightPattern())
	case *lang.ZeroOrMore:
		return true
	case *lang.Reference:
		return nullable(refs, refs[v.GetName()])
	case *lang.Not:
		return !(nullable(refs, v.GetPattern()))
	}
	panic(fmt.Sprintf("unknown typ %T", typ))
}

type Tree interface {
	Copy() Tree
	serialize.Decoder
}

func deriv(refs RefLookup, p *lang.Pattern, tree Tree) *lang.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *lang.Empty:
		return lang.NewEmptySet()
	case *lang.EmptySet:
		return lang.NewEmptySet()
	case *lang.TreeNode:

	case *lang.LeafNode:
		f, err := compose.NewBool(v.GetExpr())
		if err != nil {
			panic(err)
		}
		res, err := f.Eval(tree)
		if err != nil {
			panic(err)
		}
		if res {
			return lang.NewEmpty()
		}
		return lang.NewEmptySet()
	case *lang.Concat:
		leftDeriv := lang.NewConcat(deriv(refs, v.GetLeftPattern(), tree), v.GetRightPattern())
		if nullable(refs, v.GetLeftPattern()) {
			return lang.NewOr(
				leftDeriv,
				deriv(refs, v.GetRightPattern(), tree.Copy()),
			)
		} else {
			return leftDeriv
		}
	case *lang.Or:
		return lang.NewOr(
			deriv(refs, v.GetLeftPattern(), tree),
			deriv(refs, v.GetRightPattern(), tree.Copy()),
		)
	case *lang.And:
		return lang.NewAnd(
			deriv(refs, v.GetLeftPattern(), tree),
			deriv(refs, v.GetRightPattern(), tree.Copy()),
		)
	case *lang.ZeroOrMore:
		return lang.NewConcat(deriv(refs, v.Pattern, tree), lang.NewZeroOrMore(v.Pattern))
	case *lang.Reference:
		return deriv(refs, refs[v.GetName()], tree)
	case *lang.Not:
		return lang.NewNot(deriv(refs, v.GetPattern(), tree))
	}
	panic(fmt.Sprintf("unknown typ %T", typ))
}

// deriv :: RefLookup -> Pattern -> Tree Alpha -> Pattern
// deriv _ Empty _ = EmptySet
// deriv _ EmptySet _ = EmptySet
// deriv refs (TreeNode a p) (Node b children) =
//   if a == b && nullable refs (foldl (deriv refs) p children)
//     then Empty
//     else EmptySet
// deriv _ (LeafNode a) (Node b []) =
//   if a == b then Empty else EmptySet
// deriv _ (LeafNode _) _ = EmptySet
// deriv refs (Concat a b) t = if nullable refs a
//   then Or (Concat (deriv refs a t) b) (deriv refs b t)
//   else Concat (deriv refs a t) b
// deriv refs (Or a b) t = Or (deriv refs a t) (deriv refs b t)
// deriv refs (And a b) t = And (deriv refs a t) (deriv refs b t)
// deriv refs (ZeroOrMore a) t = Concat (deriv refs a t) (ZeroOrMore a)
// deriv refs (Not a) t = Not (deriv refs a t)
// deriv refs (Interleave a b) t =
//   Or (Concat (deriv refs a t) b) (Concat (deriv refs b t) a)
// deriv refs (Reference r) t = deriv refs (refs r) t
