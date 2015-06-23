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

package asm

import (
	"github.com/katydid/katydid/expr/ast"
)

func (this *Rule) GetAttachedComment() expr.Comment {
	if this.Root != nil {
		return this.Root.GetAttachedComment()
	}
	if this.Init != nil {
		return this.Init.GetAttachedComment()
	}
	if this.Final != nil {
		return this.Final.GetAttachedComment()
	}
	if this.Transition != nil {
		return this.Transition.GetAttachedComment()
	}
	if this.FunctionDecl != nil {
		return this.FunctionDecl.GetAttachedComment()
	}
	panic("unreachable")
}

func (this *Root) GetAttachedComment() expr.Comment {
	return this.Before.GetAttachedComment()
}

func (this *Init) GetAttachedComment() expr.Comment {
	return this.Before.GetAttachedComment()
}

func (this *Transition) GetAttachedComment() expr.Comment {
	return this.Before.GetAttachedComment()
}

func (this *FunctionDecl) GetAttachedComment() expr.Comment {
	return this.Before.GetAttachedComment()
}

func (this *Final) GetAttachedComment() expr.Comment {
	return this.Before.GetAttachedComment()
}
