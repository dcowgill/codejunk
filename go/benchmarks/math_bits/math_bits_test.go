package math_bits

import (
	"math/bits"
	"testing"
)

var uint64s = []uint64{
	5831,
	90128,
	760816297579193874,
	0,
	128,
	9223372036854775807,
	32,
	31,
	293084,
	8798623486291,
	127,
	20394823,
	3480834,
	1,
}

func BenchmarkMathBitsOnesCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := 0
		for _, v := range uint64s {
			n += bits.OnesCount64(v)
		}
	}
}

func BenchmarkKernighanOnesCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := 0
		for _, v := range uint64s {
			n += kernighanOnesCount64(v)
		}
	}
}

func BenchmarkNaiveBitCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := 0
		for _, v := range uint64s {
			n += naiveOnesCount64(v)
		}
	}
}

func kernighanOnesCount64(x uint64) int {
	n := 0
	for x != 0 {
		x &= (x - 1)
		n++
	}
	return n
}

func naiveOnesCount64(x uint64) int {
	var c, i uint64
	for i = 0; i < 64; i++ {
		c += (x >> i) & 1
	}
	return int(c)
}
