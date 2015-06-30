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

package expr

import (
	"errors"
)

func (this *Expr) GetTerminalVariable() (*Variable, error) {
	if this.Terminal != nil {
		return this.GetTerminal().GetVariable(), nil
	}
	if this.List != nil {
		return this.GetList().GetTerminalVariable()
	}
	return this.GetFunction().GetTerminalVariable()
}

func (this *List) GetTerminalVariable() (*Variable, error) {
	var res *Variable
	for _, e := range this.GetElems() {
		if v, err := e.GetTerminalVariable(); err != nil {
			return nil, err
		} else if v != nil {
			if res != nil {
				return nil, errors.New("Two Terminal Variables")
			}
			res = v
		}
	}
	return res, nil
}

func (this *Function) GetTerminalVariable() (*Variable, error) {
	var res *Variable
	for _, p := range this.GetParams() {
		if v, err := p.GetTerminalVariable(); err != nil {
			return nil, err
		} else if v != nil {
			if res != nil {
				return nil, errors.New("Two Terminal Variables")
			}
			res = v
		}
	}
	return res, nil
}
