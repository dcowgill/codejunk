package p001

import "testing"

func TestSolution(t *testing.T) {
	const expected = 233168
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
