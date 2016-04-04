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

package relapse

//NewEqual returns an builtin equal expression.
func NewEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newEqualEqual(),
			Expr:   e,
		},
	}
}

//NewNotEqual returns an builtin not equal expression.
func NewNotEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newExclamationEqual(),
			Expr:   e,
		},
	}
}

//NewLessThan returns an builtin less than expression.
func NewLessThan(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newLessThan(),
			Expr:   e,
		},
	}
}

//NewGreaterThan returns an builtin greater than expression.
func NewGreaterThan(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newGreaterThan(),
			Expr:   e,
		},
	}
}

//NewLessEqual returns an builtin less than or equal expression.
func NewLessEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newLessEqual(),
			Expr:   e,
		},
	}
}

//NewGreaterEqual returns an builtin greater than or equal expression.
func NewGreaterEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newGreaterEqual(),
			Expr:   e,
		},
	}
}

//NewRegex returns an builtin regular expression.
func NewRegex(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newTildeEqual(),
			Expr:   e,
		},
	}
}

//NewHasElem returns an builtin contains expression.
func NewHasElem(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newStarEqual(),
			Expr:   e,
		},
	}
}

//NewHasPrefix returns an builtin has prefix expression.
func NewHasPrefix(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newCaretEqual(),
			Expr:   e,
		},
	}
}

//NewHasSuffix returns an builtin has suffix expression.
func NewHasSuffix(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newDollarEqual(),
			Expr:   e,
		},
	}
}

//NewType returns an builtin type expression.
func NewType(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newColonColon(),
			Expr:   e,
		},
	}
}
