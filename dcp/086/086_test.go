package dcp086

import (
	"log"
	"testing"
)

func TestMinParensToRemove(t *testing.T) {
	var tests = []struct {
		s string // input
		n int    // expected result
	}{
		{"()())()", 1},
		{")(", 2},
		{"", 0},
		{"(((", 3},
		{")))", 3},
		{")()(", 2},
	}
	for _, tt := range tests {
		if n := minParensToRemove(tt.s); n != tt.n {
			log.Fatalf("minParensToRemove(%q) returned %d, want %d", tt.s, n, tt.n)
		}
	}
}
