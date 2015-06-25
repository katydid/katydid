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

package serialize

type Decoder interface {
	Float64() (float64, error)
	Int64() (int64, error)
	Uint64() (uint64, error)
	Bool() (bool, error)
	String() (string, error)
	Bytes() ([]byte, error)
}

type errValue struct{}

type Scanner interface {
	Copy() Scanner
	Next() error
	IsLeaf() bool
	Name() string
	Up()
	Down()
	Decoder
}
