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

	LookupKey(src int, key uint64) (int, string, bool)
	LookupType(src int) descriptor.FieldDescriptorProto_Type
	LookupName(src int) string

	Dot() string
}

type protoTokens struct {
	transKey    []maps.Uint64ToInt
	transName   map[int]map[uint64]string
	trans       map[int]map[string]int
	tokenToName map[int]string
	nameToToken map[string]int
	leaf        []bool
	types       []descriptor.FieldDescriptorProto_Type
}

func New(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet) (ProtoTokens, error) {
	visitor := newVisitor()
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

func (this *protoTokens) LookupName(src int) string {
	return this.tokenToName[src]
}

func (this *protoTokens) LookupKey(src int, key uint64) (int, string, bool) {
	if src >= len(this.transKey) {
		return 0, "", false
	}
	i, ok := this.transKey[src].Lookup(key)
	if !ok {
		return i, "", ok
	}
	return i, this.transName[src][key], ok
}

func (this *protoTokens) getToken(tokenStr string) int {
	s, ok := this.nameToToken[tokenStr]
	if ok {
		return s
	}
	s = len(this.nameToToken)
	this.nameToToken[tokenStr] = s
	this.tokenToName[s] = tokenStr
	return s
}

func (this *protoTokens) Lookup(parentToken int, name []byte) int {
	strName := string(name)
	token, ok := this.trans[parentToken][strName]
	if ok {
		return token
	}
	parentName, ok := this.tokenToName[parentToken]
	if !ok {
		panic("no parent")
	}
	fullName := parentName + "." + strName
	newToken := this.getToken(fullName)
	this.trans[parentToken][strName] = newToken
	return newToken
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
	transKeys := make([]string, 0, len(this.transKey))
	for src, transKey := range this.transKey {
		if transKey == nil {
			continue
		}
		for input, dst := range transKey.ToMap() {
			transKeys = append(transKeys, fmt.Sprintf("\tnode%d -> node%d [label=\"%d={wire=%d,field=%d}\"]\n", src, dst, input, int(input&0x7), int32(input>>3)))
		}
	}
	return "digraph tokens {\n" + strings.Join(nodes, "") + "\n" + strings.Join(transKeys, "") + "}"
}
