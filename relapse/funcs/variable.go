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

package funcs

import (
	"github.com/katydid/katydid/parser"
)

//Variable is an interface that when implemented represents a value that can change between different executions of the same function.
type Variable interface {
	IsVariable()
}

type aVariable interface {
	isVariable()
}

//Setter is an interface that represents a variable of which the value must be set.
type Setter interface {
	SetValue(parser.Value)
}
