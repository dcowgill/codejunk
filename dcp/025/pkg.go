/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Implement regular expression matching with the following special characters:

. (period) which matches any single character
* (asterisk) which matches zero or more of the preceding element

That is, implement a function that takes in a string and a valid regular
expression and returns whether or not the string matches the regular expression.

For example, given the regular expression "ra." and the string "ray", your
function should return true. The same regular expression on the string "raymond"
should return false.

Given the regular expression ".*at" and the string "chat", your function should
return true. The same regular expression on the string "chats" should return
false.

*/
package dcp025

// Super simple (and inefficient) regexp matching.
// Only supports the "." and "*" symbols.
func match(pattern string, input string) bool {
	i := 0
	j := 0
	for i < len(pattern) || j < len(input) {
		// Handle a wildcard in the pattern.
		if i < len(pattern)-1 && pattern[i+1] == '*' {
			for k := j; k <= len(input); k++ {
				if match(pattern[i+2:], input[k:]) {
					return true
				}
				if pattern[i] == '.' || pattern[i] == input[k] {
					continue
				}
				break
			}
			return false
		}
		// Handle a single-character match.
		if i < len(pattern) && j < len(input) {
			if pattern[i] == '.' || pattern[i] == input[j] {
				i++
				j++
				continue
			}
		}
		// No match.
		return false
	}
	// Success: exhausted both pattern and input.
	return true
}
