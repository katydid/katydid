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

// Work in Progress
package scanner

import (
	"io"
	"os"
	"path/filepath"
)

type walkCall struct {
	path string
	info os.FileInfo
	err  error
}

type walker struct {
	root  string
	calls chan walkCall
	first bool
}

func (w *walker) walk(path string, info os.FileInfo, err error) error {
	if w.first {
		w.first = false
		return nil
	}
	w.calls <- walkCall{path, info, err}
	if info.IsDir() {
		return filepath.SkipDir
	}
	return nil
}

func (w *walker) goWalk() {
	go func() {
		filepath.Walk(w.root, w.walk)
		w.calls <- walkCall{err: io.EOF}
	}()
}

func newWalker(root string) *walker {
	return &walker{root, make(chan walkCall, 0), true}
}

type state struct {
	*walker
	current walkCall
}

func NewScanner(root string) *scanner {
	w := newWalker(root)
	w.goWalk()
	return &scanner{root, state{walker: w}, make([]state, 0, 10)}
}

type scanner struct {
	root string
	state
	stack []state
}

func (s *scanner) Next() error {
	s.current = <-s.calls
	if s.current.err != nil {
		return s.current.err
	}
	return nil
}

func (s *scanner) IsLeaf() bool {
	return !s.current.info.IsDir()
}

func (s *scanner) Id() int {
	return 0 //TODO tokens
}

func (s *scanner) Value() []byte {
	return []byte(s.current.path)
}

func (s *scanner) Up() {
	top := len(s.stack) - 1
	s.state = s.stack[top]
	s.stack = s.stack[:top]
}

func (s *scanner) Down() {
	s.stack = append(s.stack, s.state)
	w := newWalker(s.current.path)
	s.state.walker = w
	s.state.walker.goWalk()
}
