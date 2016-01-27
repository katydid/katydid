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
	"github.com/katydid/katydid/relapse/ast"
)

func checkRef(refs relapse.RefLookup, p *relapse.Pattern) *relapse.Pattern {
	for name, rpat := range refs {
		if rpat.Equal(p) {
			return relapse.NewReference(name)
		}
	}
	return p
}

func Simplify(refs relapse.RefLookup, p *relapse.Pattern) *relapse.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return p
	case *relapse.TreeNode:
		return checkRef(refs, relapse.NewTreeNode(v.GetName(), Simplify(refs, v.GetPattern())))
	case *relapse.LeafNode:
		return p
	case *relapse.Concat:
		return checkRef(refs, simplifyConcat(
			Simplify(refs, v.GetLeftPattern()),
			Simplify(refs, v.GetRightPattern()),
		))
	case *relapse.Or:
		return checkRef(refs, simplifyOr(refs,
			Simplify(refs, v.GetLeftPattern()),
			Simplify(refs, v.GetRightPattern()),
		))
	case *relapse.And:
		return checkRef(refs, simplifyAnd(refs,
			Simplify(refs, v.GetLeftPattern()),
			Simplify(refs, v.GetRightPattern()),
		))
	case *relapse.ZeroOrMore:
		return checkRef(refs, simplifyZeroOrMore(Simplify(refs, v.GetPattern())))
	case *relapse.Reference:
		return p
	case *relapse.Not:
		return checkRef(refs, simplifyNot(Simplify(refs, v.GetPattern())))
	case *relapse.ZAny:
		return p
	case *relapse.Contains:
		return checkRef(refs, simplifyContains(Simplify(refs, v.GetPattern())))
	case *relapse.Optional:
		return simplifyOptional(Simplify(refs, v.GetPattern()))
	case *relapse.Interleave:
		return checkRef(refs, simplifyInterleave(refs,
			Simplify(refs, v.GetLeftPattern()),
			Simplify(refs, v.GetRightPattern()),
		))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func isNotZany(p *relapse.Pattern) bool {
	return p.Not != nil && p.GetNot().GetPattern().ZAny != nil
}

func isEmpty(p *relapse.Pattern) bool {
	return p.Empty != nil
}

func isZany(p *relapse.Pattern) bool {
	return p.ZAny != nil
}

func simplifyConcat(p1, p2 *relapse.Pattern) *relapse.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return relapse.NewNot(relapse.NewZAny())
	}
	if p1.Concat != nil {
		return simplifyConcat(
			p1.Concat.GetLeftPattern(),
			relapse.NewConcat(p1.Concat.GetRightPattern(), p2),
		)
	}
	if isEmpty(p1) {
		return p2
	}
	if isEmpty(p2) {
		return p1
	}
	if isZany(p1) && p2.Concat != nil {
		if l := p2.Concat.GetLeftPattern(); isZany(p2.Concat.GetRightPattern()) {
			return relapse.NewContains(l)
		}
	}
	return relapse.NewConcat(p1, p2)
}

func getOrs(p *relapse.Pattern) []*relapse.Pattern {
	if p.Or != nil {
		return append(getOrs(p.Or.GetLeftPattern()), getOrs(p.Or.GetRightPattern())...)
	}
	return []*relapse.Pattern{p}
}

func simplifyOr(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
	left := getOrs(p1)
	right := getOrs(p2)
	list := append(left, right...)
	list = relapse.Set(list)
	relapse.Sort(list)
	var p *relapse.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = relapse.NewOr(p, list[i])
	}
	return p
	// if isNotZany(p1) {
	// 	return p2
	// }
	// if isNotZany(p2) {
	// 	return p1
	// }
	// if isZany(p1) || isZany(p2) {
	// 	return relapse.NewZAny()
	// }
	// if isEmpty(p1) && Nullable(refs, p2) {
	// 	return p2
	// }
	// if isEmpty(p2) && Nullable(refs, p1) {
	// 	return p1
	// }
	// if p1.Or != nil {
	// 	return simplifyOr(refs, p1.Or.GetLeftPattern(), simplifyOr(refs, p1.Or.GetRightPattern(), p2))
	// }
	// if p1.Equal(p2) {
	// 	return p1
	// }
	// if p2.Or != nil {
	// 	if p1.Equal(p2.Or.GetLeftPattern()) || p1.Equal(p2.Or.GetRightPattern()) {
	// 		return simplifyOr(refs, p2.Or.GetLeftPattern(), p2.Or.GetRightPattern())
	// 	}
	// 	if p2.Or.GetLeftPattern().Less(p1) {
	// 		return simplifyOr(refs, p2.Or.GetLeftPattern(), simplifyOr(refs, p1, p2.Or.GetRightPattern()))
	// 	}
	// }
	// if p2.Less(p1) {
	// 	return relapse.NewOr(p2, p1)
	// }
	// return relapse.NewOr(p1, p2)
}

