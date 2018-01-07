// Code generated by funcs-gen. DO NOT EDIT.
package funcs

import (
	"github.com/katydid/katydid/parser"
)

type varDouble struct {
	Value parser.Value
	hash uint64
}

var _ Setter = &varDouble{}
var _ Variable = &varDouble{}

func (this *varDouble) Eval() (float64, error) {
	v, err := this.Value.Double()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (this *varDouble) Hash() uint64 {
	return this.hash
}

func (this *varDouble) IsVariable() {}

func (this *varDouble) SetValue(v parser.Value) {
	this.Value = v
}

func (this *varDouble) String() string {
	return "$double"
}

//DoubleVar returns a variable of type Double
func DoubleVar() *varDouble {
	h := uint64(17)
	h = 31*h + 2052876273
	return &varDouble{hash: h}
}

type varInt struct {
	Value parser.Value
	hash uint64
}

var _ Setter = &varInt{}
var _ Variable = &varInt{}

func (this *varInt) Eval() (int64, error) {
	v, err := this.Value.Int()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (this *varInt) Hash() uint64 {
	return this.hash
}

func (this *varInt) IsVariable() {}

func (this *varInt) SetValue(v parser.Value) {
	this.Value = v
}

func (this *varInt) String() string {
	return "$int"
}

//IntVar returns a variable of type Int
func IntVar() *varInt {
	h := uint64(17)
	h = 31*h + 73679
	return &varInt{hash: h}
}

type varUint struct {
	Value parser.Value
	hash uint64
}

var _ Setter = &varUint{}
var _ Variable = &varUint{}

func (this *varUint) Eval() (uint64, error) {
	v, err := this.Value.Uint()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (this *varUint) Hash() uint64 {
	return this.hash
}

func (this *varUint) IsVariable() {}

func (this *varUint) SetValue(v parser.Value) {
	this.Value = v
}

func (this *varUint) String() string {
	return "$uint"
}

//UintVar returns a variable of type Uint
func UintVar() *varUint {
	h := uint64(17)
	h = 31*h + 2636666
	return &varUint{hash: h}
}

type varBool struct {
	Value parser.Value
	hash uint64
}

var _ Setter = &varBool{}
var _ Variable = &varBool{}

func (this *varBool) Eval() (bool, error) {
	v, err := this.Value.Bool()
	if err != nil {
		return false, err
	}
	return v, nil
}

func (this *varBool) Hash() uint64 {
	return this.hash
}

func (this *varBool) IsVariable() {}

func (this *varBool) SetValue(v parser.Value) {
	this.Value = v
}

func (this *varBool) String() string {
	return "$bool"
}

//BoolVar returns a variable of type Bool
func BoolVar() *varBool {
	h := uint64(17)
	h = 31*h + 2076426
	return &varBool{hash: h}
}

type varString struct {
	Value parser.Value
	hash uint64
}

var _ Setter = &varString{}
var _ Variable = &varString{}

func (this *varString) Eval() (string, error) {
	v, err := this.Value.String()
	if err != nil {
		return "", err
	}
	return v, nil
}

func (this *varString) Hash() uint64 {
	return this.hash
}

func (this *varString) IsVariable() {}

func (this *varString) SetValue(v parser.Value) {
	this.Value = v
}

func (this *varString) String() string {
	return "$string"
}

//StringVar returns a variable of type String
func StringVar() *varString {
	h := uint64(17)
	h = 31*h + 2486848561
	return &varString{hash: h}
}

type varBytes struct {
	Value parser.Value
	hash uint64
}

var _ Setter = &varBytes{}
var _ Variable = &varBytes{}

func (this *varBytes) Eval() ([]byte, error) {
	v, err := this.Value.Bytes()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (this *varBytes) Hash() uint64 {
	return this.hash
}

func (this *varBytes) IsVariable() {}

func (this *varBytes) SetValue(v parser.Value) {
	this.Value = v
}

func (this *varBytes) String() string {
	return "$[]byte"
}

//BytesVar returns a variable of type Bytes
func BytesVar() *varBytes {
	h := uint64(17)
	h = 31*h + 64671819
	return &varBytes{hash: h}
}
