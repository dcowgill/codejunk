/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given a list of integers S and a target number k, write a function that returns
a subset of S that adds up to k. If such a subset cannot be made, then return
null.

Integers can appear more than once in the list. You may assume all numbers in
the list are positive.

For example, given S = [12, 1, 61, 5, 9, 2] and k = 24, return [12, 9, 2, 1]
since it sums up to 24.

*/
package dcp042

// Given 'a', a slice of non-negative ints, and a non-negative int 'k', returns
// a subset of 'a' that sums to 'k'. If no such set exists, returns nil.
func subsetSum(a []int, k int) []int {
	// Special case: k must be non-negative.
	if k < 0 {
		return nil
	}

	// Memoize shared sub-problems.
	type key struct{ i, k int }
	memo := make(map[key][]int)

	// Implementation.
	var rec func(i, k int) []int
	rec = func(i, k int) (answer []int) {
		// Handle memoization.
		if x, ok := memo[key{i, k}]; ok {
			return x
		}
		defer func() {
			memo[key{i, k}] = answer
		}()

		// Base cases.
		if k == 0 {
			return []int{}
		}
		if i == len(a) {
			return nil
		}

		// Try all subsets that contain the first element.
		if a[i] <= k {
			b := rec(i+1, k-a[i])
			if b != nil {
				return append(b, a[i])
			}
		}

		// Try all subsets that do not contain the first element.
		b := rec(i+1, k)
		if b != nil {
			return b
		}

		// There is no subset of a[i:] that sums to k.
		return nil
	}
	return rec(0, k)
}
