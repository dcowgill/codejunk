package p008

import "testing"

func TestSolution(t *testing.T) {
	const expected = 40824
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
