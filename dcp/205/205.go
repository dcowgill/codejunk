/*

Good morning! Here's your coding interview problem for today.

This problem was asked by IBM.

Given an integer, find the next permutation of it in absolute order. For
example, given 48975, the next permutation would be 49578.

*/
package dcp205

import "fmt"

// Returns the next permutation of the digits in "n" in lexicographic order.
// If "n" is non-positive, returns "n".
// Example: if "n" is 48975, returns 49578.
func nextPerm(n int) int {
	if n <= 0 {
		return n
	}
	a := toDigits(n)
	p := findLongestIncreasingPrefix(a)
	if p < 0 {
		return n // at final perm
	}
	reverse(a[:p])
	for i := p - 1; i >= 0; i-- {
		if a[i] > a[p] {
			a[i], a[p] = a[p], a[i]
			return digitsToInt(a)
		}
	}
	panic(fmt.Sprintf("nextPerm(%d): unexpected error", n))
}

// Returns the base-10 digits of "n" in reverse.
// Example: if "n" is 61945, returns [5, 4, 9, 1, 6].
func toDigits(n int) []uint8 {
	count := 1
	for x := n; x >= 10; x /= 10 {
		count++
	}
	a := make([]uint8, count)
	for i := range a {
		a[i] = uint8(n % 10)
		n /= 10
	}
	return a
}

// Performs the inverse of toDigits.
// Example: if a is [4, 3, 2, 1], returns 1234.
func digitsToInt(a []uint8) int {
	n := 0
	p := 1
	for _, x := range a {
		n += p * int(x)
		p *= 10
	}
	return n
}

// Returns the index ending the longest increasing prefix of "a".
// Example: if "a" is [1, 2, 3, 4, 2, 5], returns 4.
func findLongestIncreasingPrefix(a []uint8) int {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return i
		}
	}
	return -1
}

// Reverses "a" in place.
func reverse(a []uint8) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
