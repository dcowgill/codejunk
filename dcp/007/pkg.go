/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the
number of ways it can be decoded.

For example, the message '111' would give 3, since it could be decoded as 'aaa',
'ka', and 'ak'.

You can assume that the messages are decodable. For example, '001' is not
allowed.

*/
package dcp007

import "fmt"

var numToRune []rune // maps 1..26 to a..z

func init() {
	numToRune = make([]rune, 27)
	for i := 1; i <= 26; i++ {
		numToRune[i] = rune('a' + i - 1)
	}
}

// toDigits converts a string like "123" to []int{1, 2, 3}.
// Panics if any rune in s is not in the range ['1', '9'].
func toDigits(s string) []int {
	var digits []int
	for _, r := range s {
		if r < '1' || r > '9' {
			panic(fmt.Sprintf("invalid input rune: %v", r))
		}
		digits = append(digits, int(r-'0'))
	}
	return digits
}

// decode returns all possible interpretations of digits.
func decode(digits []int) [][]rune {
	// If there are no more digits, return one empty slice of runes, so that
	// the caller will have something to which it can prepend its heads.
	if len(digits) == 0 {
		return [][]rune{nil}
	}
	var result [][]rune
	// Append the leading digit to every tail.
	for _, tail := range decode(digits[1:]) {
		head := []rune{numToRune[digits[0]]}
		result = append(result, append(head, tail...))
	}
	// Append the leading two digits (if possible) to every tail.
	if len(digits) >= 2 {
		if n := digits[0]*10 + digits[1]; n < len(numToRune) {
			for _, tail := range decode(digits[2:]) {
				head := []rune{numToRune[n]}
				result = append(result, append(head, tail...))
			}
		}
	}
	return result
}

// runesToStrings converts every rune-slice in rrs to a string.
func runesToStrings(rrs [][]rune) []string {
	result := make([]string, len(rrs))
	for i, rr := range rrs {
		result[i] = string(rr)
	}
	return result
}
