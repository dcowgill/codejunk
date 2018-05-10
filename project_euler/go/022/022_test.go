package p022

import "testing"

func TestSolution(t *testing.T) {
	const expected = 871198282
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
