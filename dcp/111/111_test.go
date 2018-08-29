package dcp111

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllAnagramIndices(t *testing.T) {
	var tests = []struct {
		word, s string
		indices []int
	}{
		{"", "", nil},
		{"ab", "", nil},
		{"ab", "aa", nil},
		{"ab", "abab", []int{0, 1, 2}},
		{"ab", "abxaba", []int{0, 3, 4}},
		{"xyz", "xyzzyxyxz", []int{0, 3, 6}},
		{"abc", "cbaebabacd", []int{0, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.word, tt.s), func(t *testing.T) {
			indices := allAnagramIndices(tt.word, tt.s)
			if !reflect.DeepEqual(indices, tt.indices) {
				t.Fatalf("got %+v, want %+v", indices, tt.indices)
			}
		})
	}
}
