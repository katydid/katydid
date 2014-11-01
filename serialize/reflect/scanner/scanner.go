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

package scanner

import (
	"github.com/katydid/katydid/serialize"
	"io"
	"reflect"
)

type Tokens interface {
	Lookup(parentToken int, name []byte) int
}

type state struct {
	parentToken   int
	parent        reflect.Value
	typ           reflect.StructField
	value         reflect.Value
	field         int
	maxField      int
	tokenId       int
	sliceValue    reflect.Value
	sliceIndex    int
	sliceMaxIndex int
	inSlice       bool
}

type scanner struct {
	tokens Tokens
	state
	stack []state
}

func deref(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

func newState(parent reflect.Value, token int) state {
	value := deref(parent)
	return state{
		parentToken: token,
		parent:      value,
		maxField:    value.NumField(),
	}
}

func NewScanner(tokens Tokens, rootToken int) *scanner {
	return &scanner{stack: make([]state, 0, 10), tokens: tokens, state: state{parentToken: rootToken}}
}

func (s *scanner) Init(value reflect.Value) *scanner {
	s.state = newState(value, s.parentToken)
	return s
}

func isSlice(v reflect.Value) bool {
	return v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8
}

func (s *scanner) Next() error {
	if s.field >= s.maxField {
		return io.EOF
	}
	if s.inSlice {
		if s.sliceIndex >= s.sliceMaxIndex {
			s.inSlice = false
			s.field++
			return s.Next()
		}
		s.value = s.sliceValue.Index(s.sliceIndex)
		s.sliceIndex++
		return nil
	}
	s.typ = s.parent.Type().Field(s.field)
	s.value = s.parent.Field(s.field)
	s.tokenId = s.tokens.Lookup(s.parentToken, []byte(s.typ.Name))
	if s.value.Kind() == reflect.Ptr || s.value.Kind() == reflect.Slice {
		if s.value.IsNil() {
			s.field++
			return s.Next()
		}
	}
	if isSlice(s.value) {
		s.sliceValue = s.value
		s.sliceMaxIndex = s.value.Len()
		s.sliceIndex = 0
		s.inSlice = true
		return s.Next()
	}
	s.field++
	return nil
}

func (s *scanner) IsLeaf() bool {
	isLeaf := true
	if s.typ.Type.Kind() == reflect.Struct {
		isLeaf = false
	}
	if s.typ.Type.Kind() == reflect.Ptr {
		if s.typ.Type.Elem().Kind() == reflect.Struct {
			isLeaf = false
		}
	}
	if s.typ.Type.Kind() == reflect.Slice {
		if s.typ.Type.Elem().Kind() == reflect.Struct {
			isLeaf = false
		}
		if s.typ.Type.Elem().Kind() == reflect.Ptr {
			if s.typ.Type.Elem().Elem().Kind() == reflect.Struct {
				isLeaf = false
			}
		}
	}
	return isLeaf
}

func (s *scanner) getValue() reflect.Value {
	kind := s.value.Kind()
	if kind == reflect.Slice {
		childValue := s.value.Elem()
		childKind := childValue.Kind()
		switch childKind {
		case reflect.Uint8, reflect.Int8:
			return s.value
		case reflect.Slice:
			switch childValue.Elem().Kind() {
			case reflect.Uint8, reflect.Int8:
				return childValue
			}
		default:
			return childValue
		}
	}
	if kind == reflect.Ptr {
		return s.value.Elem()
	}
	return s.value
}

func (s *scanner) Float64() (float64, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.Float64:
		return value.Float(), nil
	}
	return 0, serialize.ErrNotFloat64
}

func (s *scanner) Float32() (float32, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.Float32:
		return float32(value.Float()), nil
	}
	return 0, serialize.ErrNotFloat32
}

func (s *scanner) Int64() (int64, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.Int64:
		return value.Int(), nil
	}
	return 0, serialize.ErrNotInt64
}

func (s *scanner) Uint64() (uint64, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.Uint64:
		return value.Uint(), nil
	}
	return 0, serialize.ErrNotUint64
}

func (s *scanner) Int32() (int32, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.Int32:
		return int32(value.Int()), nil
	}
	return 0, serialize.ErrNotInt32
}

func (s *scanner) Bool() (bool, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.Bool:
		return value.Bool(), nil
	}
	return false, serialize.ErrNotBool
}

func (s *scanner) String() (string, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.String:
		return value.String(), nil
	}
	return "", serialize.ErrNotString
}

func (s *scanner) Bytes() ([]byte, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.Slice, reflect.Uint8, reflect.Int8:
		return value.Bytes(), nil
	}
	return nil, serialize.ErrNotBytes
}

func (s *scanner) Uint32() (uint32, error) {
	value := s.getValue()
	switch value.Kind() {
	case reflect.Uint32:
		return uint32(value.Uint()), nil
	}
	return 0, serialize.ErrNotUint32
}

func (s *scanner) Id() int {
	return s.tokenId
}

func (s *scanner) Up() {
	top := len(s.stack) - 1
	s.state = s.stack[top]
	s.stack = s.stack[:top]
}

func (s *scanner) Down() {
	s.stack = append(s.stack, s.state)
	s.state = newState(s.state.value, s.tokenId)
}
