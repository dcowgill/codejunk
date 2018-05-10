package floor_division

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

// trunc(n/d)           if n ≥ 0 and d > 0,
// trunc((n + 1)/d) − 1 if n < 0 and d > 0,
// trunc((n − 1)/d) − 1 if n > 0 and d < 0,
// trunc(n/d)           if n ≤ 0 and d < 0

func fdivFloat(n, d int) int {
	return int(math.Floor(float64(n) / float64(d)))
}

func fdivCond(n, d int) int {
	switch {
	case (n >= 0 && d > 0) || (n <= 0 && d < 0):
		return n / d
	case n < 0 && d > 0:
		return (n+1)/d - 1
	case n > 0 && d < 0:
		return (n-1)/d - 1
	}
	panic("impossible")
}

var tests = []struct{ n, d, q int }{
	{5, 1, 5}, {-5, -1, 5}, {5, -1, -5}, {-5, 1, -5},
	{5, 2, 2}, {-5, -2, 2}, {5, -2, -3}, {-5, 2, -3},
	{5, 3, 1}, {-5, -3, 1}, {5, -3, -2}, {-5, 3, -2},
	{5, 4, 1}, {-5, -4, 1}, {5, -4, -2}, {-5, 4, -2},
	{5, 5, 1}, {-5, -5, 1}, {5, -5, -1}, {-5, 5, -1},
	{5, 6, 0}, {-5, -6, 0}, {5, -6, -1}, {-5, 6, -1},
}

var funcs = map[string]func(n, d int) int{
	"fdivFloat": fdivFloat,
	"fdivCond":  fdivCond,
}

func TestFdiv(t *testing.T) {
	for fname, fn := range funcs {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s(%+v)", fname, tt), func(t *testing.T) {
				q := fn(tt.n, tt.d)
				if q != tt.q {
					t.Errorf("%s(%d, %d) returned %d, want %d", fname, tt.n, tt.d, q, tt.q)
				}
			})
		}
	}
}

func BenchmarkFdiv(b *testing.B) {
	var bothPos, bothNeg, negDividend, negDivisor [][2]int
	for i := 0; i < 10; i++ {
		n, d := rand.Intn(1000000)+1, rand.Intn(1000000)+1
		bothPos = append(bothPos, [2]int{n, d})
		bothNeg = append(bothNeg, [2]int{-n, -d})
		negDividend = append(negDividend, [2]int{-n, d})
		negDivisor = append(negDivisor, [2]int{n, -d})
	}
	data := map[string][][2]int{
		"bothPos":     bothPos,
		"bothNeg":     bothNeg,
		"negDividend": negDividend,
		"negDivisor":  negDivisor,
	}
	for fname, fn := range funcs {
		for kind, pairs := range data {
			b.Run(fmt.Sprintf("%s %s", fname, kind), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, p := range pairs {
						fn(p[0], p[1])
					}
				}
			})
		}
	}
}
