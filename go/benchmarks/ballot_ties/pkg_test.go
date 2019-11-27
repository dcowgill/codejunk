package ballot_ties

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	numBallotsPerLen = 10
)

var (
	sizes       = []int{2, 5, 20, 100}
	testBallots [][]Ballot
)

func init() {
	randBallot := func(n int) Ballot {
		b := make(Ballot, n)
		for i := 0; i < n; i++ {
			b[i] = Rank(i + 1)
		}
		rand.Shuffle(len(b), func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		return b
	}
	testBallots = make([][]Ballot, len(sizes))
	for k, size := range sizes {
		for i := 0; i < numBallotsPerLen; i++ {
			testBallots[k] = append(testBallots[k], randBallot(size))
		}
	}
}

func TestWhatever(t *testing.T) {
	for _, b := range testBallots[2] {
		fmt.Printf("%+v\n", b)
	}
}

func BenchmarkHasTies(b *testing.B) {
	fns := []struct {
		name   string
		method func(b Ballot) bool
	}{
		{"brute", Ballot.HasTiesBruteForce},
		{"bits", Ballot.HasTiesBits},
		{"table", Ballot.HasTiesTable},
		{"map", Ballot.HasTiesMap},
	}
	breakopt := 0
	for _, fn := range fns {
		for _, bs := range testBallots {
			b.Run(fmt.Sprintf("%s %03d", fn.name, len(bs[0])), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, b := range bs {
						if fn.method(b) {
							breakopt++
						}
					}
				}
			})
		}
	}
	if breakopt != 0 {
		panic("oops")
	}
}
