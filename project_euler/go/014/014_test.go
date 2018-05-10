package p014

import "testing"

func TestSolution(t *testing.T) {
	const expected = 837799
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
