//  Copyright 2016 Walter Schulze
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

//Package reflect encodes a parser.Interface into a reflected go structure.
//This can be used to unmarshal or copy.
package reflect

import (
	"fmt"
	"github.com/katydid/katydid/parser"
	"io"
	"reflect"
)

//Encode encodes a parser.Interface into a go structure value.
func Encode(p parser.Interface, v interface{}) error {
	r := reflect.ValueOf(v)
	return encodeStruct(p, r)
}

func encodeStruct(p parser.Interface, v reflect.Value) error {
	kind := v.Type().Kind()
	isPtr := false
	if kind == reflect.Ptr {
		kind = v.Type().Elem().Kind()
		isPtr = true
	}
	if kind != reflect.Struct {
		return fmt.Errorf("expected struct")
	}
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				return nil
			} else {
				panic(err)
			}
		}
		if p.IsLeaf() {
			return fmt.Errorf("struct: did not expect leaf")
		}
		name, err := p.String()
		if err != nil {
			return err
		}
		strct := v
		if isPtr {
			strct = v.Elem()
		}
		field := strct.FieldByName(name)
		if !field.IsValid() {
			continue //skip field
		}
		fieldKind := field.Kind()
		fieldIsPtr := false
		fieldType := field.Type()
		if fieldKind == reflect.Ptr {
			fieldIsPtr = true
			fieldKind = field.Type().Elem().Kind()
			fieldType = field.Type().Elem()
		}
		p.Down()
		switch fieldKind {
		case reflect.Struct:
			if !p.IsLeaf() {
				if fieldIsPtr {
					field.Set(reflect.New(field.Type().Elem()))
				}
				if err := encodeStruct(p, field); err != nil {
					return err
				}
			}
		case reflect.Slice:
			list, err := newList(p, fieldType)
			if err != nil {
				return err
			}
			field.Set(list)
		default:

			value, err := newValue(p, fieldType)
			if err != nil {
				return err
			}
			field.Set(value)
		}
		p.Up()
	}
	panic("unreachable")
}

func newList(p parser.Interface, typ reflect.Type) (reflect.Value, error) {
	list := reflect.MakeSlice(typ, 0, 0)
	for {
		if err := p.Next(); err != nil {
			if err == io.EOF {
				return list, nil
			} else {
				panic(err)
			}
		}
		if p.IsLeaf() {
			return reflect.ValueOf(nil), fmt.Errorf("list: did not expect leaf")
		}
		_, err := p.Int()
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		elemType := typ.Elem()
		elemKind := elemType.Kind()
		elemIsPtr := false
		if elemKind == reflect.Ptr {
			elemIsPtr = true
			elemType = typ.Elem().Elem()
			elemKind = elemType.Kind()
		}
		p.Down()
		switch elemKind {
		case reflect.Struct:
			if !p.IsLeaf() {
				elem := reflect.New(elemType).Elem()
				if err := encodeStruct(p, elem); err != nil {
					return reflect.ValueOf(nil), err
				}
				if elemIsPtr {
					elem = elem.Addr()
				}
				list = reflect.Append(list, elem)
			} else {
				list = reflect.Append(list, reflect.Zero(typ.Elem()))
			}
		case reflect.Slice:
			newList, err := newList(p, elemType)
			if err != nil {
				return reflect.ValueOf(nil), err
			}
			list = reflect.Append(list, newList)
		default:
			elem, err := newValue(p, typ.Elem())
			if err != nil {
				return reflect.ValueOf(nil), err
			}
			list = reflect.Append(list, elem)
		}
		p.Up()
	}
	panic("unreachable")
}

func newValue(p parser.Interface, typ reflect.Type) (reflect.Value, error) {
	if err := p.Next(); err != nil {
		if err == io.EOF {
			return reflect.ValueOf(nil), nil
		} else {
			panic(err)
		}
	}
	if !p.IsLeaf() {
		return reflect.ValueOf(nil), fmt.Errorf("expected leaf")
	}
	if value, err := p.Int(); err == nil {
		return reflect.ValueOf(value), nil
	}
	if value, err := p.Uint(); err == nil {
		return reflect.ValueOf(value), nil
	}
	if value, err := p.Double(); err == nil {
		return reflect.ValueOf(value), nil
	}
	if value, err := p.Bool(); err == nil {
		return reflect.ValueOf(value), nil
	}
	if value, err := p.String(); err == nil {
		return reflect.ValueOf(value), nil
	}
	if value, err := p.Bytes(); err == nil {
		return reflect.ValueOf(value), nil
	}
	return reflect.ValueOf(nil), nil
}
