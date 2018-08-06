/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given a string of parentheses, write a function to compute the minimum number of
parentheses to be removed to make the string valid (i.e. each open parenthesis
is eventually closed).

For example, given the string "()())()", you should return 1. Given the string
")(", you should return 2, since we must remove all of them.

*/
package dcp086

func minParensToRemove(s string) int {
	var (
		n int // number of closing-parens that do not balance an opening-paren
		d int // current depth of valid parenthesis nesting
	)
	for _, ch := range s {
		switch ch {
		case '(':
			d++
		case ')':
			if d > 0 {
				d--
			} else {
				n++
			}
		}
	}
	return n + d
}
