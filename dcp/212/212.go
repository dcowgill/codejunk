/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Dropbox.

Spreadsheets often use this alphabetical encoding for its columns: "A", "B",
"C", ..., "AA", "AB", ..., "ZZ", "AAA", "AAB", ....

Given a column number, return its alphabetical column id. For example, given 1,
return "A". Given 27, return "AA".

*/
package dcp212

func colNumToID(n int) string {
	if n <= 0 {
		return ""
	}
	// First calculate the number of letters in the ID.
	l := 0
	for x := n; x > 0; x = (x - 1) / 26 {
		l++
	}
	// Store the letters in a rune slice.
	a := make([]rune, l)
	i := l - 1
	for x := n; x > 0; x = (x - 1) / 26 {
		a[i] = 'A' + rune((x-1)%26)
		i--
	}
	return string(a)
}
