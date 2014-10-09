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

package scanner

import (
	"fmt"
	"io"
	"testing"
)

func print(s *scanner, tabs string) {
	err := s.Next()
	for err == nil {
		fmt.Printf("%s%s\n", tabs, string(s.Value()))
		if !s.IsLeaf() {
			s.Down()
			print(s, tabs+"\t")
			s.Up()
		}
		err = s.Next()
	}
	if err == io.EOF {
		return
	}
	if err != nil {
		panic(err)
	}
}

func TestPrint(t *testing.T) {
	print(NewScanner("."), "")
}
