package dcp070

import (
	"testing"
)

func TestNth(t *testing.T) {
	// Use brute force, generate a list of the first N numbers whose digits sum to 10.
	const N = 100
	var answers []int
	for i := 0; len(answers) < N; i++ {
		if digitsSumTo10(i) {
			answers = append(answers, i)
		}
	}
	// Test the answers against the nth() function.
	for i, x := range answers {
		y := nth(i + 1)
		if x != y {
			t.Fatalf("nth(%d) returned %d, want %d", i+1, y, x)
		}
	}
}
