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

package eval

import (
	"fmt"
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/serialize"
	"github.com/katydid/katydid/viper/ast"
	"io"
	"strings"
)

func getStart(rules *viper.Rules) string {
	for _, rule := range rules.Rules {
		if rule.Start == nil {
			continue
		}
		return rule.Start.State.Name
	}
	panic("no start state")
}

func getFinals(rules *viper.Rules) []string {
	finals := []string{}
	for i, rule := range rules.Rules {
		if rule.Final == nil {
			continue
		}
		finals = append(finals, rules.Rules[i].Final.State.Name)
	}
	return finals
}

func isFinal(rules *viper.Rules, f string) bool {
	finals := getFinals(rules)
	for _, ff := range finals {
		if ff == f {
			return true
		}
	}
	return false
}

func Eval(rules *viper.Rules, tree serialize.Scanner) bool {
	return isFinal(rules, eval(rules, tree, getStart(rules)))
}

func evalName(e *expr.Expr, name string) bool {
	return evalExpr(e, serialize.NewStringValue(name))
}

func evalReturn(rules *viper.Rules, srcParent, srcChild string, name string) (dst string) {
	for _, rule := range rules.Rules {
		if rule.Return == nil {
			continue
		}
		if rule.Return.ParentSrc.Name == srcParent && rule.Return.ChildSrc.Name == srcChild {
			if evalName(rule.Return.Expr, name) {
				fmt.Printf("%s\n", strings.TrimSpace(rule.String()))
				return rule.Return.Dst.Name
			}
		}
	}
	panic("unknown return " + srcParent + " " + srcChild + " " + name)
}

func evalCall(rules *viper.Rules, src string, name string) (parentDst string, childDst string) {
	for _, rule := range rules.Rules {
		if rule.Call == nil {
			continue
		}
		if rule.Call.Src.Name == src {
			if evalName(rule.Call.Expr, name) {
				fmt.Printf("%s\n", strings.TrimSpace(rule.String()))
				return rule.Call.ParentDst.Name, rule.Call.ChildDst.Name
			}
		}
	}
	panic("unknown call " + src + " " + name)
}

func evalExpr(expr *expr.Expr, value serialize.Decoder) bool {
	f, err := compose.NewBool(expr)
	if err != nil {
		panic(err)
	}
	res, err := f.Eval(value)
	if err != nil {
		//TODO think about this again and write some tests
		return false
	}
	return res
}

func evalInternal(rules *viper.Rules, src string, value serialize.Decoder) (dst string) {
	for _, rule := range rules.Rules {
		if rule.Internal == nil {
			continue
		}
		if rule.Internal.Src.Name == src {
			if evalExpr(rule.Internal.Expr, value) {
				fmt.Printf("%s\n", strings.TrimSpace(rule.String()))
				return rule.Internal.Dst.Name
			}
		}
	}
	panic("unknown internal " + src + " " + serialize.Sprint(value))
}

func eval(rules *viper.Rules, tree serialize.Scanner, current string) string {
	for {
		fmt.Printf("state = %s\n", current)
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		name := tree.Name()
		parent, child := evalCall(rules, current, name)
		if tree.IsLeaf() {
			child = evalInternal(rules, child, tree)
		} else {
			tree.Down()
			child = eval(rules, tree, child)
			tree.Up()
		}
		current = evalReturn(rules, parent, child, name)
	}
	return current
}
