package dcp172

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	var tests = []struct {
		s     string
		words []string
		idxs  []int
	}{
		{"", nil, nil},
		{"hello", nil, nil},
		{"dogcatcatcodecatdog", []string{"cat", "dog"}, []int{0, 13}},
		{"barfoobazbitbyte", []string{"cat", "dog"}, nil},
		{"abcxyzxyzabc", []string{"xyz", "abc", "xyz"}, []int{0, 3}},
		{"barfoobazbitbyte", []string{"byt", "bar", "baz", "foo", "bit"}, []int{0}},
		{"aaabbbcccdddeeefff", []string{"ddd", "eee", "ccc", "bbb"}, []int{3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %v", tt.s, tt.words), func(t *testing.T) {
			idxs := findWords(tt.s, tt.words)
			if !reflect.DeepEqual(idxs, tt.idxs) {
				t.Fatalf("got %+v, want %+v", idxs, tt.idxs)
			}
		})
	}
}
