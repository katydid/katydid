//  Copyright 2017 Walter Schulze
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

package intern

import (
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
)

func (c *construct) SetContext(context *funcs.Context) {
	c.context = context
	for _, p := range c.refs {
		c.setContext(p)
	}
}

func (c *construct) setContext(p *Pattern) {
	switch p.Type {
	case Empty:
		return
	case Node:
		compose.SetContext(p.Func, c.context)
		c.setContext(p.Patterns[0])
	case Concat:
		for _, pp := range p.Patterns {
			c.setContext(pp)
		}
	case Or:
		for _, pp := range p.Patterns {
			c.setContext(pp)
		}
	case And:
		for _, pp := range p.Patterns {
			c.setContext(pp)
		}
	case ZeroOrMore:
		return
	case Reference:
		return
	case Not:
		c.setContext(p.Patterns[0])
	case Contains:
		c.setContext(p.Patterns[0])
	case Optional:
		c.setContext(p.Patterns[0])
	case Interleave:
		for _, pp := range p.Patterns {
			c.setContext(pp)
		}
	}
}
