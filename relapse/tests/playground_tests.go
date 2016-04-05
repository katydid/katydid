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

package tests

import (
	"encoding/json"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/parser"
)

var playgroundTrueStr = `
(
	.WhatsUp == "E" &
	.Survived >= 1000000 /*years*/ & 
	.DragonsExist != true &
	.MonkeysSmart :: $bool &
	.History [
		*,
		_ == "Katydids Alive"
	] &
	.FeatureRequests._ {
		Name *= "art";
		*;
		Anatomy $= "omen";
	}
)
`

var playgroundFalseStr = `
(
	.WhatsUp == "F" &
	.Survived >= 1000000 /*years*/ & 
	.DragonsExist != true &
	.MonkeysSmart :: $bool &
	.History [
		*,
		_ == "Katydids Alive"
	] &
	.FeatureRequests._ {
		Name *= "art";
		*;
		Anatomy $= "omen";
	}
)
`

var playgroundJson = `
{
    "WhatsUp": "E",
    "Survived": 100000000,
    "History": [
        "Giant Lizards",
        "Meteor",
        "Lizards Dead",
        "Katydids Alive"
    ],
    "FeatureRequests": [
        {
            "Name": "Dart",
            "Properties": [
                "Poison",
                "Projectile"
            ],
            "Anatomy": "Abdomen"
        },
        {
            "Name": "Fire Breath",
            "Properties": [
                "Fire",
                "Vapor"
            ],
            "Anatomy": "Mouth"
        }
    ],
    "DragonsExist": false,
    "MonkeysSmart": true,
    "Family": {
        "Class": "Insecta",
        "Order": {
            "Superorder": {
                "Subclass": "Pterygota",
                "Infraclass": "Polyneoptera"
            },
            "Order": "Orthoptera"
        },
        "Suborder": "Ensifera",
        "Family": "Tettigoniidae"
    }
}
`

func init() {
	playgroundTrueGrammar, err := parser.ParseGrammar(playgroundTrueStr)
	if err != nil {
		panic(err)
	}
	playgroundFalseGrammar, err := parser.ParseGrammar(playgroundFalseStr)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(playgroundJson), &m); err != nil {
		panic(err)
	}
	trueG := combinator.G(ast.NewRefLookup(playgroundTrueGrammar))
	falseG := combinator.G(ast.NewRefLookup(playgroundFalseGrammar))
	Validate("PlaygroundTrue", trueG, Json(m), true)
	Validate("PlaygroundFalse", falseG, Json(m), false)
}
