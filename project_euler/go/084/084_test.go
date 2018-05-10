package p084

import "testing"

func TestSolution(t *testing.T) {
	const expected = "101524"
	if answer := solve(); answer != expected {
		t.Fatalf("solve() returned %q, want %q", answer, expected)
	}
}
