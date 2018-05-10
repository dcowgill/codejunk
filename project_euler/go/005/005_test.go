package p005

import "testing"

func TestSolution(t *testing.T) {
	const expected = 232792560
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
