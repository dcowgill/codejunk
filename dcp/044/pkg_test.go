package dcp044

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNumInversions(t *testing.T) {
	var tests = []struct {
		a []int
		n int
	}{
		{nil, 0},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{2, 4, 1, 3, 5}, 3},
		{[]int{5, 4, 3, 2, 1}, 10},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.a), func(t *testing.T) {
			n := numInversions(tt.a)
			if n != tt.n {
				t.Fatalf("got %d, want %d", n, tt.n)
			}
		})
	}
}

func TestRandomData(t *testing.T) {
	const (
		ntrials = 10000
		size    = 100
	)
	for i := 0; i < ntrials; i++ {
		var (
			a        = randInts(size)
			expected = refImpl(a)
			actual   = numInversions(a)
		)
		if actual != expected {
			t.Fatalf("got %d, want %d", actual, expected)
		}
	}
}

// O(N^2) impl. for testing
func refImpl(a []int) int {
	n := 0
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				n++
			}
		}
	}
	return n
}

// Returns a slice of n random ints.
func randInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rand.Int()
	}
	return a
}
