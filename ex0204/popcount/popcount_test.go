package popcount

import (
	"math"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PopCount(math.MaxUint64)
	}
}

func BenchmarkPopCountSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PopCountSimple(math.MaxUint64)
	}
}
