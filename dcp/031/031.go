/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

The edit distance between two strings refers to the minimum number of character
insertions, deletions, and substitutions required to change one string to the
other. For example, the edit distance between “kitten” and “sitting” is three:
substitute the “k” for “s”, substitute the “e” for “i”, and append a “g”.

Given two strings, compute the edit distance between them.

*/
package dcp031

// Computes the Levenshtein distance between two strings.
func levenshtein(s, t string) int {
	type pair [2]int
	var memo = make(map[pair]int)
	var dist func(i, j int) int
	dist = func(i, j int) int {
		if r, ok := memo[pair{i, j}]; ok {
			return r
		}
		switch {
		case i == len(s):
			return len(t) - j
		case j == len(t):
			return len(s) - i
		}
		cost := 0
		if s[i] != t[j] {
			cost = 1
		}
		d := min(min(dist(i+1, j)+1, dist(i, j+1)+1), dist(i+1, j+1)+cost)
		memo[pair{i, j}] = d
		return d
	}
	return dist(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
