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

package debug

import (
	"fmt"
	"github.com/katydid/katydid/serialize"
	"io"
)

func getValue(parser serialize.Parser) interface{} {
	var v interface{}
	var err error
	v, err = parser.Int()
	if err == nil {
		return v
	}
	v, err = parser.Uint()
	if err == nil {
		return v
	}
	v, err = parser.Double()
	if err == nil {
		return v
	}
	v, err = parser.Bool()
	if err == nil {
		return v
	}
	v, err = parser.String()
	if err == nil {
		return v
	}
	v, err = parser.Bytes()
	if err == nil {
		return v
	}
	return nil
}

func Walk(parser serialize.Parser) interface{} {
	m := make(map[string][]interface{})
	a := make([]interface{}, 0)
	for {
		if err := parser.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		value := getValue(parser)
		if parser.IsLeaf() {
			a = append(a, value)
		} else {
			name := fmt.Sprintf("%s", value)
			parser.Down()
			v := Walk(parser)
			parser.Up()
			_, ok := m[name]
			if !ok {
				m[name] = make([]interface{}, 0, 1)
			}
			if v != nil {
				m[name] = append(m[name], v)
			}
		}
	}
	if len(m) > 0 {
		return m
	}
	if len(a) == 1 {
		return a[0]
	}
	return a
}
