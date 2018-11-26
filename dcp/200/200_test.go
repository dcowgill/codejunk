package dcp200

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinStabSet(t *testing.T) {
	var tests = []struct {
		intervals []interval
		minstab   []int
	}{
		{nil, nil},
		{[]interval{{1, 3}}, []int{3}},
		{[]interval{{9, 12}, {4, 5}, {1, 4}, {7, 9}}, []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.intervals), func(t *testing.T) {
			minstab := minStabSet(tt.intervals)
			if !reflect.DeepEqual(minstab, tt.minstab) {
				t.Fatalf("got %+v, want %+v", minstab, tt.minstab)
			}
		})
	}
}
