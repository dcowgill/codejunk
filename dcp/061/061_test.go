package dcp061

import (
	"math"
	"math/big"
	"testing"
)

const MAX = 15 // 15*15 = 8785478146473916433 = 2^62.9 (fits in int64)

func TestFastPow(t *testing.T) {
	for x := 0; x <= MAX; x++ {
		for y := 0; y <= MAX; y++ {
			expected := naive(x, y)
			actual := fastpow(x, y)
			if actual != expected {
				t.Fatalf("fastpow(%d, %d) returned %d, want %d", x, y, actual, expected)
			}
		}
	}
}

func BenchmarkPow(b *testing.B) {
	funcs := []struct {
		name string
		fn   func(int, int) int
	}{
		{"builtin", builtin},
		{"naive", naive},
		{"fastpow", fastpow},
		{"fastpowRec", fastpowRec},
	}
	var ext = 0
	for _, def := range funcs {
		b.Run(def.name, func(b *testing.B) {
			ext = 0
			for i := 0; i < b.N; i++ {
				for x := 0; x < MAX; x++ {
					for y := 0; y < MAX; y++ {
						ext += def.fn(x, y)
					}
				}
			}
		})
	}
}

func BenchmarkBigPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < MAX; x++ {
			for y := 0; y < MAX; y++ {
				bigPow(x, y)
			}
		}
	}

}

func builtin(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func naive(x, y int) int {
	r := 1
	for i := 0; i < y; i++ {
		r *= x
	}
	return r
}

func bigPow(x, y int) *big.Int {
	r := big.NewInt(1)
	b := big.NewInt(int64(x))
	for i := 0; i < y; i++ {
		r = r.Mul(r, b)
	}
	return r
}
