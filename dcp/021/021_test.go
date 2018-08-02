package dcp021

import (
	"fmt"
	"testing"
)

func TestMinRooms(t *testing.T) {
	var tests = []struct {
		intervals []interval
		n         int // minimum rooms
	}{
		{[]interval{{30, 75}, {0, 50}, {60, 150}}, 2},
		{[]interval{{30, 75}, {0, 50}, {60, 150}, {40, 70}}, 3},
		{[]interval{{30, 75}, {0, 50}, {60, 150}, {0, 20}, {80, 200}}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.intervals), func(t *testing.T) {
			n := minRooms(tt.intervals)
			if n != tt.n {
				t.Fatalf("minRooms returned %d, want %d", n, tt.n)
			}
		})
	}
}
