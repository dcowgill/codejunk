package p032

import "testing"

func TestSolution(t *testing.T) {
	const expected = 45228
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
