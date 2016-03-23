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

package xml

import (
	"encoding/json"
	"github.com/katydid/katydid/parser/debug"
	"testing"
)

func testXML(t *testing.T, s string) {
	x := NewXMLParser()
	if err := x.Init([]byte(s)); err != nil {
		panic(err)
	}
	m := debug.Walk(debug.NewLogger(x, debug.NewLineLogger()))
	data, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(string(data))
}

func TestExample(t *testing.T) {
	example := `
		<Top>
			<Name>Katydid</Name>
			<Dragons alive="false">
				<Fire>true</Fire>
			</Dragons>
			<Empty></Empty>
		</Top>`
	testXML(t, example)
}

func TestPudding(t *testing.T) {
	pudding := `
	<FinanceJudo>
	<SaladWorry><MeasureGrade></MeasureGrade><MagazineFrame>a</MagazineFrame><MagazineFrame>b</MagazineFrame>
		<XrayPilot><AnkleCoat>2</AnkleCoat><XXX_unrecognized></XXX_unrecognized></XrayPilot><XXX_unrecognized></XXX_unrecognized>
	</SaladWorry>
	<RumourSpirit>1</RumourSpirit><XXX_unrecognized></XXX_unrecognized>
	</FinanceJudo>
	`
	testXML(t, pudding)
}

func TestPerson(t *testing.T) {
	person := `<Person>
		<Name>Robert</Name>
		<Addresses>
				<Number>456</Number>
				<Street>TheStreet</Street>
		</Addresses>
		<Telephone>0127897897</Telephone>
		<XXX_unrecognized/>
	</Person>`
	testXML(t, person)
}

func TestAB(t *testing.T) {
	testXML(t, `<A>B</A>`)
}
