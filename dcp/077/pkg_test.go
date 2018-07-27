package dcp077

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerged(t *testing.T) {
	var tests = []struct {
		spans  []span
		merged []span
	}{
		{nil, nil},
		{[]span{}, nil},
		{[]span{{42, 99}}, []span{{42, 99}}},
		{
			[]span{{20, 25}, {1, 3}, {5, 8}, {4, 10}},
			[]span{{1, 3}, {4, 10}, {20, 25}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.spans), func(t *testing.T) {
			result := merged(tt.spans)
			if !reflect.DeepEqual(result, tt.merged) {
				t.Fatalf("merged returned %+v, want %+v", result, tt.merged)
			}
		})
	}
}
