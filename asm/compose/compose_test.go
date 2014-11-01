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
	"github.com/katydid/katydid/asm/ast"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/serialize"
	"github.com/katydid/katydid/types"
	"strings"
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
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	str := funcs.Sprint(b.Func)
	if str != "true" {
		t.Fatalf("trimming did not work: %s", str)
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
								Terminal: &ast.Terminal{
									Variable: &ast.Variable{
										Name: "a.a.a",
										Type: types.SINGLE_STRING,
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
	r, err := b.Eval(serialize.NewStringValue("TheStreet"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	r, err = b.Eval(serialize.NewStringValue("ThatStreet"))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
	if strings.Contains(funcs.Sprint(b.Func), "nfkc(`TheStreet`)") {
		t.Fatalf("trimming did not work")
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
								Terminal: &ast.Terminal{
									Variable: &ast.Variable{
										Name: "a.a.a",
										Type: types.SINGLE_STRING,
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
	r, err := b.Eval(serialize.NewStringValue("TheStreet"))
	if err != nil {
		panic(err)
	}
	if r != true {
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
								List: &ast.List{
									Type: types.LIST_BOOL,
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
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	str := funcs.Sprint(b.Func)
	if str != "true" {
		t.Fatalf("trimming did not work: %s", str)
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
												Type: types.LIST_INT64,
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
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	t.Logf("%s", funcs.Sprint(b.Func))
}

func TestComposeRegex(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "regex",
			Params: []*ast.Expr{
				{
					Terminal: &ast.Terminal{
						StringValue: proto.String("ab"),
					},
				},
				{
					Terminal: &ast.Terminal{
						Variable: &ast.Variable{
							Name: "a.a.a",
							Type: types.SINGLE_BYTES,
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
	r, err := b.Eval(serialize.NewBytesValue([]byte("a")))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
}

func TestConst(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "regex",
			Params: []*ast.Expr{
				{
					Terminal: &ast.Terminal{
						Variable: &ast.Variable{
							Name: "a.a.a",
							Type: types.SINGLE_STRING,
						},
					},
				},
				{
					Terminal: &ast.Terminal{
						BytesValue: []byte{1, 2},
					},
				},
			},
		},
	}
	_, err := NewBool(expr)
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "regex has constant") || !strings.Contains(err.Error(), "has a variable parameter") {
		t.Fatalf("expected more specific error %s", err.Error())
	}
}

func TestTrimInit(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "regex",
			Params: []*ast.Expr{
				{
					Terminal: &ast.Terminal{
						StringValue: proto.String("ab"),
					},
				},
				{
					Terminal: &ast.Terminal{
						BytesValue: []byte{'a', 'b'},
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestNoTrim(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "eq",
			Params: []*ast.Expr{
				{
					Function: &ast.Function{
						Name: "elem",
						Params: []*ast.Expr{
							{
								List: &ast.List{
									Type: types.LIST_STRING,
									Elems: []*ast.Expr{
										{
											Function: &ast.Function{
												Name: "print",
												Params: []*ast.Expr{
													{
														Terminal: &ast.Terminal{
															Variable: &ast.Variable{
																Name: "a.a.a",
																Type: types.SINGLE_STRING,
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
									Int64Value: proto.Int64(0),
								},
							},
						},
					},
				},
				{
					Terminal: &ast.Terminal{
						StringValue: proto.String("abc"),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	str := funcs.Sprint(b.Func)
	if str == "false" {
		t.Fatalf("too much trimming")
	}
	t.Logf("trimmed = %s", str)
	r, err := b.Eval(serialize.NewStringValue("abc"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestList(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "eq",
			Params: []*ast.Expr{
				{
					Function: &ast.Function{
						Name: "elem",
						Params: []*ast.Expr{
							{
								List: &ast.List{
									Type: types.LIST_STRING,
									Elems: []*ast.Expr{
										{
											Terminal: &ast.Terminal{
												StringValue: proto.String("abc"),
											},
										},
									},
								},
							},
							{
								Terminal: &ast.Terminal{
									Int64Value: proto.Int64(0),
								},
							},
						},
					},
				},
				{
					Terminal: &ast.Terminal{
						StringValue: proto.String("abc"),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	str := funcs.Sprint(b.Func)
	if str != "true" {
		t.Fatalf("not enough trimming on %s", str)
	}
	r, err := b.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}
