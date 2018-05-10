package p027

import "testing"

func TestSolution(t *testing.T) {
	const expected = -59231
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
