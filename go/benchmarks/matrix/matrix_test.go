package matrix

import (
	"testing"
)

var allocators = []struct {
	name string
	newf func(r, c int) Matrix
}{
	{"simple", newSimpleMatrix},
	{"dense", newDenseMatrix},
	{"flat", newFlatMatrix},
}

func BenchmarkNewMatrix(b *testing.B) {
	const (
		rows    = 23
		cols    = 17
		nallocs = 10
	)
	matrices := make([]Matrix, nallocs)
	for _, tt := range allocators {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := range matrices {
					matrices[j] = tt.newf(rows, cols)
				}
			}
		})
	}
}

func BenchmarkMatrixMul(b *testing.B) {
	const (
		n = 17
		m = 23
		p = 11
	)
	for _, tt := range allocators {
		b.Run(tt.name, func(b *testing.B) {
			A := tt.newf(n, m)
			B := tt.newf(m, p)
			C := tt.newf(n, p)
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					A.Set(i, j, i*j)
				}
			}
			for i := 0; i < m; i++ {
				for j := 0; j < p; j++ {
					B.Set(i, j, i*j)
				}
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				matrixMul(C, A, B)
			}
		})
	}
}

// Sets c = a*b.
//
// a = n × m
// b = m × p
// c = n × p
//
func matrixMul(c, a, b Matrix) {
	n := a.Rows()
	m := a.Cols()
	p := b.Cols()
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			var x int
			for k := 0; k < m; k++ {
				x += a.At(i, k) * b.At(k, j)
			}
			c.Set(i, j, x)
		}
	}
}
