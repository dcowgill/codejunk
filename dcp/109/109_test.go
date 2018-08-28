package dcp109

import (
	"strconv"
	"testing"
)

func TestSwapBitPairs(t *testing.T) {
	var tests = []struct {
		sx, sy string // sx=input, sy=swapped
	}{
		{"00000000", "00000000"},
		{"11111111", "11111111"},
		{"11110000", "11110000"},
		{"00001111", "00001111"},
		{"10101010", "01010101"},
		{"10000000", "01000000"},
		{"10010110", "01101001"},
		{"10000001", "01000010"},
		{"10011001", "01100110"},
	}
	for _, tt := range tests {
		t.Run(tt.sx, func(t *testing.T) {
			x1 := b(tt.sx)
			y1 := b(tt.sy)
			y2 := swapBitPairs(x1)
			x2 := swapBitPairs(y2) // ensure the inverse works, too
			if y1 != y2 {
				t.Fatalf("swapBitPairs(%08b) returned %08b, want %08b", x1, y2, y1)
			}
			if x1 != x2 {
				t.Fatalf("swapBitPairs(%08b) returned %08b, want %08b", y2, x2, x1)
			}
		})
	}
}

// Parses an 8-bit binary number. Panics if the input is invalid.
func b(s string) uint8 {
	n, err := strconv.ParseUint(s, 2, 8)
	if err != nil {
		panic(err)
	}
	return uint8(n)
}
