package dcp173

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	var tests = []struct {
		input  M
		result map[string]int
	}{
		{nil, map[string]int{}},
		{M{}, map[string]int{}},
		{M{"a": 99}, map[string]int{"a": 99}},
		{M{"a": M{"b": M{"c": 42}}}, map[string]int{"a.b.c": 42}},
		{
			M{"key": 3, "foo": M{"a": 5, "bar": M{"baz": 8}}},
			map[string]int{"foo.a": 5, "foo.bar.baz": 8, "key": 3},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.input), func(t *testing.T) {
			result := flatten(tt.input)
			if !reflect.DeepEqual(result, tt.result) {
				t.Fatalf("got %+v, want %+v", result, tt.result)
			}
		})
	}
}
