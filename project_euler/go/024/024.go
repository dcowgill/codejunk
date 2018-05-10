/*

A permutation is an ordered arrangement of objects. For example, 3124 is
one possible permutation of the digits 1, 2, 3 and 4. If all of the
permutations are listed numerically or alphabetically, we call it
lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits
0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

*/
package p024

import (
	"sort"
	"strconv"
)

const N = 1000 * 1000

func solve() string {
	var perm []int
	n := 0
	permutations(
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		func(a []int) bool {
			n++
			if n == N {
				perm = a
				return false
			}
			return true
		})
	s := ""
	for _, i := range perm {
		s += strconv.Itoa(i)
	}
	return s
}

// Calls visit for each permutation of a, stopping if visit returns false.
// http://en.wikipedia.org/wiki/Permutations#Generation_in_lexicographic_order
func permutations(a []int, visit func(a []int) bool) {
	sort.Ints(a)
	n := len(a) - 1
	for {
		if !visit(a) {
			break
		}

		var k int
		for k = n - 1; k >= 0; k-- {
			if a[k] < a[k+1] {
				break
			}
		}
		if k < 0 {
			break
		}

		var l int
		for l = n; ; l-- {
			if a[k] < a[l] {
				break
			}
		}

		a[k], a[l] = a[l], a[k]

		for i, j := k+1, n; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	}
}
