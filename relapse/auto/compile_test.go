package auto

import (
	"github.com/katydid/katydid/relapse/parser"
	"testing"
)

func benchCompile(b *testing.B, str string) {
	st, err := parser.ParseGrammar(str)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Compile(st)
	}
}

func BenchmarkCompileSmall(b *testing.B) {
	benchCompile(b, "*")
}

func BenchmarkCompileMedium(b *testing.B) {
	benchCompile(b, ".A.B.C: (A:* & B:* & C:* & D:*)")
}

func BenchmarkCompileLarge(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* & B:* & C:* & D:*) | {A.B.C:* ; D:* ; C:*})")
}

func BenchmarkCompileLarger(b *testing.B) {
	benchCompile(b, ".A.B.C: ((A:* & B:* & C:* & D:*) | {A.B.C:* ; D:* ; C:* ; F.D:* ; G:* ; H:*; I:*})")
}
