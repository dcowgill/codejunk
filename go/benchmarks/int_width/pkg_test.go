package int_width

import (
	"fmt"
	"testing"
)

type widthFn func(int) int

var funcs = []struct {
	name string
	fn   widthFn
}{
	{"log10", log10},
	{"itoa", itoa},
}

func TestAll(t *testing.T) {
	var tests = []struct {
		n, w int
	}{
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{50, 2},
		{99, 2},
		{100, 3},
		{999, 3},
		{1000, 4},
	}
	for _, fn := range funcs {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s(%d)", fn.name, tt.n), func(t *testing.T) {
				w := fn.fn(tt.n)
				if w != tt.w {
					t.Fatalf("%s(%d) returned %d, want %d", fn.name, tt.n, w, tt.w)
				}
			})
		}
	}
}

var g = 0

func BenchmarkAll(b *testing.B) {
	var tests = []int{0, 1, 15, 50, 150, 500, 1500}
	for _, fn := range funcs {
		for _, n := range tests {
			b.Run(fmt.Sprintf("%s(%d)", fn.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					g += fn.fn(n)
				}
			})
		}
	}
}
