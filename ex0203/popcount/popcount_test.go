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

func BenchmarkPopCountByLoop(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PopCountByLoop(math.MaxUint64)
	}
}
