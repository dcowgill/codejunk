/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a string of round, curly, and square open and closing brackets, return
whether the brackets are balanced (well-formed).

For example, given the string "([])[]({})", you should return true.

Given the string "([)]" or "((()", you should return false.

*/
package dcp027

import "fmt"

func isBalanced(input string) bool {
	var stack runeStack
	for _, ch := range input {
		switch ch {
		case '(', '[', '{':
			stack.push(ch)
		case ')':
			if stack.pop() != '(' {
				return false
			}
		case ']':
			if stack.pop() != '[' {
				return false
			}
		case '}':
			if stack.pop() != '{' {
				return false
			}
		default:
			// It's not clear what to do here, but a simple approach is to treat
			// non-bracket characters as programming errors.
			panic(fmt.Sprintf("unexpected rune in input: %c", ch))
		}
	}
	return len(stack) == 0
}

type runeStack []rune

func (s *runeStack) push(r rune) {
	*s = append(*s, r)
}

func (s *runeStack) pop() rune {
	old := *s
	n := len(old)
	if n == 0 {
		return 0
	}
	r := old[n-1]
	*s = old[:n-1]
	return r
}
