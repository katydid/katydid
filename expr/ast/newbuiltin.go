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

package expr

func NewEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newEqualEqual(),
			Expr:   e,
		},
	}
}

func NewNotEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newExclamationEqual(),
			Expr:   e,
		},
	}
}

func NewLessThan(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newLessThan(),
			Expr:   e,
		},
	}
}

func NewGreaterThan(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newGreaterThan(),
			Expr:   e,
		},
	}
}

func NewLessEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newLessEqual(),
			Expr:   e,
		},
	}
}

func NewGreaterEqual(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newGreaterEqual(),
			Expr:   e,
		},
	}
}

func NewRegex(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newTildeEqual(),
			Expr:   e,
		},
	}
}

func NewContains(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newStarEqual(),
			Expr:   e,
		},
	}
}

func NewHasPrefix(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newCaretEqual(),
			Expr:   e,
		},
	}
}

func NewHasSuffix(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newDollarEqual(),
			Expr:   e,
		},
	}
}

func NewType(e *Expr) *Expr {
	return &Expr{
		BuiltIn: &BuiltIn{
			Symbol: newColonColon(),
			Expr:   e,
		},
	}
}
