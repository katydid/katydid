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

package compose

import (
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/awalterschulze/katydid/asm/ast"
	"github.com/awalterschulze/katydid/types"
	"testing"
)

func TestComposeNot(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "not",
			Params: []*ast.Expr{
				{
					Terminal: &ast.Terminal{
						BoolValue: proto.Bool(false),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	if b.Eval(nil) != true {
		t.Fatalf("expected true")
	}
}

//contains(nfkc(decString(test.Address.Street.value)), nfkc("TheStreet"))

func TestComposeContains(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "contains",
			Params: []*ast.Expr{
				{
					Function: &ast.Function{
						Name: "nfkc",
						Params: []*ast.Expr{
							{
								Function: &ast.Function{
									Name: "decString",
									Params: []*ast.Expr{
										{
											Terminal: &ast.Terminal{
												Variable: &ast.Variable{
													Package: "a",
													Message: "a",
													Field:   "a",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Function: &ast.Function{
						Name: "nfkc",
						Params: []*ast.Expr{
							{
								Terminal: &ast.Terminal{
									StringValue: proto.String("TheStreet"),
								},
							},
						},
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	if b.Eval([]byte("TheStreet")) != true {
		t.Fatalf("expected true")
	}
	if b.Eval([]byte("ThatStreet")) != false {
		t.Fatalf("expected false")
	}
}

func TestComposeStringEq(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "eq",
			Params: []*ast.Expr{
				{
					Function: &ast.Function{
						Name: "nfkc",
						Params: []*ast.Expr{
							{
								Function: &ast.Function{
									Name: "decString",
									Params: []*ast.Expr{
										{
											Terminal: &ast.Terminal{
												Variable: &ast.Variable{
													Package: "a",
													Message: "a",
													Field:   "a",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Function: &ast.Function{
						Name: "nfkc",
						Params: []*ast.Expr{
							{
								Terminal: &ast.Terminal{
									StringValue: proto.String("TheStreet"),
								},
							},
						},
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	if b.Eval([]byte("TheStreet")) != true {
		t.Fatalf("expected true")
	}
}

func TestComposeListBool(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "eq",
			Params: []*ast.Expr{
				{
					Function: &ast.Function{
						Name: "length",
						Params: []*ast.Expr{
							{
								Function: &ast.Function{
									Name: "print",
									Params: []*ast.Expr{
										{
											List: &ast.List{
												Type: types.LIST_BOOL.Enum(),
												Elems: []*ast.Expr{
													{
														Terminal: &ast.Terminal{
															BoolValue: proto.Bool(true),
														},
													},
													{
														Terminal: &ast.Terminal{
															BoolValue: proto.Bool(false),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Terminal: &ast.Terminal{
						Int64Value: proto.Int64(2),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	if b.Eval(nil) != true {
		t.Fatalf("expected true")
	}
}

func TestComposeListInt64(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "eq",
			Params: []*ast.Expr{
				{
					Function: &ast.Function{
						Name: "elem",
						Params: []*ast.Expr{
							{
								Function: &ast.Function{
									Name: "print",
									Params: []*ast.Expr{
										{
											List: &ast.List{
												Elems: []*ast.Expr{
													{
														Terminal: &ast.Terminal{
															Int64Value: proto.Int64(1),
														},
													},
													{
														Terminal: &ast.Terminal{
															Int64Value: proto.Int64(2),
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Terminal: &ast.Terminal{
									Int64Value: proto.Int64(1),
								},
							},
						},
					},
				},
				{
					Terminal: &ast.Terminal{
						Int64Value: proto.Int64(2),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	if b.Eval(nil) != true {
		t.Fatalf("expected true")
	}
}
