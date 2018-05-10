package p003

import "testing"

func TestSolution(t *testing.T) {
	const expected = 6857
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
