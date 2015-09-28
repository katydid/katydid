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
	"github.com/katydid/katydid/serialize"
	"io"
)

func getValue(parser serialize.Parser) interface{} {
	var v interface{}
	var err error
	v, err = parser.Bool()
	if err == nil {
		return v
	}
	v, err = parser.Bytes()
	if err == nil {
		return v
	}
	v, err = parser.String()
	if err == nil {
		return v
	}
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
	return nil
}

func Walk(parser serialize.Parser) map[string][]interface{} {
	m := make(map[string][]interface{})
	for {
		if err := parser.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		name := parser.Name()
		var v interface{}
		if !parser.IsLeaf() {
			parser.Down()
			mm := Walk(parser)
			parser.Up()
			v = mm
		} else {
			v = getValue(parser)
		}
		_, ok := m[name]
		if !ok {
			m[name] = make([]interface{}, 0, 1)
		}
		if v != nil {
			m[name] = append(m[name], v)
		}
	}
	return m
}
