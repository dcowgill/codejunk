package dcp088

import (
	"fmt"
	"testing"
)

var tests = [][2]int32{
	{27, 4},
	{99, 42},
	{29873, 198},
	{298734, 1736},
	{129140163, 17},
	{2147483647, 2},
	{5764801, 999},
}

func TestDiv(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			var (
				a  = tt[0]
				b  = tt[1]
				q1 = a / b
			)
			if q2 := slowdiv(a, b); q1 != q2 {
				t.Errorf("slowdiv(%d, %d) returned %d, want %d", a, b, q2, q1)
			}
			if q2 := fastdiv(a, b); q1 != q2 {
				t.Errorf("fastdiv(%d, %d) returned %d, want %d", a, b, q2, q1)
			}
		})
	}
}

func BenchmarkDiv(b *testing.B) {
	var funcs = []struct {
		name string
		call func(a, b int32) int32
	}{
		{"builtin", func(a, b int32) int32 { return a / b }},
		{"slow", slowdiv},
		{"fast", fastdiv},
	}
	for _, fn := range funcs {
		b.Run(fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tt := range tests {
					fn.call(tt[0], tt[1])
				}
			}
		})
	}
}
