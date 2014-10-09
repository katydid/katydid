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

package main_test

import (
	"testing"
)

var preserve = `//Has this person ever lived at 456 TheStreet

root = main.Person
main.Person = start //asdf
start /*a*/ numberAndStreet = accept
start _ = start

/*asdf*/
//adsf
main.Address = address
address number = number



address street = street
address _ = address
number/*a*/street = numberAndStreet
number _ = //adsf
number
street number = numberAndStreet
street _ = street

if eq($int64(main.Address.Number), int64(456)) 






then/*a*/number else //adsf
noNumber
if contains(nfkc($string(main.Address.Street)/*asdf*/), nfkc("TheStreet"//asdf
	)) then street else noStreet`

func TestPreserve(t *testing.T) {
	test(t, "person.proto", david, preserve, false)
}
