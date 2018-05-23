/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

Given an integer k and a string s, find the length of the longest substring that
contains at most k distinct characters.

For example, given s = "abcba" and k = 2, the longest substring with k distinct
characters is "bcb".

*/
package dcp013

// Reports the range [begin, end) corresponding to the longest substring of s
// that contains at most k distinct characters.
func longestSubstringOfAtMostKDistinctRunes(s string, k int) (int, int) {
	hist := newHistogram(k)
	var curr, best window
	rs := []rune(s)
	for i, chr := range rs {
		hist.add(chr)
		for !hist.valid() {
			hist.remove(rs[curr.begin])
			curr.begin++
		}
		curr.end = i + 1
		if curr.size() > best.size() {
			best = curr
		}
	}
	return best.begin, best.end
}

// Histogram of runes.
// Valid as long as it contains k or fewer distinct runes.
type histogram struct {
	m map[rune]int
	k int
}

func newHistogram(k int) *histogram { return &histogram{make(map[rune]int), k} }
func (h *histogram) valid() bool    { return len(h.m) <= h.k }
func (h *histogram) add(r rune)     { h.m[r]++ }
func (h *histogram) remove(r rune) {
	if h.m[r] == 1 {
		delete(h.m, r)
	} else {
		h.m[r]--
	}
}

// A window represents the half-open range [begin, end).
type window struct{ begin, end int }

// Reports the window size.
func (w window) size() int { return w.end - w.begin }
