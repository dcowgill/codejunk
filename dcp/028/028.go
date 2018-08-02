/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Palantir.

Write an algorithm to justify text. Given a sequence of words and an integer
line length k, return a list of strings which represents each line, fully
justified.

More specifically, you should have as many words as possible in each line. There
should be at least one space between each word. Pad extra spaces when necessary
so that each line has exactly length k. Spaces should be distributed as equally
as possible, with the extra spaces, if any, distributed starting from the left.

If you can only fit one word on a line, then you should pad the right-hand side
with spaces.

Each word is guaranteed not to be longer than k.

For example, given the list of words ["the", "quick", "brown", "fox", "jumps",
"over", "the", "lazy", "dog"] and k = 16, you should return the following:

["the  quick brown", # 1 extra space on the left
"fox  jumps  over", # 2 extra spaces distributed evenly
"the   lazy   dog"] # 4 extra spaces distributed evenly

*/
package dcp028

import (
	"strings"
)

// Distributes the words across lines such that each line has length k.
func justify(words []string, k int) []string {
	var acc, result []string
	for _, s := range words {
		n := totalLen(acc) + len(s) + len(acc)
		if n > k && len(acc) != 0 { // overflow
			result = append(result, join(acc, k))
			acc = nil
		}
		acc = append(acc, s)
	}
	if len(acc) != 0 {
		result = append(result, join(acc, k))
	}
	return result
}

// Joins the strings in ss with a variable number of spaces, such that the total
// length of the resulting string is k.
func join(ss []string, k int) string {
	if len(ss) == 1 {
		return rpad(ss[0], k)
	}
	var (
		w = k - totalLen(ss) // total number of whitespace runes needed
		g = len(ss) - 1      // number of gaps to fill
		n = w / g            // number of spaces per gap
		r = w % g            // extra spaces to distribute among gaps
	)
	if w <= 0 {
		return strings.Join(ss, "")
	}
	var b strings.Builder
	for _, s := range ss[:g] {
		b.WriteString(s)
		for i := 0; i < n; i++ {
			b.WriteByte(' ')
		}
		if r > 0 {
			b.WriteByte(' ')
			r--
		}
	}
	b.WriteString(ss[g])
	return b.String()
}

// Reports the sum of the lengths of the strings in a.
func totalLen(a []string) int {
	n := 0
	for _, s := range a {
		n += len(s)
	}
	return n
}

// Pads s on the right with k-len(s) spaces.
func rpad(s string, k int) string {
	if x := k - len(s); x > 0 {
		return s + strings.Repeat(" ", x)
	}
	return s
}
