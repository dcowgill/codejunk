package p011

import "testing"

func TestSolution(t *testing.T) {
	const expected = 70600674
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
