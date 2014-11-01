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
	"code.google.com/p/gogoprotobuf/proto"
	main "github.com/katydid/katydid/asm/test"
	"testing"
)

var bridgePepper = `
root = main.PuddingMilkshake
main.PuddingMilkshake = start
start accept = accept
start _ = start
accept _ = accept
notfound accept = accept
notfound _ = notfound

main.FinanceJudo = notfound
main.SaladWorry = notfound
main.SpyCarpenter = notfound

if contains($string(main.SpyCarpenter.BridgePepper), "a") then accept else notfound
`

func TestBridgePepper(t *testing.T) {
	m := &main.PuddingMilkshake{FinanceJudo: &main.FinanceJudo{SaladWorry: &main.SaladWorry{SpyCarpenter: &main.SpyCarpenter{BridgePepper: []string{"a"}}}}}
	test(t, "puddingmilkshake.proto", m, bridgePepper, true)
}

func BenchmarkBridgePepper(b *testing.B) {
	newBench("puddingmilkshake.proto", bridgePepper).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPuddingMilkshake(r, easy)
	})
}

var bridgePepperAndFountainTarget = `
root = main.PuddingMilkshake

main.PuddingMilkshake = start
start accept = accept
start _ = start
accept _ = accept

main.FinanceJudo = start
main.SaladWorry = start
main.SpyCarpenter = start

start foundOrigin = foundOrigin
start foundContrib = foundContrib
start _ = start
foundOrigin _ = foundOrigin
foundOrigin foundContrib = accept
foundContrib _ = foundContrib
foundContrib foundOrigin = accept
accept _ = accept

if contains($string(main.SpyCarpenter.BridgePepper), "a") then foundOrigin else notfound
if contains($string(main.SpyCarpenter.FountainTarget), "a") then foundContrib else notfound
`

func TestBridgePepperAndFountainTarget(t *testing.T) {
	m := &main.PuddingMilkshake{FinanceJudo: &main.FinanceJudo{SaladWorry: &main.SaladWorry{SpyCarpenter: &main.SpyCarpenter{BridgePepper: []string{"a"}, FountainTarget: []string{"a"}}}}}
	test(t, "puddingmilkshake.proto", m, bridgePepperAndFountainTarget, true)
}

func BenchmarkBridgePepperAndFountainTarget(b *testing.B) {
	newBench("puddingmilkshake.proto", bridgePepperAndFountainTarget).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPuddingMilkshake(r, easy)
	})
}

var bridgePepperOrFountainTarget = `
root = main.PuddingMilkshake

main.PuddingMilkshake = start
start accept = accept
start _ = start
accept _ = accept

main.FinanceJudo = start
main.SaladWorry = start
main.SpyCarpenter = start

start foundOrigin = accept
start foundContrib = accept
start _ = start
accept _ = accept

if contains($string(main.SpyCarpenter.BridgePepper), "a") then foundOrigin else notfound
if contains($string(main.SpyCarpenter.FountainTarget), "a") then foundContrib else notfound
`

func TestBridgePepperOrFountainTarget(t *testing.T) {
	m := &main.PuddingMilkshake{FinanceJudo: &main.FinanceJudo{SaladWorry: &main.SaladWorry{SpyCarpenter: &main.SpyCarpenter{BridgePepper: []string{"b"}, FountainTarget: []string{"a"}}}}}
	test(t, "puddingmilkshake.proto", m, bridgePepperOrFountainTarget, true)
}

func BenchmarkBridgePepperOrFountainTarget(b *testing.B) {
	newBench("puddingmilkshake.proto", bridgePepperOrFountainTarget).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedPuddingMilkshake(r, easy)
	})
}
