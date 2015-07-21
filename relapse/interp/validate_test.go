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

package interp_test

import (
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/tests"
	"log"
	"testing"
)

func TestValidators(t *testing.T) {
	for _, v := range tests.Validators {
		for scannerName, s := range v.Scanners {
			log.Printf("TESTING %s with Codec %s", v.Name, scannerName)
			match := interp.Interpret(v.Grammar, s())
			if match != v.Expected {
				t.Errorf("%s: Expected a %v match given %s scanner from \n%v \non \n%#v", v.Name, v.Expected, scannerName, v.Grammar.String(), v.Description)
			}
		}
	}
}
