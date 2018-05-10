package fulcrum

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFulcrum(t *testing.T) {
	var tests = []struct {
		seq []int
		idx int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 4},
		{[]int{1000, 1, 2, 1000}, 2},
		// TODO...
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v\n", tt.seq), func(t *testing.T) {
			i := Fulcrum(tt.seq)
			if i != tt.idx {
				t.Fatalf("Fulcrum returned %d, want %d", i, tt.idx)
			}
		})
	}
}

// Generate many random sequences, then exhaustively compare every possible
// fulcrum of the sequence to the return value of Fulcrum(sequence).
func TestFulcrumRand(t *testing.T) {
	const ntrials = 1000
	for i := 0; i < ntrials; i++ {
		xs := randInts(rand.Intn(1000) + 1)
		md := diff(0, sum(xs))
		mj := 0
		for j := 0; j < len(xs); j++ { // try every fulcrum
			d := diff(sum(xs[:j]), sum(xs[j:]))
			if d < md {
				md, mj = d, j
			}
		}
		if j := Fulcrum(xs); mj != j {
			t.Fatalf("Fulcrum(%+v) returned %d, want %d", xs, j, mj)
		}
	}
}

func randInts(n int) []int {
	const N = 1000 * 1000 * 1000
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = rand.Intn(N) - N/2
		// xs[i] = int(rand.Uint64())
	}
	return xs
}

func sum(xs []int) int64 {
	var s int64
	for _, x := range xs {
		s = add64(s, int64(x))
	}
	return s
}
