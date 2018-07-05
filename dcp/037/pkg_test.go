package dcp037

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestPowerset(t *testing.T) {
	var tests = []struct {
		set []int
		ps  [][]int
	}{
		{nil, [][]int{{}}},
		{[]int{1}, [][]int{{}, {1}}},
		{[]int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.set), func(t *testing.T) {
			ps := powerset(tt.set)
			sortSets(ps)
			if !reflect.DeepEqual(ps, tt.ps) {
				t.Fatalf("powerset returned %#v, want %#v", ps, tt.ps)
			}
		})
	}
}

// To simplify equality comparisons of sets of sets.
func sortSets(a [][]int) {
	sort.Slice(a, func(i, j int) bool {
		x := a[i]
		y := a[j]
		switch {
		case len(x) < len(y):
			return true
		case len(x) > len(y):
			return false
		default:
			for k := range x {
				switch {
				case x[k] < y[k]:
					return true
				case x[k] > y[k]:
					return false
				}
			}
			return false
		}
	})
}
