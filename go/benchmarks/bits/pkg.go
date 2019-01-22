package bits

import "math/bits"

//====================

type bitIter1 uint64

func makeIter1(x uint64) bitIter1 { return bitIter1(x) }

func (it *bitIter1) next() int {
	x := uint64(*it)
	if x == 0 {
		return -1
	}
	n := bits.TrailingZeros64(x)
	*it &= bitIter1(x - 1)
	return n
}

//====================

type bitIter2 struct {
	set uint64
	cur int
	end int
}

func makeIter2(x uint64, high int) bitIter2 {
	return bitIter2{set: x, cur: 0, end: high + 1}
}

func (it *bitIter2) next() int {
	for it.cur < it.end {
		if it.set&(1<<uint64(it.cur)) != 0 {
			x := it.cur
			it.cur++
			return x
		}
		it.cur++
	}
	return -1
}

//====================

type sliceIter struct {
	a   []int
	pos int
}

func makeSliceIter(a []int) sliceIter {
	return sliceIter{a: a}
}

func (it *sliceIter) next() int {
	if it.pos < len(it.a) {
		x := it.a[it.pos]
		it.pos++
		return x
	}
	return -1
}
