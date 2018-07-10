/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

Given a string, find the longest palindromic contiguous substring. If there are
more than one with the maximum length, return any one.

For example, the longest palindromic substring of "aabcdcb" is "bcdcb". The
longest palindromic substring of "bananas" is "anana".

*/
package dcp046

func longestPalindrome(s string) string {
	type pair [2]int
	var (
		memo = make(map[pair]pair)
		rec  func(l, r int) (int, int)
	)
	rec = func(l, r int) (int, int) {
		if r, ok := memo[pair{l, r}]; ok {
			return r[0], r[1]
		}
		if l >= r {
			return l, r
		}
		x, y := rec(l+1, r-1)
		if s[l] == s[r] && x == l+1 && y == r-1 {
			x, y = l, r
		}
		{
			l1, r1 := rec(l, r-1)
			if r1-l1 > y-x {
				x, y = l1, r1
			}
		}
		{
			l1, r1 := rec(l+1, r)
			if r1-l1 > y-x {
				x, y = l1, r1
			}
		}
		memo[pair{l, r}] = pair{x, y}
		return x, y
	}
	l, r := rec(0, len(s)-1)
	if l <= r {
		return s[l : r+1]
	}
	return ""
}
