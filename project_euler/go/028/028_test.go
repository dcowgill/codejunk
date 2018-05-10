package p028

import "testing"

func TestSolution(t *testing.T) {
	const expected = 669171001
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
