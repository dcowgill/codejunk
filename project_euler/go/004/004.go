/*

A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 x 99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/
package p004

import (
	"strconv"
)

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func solve() int {
	max := 0
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			p := i * j
			if p > max && isPalindrome(p) {
				max = p
			}
		}
	}
	return max
}
