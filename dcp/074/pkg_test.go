package dcp074

import "testing"

func TestNumTimes(t *testing.T) {
	const (
		maxN = 10
		maxX = 2 * maxN * maxN
	)
	for n := 1; n <= maxN; n++ {
		for x := 1; x <= maxX; x++ {
			expected := bruteForce(n, x)
			actual := numTimes(n, x)
			if actual != expected {
				t.Errorf("numTimes(%d, %d) returned %d, want %d", n, x, actual, expected)
			}
		}
	}
}

// Straightforward implementation for testing.
func bruteForce(n, x int) int {
	c := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i*j == x {
				c++
			}
		}
	}
	return c
}
