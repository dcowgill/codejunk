/*

Good morning! Here's your coding interview problem for today.

This problem was asked Microsoft.

Using a read7() method that returns 7 characters from a file, implement readN(n)
which reads n characters.

For example, given a file with the content “Hello world”, three read7() returns
“Hello w”, “orld” and then “”.

*/
package dcp082

import (
	"bufio"
	"io"
)

// Given a function that reads 7 runes at a time from some source, returns a
// function that accepts an int n and reads n runes at a time.
func makeReadN(read7 func() []rune) func(int) []rune {
	var buf []rune
	return func(n int) []rune {
		for len(buf) < n {
			runes := read7()
			if len(runes) == 0 {
				break
			}
			buf = append(buf, runes...)
		}
		if len(buf) < n {
			res := buf
			buf = nil
			return res
		}
		res := buf[:n]
		buf = buf[n:]
		return res
	}
}

// Returns a function that reads up to 7 characters at a time from r. Once r is
// exhausted, returns an empty slice.
func makeRead7(r io.Reader) func() []rune {
	br := bufio.NewReader(r)
	return func() []rune {
		runes := make([]rune, 0, 7)
		for len(runes) < 7 {
			r, _, err := br.ReadRune()
			if err != nil {
				break
			}
			runes = append(runes, r)
		}
		return runes
	}
}
