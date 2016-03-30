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
	"io"
)

//NoLatentAppendingOrMerging returns whether the current parser has some latent fields.
//Latent fields are those fields you have already seen on your walk, but then after seeing a different field you see this field again.
//This typically happens when the protocol buffer user created an object marshaled it and then merged it with another value.
func NoLatentAppendingOrMerging(parser ProtoParser) error {
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
				if fieldNum, err := parser.Uint(); err == nil {
					if _, ok := seen[fieldNum]; ok {
						return fmt.Errorf("%s requires merging", parser.Field().GetName())
					}
					seen[fieldNum] = true
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
