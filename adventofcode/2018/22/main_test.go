package main

import "testing"

func TestErosionLevel(t *testing.T) {
	cs := newCaveSystem(510, Point{}, Point{10, 10})
	var tests = []struct {
		p  Point
		el int
	}{
		{Point{0, 0}, 510},
		{Point{1, 0}, 17317},
		{Point{0, 1}, 8415},
		{Point{1, 1}, 1805},
		{Point{10, 10}, 510},
	}
	for _, tt := range tests {
		el := cs.erosionLevel(tt.p)
		if el != tt.el {
			t.Fatalf("erosionLevel(%+v) returned %d, want %d", tt.p, el, tt.el)
		}
	}
}
