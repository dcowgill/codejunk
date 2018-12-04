package dcp206

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestApplyPermInPlaceRandomly(t *testing.T) {
	const trials = 10000
	const size = 5
	for i := 0; i < trials; i++ {
		var (
			a  = letters(size)
			p  = randPerm(size)
			b  = applyPerm(a, p)
			a1 = copyStrings(a)
			p1 = copyInts(p)
		)
		applyPermInPlace(a1, p1)
		if !reflect.DeepEqual(a1, b) {
			t.Fatalf("applyPermInPlace(%q, %+v): got %+v, want %+v", a, p, a1, b)
		}
	}
}

// Returns a slice containing the first "n" capital ascii letters.
func letters(n int) []string {
	if n > 26 {
		panic("n must be <=26")
	}
	a := make([]string, n)
	for i := range a {
		a[i] = fmt.Sprintf("%c", 'A'+i)
	}
	return a
}

// Generates a random permutation of [0, n-1].
func randPerm(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func copyStrings(a []string) []string {
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func copyInts(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
