package p013

import "testing"

func TestSolution(t *testing.T) {
	const expected = "5537376230"
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %q, want %q", answer, expected)
	}
}
