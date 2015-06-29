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
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/fastmaps"
)

type Visitor interface {
	GetToken(pkg, msg, field string) int
	AddTrans(srcToken int, field *descriptor.FieldDescriptorProto, dstToken int)
}

type visitor struct {
	transKey    map[int]map[uint64]int
	trans       map[int]map[string]int
	transName   map[int]map[uint64]string
	nameToToken map[string]int
	tokenToName map[int]string
	leaf        map[int]bool
	types       map[int]descriptor.FieldDescriptorProto_Type
}

func newVisitor() *visitor {
	visitor := &visitor{
		transKey:    make(map[int]map[uint64]int),
		trans:       make(map[int]map[string]int),
		transName:   make(map[int]map[uint64]string),
		nameToToken: make(map[string]int),
		tokenToName: make(map[int]string),
		leaf:        make(map[int]bool),
		types:       make(map[int]descriptor.FieldDescriptorProto_Type),
	}
	return visitor
}

func tokenName(pkg, msg, field string) string {
	key := pkg + "." + msg
	if len(field) > 0 {
		key += "." + field
	}
	return key
}

func (this *visitor) GetToken(pkg, msg, field string) int {
	key := tokenName(pkg, msg, field)
	s, ok := this.nameToToken[key]
	if ok {
		return s
	}
	s = len(this.nameToToken)
	this.nameToToken[key] = s
	this.tokenToName[s] = key
	if len(field) > 0 {
		this.leaf[s] = true
	}
	return s
}

func (this *visitor) AddTrans(srcToken int, field *descriptor.FieldDescriptorProto, dstToken int) {
	input := field.GetKeyUint64()
	if _, ok := this.transKey[srcToken]; !ok {
		this.transKey[srcToken] = make(map[uint64]int)
	}
	if _, ok := this.transKey[srcToken][input]; ok {
		panic("transition already exists")
	}
	this.transKey[srcToken][input] = dstToken
	this.types[dstToken] = *field.Type

	inputName := field.GetName()
	if _, ok := this.trans[srcToken]; !ok {
		this.trans[srcToken] = make(map[string]int)
	}
	if _, ok := this.trans[srcToken][inputName]; ok {
		panic("transition already exists")
	}
	this.trans[srcToken][inputName] = dstToken

	if _, ok := this.transName[srcToken]; !ok {
		this.transName[srcToken] = make(map[uint64]string)
	}
	if _, ok := this.transName[srcToken][input]; ok {
		panic("transition already exists")
	}
	this.transName[srcToken][input] = inputName
}

func (this *visitor) finalize() *protoTokens {
	numStates := len(this.tokenToName)
	m := &protoTokens{
		transKey:    make([]fastmaps.Uint64ToInt, numStates),
		trans:       make(map[int]map[string]int),
		tokenToName: make(map[int]string),
		transName:   this.transName,
		nameToToken: this.nameToToken,
		leaf:        make([]bool, numStates),
		types:       make([]descriptor.FieldDescriptorProto_Type, numStates),
	}
	for i := 0; i < numStates; i++ {
		if len(this.transKey[i]) > 0 {
			m.transKey[i] = fastmaps.NewUint64ToInt(this.transKey[i])
		}
		m.tokenToName[i] = this.tokenToName[i]
		m.leaf[i] = this.leaf[i]
		if t, ok := this.types[i]; ok {
			m.types[i] = t
		}
		m.trans[i] = make(map[string]int)
		for k, v := range this.trans[i] {
			m.trans[i][k] = v
		}
	}
	return m
}

func Walk(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet, visitor Visitor) error {
	w := &walker{
		visited: make(map[int]bool),
		desc:    desc,
		visitor: visitor,
	}
	return w.visit(srcPackage, srcMessage)
}

type walker struct {
	visited map[int]bool
	desc    *descriptor.FileDescriptorSet
	visitor Visitor
}

func (this *walker) isVisited(pkg, msg string) bool {
	s := this.visitor.GetToken(pkg, msg, "")
	_, ok := this.visited[s]
	return ok
}

func (this *walker) visit(srcPackage string, srcMessage string) error {
	msg := this.desc.GetMessage(srcPackage, srcMessage)
	if msg == nil {
		return &errUnknown{srcPackage + "." + srcMessage}
	}
	currentState := this.visitor.GetToken(srcPackage, srcMessage, "")
	this.visited[currentState] = true
	for _, f := range msg.GetField() {
		if f.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			newPackage, newMessage := this.desc.FindMessage(srcPackage, srcMessage, f.GetName())
			if len(newMessage) == 0 {
				return &errUnknown{srcPackage + "." + srcMessage + "." + f.GetName()}
			}
			dstState := this.visitor.GetToken(newPackage, newMessage, "")
			if _, ok := this.visited[dstState]; !ok {
				if err := this.visit(newPackage, newMessage); err != nil {
					return err
				}
			}
			this.visitor.AddTrans(currentState, f, dstState)
		} else {
			dstState := this.visitor.GetToken(srcPackage, srcMessage, f.GetName())
			this.visitor.AddTrans(currentState, f, dstState)
		}
	}
	return nil
}
