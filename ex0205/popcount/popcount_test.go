package popcount

import (
	"math"
	"math/bits"
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

func BenchmarkPopCountClear(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PopCountClear(math.MaxUint64)
	}
}

func BenchmarkBuiltin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bits.OnesCount64(math.MaxUint64)
	}
}
