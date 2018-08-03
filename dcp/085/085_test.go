package dcp085

import (
	"math/rand"
	"testing"
)

func TestBittyWithRandomValues(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		x := rand.Int31()
		y := rand.Int31()
		if v := bitty(x, y, 1); v != x {
			t.Fatalf("bitty(%d, %d, 1) returned %d, want %d", x, y, v, x)
		}
		if v := bitty(x, y, 0); v != y {
			t.Fatalf("bitty(%d, %d, 0) returned %d, want %d", x, y, v, y)
		}
	}
}
