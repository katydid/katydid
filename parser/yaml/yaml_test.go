//  Copyright 2015 Elphas Tori
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

package yaml

import (
	"testing"

	"github.com/katydid/katydid/parser/debug"
)

func testYaml(t *testing.T, s string, output debug.Nodes) {
	p := NewYamlParser()
	if err := p.Init([]byte(s)); err != nil {
		t.Fatal(err)
	}
	m := debug.Walk(p)
	if !m.Equal(output) {
		t.Fatalf("expected %s but got %s", output, m)
	}
}

//debugInput is a serialized YAML equivalent of the sample Debug struct in debug.Input
var debugInput = `
		A: 1
		B:
			- b2
			- b3
		C:
			A: 2
			D: 3
			E:
				- B:
					- b4

				- B:
					- b5
		D: 4
		F:
			- 5`

//A simplified version of debug.Output which excludes complex array objects, a known limitation
//of the current YAML parser implementation
var debugOutput = debug.Nodes{
	debug.Field(`A`, `1`),
	debug.Nested(`B`,
		debug.Field(`0`, `b2`),
		debug.Field(`1`, `b3`),
	),
	debug.Nested(`C`,
		debug.Field(`A`, `2`),
		debug.Field(`D`, `3`),
		debug.Nested(`E`,
			debug.Nested(`0`,
				debug.Nested(`B`,
					debug.Field(`0`, `b4`),
				),
			),
			debug.Nested(`1`,
				debug.Nested(`B`,
					debug.Field(`0`, `b5`),
				),
			),
		),
	),
	debug.Field(`D`, `4`),
	debug.Nested(`F`,
		debug.Field(`0`, `5`),
	),
}

func TestDebug(t *testing.T) {
	testYaml(t, debugInput, debugOutput)
}

func TestRandomDebug(t *testing.T) {
	p := NewYamlParser()
	for i := 0; i < 10; i++ {
		if err := p.Init([]byte(debugInput)); err != nil {
			t.Fatal(err)
		}
		debug.RandomWalk(p, debug.NewRand(), 10, 3)
	}
}

func TestExample(t *testing.T) {
	input := `
	Top:
		Name: Katydid
		Dragons:
			Fire: true
		Empty:`

	output := debug.Nodes{
		debug.Nested(`Top`,
			debug.Field(`Name`, `Katydid`),
			debug.Nested(`Dragons`,
				debug.Field(`Fire`, `true`),
			),
			debug.Field(`Empty`, `<nil>`),
		),
	}

	testYaml(t, input, output)
}

func TestPudding(t *testing.T) {
	input := `
FinanceJudo:
	SaladWorry:
		MeasureGrade
		MagazineFrame:
			- a
			- b
		XrayPilot:
			AnkleCoat: 2
			XXX_unrecognized:
		XXX_unrecognized:
	RumourSpirit: 1
	XXX_unrecognized:
	`
	output := debug.Nodes{
		debug.Nested(`FinanceJudo`,
			debug.Nested(`SaladWorry`,
				debug.Field(`MeasureGrade`, `<nil>`),
				debug.Nested(`MagazineFrame`,
					debug.Field(`0`, `a`),
					debug.Field(`1`, `b`),
				),
				debug.Nested(`XrayPilot`,
					debug.Field(`AnkleCoat`, `2`),
					debug.Field(`XXX_unrecognized`, `<nil>`),
				),
				debug.Field(`XXX_unrecognized`, `<nil>`),
			),
			debug.Field(`RumourSpirit`, `1`),
			debug.Field(`XXX_unrecognized`, `<nil>`),
		),
	}

	testYaml(t, input, output)
}

func TestPerson(t *testing.T) {
	input := `
	Person:
		Name: Robert
		Addresses:
			Number: 456
			Street: TheStreet
		Telephone: 0127897897
		XXX_unrecognized`

	output := debug.Nodes{
		debug.Nested(`Person`,
			debug.Field(`Name`, `Robert`),
			debug.Nested(`Addresses`,
				debug.Field(`Number`, `456`),
				debug.Field(`Street`, `TheStreet`),
			),
			debug.Field(`Telephone`, `0127897897`),
			debug.Field(`XXX_unrecognized`, `<nil>`),
		),
	}

	testYaml(t, input, output)
}

func TestAB(t *testing.T) {
	input := `
	A:
		B`

	output := debug.Nodes{
		debug.Nested(`A`,
			debug.Field(`B`, `<nil>`),
		),
	}

	testYaml(t, input, output)
}
