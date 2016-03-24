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

func newOpenParen() *Keyword {
	return &Keyword{
		Value: "(",
	}
}

func newCloseParen() *Keyword {
	return &Keyword{
		Value: ")",
	}
}

func newEqualEqual() *Keyword {
	return &Keyword{
		Value: "==",
	}
}

func newExclamationEqual() *Keyword {
	return &Keyword{
		Value: "!=",
	}
}

func newLessThan() *Keyword {
	return &Keyword{
		Value: "<",
	}
}

func newGreaterThan() *Keyword {
	return &Keyword{
		Value: ">",
	}
}

func newLessEqual() *Keyword {
	return &Keyword{
		Value: "<=",
	}
}

func newGreaterEqual() *Keyword {
	return &Keyword{
		Value: ">=",
	}
}

func newTildeEqual() *Keyword {
	return &Keyword{
		Value: "~=",
	}
}

func newStarEqual() *Keyword {
	return &Keyword{
		Value: "*=",
	}
}

func newCaretEqual() *Keyword {
	return &Keyword{
		Value: "^=",
	}
}

func newDollarEqual() *Keyword {
	return &Keyword{
		Value: "*=",
	}
}

func newColonColon() *Keyword {
	return &Keyword{
		Value: "::",
	}
}

func newUnderscore() *Keyword {
	return &Keyword{
		Value: "_",
	}
}

func newExclamation() *Keyword {
	return &Keyword{
		Value: "!",
	}
}

func newPipe() *Keyword {
	return &Keyword{
		Value: "|",
	}
}
