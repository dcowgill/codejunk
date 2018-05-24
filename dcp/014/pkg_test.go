package dcp014

import (
	"fmt"
	"testing"
)

func TestEstimatePi(t *testing.T) {
	const (
		expected = "3.141"
		trials   = 1000000
	)
	s := fmt.Sprintf("%.3f", estimatePi(trials))
	if s != expected {
		t.Fatalf("estimatePi(%d) returned %q, want %q", trials, s, expected)
	}
}
