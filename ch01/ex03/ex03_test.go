// 非効率な可能性のあるバージョンと strings.Join を使ったバージョンとで、
// 実行時間の差を計測しなさい（1.6 節は time パッケージの一部を説明していますし、11.4 節では体系的に
// 性能評価を行うためのベンチマークテストの書き方を説明しています）。

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
