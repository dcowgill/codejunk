/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given an array of integers where every integer occurs three times except for one
integer, which only occurs once, find and return the non-duplicated integer.

For example, given [6, 1, 3, 3, 3, 6, 6], return 1.
Given [13, 19, 13, 13], return 19.

Do this in O(N) time and O(1) space.

*/
package dcp040

func xor3(x, y int) int {
	n := 0
	p := 1
	for x != 0 || y != 0 {
		n += p * ((x%3 + y%3) % 3)
		p *= 3
		x /= 3
		y /= 3
	}
	return n
}

func findNonDuplicate(a []int) int {
	if len(a) == 0 {
		return 0
	}
	n := a[0]
	for i := 1; i < len(a); i++ {
		n = xor3(n, a[i])
	}
	return n
}
