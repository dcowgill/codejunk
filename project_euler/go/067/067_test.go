package p067

import "testing"

func TestSolution(t *testing.T) {
	const expected = 7273
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
