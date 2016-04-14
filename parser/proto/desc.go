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

package proto

import (
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"strings"
)

type errUnknown struct {
	msg string
}

func (this *errUnknown) Error() string {
	return "Could not find " + this.msg
}

//DescMap is a map of the descriptor.FileDescriptorSet
type DescMap interface {
	//GetRoot returns the root message that was used to create this map.
	GetRoot() *descriptor.DescriptorProto
	//LookupMessage returns the message descriptor of the field type.
	LookupMessage(field *descriptor.FieldDescriptorProto) *descriptor.DescriptorProto
	//LookupField returns the field in the message which has the corresponding key.
	LookupField(msg *descriptor.DescriptorProto, key uint64) (*descriptor.FieldDescriptorProto, bool)
}

type descMap struct {
	desc       *descriptor.FileDescriptorSet
	root       *descriptor.DescriptorProto
	fieldToMsg map[*descriptor.FieldDescriptorProto]*descriptor.DescriptorProto
	msgToField map[*descriptor.DescriptorProto]map[uint64]*descriptor.FieldDescriptorProto
}

//NewDescriptorMap returns a map of the FileDescriptorSet starting at the message represented by the package name and message name.
func NewDescriptorMap(pkgName, msgName string, desc *descriptor.FileDescriptorSet) (DescMap, error) {
	root := desc.GetMessage(pkgName, msgName)
	if root == nil {
		return nil, &errUnknown{pkgName + "." + msgName}
	}
	d := &descMap{
		desc:       desc,
		root:       root,
		fieldToMsg: make(map[*descriptor.FieldDescriptorProto]*descriptor.DescriptorProto),
		msgToField: make(map[*descriptor.DescriptorProto]map[uint64]*descriptor.FieldDescriptorProto),
	}
	err := d.visit(pkgName, root)
	return d, err
}

func dotToUnderscore(r rune) rune {
	if r == '.' {
		return '_'
	}
	return r
}

func (this *descMap) findExts(pkgName string, msgName string) []*descriptor.FieldDescriptorProto {
	exts := []*descriptor.FieldDescriptorProto{}
	extendee := "." + pkgName + "." + msgName
	for _, file := range this.desc.GetFile() {
		for i, ext := range file.GetExtension() {
			if strings.Map(dotToUnderscore, file.GetPackage()) == strings.Map(dotToUnderscore, pkgName) {
				if !(ext.GetExtendee() == msgName || ext.GetExtendee() == extendee) {
					continue
				}
			} else {
				if ext.GetExtendee() != extendee {
					continue
				}
			}
			exts = append(exts, file.Extension[i])
		}
	}
	return exts
}

func (this *descMap) visit(pkgName string, msg *descriptor.DescriptorProto) error {
	if _, ok := this.msgToField[msg]; ok {
		return nil
	}
	fields := msg.GetField()
	if msg.HasExtension() {
		exts := this.findExts(pkgName, msg.GetName())
		fields = append(fields, exts...)
	}
	for i, f := range fields {
		if _, ok := this.msgToField[msg]; !ok {
			this.msgToField[msg] = make(map[uint64]*descriptor.FieldDescriptorProto)
		}
		this.msgToField[msg][f.GetKeyUint64()] = fields[i]
		if f.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			newPkgName, newMsgName := this.desc.FindMessage(pkgName, msg.GetName(), f.GetName())
			if len(newMsgName) == 0 {
				return &errUnknown{pkgName + "." + msg.GetName() + "." + f.GetName()}
			}
			newMsg := this.desc.GetMessage(newPkgName, newMsgName)
			if newMsg == nil {
				return &errUnknown{newPkgName + "." + newMsgName}
			}
			this.fieldToMsg[fields[i]] = newMsg
			if _, ok := this.msgToField[newMsg]; ok {
				continue
			}
			if err := this.visit(newPkgName, newMsg); err != nil {
				return err
			}
		}
	}
	for i := range msg.GetNestedType() {
		this.visit(pkgName, msg.GetNestedType()[i])
	}
	return nil
}

func (this *descMap) GetRoot() *descriptor.DescriptorProto {
	return this.root
}

func (this *descMap) LookupMessage(field *descriptor.FieldDescriptorProto) *descriptor.DescriptorProto {
	return this.fieldToMsg[field]
}

func (this *descMap) LookupField(msg *descriptor.DescriptorProto, key uint64) (*descriptor.FieldDescriptorProto, bool) {
	fields, ok := this.msgToField[msg]
	if !ok {
		return nil, false
	}
	f, ok := fields[key]
	return f, ok
}
