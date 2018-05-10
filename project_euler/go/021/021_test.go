package p021

import "testing"

func TestSolution(t *testing.T) {
	const expected = 31626
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
