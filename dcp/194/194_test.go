package dcp194

import (
	"math/rand"
	"testing"
)

func TestRandomExamples(t *testing.T) {
	const (
		trials = 10000
		size   = 100
	)
	for i := 0; i < trials; i++ {
		var (
			a = shuffle(oneToN(size))
			b = shuffle(oneToN(size))
			x = slow(a, b)
			y = fast(a, b)
		)
		if x != y {
			t.Fatalf("got %d, want %d (a = %+v, b = %+v)", x, y, a, b)
		}
	}
}

func oneToN(n int) []int {
	a := make([]int, n)
	for j := 0; j < n; j++ {
		a[j] = j + 1
	}
	return a
}

func shuffle(a []int) []int {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}
