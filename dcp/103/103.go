/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Square.

Given a string and a set of characters, return the shortest substring containing
all the characters in the set.

For example, given the string "figehaeci" and the set of characters {a, e, i},
you should return "aeci".

If there is no substring containing all the characters in the set, return null.

*/
package dcp103

// runeHistogram is a histogram of runes.
//
// The histogram has a special constraint: to be considered "valid", every rune
// being tracked (see newRuneHistogram) must appear at least once.
type runeHistogram struct {
	set   map[rune]int
	valid bool
}

// Creates a new runeHistogram.
//
// tracking contains the set of runes we're interested in counting. Future calls
// to addRune/removeRune with a rune not in this set will have no effect.
//
// init is an optional list of runes to add. (It's more efficient to add a bunch
// of runes all at once than it is to repeatedly call addRune.)
func newRuneHistogram(tracking, init []rune) *runeHistogram {
	h := &runeHistogram{set: make(map[rune]int)}
	for _, ch := range tracking {
		h.set[ch] = 0
	}
	if len(init) > 0 {
		for _, ch := range init {
			if _, ok := h.set[ch]; ok {
				h.set[ch]++
			}
		}
		h.refreshValidity()
	}
	return h
}

// addRune increments the count for a character.
// This may cause the histogram to become valid.
func (h *runeHistogram) addRune(r rune) {
	oldval, ok := h.set[r]
	if !ok {
		return // not tracking this rune
	}
	h.set[r]++

	// If the count for r just went from 0 to 1, the histogram might have become
	// valid. (But there is no need to check if it is already valid.)
	if oldval == 0 && !h.valid {
		h.refreshValidity()
	}
}

// removeRune decrements the count for a character.
// This may cause the histogram to become invalid.
func (h *runeHistogram) removeRune(r rune) {
	oldval, ok := h.set[r]
	if !ok {
		return // not tracking this rune
	}
	h.set[r]--

	// If the count for r just went from 1 to 0, the histogram is certainly
	// invalid now (whether or not it was valid before).
	if oldval == 1 {
		h.valid = false
	}
}

// isValid reports whether the histogram is currently considered valid. That is,
// whether it contains at least one of each character being tracked.
func (h *runeHistogram) isValid() bool { return h.valid }

// Sets h.valid to its correct value; see isValid.
func (h *runeHistogram) refreshValidity() {
	h.valid = true
	for _, n := range h.set {
		if n == 0 {
			h.valid = false
			return
		}
	}
}

// Returns the shortest substring of s that contains every rune in chars at
// least once. chars is treated like a set, so duplicate runes are ignored.
// Returns an empty string if there is no such substring of s.
func shortestSubstringContainingRunes(s, chars string) string {
	if len(s) == 0 || len(chars) == 0 {
		return ""
	}

	// Create a histogram of runes and initialize it with s.
	a := []rune(s)
	h := newRuneHistogram([]rune(chars), a)
	if !h.isValid() {
		return "" // there is no solution
	}

	// Find the index of the last rune in the string which we cannot remove from
	// the histogram without making it invalid.
	last := len(a) - 1
	for ; last >= 0; last-- {
		h.removeRune(a[last])
		if !h.isValid() {
			h.addRune(a[last]) // put back to restore validity
			break
		}
	}

	// For each index in the string starting at 0, remove the corresponding rune
	// from the histogram, then add runes beyond LAST as long as the histogram
	// is invalid. If the new span is shorter than the old one, remember it.
	best := Range{0, last}
	for i, r := range a {
		h.removeRune(r)
		for !h.isValid() && last < len(a)-1 {
			last++
			h.addRune(a[last])
		}
		if !h.isValid() {
			break
		}
		span := Range{first: i + 1, last: last} // N.B. we removed a[i]
		if span.length() < best.length() {
			best = span
		}
	}

	// Extract and return the substring corresponding to the best span, or an
	// empty string if the spam is empty.
	return s[best.first : best.last+1]
}

// Small helper type to improve readability.
type Range struct {
	first, last int
}

// Reports the number of runes in the span.
func (r Range) length() int {
	return r.last - r.first
}
