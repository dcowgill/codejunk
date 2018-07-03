package dcp035

import "testing"

func TestSortRGB(t *testing.T) {
	var tests = []struct {
		p, q string // input and result
	}{
		{"", ""},
		{"R", "R"},
		{"G", "G"},
		{"B", "B"},
		{"BGR", "RGB"},
		{"GBRRBRG", "RRRGGBB"},
		{"BBGRRBGRRRBGRBGRBGRGRRGRGRBRGBR", "RRRRRRRRRRRRRRGGGGGGGGGBBBBBBBB"},
	}
	for _, tt := range tests {
		t.Run(tt.p, func(t *testing.T) {
			a := []rune(tt.p)
			sortRGB(a)
			q := string(a)
			if q != tt.q {
				t.Fatalf("sortRGB returned %q, want %q", q, tt.q)
			}
		})
	}
}
