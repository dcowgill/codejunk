package uint8_set

import (
	"math/rand"
	"testing"
)

type iset interface {
	set(n uint8)
	has(n uint8) bool
}

var kinds = []struct {
	name string
	cons func() iset
}{
	{"set1", func() iset { return newSet1() }},
	{"set2", func() iset { return newSet2() }},
	{"set3", func() iset { return newSet3() }},
	{"set4", func() iset { return newSet4() }},
	{"set5", func() iset { return newSet5() }},
}

func TestSets(t *testing.T) {
	for _, k := range kinds {
		t.Run(k.name, func(t *testing.T) {
			const N = 20
			xs := make([]uint8, N)
			for i := 0; i < N; i++ {
				xs[i] = uint8(rand.Intn(256))
			}
			m := make(map[uint8]bool)
			s := k.cons()
			for _, x := range xs {
				m[x] = true
				s.set(x)
			}
			for i := 0; i < 256; i++ {
				x := uint8(i)
				if s.has(x) != m[x] {
					t.Fatalf("s.has(%d) returned %v, want %v", x, s.has(x), m[x])
				}
			}
		})
	}
}

func BenchmarkSets(b *testing.B) {
	const N = 20
	xs := make([]uint8, N)
	for i := 0; i < N; i++ {
		xs[i] = uint8(rand.Intn(256))
	}
	for _, k := range kinds {
		b.Run(k.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := k.cons()
				for _, x := range xs {
					s.set(x)
				}
				n := 0
				for i := 0; i < 256; i++ {
					if s.has(uint8(i)) {
						n++
					}
				}
				if n != len(xs) {
					b.Fatalf("counted %d values, expected %d", n, len(xs))
				}
			}
		})
	}
}
