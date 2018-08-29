/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given a word W and a string S, find all starting indices in S which are anagrams of W.

For example, given that W is "ab", and S is "abxaba", return 0, 3, and 4.

*/
package dcp111

func allAnagramIndices(W, S string) []int {
	// Convert the input strings to slices of runes. (Anagrams of their utf8
	// bytes might not give the correct answer.)
	var (
		s = []rune(S)
		w = []rune(W)
		n = len(w)
	)
	if len(w) == 0 || len(s) < n {
		return nil
	}

	// Create two histograms: one of the runes in the word (w),
	// and another of the first n runes in the string (s).
	runesToHistogram := func(a []rune) map[rune]int {
		m := make(map[rune]int)
		for _, r := range a {
			m[r]++
		}
		return m
	}
	wordHist := runesToHistogram(w)
	currHist := runesToHistogram(s[:n])

	// If the two histograms are ever identical, there is an anagram of w
	// beginning at the current position in s.
	anagramHere := func() (x bool) {
		for r, n := range wordHist {
			if n != currHist[r] {
				return false
			}
		}
		return true
	}

	// For each index in s starting from zero (until we are fewer than n runes
	// from the end of s, beyond which there isn't space for an anagram of w),
	// remove (from currHist) the rune at the front and add the rune at the end.
	// Whenever the histograms are identical, record the current index.
	var result []int
	i := 0
	for {
		if anagramHere() {
			result = append(result, i)
		}
		if i >= len(s)-n {
			break
		}
		currHist[s[i]]--
		currHist[s[i+n]]++
		i++
	}
	return result
}
