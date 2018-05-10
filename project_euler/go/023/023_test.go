package p023

import "testing"

func TestSolution(t *testing.T) {
	const expected = 4179871
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
