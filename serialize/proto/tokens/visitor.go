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
	"github.com/katydid/katydid/asm/maps"
	"strings"
)

type Visitor interface {
	GetToken(pkg, msg, field string) int
	AddTrans(srcToken int, field *descriptor.FieldDescriptorProto, dstToken int)
	StringTokens() []string
}

type visitor struct {
	trans        map[int]map[uint64]int
	nameToToken  map[string]int
	tokenToName  map[int]string
	leaf         map[int]bool
	types        map[int]descriptor.FieldDescriptorProto_Type
	stringTokens []string
}

func newVisitor(stringTokens []string) *visitor {
	visitor := &visitor{
		trans:        make(map[int]map[uint64]int),
		nameToToken:  make(map[string]int),
		tokenToName:  make(map[int]string),
		leaf:         make(map[int]bool),
		types:        make(map[int]descriptor.FieldDescriptorProto_Type),
		stringTokens: stringTokens,
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
	if _, ok := this.trans[srcToken]; !ok {
		this.trans[srcToken] = make(map[uint64]int)
	}
	if _, ok := this.trans[srcToken][input]; ok {
		panic("transition already exists")
	}
	this.trans[srcToken][input] = dstToken
	this.types[dstToken] = *field.Type
}

func (this *visitor) StringTokens() []string {
	return this.stringTokens
}

func (this *visitor) finalize() *protoTokens {
	numStates := len(this.tokenToName)
	m := &protoTokens{
		trans:       make([]maps.Uint64ToInt, numStates),
		tokenToName: make([]string, numStates),
		nameToToken: this.nameToToken,
		leaf:        make([]bool, numStates),
		types:       make([]descriptor.FieldDescriptorProto_Type, numStates),
	}
	for i := 0; i < numStates; i++ {
		if len(this.trans[i]) > 0 {
			m.trans[i] = maps.NewUint64ToInt(this.trans[i])
		}
		m.tokenToName[i] = this.tokenToName[i]
		m.leaf[i] = this.leaf[i]
		if t, ok := this.types[i]; ok {
			m.types[i] = t
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

func included(stringTokens []string, pkg string, msg string, field string) bool {
	if stringTokens == nil {
		return true
	}
	for _, inc := range stringTokens {
		incs := strings.Split(inc, ".")
		if incs[0] == pkg {
			if incs[1] == msg {
				if len(incs) > 2 {
					if incs[2] == field {
						return true
					}
				} else {
					if len(field) == 0 {
						return true
					}
				}
			}
		}
	}
	return false
}

func (this *walker) visit(srcPackage string, srcMessage string) error {
	if !included(this.visitor.StringTokens(), srcPackage, srcMessage, "") {
		return nil
	}
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
			if !included(this.visitor.StringTokens(), srcPackage, srcMessage, f.GetName()) {
				continue
			}
			dstState := this.visitor.GetToken(srcPackage, srcMessage, f.GetName())
			this.visitor.AddTrans(currentState, f, dstState)
		}
	}
	return nil
}
