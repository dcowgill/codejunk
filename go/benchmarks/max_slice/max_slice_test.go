package max_slice

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkMaxSlice(b *testing.B) {
	sizes := []int{1, 2, 8, 32, 128}
	tests := make([][]int, len(sizes))
	for i, n := range sizes {
		a := make([]int, n)
		for j := range a {
			a[j] = rand.Int()
		}
		tests[i] = a
	}
	var funcs = []struct {
		name string
		call func([]int) int
	}{
		{"f1", f1},
		{"f2", f2},
		{"f3", f3},
		{"f4", f4},
		{"f5", f5},
	}
	for _, a := range tests {
		for _, fn := range funcs {
			b.Run(fmt.Sprintf("%s %d", fn.name, len(a)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					fn.call(a)
				}
			})
		}
	}
}
