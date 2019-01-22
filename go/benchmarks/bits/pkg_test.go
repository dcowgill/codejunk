package bits

import (
	"fmt"
	"math/bits"
	"testing"
)

// [1, 8]
func BenchmarkIteratorsLowValues(b *testing.B) {
	var xs = []uint64{0, 1, 8, 18, 64, 129, 255}
	for _, x := range xs {
		binary := fmt.Sprintf("%08b", uint8(x))
		n := 0
		b.Run("pkg bits "+binary, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := makeIter1(x)
				for v := it.next(); v >= 0; v = it.next() {
					n += v
				}
			}
		})
		b.Run("loop "+binary, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := makeIter2(x, 8)
				for v := it.next(); v >= 0; v = it.next() {
					n += v
				}
			}
		})
		// Put positions into a slice
		var a []int
		{
			it := makeIter1(x)
			for v := it.next(); v >= 0; v = it.next() {
				a = append(a, v)
			}
		}
		b.Run("slice "+binary, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := makeSliceIter(a)
				for v := it.next(); v >= 0; v = it.next() {
					n += v
				}
			}
		})
	}
}

func BenchmarkZeros8(b *testing.B) {
	var xs = []uint8{0, 1, 2, 8, 128}
	n := 0
	for _, x := range xs {
		b.Run(fmt.Sprintf("leading %08b", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n += bits.LeadingZeros8(x)
			}
		})
		b.Run(fmt.Sprintf("trailing %08b", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n += bits.TrailingZeros8(x)
			}
		})
	}
}

func BenchmarkZeros16(b *testing.B) {
	var xs = []uint16{0, 1, 2, 8, 32, 256, 32767}
	n := 0
	for _, x := range xs {
		b.Run(fmt.Sprintf("leading %016b", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n += bits.LeadingZeros16(x)
			}
		})
		b.Run(fmt.Sprintf("trailing %016b", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n += bits.TrailingZeros16(x)
			}
		})
	}
}

func BenchmarkZeros32(b *testing.B) {
	var xs = []uint32{0, 1, 2, 256, 32767, 1048575, 1073741823}
	n := 0
	for _, x := range xs {
		b.Run(fmt.Sprintf("leading %032b", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n += bits.LeadingZeros32(x)
			}
		})
		b.Run(fmt.Sprintf("trailing %032b", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n += bits.TrailingZeros32(x)
			}
		})
	}
}

func BenchmarkZeros64(b *testing.B) {
	var xs = []uint64{0, 1, 2, 256, 32767, 1048575, 1073741823, 1099511627775, 1152921504606846975}
	n := 0
	for _, x := range xs {
		b.Run(fmt.Sprintf("leading %064b", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n += bits.LeadingZeros64(x)
			}
		})
		b.Run(fmt.Sprintf("trailing %064b", x), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				n += bits.TrailingZeros64(x)
			}
		})
	}
}
