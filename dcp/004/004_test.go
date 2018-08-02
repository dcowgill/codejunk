package dcp004

import (
	"fmt"
	"testing"
)

var tests = []struct {
	a []int
	n int
}{
	{[]int{3, 4, -1, 1}, 2},
	{[]int{1, 2, 0}, 3},
	{[]int{2, 4, 3, 4, 2, 4, 3, 2}, 1},
	{[]int{0, -1, -2, -3, 0, -4}, 1},
}

var funcs = []struct {
	name string
	f    func([]int) int
}{
	{"refImpl", refImpl},
	{"linear", linear},
}

func TestAll(t *testing.T) {
	for _, fn := range funcs {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s(%+v)", fn.name, tt.a), func(t *testing.T) {
				a := append(tt.a[:0:0], tt.a...) // make a copy
				n := fn.f(a)
				if n != tt.n {
					t.Fatalf("got %d, want %d", n, tt.n)
				}
			})
		}
	}
}
