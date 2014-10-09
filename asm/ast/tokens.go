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

func GetStringTokens(rules *Rules) ([]string, error) {
	s := make([]string, 0, 1+len(rules.Rules))
	s = append(s, rules.GetRoot().GetPackage()+"."+rules.GetRoot().GetMessage())
	for _, rule := range rules.Rules {
		if init := rule.Init; init != nil {
			s = append(s, init.GetPackage()+"."+init.GetMessage())
		}
	}
	for _, rule := range rules.Rules {
		if ifExpr := rule.IfExpr; ifExpr != nil {
			v, err := ifExpr.GetVariable()
			if err != nil {
				return nil, err
			}
			s = append(s, v.Name)
		}
	}
	return s, nil
}

type errUnmatchedVariables struct {
	var1 string
	var2 string
}

func newErrUnmatchedVariables(var1 *Variable, var2 *Variable) error {
	return &errUnmatchedVariables{
		var1: var1.Name,
		var2: var2.Name,
	}
}

func (this *errUnmatchedVariables) Error() string {
	return "Variables in if expr do not match " + this.var1 + " != " + this.var2
}

func (ifExpr *IfExpr) GetVariable() (*Variable, error) {
	var1, err := ifExpr.GetCondition().getVariable()
	if err != nil {
		return nil, err
	}
	if ifExpr.GetThenClause().IfExpr != nil {
		var2, err := ifExpr.GetThenClause().GetIfExpr().GetVariable()
		if err != nil {
			return nil, err
		}
		if var1 == nil {
			var1 = var2
		} else if var2 != nil {
			if var1.Name != var2.Name {
				return nil, newErrUnmatchedVariables(var1, var2)
			}
		}
	}
	if ifExpr.GetElseClause().IfExpr != nil {
		var2, err := ifExpr.GetElseClause().GetIfExpr().GetVariable()
		if err != nil {
			return nil, err
		}
		if var1 == nil {
			var1 = var2
		} else if var2 != nil {
			if var1.Name != var2.Name {
				return nil, newErrUnmatchedVariables(var1, var2)
			}
		}
	}
	return var1, nil
}

func (expr *Expr) getVariable() (*Variable, error) {
	if expr.Terminal != nil {
		return expr.GetTerminal().GetVariable(), nil
	}
	if expr.List != nil {
		return getVariables(expr.GetList().GetElems())
	}
	return getVariables(expr.GetFunction().GetParams())
}

func getVariables(exprs []*Expr) (*Variable, error) {
	var var1 *Variable
	for _, expr := range exprs {
		var2, err := expr.getVariable()
		if err != nil {
			return nil, err
		}
		if var1 == nil {
			var1 = var2
		} else if var2 != nil {
			if var1.Name != var2.Name {
				return nil, newErrUnmatchedVariables(var1, var2)
			}
		}
	}
	return var1, nil
}
