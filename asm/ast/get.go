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

func (rs *Rules) HasRoot() bool {
	for _, r := range rs.Rules {
		if r.Root != nil {
			return true
		}
	}
	return false
}

func (rs *Rules) GetRoot() *Root {
	for _, r := range rs.Rules {
		if r.Root != nil {
			return r.Root
		}
	}
	return nil
}

func (rs *Rules) HasInit() bool {
	for _, r := range rs.Rules {
		if r.Init != nil {
			return true
		}
	}
	return false
}

func (rs *Rules) GetInit() *Init {
	for _, r := range rs.Rules {
		if r.Init != nil {
			return r.Init
		}
	}
	return nil
}

func (rs *Rules) GetTransitions() []*Transition {
	trans := make([]*Transition, 0, len(rs.Rules))
	for i, r := range rs.Rules {
		if r.Transition != nil {
			trans = append(trans, rs.Rules[i].Transition)
		}
	}
	return trans
}

func (rs *Rules) GetFunctionDecls() []*FunctionDecl {
	fs := make([]*FunctionDecl, 0, len(rs.Rules))
	for i, r := range rs.Rules {
		if r.FunctionDecl != nil {
			fs = append(fs, rs.Rules[i].FunctionDecl)
		}
	}
	return fs
}

func (rs *Rules) GetFinals() []*Final {
	finals := make([]*Final, 0, len(rs.Rules))
	for i, r := range rs.Rules {
		if r.Final != nil {
			finals = append(finals, rs.Rules[i].Final)
		}
	}
	return finals
}
