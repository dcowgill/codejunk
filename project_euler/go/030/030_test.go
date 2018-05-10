package p030

import "testing"

func TestSolution(t *testing.T) {
	const expected = 443839
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
