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

package ast

// func GetStringTokens(rules *Rules) ([]string, error) {
// 	s := make([]string, 0, 1+len(rules.Rules))
// 	s = append(s, rules.GetRoot().GetPackage()+"."+rules.GetRoot().GetMessage())
// 	for _, rule := range rules.Rules {
// 		if init := rule.Init; init != nil {
// 			s = append(s, init.GetPackage()+"."+init.GetMessage())
// 		}
// 	}
// 	for _, rule := range rules.Rules {
// 		if ifExpr := rule.IfExpr; ifExpr != nil {
// 			v, err := ifExpr.GetVariable()
// 			if err != nil {
// 				return nil, err
// 			}
// 			s = append(s, v.Name)
// 		}
// 	}
// 	return s, nil
// }
