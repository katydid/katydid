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

package convert

import (
	"fmt"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/serialize"
	"io"
)

type next struct {
	value   funcs.Bool
	down    pattern
	null    pattern
	notnull pattern
}

func derivUnion(left, right []next) []next {
	panic("todo")
}

func derivIntersect(left, right []next) []next {
	panic("todo")
}

func DerivEval(refs map[string]*relapse.Pattern, parser serialize.Parser) bool {
	newRefs := make(map[string]pattern)
	for name, p := range refs {
		newRefs[name] = FromPattern(p)
	}
	m := derivEval(newRefs, map[funcs.Bool]pattern{funcs.BoolConst(true): newRefs["main"]}, parser)
	for _, p := range m {
		return nullable(newRefs, p)
	}
	panic("Wtf")
}

func printMap(name string, m map[funcs.Bool]pattern) {
	fmt.Printf("<%v>\n", name)
	for k, v := range m {
		fmt.Printf("\t%s: %s\n", funcs.Sprint(k), v.String())
	}
	fmt.Printf("</%v>\n", name)
}

func derivEval(refs map[string]pattern, currents map[funcs.Bool]pattern, parser serialize.Parser) map[funcs.Bool]pattern {
	for {
		// fmt.Printf("NEXT\n")
		// printMap("CURRENTS", currents)
		if err := parser.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		downs := make(map[funcs.Bool]pattern)
		//fmt.Printf("VALUE: %v\n", getValue(parser))
		for _, current := range currents {
			thisdowns := derivCall(refs, current, parser)
			downs = unionMap(downs, thisdowns)
		}
		if parser.IsLeaf() {
			//do nothing
		} else {
			//		fmt.Printf("DOWN\n")
			parser.Down()
			downs = derivEval(refs, downs, parser)
			parser.Up()
			//		fmt.Printf("UP\n")
		}
		for currentf, current := range currents {
			thisups := derivReturn(refs, current, downs)
			currents[currentf] = thisups
		}
	}
	return currents
}

func derivReturn(refs map[string]pattern, p pattern, downs map[funcs.Bool]pattern) pattern {
	switch p.typ {
	case EmptyType:
		p, ok := downs[p.expr]
		if !ok {
			panic("wtf")
		}
		return p
	case TreeNodeType:
		child, ok := downs[p.expr]
		if !ok {
			panic("wtf")
		}
		if nullable(refs, child) {
			return newEmpty()
		} else {
			return newNot(newZany())
		}
	case LeafNodeType:
		p, ok := downs[p.expr]
		if !ok {
			panic("wtf")
		}
		return p
	case ConcatType:
		left := derivReturn(refs, p.left(), downs)
		lefts := newConcat(left, p.right())
		if !nullable(refs, p.left()) {
			return lefts
		}
		right := derivReturn(refs, p.right(), downs)
		return newOr(lefts, right)
	case OrType:
		left := derivReturn(refs, p.left(), downs)
		right := derivReturn(refs, p.right(), downs)
		return newOr(left, right)
	case AndType:
		left := derivReturn(refs, p.left(), downs)
		right := derivReturn(refs, p.right(), downs)
		return newAnd(left, right)
	case ZeroOrMoreType:
		child := derivReturn(refs, p.child(), downs)
		return newConcat(child, p)
	case ReferenceType:
		p := refs[p.name]
		d := derivReturn(refs, p, downs)
		return d
	case NotType:
		d := derivReturn(refs, p.child(), downs)
		return newNot(d)
	case ZAnyType:
		p, ok := downs[p.expr]
		if !ok {
			panic("wtf")
		}
		return p
	// case ContainsType:
	// 	return derivReturn(refs, newConcat(newZany(), newConcat(p.child(), newZany())), downs)
	// case OptionalType:
	// 	return derivReturn(refs, newOr(p.child(), newEmpty()), downs)
	case InterleaveType:
		left := derivReturn(refs, p.left(), downs)
		right := derivReturn(refs, p.right(), downs)
		lefts := newInterleave(left, p.right())
		rights := newInterleave(right, p.left())
		return newOr(lefts, rights)
	}
	panic(fmt.Sprintf("unknown pattern typ %v", p))
}

func unionMap(left, right map[funcs.Bool]pattern) map[funcs.Bool]pattern {
	n := make(map[funcs.Bool]pattern)
	for f, p := range left {
		n[f] = p
	}
	for f, p := range right {
		n[f] = p
	}
	return n
}

func derivCall(refs map[string]pattern, p pattern, value serialize.Decoder) map[funcs.Bool]pattern {
	switch p.typ {
	case EmptyType:
		return map[funcs.Bool]pattern{p.expr: newNot(newZany())}
	case TreeNodeType:
		f, err := compose.NewBoolFunc(p.expr)
		if err != nil {
			panic(err)
		}
		b, err := f.Eval(value)
		if err != nil {
			panic(err)
		}
		if b {
			return map[funcs.Bool]pattern{p.expr: p.child()}
		}
		return map[funcs.Bool]pattern{p.expr: newNot(newZany())}
	case LeafNodeType:
		f, err := compose.NewBoolFunc(p.expr)
		if err != nil {
			panic(err)
		}
		b, err := f.Eval(value)
		if err != nil {
			panic(err)
		}
		if b {
			return map[funcs.Bool]pattern{p.expr: newEmpty()}
		}
		return map[funcs.Bool]pattern{p.expr: newNot(newZany())}
	case ConcatType:
		left := derivCall(refs, p.left(), value)
		if !nullable(refs, p.left()) {
			return left
		}
		right := derivCall(refs, p.right(), value)
		return unionMap(left, right)
	case OrType:
		left := derivCall(refs, p.left(), value)
		right := derivCall(refs, p.right(), value)
		return unionMap(left, right)
	case AndType:
		left := derivCall(refs, p.left(), value)
		right := derivCall(refs, p.right(), value)
		return unionMap(left, right)
	case ZeroOrMoreType:
		return derivCall(refs, p.child(), value)
	case ReferenceType:
		p := refs[p.name]
		d := derivCall(refs, p, value)
		return d
	case NotType:
		d := derivCall(refs, p.child(), value)
		return d
	case ZAnyType:
		return map[funcs.Bool]pattern{p.expr: newZany()}
	// case ContainsType:
	// 	return derivCall(refs, newConcat(newZany(), newConcat(p.child(), newZany())), value)
	// case OptionalType:
	// 	return derivCall(refs, newOr(p.child(), newEmpty()), value)
	case InterleaveType:
		left := derivCall(refs, p.left(), value)
		right := derivCall(refs, p.right(), value)
		return unionMap(left, right)
	}
	panic(fmt.Sprintf("unknown pattern typ %v", p))
}

func nullable(refs map[string]pattern, p pattern) bool {
	switch p.typ {
	case EmptyType:
		return true
	case TreeNodeType:
		return false
	case LeafNodeType:
		return false
	case ConcatType:
		return nullable(refs, p.left()) && nullable(refs, p.right())
	case OrType:
		return nullable(refs, p.left()) || nullable(refs, p.right())
	case AndType:
		return nullable(refs, p.left()) && nullable(refs, p.right())
	case ZeroOrMoreType:
		return true
	case ReferenceType:
		return nullable(refs, refs[p.name])
	case NotType:
		return !(nullable(refs, p.left()))
	case ZAnyType:
		return true
	// case ContainsType:
	// 	return nullable(refs, p.left())
	// case OptionalType:
	// 	return true
	case InterleaveType:
		return nullable(refs, p.left()) && nullable(refs, p.right())
	}
	panic(fmt.Sprintf("unknown pattern typ %v", p))
}
