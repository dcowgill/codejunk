package dcp018

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubarrayMaxes(t *testing.T) {
	var tests = []struct {
		a []int // input array
		k int   // size of subarray
		r []int // answer
	}{
		{[]int{10, 5, 2, 7, 8, 7}, 3, []int{10, 7, 8, 8}},
		{[]int{1, 2, 3, 1, 4, 5, 2, 3, 6}, 3, []int{3, 3, 4, 5, 5, 5, 6}},
		{[]int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}, 4, []int{10, 10, 10, 15, 15, 90, 90}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v k=%d", tt.a, tt.k), func(t *testing.T) {
			var r []int
			subarrayMaxes(tt.a, tt.k, func(x int) { r = append(r, x) })
			if !reflect.DeepEqual(r, tt.r) {
				t.Fatalf("got %+v, want %+v\n", r, tt.r)
			}
		})
	}
}

// TODO: test randomly generated arrays and k-values against bruteForce()

// Reference implementation for testing.
func bruteForce(a []int, k int, emit func(int)) {
	for i := 0; i <= len(a)-k; i++ {
		emit(greatest(a[i : i+k]))
	}
}

// Reports the largest value in a.
func greatest(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}
