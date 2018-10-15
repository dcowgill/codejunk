/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

Given a string, determine whether any permutation of it is a palindrome.

For example, carrace should return true, since it can be rearranged to form
racecar, which is a palindrome. daily should return false, since there's no
rearrangement that can form a palindrome.

*/
package dcp157

import (
	"sort"
)

// hasPalindromePermutation reports whether the runes in the string can be
// rearranged to form a palindrome.
func hasPalindromePermutation(s string) bool {
	// Solution: if every rune in the string - with at most one exception -
	// appears an even number of times, the string can form a palindrome.

	// First convert the string to a slice of runes, then sort it.
	a := []rune(s)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	// Then count each run of runes in the slice, returning false if we
	// encounter a second odd-count run. If we reach the end of the string
	// having seen zero or one odd-count runs, return true.
	hasOddRun := false
	pos := 0 // current offset in a[]
	for pos < len(a) {
		i := pos
		for i < len(a) && a[i] == a[pos] {
			i++
		}
		if runLen := i - pos; runLen%2 == 1 {
			if hasOddRun {
				return false
			}
			hasOddRun = true
		}
		pos = i
	}
	return true
}
