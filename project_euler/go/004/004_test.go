package p004

import "testing"

func TestSolution(t *testing.T) {
	const expected = 906609
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
