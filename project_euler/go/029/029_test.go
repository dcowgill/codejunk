package p029

import "testing"

func TestSolution(t *testing.T) {
	const expected = 9183
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
