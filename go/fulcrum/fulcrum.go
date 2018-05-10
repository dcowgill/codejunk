package fulcrum

// TODO: handle overflow panic

import "fmt"

// Fulcrum: given a sequence of integers, returns the index i that minimizes
// |sum(seq[..i]) - sum(seq[i..])|. Does this in O(n) time and O(n) memory.
func Fulcrum(xs []int) int {
	var l, r int64
	for _, x := range xs { // O(n)
		r = add64(r, int64(x))
	}
	md := diff(l, r)
	mi := 0
	for i := 1; i <= len(xs); i++ { // O(n)
		x := int64(xs[i-1])
		l = add64(l, x)
		r = add64(r, -x)
		if d := diff(l, r); d < md {
			md, mi = d, i
		}
	}
	return mi
}

// Panics on overflow.
func add64(a, b int64) int64 {
	c := a + b
	if (c > a) == (b > 0) {
		return c
	}
	panic(fmt.Sprintf("overflow: %d + %d", a, b))
}

func diff(a, b int64) uint64 {
	if a > b {
		return uint64(a - b)
	}
	return uint64(b - a)
}
