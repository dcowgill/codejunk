package dcp102

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	var tests = []struct {
		a []int
		k int
		b []int
	}{
		{nil, 0, nil},
		{nil, 1, nil},
		{[]int{0}, 0, nil},
		{[]int{1}, 0, nil},

		{[]int{0, 3}, 3, []int{0, 3}},
		{[]int{3, 0}, 3, []int{3}},
		{[]int{1, 0, 3}, 3, []int{0, 3}},

		{[]int{7}, 6, nil},
		{[]int{7}, 7, []int{7}},
		{[]int{7}, 8, nil},

		{[]int{2, 4, 8}, 1, nil},
		{[]int{2, 4, 8}, 3, nil},
		{[]int{2, 4, 8}, 5, nil},
		{[]int{2, 4, 8}, 7, nil},
		{[]int{2, 4, 8}, 9, nil},

		{[]int{1, 2, 3}, 0, nil},
		{[]int{1, 2, 3}, 1, []int{1}},
		{[]int{1, 2, 3}, 2, []int{2}},
		{[]int{1, 2, 3}, 3, []int{1, 2}},
		{[]int{1, 2, 3}, 4, nil},
		{[]int{1, 2, 3}, 5, []int{2, 3}},
		{[]int{1, 2, 3}, 6, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 7, nil},

		{[]int{1, 2, 3, 4, 5}, 9, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 10, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 11, nil},
		{[]int{1, 2, 3, 4, 5}, 12, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 13, nil},

		{[]int{1, 2, 13}, 13, []int{13}},
		{[]int{1, 2, 13}, 14, nil},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("array=%v sum=%d", tt.a, tt.k), func(t *testing.T) {
			b := subarraySum(tt.a, tt.k)
			if !reflect.DeepEqual(b, tt.b) {
				t.Fatalf("subArraySum(%+v, %d) returned %+v, want %+v", tt.a, tt.k, b, tt.b)
			}
		})
	}
}
