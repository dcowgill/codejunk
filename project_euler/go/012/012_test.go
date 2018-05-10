package p012

import "testing"

func TestSolution(t *testing.T) {
	const expected = 76576500
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
