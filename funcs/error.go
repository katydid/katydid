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
	"fmt"
	"runtime/debug"
	"strings"
)

type SetCatcher interface {
	SetCatcher(c Catcher)
}

type Thrower struct {
	catcher Catcher
}

func (this *Thrower) SetCatcher(c Catcher) {
	this.catcher = c
}

func (this *Thrower) Throw(err error) {
	this.catcher.Catch(err)
}

func (this *Thrower) ThrowDouble(err error) float64 {
	this.catcher.Catch(err)
	return 0
}

func (this *Thrower) ThrowInt(err error) int64 {
	this.catcher.Catch(err)
	return 0
}

func (this *Thrower) ThrowUint(err error) uint64 {
	this.catcher.Catch(err)
	return 0
}

func (this *Thrower) ThrowBool(err error) bool {
	this.catcher.Catch(err)
	return false
}

func (this *Thrower) ThrowString(err error) string {
	this.catcher.Catch(err)
	return ""
}

func (this *Thrower) ThrowBytes(err error) []byte {
	this.catcher.Catch(err)
	return nil
}

func (this *Thrower) ThrowDoubles(err error) []float64 {
	this.catcher.Catch(err)
	return nil
}

func (this *Thrower) ThrowInts(err error) []int64 {
	this.catcher.Catch(err)
	return nil
}

func (this *Thrower) ThrowUints(err error) []uint64 {
	this.catcher.Catch(err)
	return nil
}

func (this *Thrower) ThrowBools(err error) []bool {
	this.catcher.Catch(err)
	return nil
}

func (this *Thrower) ThrowStrings(err error) []string {
	this.catcher.Catch(err)
	return nil
}

func (this *Thrower) ThrowListOfBytes(err error) [][]byte {
	this.catcher.Catch(err)
	return nil
}

type Catcher interface {
	Catch(error)
	GetError() error
	Clear()
}

type stackErr struct {
	err   error
	stack []byte
}

func (this *stackErr) Error() string {
	return fmt.Sprintf("error: %s\nstack: %s", this.err, this.stack)
}

func NewCatcher(debug bool) Catcher {
	return &catcher{nil, debug}
}

type catcher struct {
	errors []error
	debug  bool
}

func (this *catcher) Catch(err error) {
	if !this.debug {
		this.errors = append(this.errors, err)
		return
	}
	stack := debug.Stack()
	this.errors = append(this.errors, &stackErr{
		err:   err,
		stack: stack,
	})
}

func (this *catcher) Clear() {
	this.errors = nil
}

func (this *catcher) GetError() error {
	if len(this.errors) > 0 {
		return &caughtError{this.errors}
	}
	return nil
}

type caughtError struct {
	errors []error
}

func (this *caughtError) Error() string {
	ss := make([]string, len(this.errors))
	for i, e := range this.errors {
		ss[i] = e.Error()
	}
	return strings.Join(ss, "\n\n")
}
