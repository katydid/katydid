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

package tokens

import (
	"fmt"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/asm/maps"
	"strings"
)

type errUnknown struct {
	Token string
}

func (this *errUnknown) Error() string {
	return "Could not find " + this.Token
}

type ProtoTokens interface {
	GetTokenId(tokenString string) (int, error)
	Len() int
	IsLeaf(int) bool

	Lookup(src int, key uint64) (int, bool)
	LookupType(src int) descriptor.FieldDescriptorProto_Type

	Dot() string
}

type protoTokens struct {
	trans       []maps.Uint64ToInt
	tokenToName []string
	nameToToken map[string]int
	leaf        []bool
	types       []descriptor.FieldDescriptorProto_Type
}

func NewZipped(rules *ast.Rules, desc *descriptor.FileDescriptorSet) (ProtoTokens, error) {
	stringTokens, err := ast.GetStringTokens(rules)
	if err != nil {
		return nil, err
	}
	tokens, err := new(rules.GetRoot().GetPackage(), rules.GetRoot().GetMessage(), desc, stringTokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func New(rules *ast.Rules, desc *descriptor.FileDescriptorSet) (ProtoTokens, error) {
	srcPackage, srcMessage := rules.GetRoot().GetPackage(), rules.GetRoot().GetMessage()
	return new(srcPackage, srcMessage, desc, nil)
}

func new(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet, stringTokens []string) (ProtoTokens, error) {
	visitor := newVisitor(stringTokens)
	if err := Walk(srcPackage, srcMessage, desc, visitor); err != nil {
		return nil, err
	}
	tokens := visitor.finalize()
	return tokens, nil
}

func (this *protoTokens) GetTokenId(tokenString string) (int, error) {
	s, ok := this.nameToToken[tokenString]
	if !ok {
		return 0, &errUnknown{tokenString}
	}
	return s, nil
}

func (this *protoTokens) Len() int {
	return len(this.tokenToName)
}

func (this *protoTokens) IsLeaf(token int) bool {
	if token >= len(this.leaf) {
		return false
	}
	return this.leaf[token]
}

func (this *protoTokens) Lookup(src int, key uint64) (int, bool) {
	if src >= len(this.trans) {
		return 0, false
	}
	return this.trans[src].Lookup(key)
}

func (this *protoTokens) LookupType(src int) descriptor.FieldDescriptorProto_Type {
	if src >= len(this.types) {
		return 0
	}
	return this.types[src]
}

func (this *protoTokens) Dot() string {
	nodes := make([]string, len(this.tokenToName))
	for i, name := range this.tokenToName {
		nodes[i] = fmt.Sprintf("\tnode%d [label=\"%v\"]\n", i, name)
	}
	transs := make([]string, 0, len(this.trans))
	for src, trans := range this.trans {
		if trans == nil {
			continue
		}
		for input, dst := range trans.ToMap() {
			transs = append(transs, fmt.Sprintf("\tnode%d -> node%d [label=\"%d={wire=%d,field=%d}\"]\n", src, dst, input, int(input&0x7), int32(input>>3)))
		}
	}
	return "digraph tokens {\n" + strings.Join(nodes, "") + "\n" + strings.Join(transs, "") + "}"
}
