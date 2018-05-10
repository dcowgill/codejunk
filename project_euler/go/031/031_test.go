package p031

import "testing"

func TestSolution(t *testing.T) {
	const expected = 73682
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
