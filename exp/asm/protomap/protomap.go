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

package protomap

import (
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"fmt"
	"github.com/awalterschulze/katydid/exp/asm/maps"
	"strings"
)

type errUnknown struct {
	Package string
	Message string
	Field   string
}

func (this *errUnknown) Error() string {
	return "Could not find " + this.Package + "." + this.Message + "." + this.Field
}

type ProtoMap interface {
	State(pkg, msg, field string) (int, error)
	LenStates() int
	Trans(src int, key uint64) (int, bool)
	IsLeave(int) bool
	Dot() string
}

type protoMap struct {
	trans       []maps.Uint64ToInt
	stateToName []string
	nameToState map[string]int
	leave       []bool
}

func New(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet) (ProtoMap, error) {
	return NewZipped(srcPackage, srcMessage, desc, nil)
}

func NewZipped(srcPackage, srcMessage string, desc *descriptor.FileDescriptorSet, includes []string) (ProtoMap, error) {
	builder := &mapBuilder{
		trans:       make(map[int]map[uint64]int),
		nameToState: make(map[string]int),
		stateToName: make(map[int]string),
		visited:     make(map[int]bool),
		leave:       make(map[int]bool),
		srcPackage:  srcPackage,
		srcMessage:  srcMessage,
		includes:    includes,
		desc:        desc,
	}
	if err := builder.build(); err != nil {
		return nil, err
	}
	pm := builder.finalize()
	return pm, nil
}

func (this *protoMap) State(pkg, msg, field string) (int, error) {
	s, ok := this.nameToState[stateName(pkg, msg, field)]
	if !ok {
		return 0, &errUnknown{pkg, msg, field}
	}
	return s, nil
}

func (this *protoMap) LenStates() int {
	return len(this.stateToName)
}

func (this *protoMap) IsLeave(state int) bool {
	if state >= len(this.leave) {
		return false
	}
	return this.leave[state]
}

func (this *protoMap) Trans(src int, key uint64) (int, bool) {
	if src >= len(this.trans) {
		return 0, false
	}
	return this.trans[src].Lookup(key)
}

func (this *protoMap) Dot() string {
	nodes := make([]string, len(this.stateToName))
	for i, name := range this.stateToName {
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
	return "digraph map {\n" + strings.Join(nodes, "") + "\n" + strings.Join(transs, "") + "}"
}

func stateName(pkg, msg, field string) string {
	key := pkg + "." + msg
	if len(field) > 0 {
		key += "." + field
	}
	return key
}

type mapBuilder struct {
	trans       map[int]map[uint64]int
	nameToState map[string]int
	stateToName map[int]string
	visited     map[int]bool
	leave       map[int]bool
	srcPackage  string
	srcMessage  string
	includes    []string
	desc        *descriptor.FileDescriptorSet
}

func (this *mapBuilder) finalize() *protoMap {
	numStates := len(this.stateToName)
	m := &protoMap{
		trans:       make([]maps.Uint64ToInt, numStates),
		stateToName: make([]string, numStates),
		nameToState: this.nameToState,
		leave:       make([]bool, numStates),
	}
	for i := 0; i < numStates; i++ {
		if len(this.trans[i]) > 0 {
			m.trans[i] = maps.NewUint64ToInt(this.trans[i])
		}
		m.stateToName[i] = this.stateToName[i]
		m.leave[i] = this.leave[i]
	}
	return m
}

func (this *mapBuilder) build() error {
	return this.visit(this.srcPackage, this.srcMessage)
}

func (this *mapBuilder) isVisited(pkg, msg string) bool {
	s := this.getState(pkg, msg, "")
	_, ok := this.visited[s]
	return ok
}

func (this *mapBuilder) getState(pkg, msg, field string) int {
	key := stateName(pkg, msg, field)
	s, ok := this.nameToState[key]
	if ok {
		return s
	}
	s = len(this.nameToState)
	this.nameToState[key] = s
	this.stateToName[s] = key
	if len(field) > 0 {
		this.leave[s] = true
	}
	return s
}

func (this *mapBuilder) addTrans(srcState int, input uint64, dstState int) {
	if _, ok := this.trans[srcState]; !ok {
		this.trans[srcState] = make(map[uint64]int)
	}
	if _, ok := this.trans[srcState][input]; ok {
		panic("transition already exists")
	}
	this.trans[srcState][input] = dstState
}

func (this *mapBuilder) included(pkg string, msg string, field string) bool {
	if this.includes == nil {
		return true
	}
	for _, inc := range this.includes {
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

func (this *mapBuilder) visit(srcPackage string, srcMessage string) error {
	if !this.included(srcPackage, srcMessage, "") {
		return nil
	}
	msg := this.desc.GetMessage(srcPackage, srcMessage)
	if msg == nil {
		return &errUnknown{srcPackage, srcMessage, ""}
	}
	currentState := this.getState(srcPackage, srcMessage, "")
	this.visited[currentState] = true
	for _, f := range msg.GetField() {
		if f.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			newPackage, newMessage := this.desc.FindMessage(srcPackage, srcMessage, f.GetName())
			if len(newMessage) == 0 {
				return &errUnknown{srcPackage, srcMessage, f.GetName()}
			}
			dstState := this.getState(newPackage, newMessage, "")
			this.addTrans(currentState, f.GetKeyUint64(), dstState)
			if _, ok := this.visited[dstState]; !ok {
				if err := this.visit(newPackage, newMessage); err != nil {
					return err
				}
			}
		} else {
			if !this.included(srcPackage, srcMessage, f.GetName()) {
				continue
			}
			dstState := this.getState(srcPackage, srcMessage, f.GetName())
			this.addTrans(currentState, f.GetKeyUint64(), dstState)
		}
	}
	return nil
}
