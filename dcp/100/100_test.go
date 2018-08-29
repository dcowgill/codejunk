package dcp100

import (
	"fmt"
	"testing"
)

func TestPathSteps(t *testing.T) {
	var tests = []struct {
		path []pt
		dist int
	}{
		{pts(0, 0), 0},
		{pts(0, 0, 1, 0), 1},
		{pts(0, 0, 1, 1), 1},
		{pts(0, 0, 0, 1), 1},
		{pts(0, 0, 1, 1, 1, 2), 2},
		{pts(0, 0, 1, 1, 2, 2), 2},
		{pts(0, 0, 1, 2, 3, 4, 5, 6), 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.path), func(t *testing.T) {
			dist := pathSteps(tt.path)
			if dist != tt.dist {
				t.Fatalf("pathSteps returned %d, want %d", dist, tt.dist)
			}
		})
	}
}

// Shorthand for making points.
func pts(v ...int) []pt {
	var path []pt
	for i := 0; i < len(v); i += 2 {
		path = append(path, pt{v[i], v[i+1]})
	}
	return path
}
