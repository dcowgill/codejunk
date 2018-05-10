package p019

import "testing"

func TestSolution(t *testing.T) {
	const expected = 171
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
