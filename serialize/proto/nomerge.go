//  Copyright 2016 Walter Schulze
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

package proto

import (
	"fmt"
	"github.com/katydid/katydid/serialize"
	"io"
)

func NoLatentAppendingOrMerging(parser serialize.Parser) error {
	seen := make(map[uint64]bool)
	for {
		if err := parser.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		if !parser.IsLeaf() {
			if _, err := parser.Int(); err != nil {
				if key, err := parser.Uint(); err == nil {
					name, err := parser.String()
					if err != nil {
						return err
					}
					if _, ok := seen[key]; ok {
						return fmt.Errorf(name + " requires merging")
					}
					seen[key] = true
				} else {
					panic("not an index, field or leaf: " + err.Error())
				}
			}
			parser.Down()
			if err := NoLatentAppendingOrMerging(parser); err != nil {
				return err
			}
			parser.Up()
		}
	}
	return nil
}
