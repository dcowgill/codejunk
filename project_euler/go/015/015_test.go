package p015

import "testing"

func TestSolution(t *testing.T) {
	const expected = "137846528820"
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %q, want %q", answer, expected)
	}
}
