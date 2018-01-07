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

package funcs

//Not returns a new not function with the input function as its parameter.
func Not(v1 Bool) Bool {
	if vv, ok := v1.(*not); ok {
		return vv.V1
	}
	if vv, ok := v1.(*constBool); ok {
		return BoolConst(!vv.v)
	}
	if vv, ok := v1.(*and); ok {
		return Or(Not(vv.V1), Not(vv.V2))
	}
	if vv, ok := v1.(*or); ok {
		return And(Not(vv.V1), Not(vv.V2))
	}
	switch vv := v1.(type) {
	case *boolEq:
		return BoolNe(vv.V1, vv.V2)
	case *bytesEq:
		return BytesNe(vv.V1, vv.V2)
	case *doubleEq:
		return DoubleNe(vv.V1, vv.V2)
	case *intEq:
		return IntNe(vv.V1, vv.V2)
	case *stringEq:
		return StringNe(vv.V1, vv.V2)
	case *uintEq:
		return UintNe(vv.V1, vv.V2)
	case *boolNe:
		return BoolEq(vv.V1, vv.V2)
	case *bytesNe:
		return BytesEq(vv.V1, vv.V2)
	case *doubleNe:
		return DoubleEq(vv.V1, vv.V2)
	case *intNe:
		return IntEq(vv.V1, vv.V2)
	case *stringNe:
		return StringEq(vv.V1, vv.V2)
	case *uintNe:
		return UintEq(vv.V1, vv.V2)
	}
	return &not{v1}
}

type not struct {
	V1 Bool
}

func (this *not) Eval() (bool, error) {
	b, err := this.V1.Eval()
	if err != nil || !b {
		return true, nil
	}
	return !b, nil
}

func init() {
	Register("not", new(not))
}

//And returns a new and function with the two input functions as its parameters.
func And(v1, v2 Bool) Bool {
	if l, ok := v1.(*constBool); ok {
		if l.v == false {
			return BoolConst(false)
		} else {
			return v2
		}
	}
	if r, ok := v2.(*constBool); ok {
		if r.v == false {
			return BoolConst(false)
		} else {
			return v1
		}
	}
	if Equal(v1, v2) {
		return v1
	}
	if l, ok := v1.(*not); ok {
		if Equal(l.V1, v2) {
			return BoolConst(false)
		}
	}
	if r, ok := v2.(*not); ok {
		if Equal(v1, r.V1) {
			return BoolConst(false)
		}
	}
	switch vv1 := v1.(type) {
	case *stringEq:
		if vv2, ok := v2.(*stringEq); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return BoolConst(false)
					}
				}
			}
		}
		if vv2, ok := v2.(*stringNe); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return v1
					} else {
						return BoolConst(false)
					}
				}
			}
		}
	case *stringNe:
		if vv2, ok := v2.(*stringEq); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return v2
					} else {
						return BoolConst(false)
					}
				}
			}
		}
	case *intEq:
		if vv2, ok := v2.(*intEq); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return BoolConst(false)
					}
				}
			}
		}
		if vv2, ok := v2.(*intNe); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return v1
					} else {
						return BoolConst(false)
					}
				}
			}
		}
	case *intNe:
		if vv2, ok := v2.(*intEq); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return v2
					} else {
						return BoolConst(false)
					}
				}
			}
		}
	case *uintEq:
		if vv2, ok := v2.(*uintEq); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return BoolConst(false)
					}
				}
			}
		}
		if vv2, ok := v2.(*uintNe); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return v1
					} else {
						return BoolConst(false)
					}
				}
			}
		}
	case *uintNe:
		if vv2, ok := v2.(*uintEq); ok {
			if vvv1, ok1 := isVarConst(vv1.V1, vv1.V2); ok1 {
				if vvv2, ok2 := isVarConst(vv2.V1, vv2.V2); ok2 {
					if vvv1 != vvv2 {
						return v2
					} else {
						return BoolConst(false)
					}
				}
			}
		}
	}
	return &and{v1, v2}
}

type and struct {
	V1 Bool
	V2 Bool
}

func (this *and) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil || !v1 {
		return false, err
	}
	return this.V2.Eval()
}

func init() {
	Register("and", new(and))
}

//Or returns a new or function with the two input functions as its parameters.
func Or(v1, v2 Bool) Bool {
	if l, ok := v1.(*constBool); ok {
		if l.v == true {
			return BoolConst(true)
		} else {
			return v2
		}
	}
	if r, ok := v2.(*constBool); ok {
		if r.v == true {
			return BoolConst(true)
		} else {
			return v1
		}
	}
	if Equal(v1, v2) {
		return v1
	}
	if l, ok := v1.(*not); ok {
		if Equal(l.V1, v2) {
			return BoolConst(true)
		}
	}
	if r, ok := v2.(*not); ok {
		if Equal(v1, r.V1) {
			return BoolConst(true)
		}
	}
	return &or{v1, v2}
}

type or struct {
	V1 Bool
	V2 Bool
}

func (this *or) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err == nil && v1 {
		return true, nil
	}
	return this.V2.Eval()
}

func init() {
	Register("or", new(or))
}
