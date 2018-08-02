/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a list of integers, return the largest product that can be made by
multiplying any three integers.

For example, if the list is [-10, -10, 5, 2], we should return 500, since that's
-10 * -10 * 5.

You can assume the list has at least three integers.

*/
package dcp069

import "sort"

// This should use big.Int to avoid overflow.
func max3product(a []int) int {
	n := len(a)
	b := make([]int, n)
	copy(b, a)
	sort.Ints(b)
	return max(b[0]*b[1]*b[n-1], b[n-3]*b[n-2]*b[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
