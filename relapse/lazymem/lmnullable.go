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

package lazymem

import (
	"fmt"
)

func Nullable(p *Pattern) bool {
	visited := make(map[*Pattern]struct{})
	changed := true
	for changed {
		changed = nullable(visited, p)
	}
	return p.GetNullable()
}

func nullable(visited map[*Pattern]struct{}, p *Pattern) (changed bool) {
	if _, ok := visited[p]; ok {
		return false
	}
	visited[p] = struct{}{}
	head := p.Head().GetValue()
	switch v := head.(type) {
	case *Empty:
		changed = true
		p.SetNullable(true)
	case *Node:
		changed = true
		p.SetNullable(false)
	case *Concat:
		leftChanged := nullable(visited, v.Left)
		rightChanged := nullable(visited, v.Right)
		changed = leftChanged || rightChanged
		p.SetNullable(v.Left.GetNullable() && v.Right.GetNullable())
	case *Or:
		leftChanged := nullable(visited, v.Left)
		rightChanged := nullable(visited, v.Right)
		changed = leftChanged || rightChanged
		p.SetNullable(v.Left.GetNullable() || v.Right.GetNullable())
	case *And:
		leftChanged := nullable(visited, v.Left)
		rightChanged := nullable(visited, v.Right)
		changed = leftChanged || rightChanged
		p.SetNullable(v.Left.GetNullable() && v.Right.GetNullable())
	case *ZeroOrMore:
		changed = true
		p.SetNullable(true)
	case *Not:
		changed = nullable(visited, v.Pattern)
		if v.Pattern.HasNullable() {
			p.SetNullable(!v.Pattern.GetNullable())
		}
	case *ZAny:
		changed = true
		p.SetNullable(true)
	case *Contains:
		changed = nullable(visited, v.Pattern)
		if v.Pattern.HasNullable() {
			p.SetNullable(true)
		}
	case *Optional:
		changed = true
		p.SetNullable(true)
	case *Interleave:
		leftChanged := nullable(visited, v.Left)
		rightChanged := nullable(visited, v.Right)
		changed = leftChanged || rightChanged
		p.SetNullable(v.Left.GetNullable() && v.Right.GetNullable())
	default:
		panic(fmt.Sprintf("unknown pattern typ %T", head))
	}
	return changed
}
