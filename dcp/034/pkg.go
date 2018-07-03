/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Quora.

Given a string, find the palindrome that can be made by inserting the fewest
number of characters as possible anywhere in the word. If there is more than one
palindrome of minimum length that can be made, return the lexicographically
earliest one (the first one alphabetically).

For example, given the string "race", you should return "ecarace", since we can
add three letters to it (which is the smallest amount to make a palindrome).
There are seven other palindromes that can be made from "race" by adding three
letters, but "ecarace" comes first alphabetically.

As another example, given the string "google", you should return "elgoogle".

*/
package dcp034

// minPalindrome returns the palindrome formed by making the fewest insertions
// into s. In case of ties, chooses the lexicographically least palindrome.
func minPalindrome(s string) string {
	return rec(s, "", 0).s
}

// Recursive implementation of minPalindrome.
func rec(a string, r string, x int) result {
	n := len(a)
	if n <= 1 {
		return result{r + a + r, x}
	}
	best := rec(a[1:], a[:1], 1)
	{
		other := rec(a[:n-1], a[n-1:], 1)
		if other.beats(best) {
			best = other
		}
	}
	if a[0] == a[n-1] {
		other := rec(a[1:n-1], a[:1], 0)
		if other.beats(best) {
			best = other
		}
	}
	return result{r + best.s + r, best.n + x}
}

// Intermediate result produced by rec().
type result struct {
	s string // accumulator; always a palindrome
	n int    // number of insertions
}

// Reports whether result a is preferable to result b.
func (a result) beats(b result) bool {
	return a.n < b.n || (a.n == b.n && a.s < b.s)
}
