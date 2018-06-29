/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

Run-length encoding is a fast and simple method of encoding strings. The basic
idea is to represent repeated successive characters as a single count and
character. For example, the string "AAAABBBCCDAA" would be encoded as
"4A3B2C1D2A".

Implement run-length encoding and decoding. You can assume the string to be
encoded have no digits and consists solely of alphabetic characters. You can
assume the string to be decoded is valid.

*/
package dcp029

import (
	"strconv"
	"strings"
)

// encode compresses s using simple run-length encoding.
func encode(s string) string {
	var (
		curr rune
		n    int
		b    strings.Builder
	)
	for _, r := range s {
		if r != curr {
			if n > 0 {
				b.WriteString(strconv.Itoa(n))
				b.WriteRune(curr)
			}
			curr, n = r, 0
		}
		n++
	}
	if n > 0 {
		b.WriteString(strconv.Itoa(n))
		b.WriteRune(curr)
	}
	return b.String()
}

// decode reverses encode.
func decode(s string) string {
	var (
		count []rune
		b     strings.Builder
	)
	for _, r := range s {
		// Accumulate digits and skip to the next rune.
		if r >= '0' && r <= '9' {
			count = append(count, r)
			continue
		}
		// Convert the runes in count to an integer. This loop is equivalent to
		// strconv.Atoi(string(count)), but is faster and doesn't allocate.
		p := 1
		n := 0
		for i := len(count) - 1; i >= 0; i-- {
			n += p * int(count[i]-'0')
			p *= 10
		}
		// Write n copies of r to the output buffer.
		for i := 0; i < n; i++ {
			b.WriteRune(r)
		}
		count = count[:0] // reuse storage
	}
	return b.String()
}
