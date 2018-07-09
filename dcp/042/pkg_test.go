package dcp042

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSubsetSum(t *testing.T) {
	var tests = []struct {
		a []int
		k int
		r []int
	}{
		{[]int{-5}, -5, nil},
		{[]int{}, 42, nil},
		{[]int{1, 2, 3}, 0, []int{}},
		{[]int{12, 1, 61, 5, 9, 2}, 24, []int{1, 2, 9, 12}},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := subsetSum(tt.a, tt.k)
			if !intSetEqual(r, tt.r) {
				t.Fatalf("subsetSum(%+v, %d) returned %+v, want %+v", tt.a, tt.k, r, tt.r)
			}
		})
	}
}

// Reports whether two sets of integers are identical.
// Assumes the sets do not contain duplicates.
func intSetEqual(a, b []int) bool {
	return reflect.DeepEqual(sorted(a), sorted(b))
}

// Returns a sorted copy of a.
func sorted(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(b)
	return b
}
