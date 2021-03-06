package slice_tricks

import (
	"fmt"
	"testing"
)

type T struct{ X, Y int }

func extendOneLine(a []*T, n int) []*T {
	x := n - len(a)
	if x <= 0 {
		return a
	}
	return append(a, make([]*T, x)...)
}

func extendForLoop(a []*T, n int) []*T {
	for len(a) < n {
		a = append(a, nil)
	}
	return a
}

func extendCopy(a []*T, n int) []*T {
	if len(a) >= n {
		return a
	}
	b := make([]*T, n)
	copy(b, a)
	return b
}

var b1conf = []struct {
	a, b int
}{
	{0, 0},
	{0, 1},
	{0, 2},
	{0, 10},
	{0, 100},
	{0, 1000},
	{100, 101},
	{100, 102},
	{100, 110},
	{100, 200},
	{100, 1100},
}

func BenchmarkExtend(b *testing.B) {
	aa := make([][]*T, len(b1conf))
	for i, c := range b1conf {
		aa[i] = make([]*T, c.a)
	}
	b.ResetTimer()
	for i, c := range b1conf {
		a := aa[i]
		b.Run(fmt.Sprintf("%d to %d one line", c.a, c.b), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				extendOneLine(a, c.b)
			}
		})
		b.Run(fmt.Sprintf("%d to %d loop", c.a, c.b), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				extendForLoop(a, c.b)
			}
		})
		b.Run(fmt.Sprintf("%d to %d copy", c.a, c.b), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				extendCopy(a, c.b)
			}
		})
	}
}

func BenchmarkRepeatedExtend(b *testing.B) {
	const (
		low  = 1
		high = 500
		step = 10
	)
	b.Run("one line", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var a []*T
			for i := low; i <= high; i += step {
				a = extendOneLine(a, i)
			}
		}
	})
	b.Run("loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var a []*T
			for i := low; i <= high; i += step {
				a = extendForLoop(a, i)
			}
		}
	})
	b.Run("copy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var a []*T
			for i := low; i <= high; i += step {
				a = extendCopy(a, i)
			}
		}
	})
}
