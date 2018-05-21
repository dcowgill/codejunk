/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Airbnb.

Given a list of integers, write a function that returns the largest sum of
non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 8] should return 12, since we pick 4 and 8. [5, 1, 1, 5]
should return 10, since we pick 5 and 5.

Follow-up: Can you do this in O(N) time and constant space?

*/
package dcp009

func maxSumNonAdj(xs []int) int {
	var inc, exc int
	for _, x := range xs {
		inc, exc = max(exc, x+exc), max(exc, inc)
	}
	return max(inc, exc)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
