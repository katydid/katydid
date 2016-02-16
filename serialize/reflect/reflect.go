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

package reflect

import (
	"github.com/katydid/katydid/serialize"
	"io"
	"reflect"
)

type state struct {
	parent   reflect.Value
	typ      reflect.StructField
	value    reflect.Value
	field    int
	maxField int
	isLeaf   bool
	isArray  bool
}

type parser struct {
	state
	stack []state
}

func deref(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

func newState(val reflect.Value) state {
	value := deref(val)
	if value.Kind() == reflect.Struct {
		return state{
			parent:   value,
			maxField: value.NumField(),
		}
	}
	if isSlice(value) {
		return state{
			parent:   value,
			maxField: value.Len(),
			isArray:  true,
		}
	}
	return state{
		value:    val,
		isLeaf:   true,
		maxField: 1,
	}
}

func isSlice(v reflect.Value) bool {
	return v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8
}

func NewReflectParser() *parser {
	return &parser{stack: make([]state, 0, 10)}
}

func (s *parser) Init(value reflect.Value) *parser {
	s.state = newState(value)
	return s
}

func (s *parser) Next() error {
	if s.field >= s.maxField {
		return io.EOF
	}
	if !s.isLeaf && !s.isArray {
		s.typ = s.parent.Type().Field(s.field)
		s.value = s.parent.Field(s.field)
		if s.value.Kind() == reflect.Ptr || s.value.Kind() == reflect.Slice {
			if s.value.IsNil() {
				s.field++
				return s.Next()
			}
		}
	}
	s.field++
	return nil
}

func (s *parser) IsLeaf() bool {
	return s.isLeaf
}

func (s *parser) getValue() reflect.Value {
	return deref(s.value)
}

func (s *parser) Double() (float64, error) {
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Float64, reflect.Float32:
			return value.Float(), nil
		}
	}
	return 0, serialize.ErrNotDouble
}

func (s *parser) Int() (int64, error) {
	if s.isArray {
		return int64(s.field - 1), nil
	}
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Int64, reflect.Int32:
			return value.Int(), nil
		}
	}
	return 0, serialize.ErrNotInt
}

func (s *parser) Uint() (uint64, error) {
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Uint64, reflect.Uint32:
			return value.Uint(), nil
		}
	}
	return 0, serialize.ErrNotUint
}

func (s *parser) Bool() (bool, error) {
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Bool:
			return value.Bool(), nil
		}
	}
	return false, serialize.ErrNotBool
}

func (s *parser) String() (string, error) {
	if !s.isLeaf {
		return s.typ.Name, nil
	}
	value := s.getValue()
	switch value.Kind() {
	case reflect.String:
		return value.String(), nil
	}
	return "", serialize.ErrNotString
}

func (s *parser) Bytes() ([]byte, error) {
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Slice, reflect.Uint8, reflect.Int8:
			return value.Bytes(), nil
		}
	}
	return nil, serialize.ErrNotBytes
}

func (s *parser) Up() {
	top := len(s.stack) - 1
	s.state = s.stack[top]
	s.stack = s.stack[:top]
}

func (s *parser) Down() {
	s.stack = append(s.stack, s.state)
	if s.isArray {
		s.state = newState(s.state.parent.Index(s.field - 1))
	} else {
		s.state = newState(s.state.value)
	}
}
