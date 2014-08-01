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
	main "github.com/awalterschulze/katydid/asm/test"
	"testing"
)

var scarBusStop = `
root = main.TypewriterPrison

main.TypewriterPrison = start
start accept = accept
start reject = reject
start _ = start
accept _ = accept
reject _ = reject

main.PocketRoses = start

if contains(decString(main.PocketRoses.ScarBusStop), "a") then accept else reject
`

func TestTypewriterPrisonScarBusStop(t *testing.T) {
	m := &main.TypewriterPrison{PocketRoses: &main.PocketRoses{ScarBusStop: proto.String("a")}}
	test(t, "typewriterprison.proto", m, scarBusStop, true)
}

func BenchmarkTypewriterPrisonScarBusStop(b *testing.B) {
	newBench("typewriterprison.proto", scarBusStop).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedTypewriterPrison(r, easy)
	})
}

var daisySled = `

root = main.TypewriterPrison

main.TypewriterPrison = start
start accept = accept
start reject = reject
start _ = start
accept _ = accept
reject _ = reject

main.PocketRoses = start

if eq(decInt64(main.PocketRoses.DaisySled), int64(1)) then accept else reject
`

func TestTypewriterPrisonDaisySled(t *testing.T) {
	m := &main.TypewriterPrison{PocketRoses: &main.PocketRoses{DaisySled: proto.Int64(1)}}
	test(t, "typewriterprison.proto", m, daisySled, true)
}

func BenchmarkTypewriterPrisonDaisySled(b *testing.B) {
	newBench("typewriterprison.proto", daisySled).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedTypewriterPrison(r, easy)
	})
}

var smileLetter = `
root = main.TypewriterPrison

main.TypewriterPrison = start
start accept = accept
start reject = reject
start _ = start
accept _ = accept
reject _ = reject

main.PocketRoses = start

if decBool(main.PocketRoses.SmileLetter) then accept else reject`

func TestTypewriterPrisonSmileLetter(t *testing.T) {
	m := &main.TypewriterPrison{PocketRoses: &main.PocketRoses{SmileLetter: proto.Bool(true)}}
	test(t, "typewriterprison.proto", m, smileLetter, true)
}

func BenchmarkTypewriterPrisonSmileLetter(b *testing.B) {
	newBench("typewriterprison.proto", smileLetter).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedTypewriterPrison(r, easy)
	})
}

var menuPaperClip = `
root = main.TypewriterPrison
main.TypewriterPrison = start
start accept = accept
start _ = start

main.PocketRoses = notfound
notfound accept = accept
notfound _ = notfound
accept _ = accept

if contains(decString(main.PocketRoses.MenuPaperclip), "a") then accept else notfound`

func TestTypewriterPrisonMenuPaperclip(t *testing.T) {
	m := &main.TypewriterPrison{PocketRoses: &main.PocketRoses{MenuPaperclip: []string{"a"}}}
	test(t, "typewriterprison.proto", m, menuPaperClip, true)
}

func BenchmarkTypewriterPrisonMenuPaperclip(b *testing.B) {
	newBench("typewriterprison.proto", menuPaperClip).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedTypewriterPrison(r, easy)
	})
}

var mapShark = `
root = main.TypewriterPrison

main.TypewriterPrison = start
start accept = accept
start reject = reject
start _ = start
accept _ = accept
reject _ = reject

main.PocketRoses = start

if contains(decString(main.PocketRoses.MapShark), "a") then accept else reject
`

func TestTypewriterPrisonMapShark(t *testing.T) {
	m := &main.TypewriterPrison{PocketRoses: &main.PocketRoses{MapShark: proto.String("a")}}
	test(t, "typewriterprison.proto", m, mapShark, true)
}

func BenchmarkTypewriterPrisonMapShark(b *testing.B) {
	newBench("typewriterprison.proto", mapShark).bench(b, func(r randyTest, easy bool) proto.Message {
		return main.NewPopulatedTypewriterPrison(r, easy)
	})
}
