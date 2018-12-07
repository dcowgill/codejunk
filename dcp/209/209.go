/*

Good morning! Here's your coding interview problem for today.

This problem was asked by YouTube.

Write a program that computes the length of the longest common subsequence of
three given strings. For example, given "epidemiologist", "refrigeration", and
"supercalifragilisticexpialodocious", it should return 5, since the longest
common subsequence is "eieio".

*/
package dcp209

// lcs3 returns the longest common subsequence of three strings.
func lcs3(s, t, u string) string {
	// Create a cache for solved subproblems.
	type indices struct{ i, j, k int }
	cache := make(map[indices][]rune)

	// Handling multibyte characters in strings is tricky. Runes are easier.
	a := []rune(s)
	b := []rune(t)
	c := []rune(u)

	// f is the recursive function that does all the work.
	// i,j,k are the ending indices of a,b,c; they go from len(a,b,c)->0.
	var f func(i, j, k int) []rune

	// Name the return value so it can be accessed via defer.
	f = func(i, j, k int) (retval []rune) {
		// If any of the strings is empty (i.e. if the current index has fallen
		// off the front of the corresponding slice), there is no common
		// subsequence.
		if i < 0 || j < 0 || k < 0 {
			return nil
		}
		// Try the cache. On a miss, defer a func to populate the cache.
		key := indices{i, j, k}
		if answer, ok := cache[key]; ok {
			return answer
		}
		defer func() {
			cache[key] = retval
		}()
		// If the strings all end with the same rune, the LCS is simply the LCS
		// of the strings with that rune removed, plus the rune.
		if a[i] == b[j] && a[i] == c[k] {
			return append(f(i-1, j-1, k-1), a[i])
		}
		// Otherwise, compute three LCSes by removing the final rune from each
		// string in turn. Return the longest.
		x := f(i-1, j, k)
		y := f(i, j-1, k)
		z := f(i, j, k-1)
		if len(x) >= len(y) && len(x) >= len(z) {
			return x
		}
		if len(y) >= len(x) && len(y) >= len(z) {
			return y
		}
		return z
	}
	return string(f(len(a)-1, len(b)-1, len(c)-1))
}
