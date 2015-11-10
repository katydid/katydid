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

package protokey

import (
	"fmt"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/nameexpr"
	"github.com/katydid/katydid/serialize"
	"github.com/katydid/katydid/serialize/proto"
)

func KeyTheGrammar(pkgName, msgName string, desc *descriptor.FileDescriptorSet, grammar *relapse.Grammar) (*relapse.Grammar, error) {
	g := grammar.Clone()
	descMap, err := proto.NewDescriptorMap(pkgName, msgName, desc)
	if err != nil {
		return nil, err
	}
	root := descMap.GetRoot()
	refs := relapse.NewRefsLookup(g)
	keyer := &keyer{
		names:   make(map[*relapse.NameExpr]*relapse.NameExpr),
		msgs:    make(map[*relapse.NameExpr]*msgs),
		descMap: descMap,
		refs:    refs,
	}
	changed := true
	topMsg := &msgs{map[*descriptor.DescriptorProto]struct{}{root: struct{}{}}, nil}
	for changed {
		keyer.reset()
		changed = keyer.pattern(topMsg, refs["main"])
	}
	keyer.replace(g.TopPattern)
	for i := range g.PatternDecls {
		keyer.replace(g.PatternDecls[i].GetPattern())
	}
	return g, nil
}

func (this *keyer) replace(p *relapse.Pattern) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty, *relapse.EmptySet, *relapse.LeafNode, *relapse.ZAny:
		//do nothing
	case *relapse.TreeNode:
		name, ok := this.names[v.Name]
		if !ok {
			panic("unreachable")
		}
		v.Name = name
		this.replace(v.GetPattern())
	case *relapse.Concat:
		this.replace(v.GetLeftPattern())
		this.replace(v.GetRightPattern())
	case *relapse.Or:
		this.replace(v.GetLeftPattern())
		this.replace(v.GetRightPattern())
	case *relapse.And:
		this.replace(v.GetLeftPattern())
		this.replace(v.GetRightPattern())
	case *relapse.ZeroOrMore:
		this.replace(v.GetPattern())
	case *relapse.Reference:
		//do nothing
	case *relapse.Not:
		this.replace(v.GetPattern())
	case *relapse.WithSomeTreeNode:
		this.replace(v.GetPattern())
	case *relapse.Optional:
		this.replace(v.GetPattern())
	case *relapse.Interleave:
		this.replace(v.GetLeftPattern())
		this.replace(v.GetRightPattern())
	default:
		panic(fmt.Sprintf("unknown pattern typ %T", typ))
	}
}

type msgs struct {
	optional map[*descriptor.DescriptorProto]struct{}
	repeated map[*descriptor.DescriptorProto]struct{}
}

func (this *msgs) Len() int {
	return len(this.optional) + len(this.repeated)
}

func same(prev, next map[*descriptor.DescriptorProto]struct{}) bool {
	if len(prev) != len(next) {
		return false
	}
	for p := range prev {
		_, ok := next[p]
		if !ok {
			return false
		}
	}
	return true
}

func (this *msgs) Equal(that *msgs) bool {
	return same(this.optional, that.optional) && same(this.repeated, that.repeated)
}

type keyer struct {
	names   map[*relapse.NameExpr]*relapse.NameExpr
	msgs    map[*relapse.NameExpr]*msgs
	visited map[*relapse.Pattern]struct{}
	descMap proto.DescMap
	refs    relapse.RefLookup
}

func (this *keyer) reset() {
	this.visited = make(map[*relapse.Pattern]struct{})
}

