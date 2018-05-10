package p018

import "testing"

func TestSolution(t *testing.T) {
	const expected = 1074
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
