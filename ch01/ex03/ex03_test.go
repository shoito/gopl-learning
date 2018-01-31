package bench

import (
	"testing"
)

func BenchmarkForLoop(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ForLoop()
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StringsJoin()
	}
}
