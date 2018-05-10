package intset64

import "testing"

var (
	ranges = []Range{
		{0, 999, 5},
		{1, 15, 2},
		{16, 31, 4},
	}
	low  = 1
	high = 31
)

func BenchmarkApplyBitSet(b *testing.B) {
	x := 0
	for i := 0; i < b.N; i++ {
		x += len(applyBitSet(ranges, low, high))
	}
}

func BenchmarkApplyLinearSet(b *testing.B) {
	x := 0
	for i := 0; i < b.N; i++ {
		x += len(applyLinearSet(ranges, low, high))
	}
}

func BenchmarkApplySortedSet(b *testing.B) {
	x := 0
	for i := 0; i < b.N; i++ {
		x += len(applySortedSet(ranges, low, high))
	}
}

func BenchmarkApplySortUnique(b *testing.B) {
	x := 0
	for i := 0; i < b.N; i++ {
		x += len(applySortUnique(ranges, low, high))
	}
}

func BenchmarkApplyMap(b *testing.B) {
	x := 0
	for i := 0; i < b.N; i++ {
		x += len(applyMap(ranges, low, high))
	}
}
