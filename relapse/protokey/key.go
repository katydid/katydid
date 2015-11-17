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
		refs:    make(map[string]*context),
		descMap: descMap,
	}
	keyer.refs["main"] = &context{root, false}
	newRefs := make(map[string]*relapse.Pattern)
	oldContexts := 0
	newContexts := 1
	for oldContexts != newContexts {
		oldContexts = newContexts
		for name, context := range keyer.refs {
			newp, err := keyer.translate(context, refs[name])
			if err != nil {
				return nil, err
			}
			newRefs[name] = newp
		}
		newContexts = len(keyer.refs)
	}
	return relapse.NewGrammar(newRefs), nil
}

type keyer struct {
	refs    map[string]*context
	descMap proto.DescMap
}

type context struct {
	msg   *descriptor.DescriptorProto
	index bool
}

func (this *context) Equal(that *context) bool {
	return this.msg == that.msg && this.index == that.index
}

func anyErr(err1, err2 error) error {
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

type ErrDup struct {
	name string
	c1   *context
	c2   *context
}

func (this *ErrDup) Error() string {
	return fmt.Sprintf("Duplicate Reference Error: Name: %v, Context1: %v, Context2: %v", this.name, this.c1.msg.GetName(), this.c2.msg.GetName())
}

func (this *keyer) translate(context *context, p *relapse.Pattern) (*relapse.Pattern, error) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty, *relapse.EmptySet, *relapse.LeafNode, *relapse.ZAny:
		return p, nil
	case *relapse.TreeNode:
		return this.translateName(context, v.GetName(), v.GetPattern())
	case *relapse.Concat:
		l, err1 := this.translate(context, v.GetLeftPattern())
		r, err2 := this.translate(context, v.GetRightPattern())
		return relapse.NewConcat(l, r), anyErr(err1, err2)
	case *relapse.Or:
		l, err1 := this.translate(context, v.GetLeftPattern())
		r, err2 := this.translate(context, v.GetRightPattern())
		return relapse.NewOr(l, r), anyErr(err1, err2)
	case *relapse.And:
		l, err1 := this.translate(context, v.GetLeftPattern())
		r, err2 := this.translate(context, v.GetRightPattern())
		return relapse.NewAnd(l, r), anyErr(err1, err2)
	case *relapse.ZeroOrMore:
		p, err := this.translate(context, v.GetPattern())
		return relapse.NewZeroOrMore(p), err
	case *relapse.Reference:
		c, ok := this.refs[v.GetName()]
		if !ok {
			this.refs[v.GetName()] = context
			return p, nil
		}
		if !c.Equal(context) {
			return nil, &ErrDup{v.GetName(), c, context}
		}
		return p, nil
	case *relapse.Not:
		p, err := this.translate(context, v.GetPattern())
		return relapse.NewNot(p), err
	case *relapse.WithSomeTreeNode:
		return nil, fmt.Errorf("todo")
	case *relapse.Optional:
		p, err := this.translate(context, v.GetPattern())
		return relapse.NewOptional(p), err
	case *relapse.Interleave:
		l, err1 := this.translate(context, v.GetLeftPattern())
		r, err2 := this.translate(context, v.GetRightPattern())
		return relapse.NewInterleave(l, r), anyErr(err1, err2)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

type ErrExpectedArray struct {
	name string
	c    *context
}

func (this *ErrExpectedArray) Error() string {
	return fmt.Sprintf("Expected Array Error: Name: %v, Msg: %v", this.name, this.c.msg.GetName())
}

type ErrExpectedField struct {
	name string
	c    *context
}

func (this *ErrExpectedField) Error() string {
	return fmt.Sprintf("Expected Field Error: Name: %v, Msg: %v", this.name, this.c.msg.GetName())
}

type ErrAnyFieldNotSupported struct {
	name string
}

func (this *ErrAnyFieldNotSupported) Error() string {
	return fmt.Sprintf("Any Field Not Supported: Name: %v", this.name)
}

type ErrAnyNameExceptNotSupported struct {
	name string
}

func (this *ErrAnyNameExceptNotSupported) Error() string {
	return fmt.Sprintf("AnyNameExcept Not Supported Error: Name: %v", this.name)
}

type ErrUnknownField struct {
	name string
	c    *context
}

func (this *ErrUnknownField) Error() string {
	return fmt.Sprintf("Unknown Field Error: Name: %v, Msg: %v", this.name, this.c.msg.GetName())
}

func getField(msg *descriptor.DescriptorProto, name string) *descriptor.FieldDescriptorProto {
	for i, f := range msg.Field {
		if f.GetName() != name {
			continue
		}
		return msg.Field[i]
	}
	return nil
}

func (this *keyer) translateName(current *context, name *relapse.NameExpr, child *relapse.Pattern) (*relapse.Pattern, error) {
	switch n := name.GetValue().(type) {
	case *relapse.Name:
		if current.index {
			if n.IntValue == nil {
				return nil, &ErrExpectedArray{name.String(), current}
			}
			c := &context{current.msg, false}
			newp, err := this.translate(c, child)
			if err != nil {
				return nil, err
			}
			return relapse.NewTreeNode(name, newp), nil
		}
		if n.StringValue == nil {
			return nil, &ErrExpectedField{name.String(), current}
		}
		f := getField(current.msg, n.GetStringValue())
		if f == nil {
			return nil, &ErrUnknownField{name.String(), current}
		}
		msg := this.descMap.LookupMessage(f)
		c := &context{msg, f.IsRepeated()}
		newp, err := this.translate(c, child)
		if err != nil {
			return nil, err
		}
		newName := relapse.NewUintName(uint64(f.GetNumber()))
		return relapse.NewTreeNode(newName, newp), nil
	case *relapse.AnyName:
		if current.index {
			c := &context{current.msg, false}
			newp, err := this.translate(c, child)
			if err != nil {
				return nil, err
			}
			return relapse.NewTreeNode(name, newp), nil
		} else {
			return nil, &ErrAnyFieldNotSupported{name.String()}
		}
	case *relapse.AnyNameExcept:
		return nil, &ErrAnyNameExceptNotSupported{name.String()}
	case *relapse.NameChoice:
		l, err1 := this.translateName(current, n.GetLeft(), child)
		r, err2 := this.translateName(current, n.GetRight(), child)
		return relapse.NewOr(l, r), anyErr(err1, err2)
	}
	panic(fmt.Sprintf("unknown name typ %T", name))
}
