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

package parser

import "fmt"

//A type conforming to the parser.Interface interface, abstracts away the implementation details of a parser.
//This allows the parser to be used by other katydid functions, like relapse validation.
type Interface interface {
	Next() error
	IsLeaf() bool
	Up()
	Down()
	Value
}

//A type confirming to the parser.Value interface, repesents one native value, tree node label (field name) or some repesentation a node label.
//Typically only one of the methods returns a value without an error, but more than one method can return without an error.
//For example a positive json number can return an errorless value for the Double, Int and Uint methods.
type Value interface {
	Double() (float64, error)
	Int() (int64, error)
	Uint() (uint64, error)
	Bool() (bool, error)
	String() (string, error)
	Bytes() ([]byte, error)
}

//Sprint returns a value printed as a string.
func Sprint(value Value) string {
	return fmt.Sprintf("%#v", getValue(value))
}

func getValue(value Value) interface{} {
	var v interface{}
	var err error
	v, err = value.Bool()
	if err == nil {
		return v
	}
	v, err = value.Bytes()
	if err == nil {
		return v
	}
	v, err = value.String()
	if err == nil {
		return v
	}
	v, err = value.Int()
	if err == nil {
		return v
	}
	v, err = value.Uint()
	if err == nil {
		return v
	}
	v, err = value.Double()
	if err == nil {
		return v
	}
	return nil
}