func (this *keyer) pattern(current *msgs, p *relapse.Pattern) bool {
	if _, ok := this.visited[p]; ok {
		return false
	}
	this.visited[p] = struct{}{}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty, *relapse.EmptySet, *relapse.LeafNode, *relapse.ZAny:
		//do nothing
		return false
	case *relapse.TreeNode:
		if _, ok := this.msgs[v.Name]; ok {
			if this.msgs[v.Name].Equal(current) {
				return false
			}
		}
		if current.Len() == 0 {
			//the query has specified an unreachable path
			this.msgs[p.TreeNode.Name] = &msgs{}
			if evalAnyInt(p.TreeNode.Name) {
				this.names[p.TreeNode.Name] = p.TreeNode.Name
			} else {
				this.names[p.TreeNode.Name] = relapse.NewAnyNameExcept(relapse.NewAnyName())
			}
			this.pattern(&msgs{}, p.TreeNode.Pattern)
			return true
		}
		this.msgs[p.TreeNode.Name] = current
		fields := fieldsMap(current.optional)
		newName := keyTheName(fields, p.TreeNode.Name)
		this.names[p.TreeNode.Name] = newName
		nextMsgs := getMsgs(this.descMap, current, newName)
		fmt.Printf("%d:%v:%d:%v\n", current.Len(), p.TreeNode.Name, nextMsgs.Len(), newName)
		this.pattern(nextMsgs, p.TreeNode.Pattern)
		return true
	case *relapse.Concat:
		leftChanged := this.pattern(current, v.GetLeftPattern())
		rightChanged := this.pattern(current, v.GetRightPattern())
		return leftChanged || rightChanged
	case *relapse.Or:
		leftChanged := this.pattern(current, v.GetLeftPattern())
		rightChanged := this.pattern(current, v.GetRightPattern())
		return leftChanged || rightChanged
	case *relapse.And:
		leftChanged := this.pattern(current, v.GetLeftPattern())
		rightChanged := this.pattern(current, v.GetRightPattern())
		return leftChanged || rightChanged
	case *relapse.ZeroOrMore:
		return this.pattern(current, v.GetPattern())
	case *relapse.Reference:
		return this.pattern(current, this.refs[v.GetName()])
	case *relapse.Not:
		return this.pattern(current, v.GetPattern())
	case *relapse.WithSomeTreeNode:
		return this.pattern(current, v.GetPattern())
	case *relapse.Optional:
		return this.pattern(current, v.GetPattern())
	case *relapse.Interleave:
		leftChanged := this.pattern(current, v.GetLeftPattern())
		rightChanged := this.pattern(current, v.GetRightPattern())
		return leftChanged || rightChanged
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func fieldsMap(msgs map[*descriptor.DescriptorProto]struct{}) map[string][]*relapse.NameExpr {
	fields := make(map[string][]*relapse.NameExpr)
	for msg, _ := range msgs {
		for _, field := range msg.Field {
			_, ok := fields[field.GetName()]
			if !ok {
				fields[field.GetName()] = []*relapse.NameExpr{}
			}
			fields[field.GetName()] = append(fields[field.GetName()], relapse.NewUintName(uint64(field.GetNumber())))
		}
	}
	return fields
}

func keyTheName(msg map[string][]*relapse.NameExpr, name *relapse.NameExpr) *relapse.NameExpr {
	v := name.GetValue()
	switch n := v.(type) {
	case *relapse.Name:
		if n.StringValue == nil {
			//do nothing
			return name
		}
		keys, found := msg[n.GetStringValue()]
		if found {
			return relapse.NewNameChoice(keys...)
		}
		return relapse.NewAnyNameExcept(relapse.NewAnyName())
	case *relapse.AnyName:
		return relapse.NewAnyName()
	case *relapse.AnyNameExcept:
		return relapse.NewAnyNameExcept(keyTheName(msg, n.GetExcept()))
	case *relapse.NameChoice:
		return relapse.NewNameChoice(
			keyTheName(msg, n.GetLeft()),
			keyTheName(msg, n.GetRight()),
		)
	}
	panic(fmt.Sprintf("unknown name expression typ %T", v))
}

func evalAnyInt(name *relapse.NameExpr) bool {
	v := name.GetValue()
	switch n := v.(type) {
	case *relapse.Name:
		return n.IntValue != nil
	case *relapse.AnyName:
		return true
	case *relapse.AnyNameExcept:
		return !evalAnyInt(n.GetExcept())
	case *relapse.NameChoice:
		return evalAnyInt(n.GetLeft()) || evalAnyInt(n.GetRight())
	}
	panic(fmt.Sprintf("unknown name expression typ %T", v))
}

func getMsgs(descMap proto.DescMap, oldMsgs *msgs, name *relapse.NameExpr) *msgs {
	optMsgs := make(map[*descriptor.DescriptorProto]struct{})
	repMsgs := make(map[*descriptor.DescriptorProto]struct{})
	for oldMsg := range oldMsgs.optional {
		for i, field := range oldMsg.GetField() {
			if field.IsMessage() {
				if nameexpr.EvalName(name, serialize.NewUintValue(uint64(field.GetNumber()))) {
					msg := descMap.LookupMessage(oldMsg.Field[i])
					if msg == nil {
						panic("unreachable")
					}
					if !field.IsRepeated() {
						optMsgs[msg] = struct{}{}
					} else {
						repMsgs[msg] = struct{}{}
					}

				}
			}
		}
	}
	for msg := range oldMsgs.repeated {
		if evalAnyInt(name) {
			optMsgs[msg] = struct{}{}
		}
	}
	return &msgs{optMsgs, repMsgs}
}
