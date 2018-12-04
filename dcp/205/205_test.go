package dcp205

import (
	"strconv"
	"testing"
)

func TestNextPerm(t *testing.T) {
	var tests = []struct {
		x int
		y int // must == nextPerm(x)
	}{
		{48975, 49578},
		{48957, 48975},
		{49875, 54789},
		{0, 0},
		{1, 1},
		{-123, -123},
		{8421, 8421},
		{123456789, 123456798},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.x), func(t *testing.T) {
			y := nextPerm(tt.x)
			if y != tt.y {
				t.Fatalf("got %d, want %d", y, tt.y)
			}
		})
	}
}
