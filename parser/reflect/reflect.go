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

//Package reflect contains an implementation of a parser for a reflected go structure.
package reflect

import (
	"io"
	"reflect"

	"github.com/katydid/katydid/parser"
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

type reflectParser struct {
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

//ReflectParser is a parser for a reflected go structure.
type ReflectParser interface {
	parser.Interface
	//Init initialises the parser with a value of reflected go structure.
	Init(value reflect.Value) ReflectParser
	Reset() error
}

//NewReflectParser returns a new reflect parser.
func NewReflectParser() ReflectParser {
	return &reflectParser{stack: make([]state, 0, 10)}
}

func (s *reflectParser) Init(value reflect.Value) ReflectParser {
	s.state = newState(value)
	return s
}

func (s *reflectParser) Reset() error {
	s.stack = s.stack[:0]
	s.state = newState(s.state.value)
	return nil
}

func (s *reflectParser) Next() error {
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

func (s *reflectParser) IsLeaf() bool {
	return s.isLeaf
}

func (s *reflectParser) getValue() reflect.Value {
	return deref(s.value)
}

func (s *reflectParser) Double() (float64, error) {
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Float64, reflect.Float32:
			return value.Float(), nil
		}
	}
	return 0, parser.ErrNotDouble
}

func (s *reflectParser) Int() (int64, error) {
	if s.isArray {
		return int64(s.field - 1), nil
	}
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Int64, reflect.Int32, reflect.Int:
			return value.Int(), nil
		}
	}
	return 0, parser.ErrNotInt
}

func (s *reflectParser) Uint() (uint64, error) {
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Uint64, reflect.Uint32:
			return value.Uint(), nil
		}
	}
	return 0, parser.ErrNotUint
}

func (s *reflectParser) Bool() (bool, error) {
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Bool:
			return value.Bool(), nil
		}
	}
	return false, parser.ErrNotBool
}

func (s *reflectParser) String() (string, error) {
	if !s.isLeaf {
		return s.typ.Name, nil
	}
	value := s.getValue()
	switch value.Kind() {
	case reflect.String:
		return value.String(), nil
	}
	return "", parser.ErrNotString
}

func (s *reflectParser) Bytes() ([]byte, error) {
	if s.isLeaf {
		value := s.getValue()
		switch value.Kind() {
		case reflect.Slice, reflect.Uint8, reflect.Int8:
			return value.Bytes(), nil
		}
	}
	return nil, parser.ErrNotBytes
}

func (s *reflectParser) Up() {
	top := len(s.stack) - 1
	s.state = s.stack[top]
	s.stack = s.stack[:top]
}

func (s *reflectParser) Down() {
	s.stack = append(s.stack, s.state)
	if s.isArray {
		s.state = newState(s.state.parent.Index(s.field - 1))
	} else {
		s.state = newState(s.state.value)
	}
}
