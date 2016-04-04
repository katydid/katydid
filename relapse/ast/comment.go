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

import (
	"strings"
)

type Comment string

//GetContent returns the content of the comment excluding the // or /* */
func (this Comment) GetContent() string {
	if len(this) == 0 {
		return ""
	}
	s := []byte(this)
	if isLineComment(string(this)) {
		return string(s[2 : len(s)-1])
	}
	return string(s[2 : len(s)-2])
}

func isBlockComment(s string) bool {
	return strings.HasPrefix(s, `/*`) && strings.HasSuffix(s, `*/`)
}

func isLineComment(s string) bool {
	return strings.HasPrefix(s, `//`) && strings.HasSuffix(s, "\n")
}

func isComment(s string) bool {
	return isBlockComment(s) || isLineComment(s)
}

//HasComment returns whether the white space contains a comment.
func (this *Space) HasComment() bool {
	if this == nil {
		return false
	}
	for _, s := range this.Space {
		if isComment(s) {
			return true
		}
	}
	return false
}

//GetComments returns a list of comments contained in the white space.
func (this *Space) GetComments() []Comment {
	if this == nil {
		return nil
	}
	comments := []Comment{}
	for i, s := range this.Space {
		if isComment(s) {
			comments = append(comments, Comment(this.Space[i]))
		}
	}
	return comments
}

//HasAttachedComment returns whether the white space has any attached comments.
//An attached comment is one that does not contain a newline at the end of the white space.
func (this *Space) HasAttachedComment() bool {
	return len(this.GetAttachedComment()) > 0
}

//GetAttachedComment returns the last comment in the whitespace that does not have a newline at the end of the comment.
//If there is a newline at the end of the comment, this function returns an empty Comment.
func (this *Space) GetAttachedComment() Comment {
	if this == nil {
		return ""
	}
	newlines := 2
	var comment Comment
	for i, s := range this.Space {
		if isLineComment(s) {
			newlines = 1
			comment = Comment(this.Space[i])
		} else if isBlockComment(s) {
			newlines = 0
			comment = Comment(this.Space[i])
		} else if strings.Contains(s, "\n") {
			newlines++
		}
	}
	if newlines < 2 {
		return comment
	}
	return ""
}
