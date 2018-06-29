package dcp029

import (
	"math/rand"
	"testing"
)

func TestEncode(t *testing.T) {
	var tests = []struct {
		x, y string // original, encoded
	}{
		{"AAAABBBCCDAA", "4A3B2C1D2A"},
		{"A", "1A"},
		{"XY", "1X1Y"},
		{"qqqqqqqqqqqqqqqqq", "17q"},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.x, func(t *testing.T) {
			y := encode(tt.x)
			if y != tt.y {
				t.Fatalf("encode(%q) returned %q, want %q", tt.x, y, tt.y)
			}
		})
	}
}

func TestRandomDecodeEncode(t *testing.T) {
	const numTrials = 10000
	for i := 0; i < numTrials; i++ {
		a := randInput()
		b := decode(encode(a))
		if a != b {
			t.Fatalf("decode(encode(%q)) returned %q", a, b)
		}
	}
}

func randInput() string {
	var a []rune
	for i := 0; i < 10; i++ {
		r := 'A' + rand.Intn(26)
		for n := rand.Intn(20); n >= 0; n-- {
			a = append(a, rune(r))
		}
	}
	return string(a)
}
