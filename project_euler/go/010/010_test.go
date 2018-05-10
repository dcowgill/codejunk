package p010

import "testing"

func TestSolution(t *testing.T) {
	const expected = 142913828922
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %d, want %d", answer, expected)
	}
}
