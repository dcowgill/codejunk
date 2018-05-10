package atomics

import (
	"sync/atomic"
	"testing"
)

func BenchmarkAddInt32(b *testing.B) {
	v := new(int32)
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(v, 1)
	}
}

func BenchmarkAddInt64(b *testing.B) {
	v := new(int64)
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(v, 1)
	}
}
