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

package relapse

import (
	"github.com/katydid/katydid/expr/ast"
)

//Work in Progress: Like gofmt, but for relapse
func Format(g *Grammar) {
	first := true
	if g.TopPattern == nil {
		for i := range g.PatternDecls {
			if g.PatternDecls[i].Name == "main" {
				if !g.PatternDecls[i].Before.HasComment() &&
					!g.PatternDecls[i].Hash.Before.HasComment() &&
					!g.PatternDecls[i].Eq.Before.HasComment() {
					g.TopPattern = g.PatternDecls[i].Pattern
					g.PatternDecls = append(g.PatternDecls[:i], g.PatternDecls[i+1:]...)
					break
				}
			}
		}
	}
	if g.TopPattern != nil {
		formatPattern(g.TopPattern, true, 0)
	}
	for i := range g.PatternDecls {
		formatPatternDecl(g.PatternDecls[i], first)
		first = false
	}
	if g.After == nil {
		g.After = &expr.Space{}
	}
	formatTrim(g.After, true, 0)
}

func formatPatternDecl(patternDecl *PatternDecl, first bool) {
	if patternDecl.Before == nil {
		patternDecl.Before = &expr.Space{}
	}
	if first {
		formatTrim(patternDecl.Before, true, 0)
	} else {
		formatTrim(patternDecl.Before, false, 0)
	}
	formatKeyword(patternDecl.Eq, 0)
	formatPattern(patternDecl.Pattern, true, 0)
}

func formatPattern(pattern *Pattern, first bool, tabs int) {
	if pattern.Empty != nil {
		formatEmpty(pattern.Empty, first, tabs)
	}
	if pattern.TreeNode != nil {
		formatTreeNode(pattern.TreeNode, true, tabs)
	}
	if pattern.ZAny != nil {
		if !pattern.ZAny.Star.Before.HasComment() {
			pattern.ZAny.Star.Before = nil
		}
	}
}

func formatEmpty(e *Empty, first bool, tabs int) {
	formatKeyword(e.Empty, tabs)
}

func formatTreeNode(t *TreeNode, first bool, tabs int) {
	formatNameExpr(t.Name, first, tabs)
	if t.Colon != nil {
		formatKeyword(t.Colon, tabs)
	}
	formatPattern(t.Pattern, true, tabs)
}

func formatNameExpr(nameExpr *expr.NameExpr, first bool, tabs int) {
	if nameExpr.Name != nil {
		formatName(nameExpr.Name, first, tabs)
	}
}

func formatName(name *expr.Name, first bool, tabs int) {
	if name.Before == nil {
		name.Before = &expr.Space{}
	}
	formatTrim(name.Before, first, tabs)
	if first {
		addSpace(name.Before)
	}
}

// func formatNewline(space *expr.Space, tabs int) {
// 	formatTrim(space, false, tabs)
// 	space.Space = append([]string{"\n"}, space.Space...)
// }

func newTabs(tabs int) string {
	ss := make([]rune, tabs)
	for i := 0; i < tabs; i++ {
		ss[i] = '\t'
	}
	return string(ss)
}

func formatTrim(space *expr.Space, first bool, tabs int) {
	comments := space.GetComments()
	hadAttachedComment := space.HasAttachedComment()
	space.Space = make([]string, 0, len(comments)*2)
	firstComment := true
	for i := range comments {
		if firstComment {
			space.Space = append(space.Space, string(comments[i]))
		} else {
			space.Space = append(space.Space, newTabs(tabs)+string(comments[i]))
		}
		firstComment = false
	}
	if len(comments) > 0 && !hadAttachedComment {
		space.Space = append(space.Space, "\n")
	}
	if tabs == 0 {
		if !first {
			space.Space = append(space.Space, "\n")
		}
	} else {
		space.Space = append(space.Space, newTabs(tabs))
	}
}

func addSpace(space *expr.Space) {
	space.Space = append(space.Space, " ")
}

func formatKeyword(keyword *expr.Keyword, tabs int) {
	if keyword.Before == nil {
		keyword.Before = &expr.Space{}
	}
	formatTrim(keyword.Before, true, tabs)
	if keyword.GetValue() == "=" {
		addSpace(keyword.Before)
	}
}
