/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Suppose you are given two lists of n points, one list p1, p2, ..., pn on the
line y = 0 and the other list q1, q2, ..., qn on the line y = 1. Imagine a set
of n line segments connecting each point pi to qi. Write an algorithm to
determine how many pairs of the line segments intersect.

*/
package dcp194

import "sort"

// Exhaustively checks all pairs of line: O(N**2)
// For testing.
func slow(ps, qs []int) int {
	n := 0
	for i := 0; i < len(ps)-1; i++ {
		for j := i + 1; j < len(ps); j++ {
			if i != j {
				if (ps[i] > ps[j] && qs[i] < qs[j]) || ps[i] < ps[j] && qs[i] > qs[j] {
					n++
				}
			}
		}
	}
	return n
}

// Divide-and-conquer solution: O(NlogN)
// N.B. sorts both "ps" and "qs" in place.
func fast(ps, qs []int) int {
	sort.Sort(slicePair{ps, qs})
	return mergesort(qs, make([]int, len(qs)))
}

// Sorts "a", using "b" as a work array.
// Returns the number of inversions performed.
func mergesort(a, b []int) int {
	if len(a) <= 1 {
		return 0
	}
	copy(b, a)
	var (
		m = len(b) / 2          // pivot index
		l = b[:m]               // left half
		r = b[m:]               // right half
		x = mergesort(l, a[:m]) // inversions on left
		y = mergesort(r, a[m:]) // inversions on right
		z = merge(a, l, r)      // merge inversions
	)
	return x + y + z
}

// Merges "a" and "b" into "d"; len(d) must equal len(a)+len(b).
// Assumes both "a" and "b" are sorted.
// Returns the number of inversions performed.
func merge(d, a, b []int) int {
	var (
		m = len(a)
		n = len(b)
		i = 0 // current index in a
		j = 0 // current index in b
		c = 0 // inversion count
	)
	for k := range d {
		if i < m && (j == n || a[i] < b[j]) {
			d[k] = a[i]
			i++
		} else {
			d[k] = b[j]
			j++
			if i < m {
				c += m - i
			}
		}
	}
	return c
}

// Sorts "ps" in ascending order, while mirroring the swaps to "qs", so that the
// original pairings of [ps[i], qs[i]] are preserved. Example: given
// ps=[3,2,1,4] and qs=[4,5,6,7], sorting ps causes qs to become [6,5,4,7].
type slicePair struct{ ps, qs []int }

func (a slicePair) Len() int           { return len(a.ps) }
func (a slicePair) Less(i, j int) bool { return a.ps[i] < a.ps[j] }
func (a slicePair) Swap(i, j int) {
	a.ps[i], a.ps[j] = a.ps[j], a.ps[i]
	a.qs[i], a.qs[j] = a.qs[j], a.qs[i]
}
