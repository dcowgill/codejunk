package dcp069

import (
	"math/rand"
	"testing"
)

func TestRandom(t *testing.T) {
	const ntrials = 1000
	for i := 0; i < ntrials; i++ {
		a := randomInts(50)
		expected := bruteForce(a)
		actual := max3product(a)
		if actual != expected {
			t.Fatalf("max3product(%+v) returned %d, want %d", a, actual, expected)
		}
	}
}

// For comparison testing.
func bruteForce(a []int) int {
	x := 0
	for i := range a {
		for j := range a {
			if i != j {
				for k := range a {
					if i != k && j != k {
						x = max(x, a[i]*a[j]*a[k])
					}
				}
			}
		}
	}
	return x
}

// Generates a slice of n uniformly random ints in [-1000, 1000].
func randomInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rand.Intn(2001) - 1000
	}
	return a
}
