/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in
linear time and constant space. In other words, find the lowest positive
integer that does not exist in the array. The array can contain
duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0]
should give 3.

You can modify the input array in-place.

*/
package dcp004

import (
	"sort"
)

// Ugly but straightforward. Sorts the list, so O(n*log(n)).
func refImpl(a []int) int {
	sort.Ints(a)
	a = posTail(a)
	x := 1
	i := 0
	for i < len(a) {
		j := i
		for j < len(a) && a[j] == x {
			j++
		}
		if j == i {
			return x
		}
		i = j
		x = x + 1
	}
	return x
}

// Assumes a is sorted.
// Returns the tail of a containing positive ints.
func posTail(a []int) []int {
	for i, x := range a {
		if x > 0 {
			return a[i:]
		}
	}
	return nil
}

// We know a priori the answer must be in the range [1, len(a)].
// Therefore we can use the array itself to keep track of whether we
// have encountered each integer. Care must be taken to ignore both
// non-positive and duplicate integers.
func linear(a []int) int {
	for i := 0; i < len(a); i++ {
		x := a[i]
		if x <= 0 || x >= len(a) {
			continue // value not in range; ignore
		}
		j := x - 1 // target slot for a[i]
		if a[j] == x {
			// The target is already set to its correct value, which
			// means a[i] is a duplicate.
			continue
		}
		a[j], a[i] = a[i], a[j]
		i-- // restart loop at same position
	}
	// Report the first slot that does not contain its corresponding value.
	for i, x := range a {
		if x != i+1 {
			return i + 1
		}
	}
	return 0
}
