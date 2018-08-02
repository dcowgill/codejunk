package dcp002

import (
	"fmt"
	"reflect"
	"testing"
)

var tests = []struct {
	a []int
	b []int
}{
	// Examples provided in problem description.
	{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	{[]int{3, 2, 1}, []int{2, 3, 6}},

	// Test empty inputs.
	{nil, []int{}},
	{[]int{}, []int{}},
}

var funcs = []struct {
	name string
	f    func([]int) []int
}{
	{"quadratic", quadratic},
	{"linearWithDivision", linearWithDivision},
	{"linearNoDivision", linearNoDivision},
}

func TestAll(t *testing.T) {
	for _, fn := range funcs {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s(%+v)", fn.name, tt.a), func(t *testing.T) {
				b := fn.f(tt.a)
				if !reflect.DeepEqual(b, tt.b) {
					t.Fatalf("got %+v, want %+v", b, tt.b)
				}
			})
		}
	}
}
