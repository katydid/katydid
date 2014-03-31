package funcs

type Bool interface {
	Eval([]byte) bool
}

type String interface {
	Eval([]byte) string
}

type Int64 interface {
	Eval([]byte) int64
}

type constBool struct {
	v bool
}

func NewBool(v bool) Bool {
	return &constBool{v}
}

func (this *constBool) Eval(buf []byte) bool {
	return this.v
}

type constString struct {
	v string
}

func NewString(v string) String {
	return &constString{v}
}

func (this *constString) Eval(buf []byte) string {
	return this.v
}

type constInt64 struct {
	v int64
}

func NewInt64(v int64) Int64 {
	return &constInt64{v}
}

func (this *constInt64) Eval(buf []byte) int64 {
	return this.v
}