func getAnds(p *relapse.Pattern) []*relapse.Pattern {
	if p.And != nil {
		return append(getOrs(p.And.GetLeftPattern()), getOrs(p.And.GetRightPattern())...)
	}
	return []*relapse.Pattern{p}
}

func simplifyAnd(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
	// if isNotZany(p1) || isNotZany(p2) {
	// 	return relapse.NewNot(relapse.NewZAny())
	// }
	// if isZany(p1) {
	// 	return p2
	// }
	// if isZany(p2) {
	// 	return p1
	// }
	// if isEmpty(p1) {
	// 	if Nullable(refs, p2) {
	// 		return relapse.NewEmpty()
	// 	} else {
	// 		return relapse.NewNot(relapse.NewZAny())
	// 	}
	// }
	// if isEmpty(p2) {
	// 	if Nullable(refs, p1) {
	// 		return relapse.NewEmpty()
	// 	} else {
	// 		return relapse.NewNot(relapse.NewZAny())
	// 	}
	// }
	left := getAnds(p1)
	right := getAnds(p2)
	list := append(left, right...)
	list = relapse.Set(list)
	relapse.Sort(list)
	var p *relapse.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = relapse.NewAnd(p, list[i])
	}
	return p
}

// func simplifyAnd(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
// 	if isNotZany(p1) || isNotZany(p2) {
// 		return relapse.NewNot(relapse.NewZAny())
// 	}
// 	if isZany(p1) {
// 		return p2
// 	}
// 	if isZany(p2) {
// 		return p1
// 	}
// 	if isEmpty(p1) {
// 		if Nullable(refs, p2) {
// 			return relapse.NewEmpty()
// 		} else {
// 			return relapse.NewNot(relapse.NewZAny())
// 		}
// 	}
// 	if isEmpty(p2) {
// 		if Nullable(refs, p1) {
// 			return relapse.NewEmpty()
// 		} else {
// 			return relapse.NewNot(relapse.NewZAny())
// 		}
// 	}
// 	if p1.Equal(p2) {
// 		return p1
// 	}
// 	return relapse.NewAnd(p1, p2)
// }

func simplifyZeroOrMore(p *relapse.Pattern) *relapse.Pattern {
	if p.ZeroOrMore != nil {
		return p
	}
	return relapse.NewZeroOrMore(p)
}

func simplifyNot(p *relapse.Pattern) *relapse.Pattern {
	if p.Not != nil {
		return p.Not.GetPattern()
	}
	return relapse.NewNot(p)
}

//TODO we can do better
func simplifyOptional(p *relapse.Pattern) *relapse.Pattern {
	if isEmpty(p) {
		return relapse.NewEmpty()
	}
	return relapse.NewOptional(p)
}

func getInterleaves(p *relapse.Pattern) []*relapse.Pattern {
	if p.Interleave != nil {
		return append(getOrs(p.Interleave.GetLeftPattern()), getOrs(p.Interleave.GetRightPattern())...)
	}
	return []*relapse.Pattern{p}
}

func simplifyInterleave(refs relapse.RefLookup, p1, p2 *relapse.Pattern) *relapse.Pattern {
	if isNotZany(p1) || isNotZany(p2) {
		return relapse.NewNot(relapse.NewZAny())
	}
	if isEmpty(p1) {
		return p2
	}
	if isEmpty(p2) {
		return p1
	}
	if isZany(p1) && isZany(p2) {
		return p1
	}
	left := getInterleaves(p1)
	right := getInterleaves(p2)
	list := append(left, right...)
	list = relapse.Set(list)
	relapse.Sort(list)
	var p *relapse.Pattern = list[0]
	for i := range list {
		if i == 0 {
			continue
		}
		p = relapse.NewInterleave(p, list[i])
	}
	return p
}

func simplifyContains(p *relapse.Pattern) *relapse.Pattern {
	if isEmpty(p) || isZany(p) {
		return relapse.NewZAny()
	}
	if isNotZany(p) {
		return p
	}
	return relapse.NewContains(p)
}
