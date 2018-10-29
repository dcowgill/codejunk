/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Dropbox.

Given a string s and a list of words words, where each word is the same length,
find all starting indices of substrings in s that is a concatenation of every
word in words exactly once.

For example, given s = "dogcatcatcodecatdog" and words = ["cat", "dog"], return
[0, 13], since "dogcat" starts at index 0 and "catdog" starts at index 13.

Given s = "barfoobazbitbyte" and words = ["dog", "cat"], return [] since there
are no substrings composed of "dog" and "cat" in s.

The order of the indices does not matter.

*/
package dcp172

import "fmt"

// findWords returns the indices in "s", in ascending order, at which begins a
// substring containing every string in "words", in any order, exactly once.
// All of the strings in "words" must have equal length and none may be empty;
// otherwise, the function will panic.
func findWords(s string, words []string) []int {
	// It's OK to provide an empty list of words; don't panic.
	if len(words) == 0 {
		return nil
	}

	// Since the strings in "words" have equal length, we can easily compute the
	// length of the full substring we're searching for in "s".
	wordLen := len(words[0])
	allWordsLen := wordLen * len(words)

	// If the word length is zero, the loop below will not exit; panicking is
	// better. Also reject the input if any two words have unequal length.
	if wordLen == 0 {
		panic("words[0] has length zero")
	}
	for i, w := range words {
		if len(w) != wordLen {
			panic(fmt.Sprintf("words[%d] has length %d, want %d", i, len(w), wordLen))
		}
	}

	// Count the number of occurrences of each string in "words".
	wordCount := newWordCounter(words)

	// We can stop checking possible starting indices in "s" before the end: we
	// know the substring is "allWordsLen" bytes long.
	var answer []int
nextStartPos:
	for i := 0; i <= len(s)-allWordsLen; i++ {
		numWordsLeft := len(words)
		wordCount := wordCount.clone() // Creates a copy of "wordCount" for this loop iteration.

		// This loop breaks the substring s[i : i+allWordsLen] into chunks of
		// length "wordLen". If a chunk is not in "wordCount" or has a count of
		// zero, quit and restart the loop at the next starting position in "s".
		// Otherwise, decrement its count and continue.
		for j := i; j < i+allWordsLen; j += wordLen {
			wordHere := s[j : j+wordLen]
			if n := wordCount[wordHere]; n <= 0 {
				continue nextStartPos
			}
			wordCount[wordHere]--
			numWordsLeft--
		}

		// Did we find an instance of every string in "words"?
		if numWordsLeft == 0 {
			answer = append(answer, i)
		}
	}

	// Done.
	return answer
}

// Helper type.
type WordCounter map[string]int

// Counts the occurrences of each string in "words".
func newWordCounter(words []string) WordCounter {
	m := make(WordCounter, len(words))
	for _, s := range words {
		m[s]++
	}
	return m
}

// Returns a deep copy of "m".
func (m WordCounter) clone() WordCounter {
	copy := make(WordCounter, len(m))
	for k, v := range m {
		copy[k] = v
	}
	return copy
}
