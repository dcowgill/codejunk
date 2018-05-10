/*

We shall say that an n-digit number is pandigital if it makes use of all
the digits 1 to n exactly once; for example, the 5-digit number, 15234,
is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 x 186 = 7254,
containing multiplicand, multiplier, and product is 1 through 9
pandigital.

Find the sum of all products whose multiplicand/multiplier/product
identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to
only include it once in your sum.

*/
package p032

import (
	"sort"
)

func solve() int {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	pandigitalProducts := make(map[int]bool)
	permutations(
		digits,
		func(a []int) bool {
			a2n := func(a []int) int {
				p := 1
				n := 0
				for i := len(a) - 1; i >= 0; i-- {
					n += p * a[i]
					p *= 10
				}
				return n
			}
			for i := 1; i < len(a)-2; i++ {
				for j := i + 1; j < len(a)-1; j++ {
					x, y, z := a2n(a[:i]), a2n(a[i:j]), a2n(a[j:])
					if x*y == z {
						pandigitalProducts[z] = true
					}
				}
			}
			return true
		})

	sum := 0
	for k, _ := range pandigitalProducts {
		sum += k
	}
	return sum
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
