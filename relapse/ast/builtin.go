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

package ast

//BuiltInFunctionName returns the full function name for the given builtin symbol.
//The function returns an empty string when the symbol is not a known builtin symbol.
func BuiltInFunctionName(symbol string) string {
	switch symbol {
	case "==":
		return "eq"
	case "!=":
		return "ne"
	case "<":
		return "lt"
	case ">":
		return "gt"
	case "<=":
		return "le"
	case ">=":
		return "ge"
	case "~=":
		return "regex"
	case "*=":
		return "contains"
	case "^=":
		return "hasPrefix"
	case "$=":
		return "hasSuffix"
	case "::":
		return "type"
	}
	return ""
}

// FunctionNameToBuiltIn returns the builtin constructor for a given function name.
// If a symbol does not exist, nil is returned.
func FunctionNameToBuiltIn(symbol string) func(*Expr) *Expr {
	switch symbol {
	case "eq":
		return NewEqual
	case "ne":
		return NewNotEqual
	case "lt":
		return NewLessThan
	case "gt":
		return NewGreaterThan
	case "le":
		return NewLessEqual
	case "ge":
		return NewGreaterEqual
	case "regex":
		return NewRegex
	case "contains":
		return NewHasElem
	case "hasPrefix":
		return NewHasPrefix
	case "hasSuffix":
		return NewHasSuffix
	case "type":
		return NewType
	}
	return nil
}

//NewEqual returns an builtin equal expression.
//  == e
func NewEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newEqualEqual(),
			Expr:   e,
		},
	}
}

//NewNotEqual returns an builtin not equal expression.
//  != e
func NewNotEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newExclamationEqual(),
			Expr:   e,
		},
	}
}

//NewLessThan returns an builtin less than expression.
//  < e
func NewLessThan(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newLessThan(),
			Expr:   e,
		},
	}
}

//NewGreaterThan returns an builtin greater than expression.
//  > e
func NewGreaterThan(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newGreaterThan(),
			Expr:   e,
		},
	}
}

//NewLessEqual returns an builtin less than or equal expression.
//  <= e
func NewLessEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newLessEqual(),
			Expr:   e,
		},
	}
}

//NewGreaterEqual returns an builtin greater than or equal expression.
//  >= e
func NewGreaterEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newGreaterEqual(),
			Expr:   e,
		},
	}
}

//NewRegex returns an builtin regular expression.
//  ~= e
func NewRegex(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newTildeEqual(),
			Expr:   e,
		},
	}
}

//NewHasElem returns an builtin contains expression.
//  *= e
func NewHasElem(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newStarEqual(),
			Expr:   e,
		},
	}
}

//NewHasPrefix returns an builtin has prefix expression.
//  ^= e
func NewHasPrefix(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newCaretEqual(),
			Expr:   e,
		},
	}
}

//NewHasSuffix returns an builtin has suffix expression.
//  $= e
func NewHasSuffix(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newDollarEqual(),
			Expr:   e,
		},
	}
}

//NewType returns an builtin type expression.
//  :: e
func NewType(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newColonColon(),
			Expr:   e,
		},
	}
}
