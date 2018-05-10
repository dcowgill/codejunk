package p024

import "testing"

func TestSolution(t *testing.T) {
	const expected = "2783915460"
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %q, want %q", answer, expected)
	}
}
