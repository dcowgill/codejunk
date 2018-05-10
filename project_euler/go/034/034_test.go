package p034

import "testing"

func TestSolution(t *testing.T) {
	const expected = 40730
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
