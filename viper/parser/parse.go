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

package parser

import (
	"fmt"
	"github.com/katydid/katydid/viper/ast"
	"github.com/katydid/katydid/viper/lexer"
)

type ErrWrongType struct {
	typ string
	res interface{}
}

func (this *ErrWrongType) Error() string {
	return fmt.Sprintf("expected %s, but got %#v", this.typ, this.res)
}

func (this *Parser) ParseViper(s string) (res *viper.Rules, err error) {
	scanner := lexer.NewLexer([]byte(s))
	st, err := this.Parse(scanner)
	if err != nil {
		return nil, err
	}
	vst, ok := st.(*viper.Rules)
	if !ok {
		return nil, &ErrWrongType{"*viper.Rules", st}
	}
	return vst, nil
}
