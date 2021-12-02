package d10

import (
	"adventofcode2021/lib"
	"fmt"
	"sort"
)

func Run() {
	lib.Run(10, part1, part2)
}

func part1() int64 {
	var pointsByDelim = map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	computeLineScore := func(line string) int {
		var stack bytestack
		for i := 0; i < len(line); i++ {
			switch c := line[i]; c {
			case '(', '[', '{', '<':
				stack.push(c)
			case ')', ']', '}', '>':
				expect := invertDelim(c)
				if stack.top() != expect {
					return pointsByDelim[c]
				}
				stack.pop()
			}
		}
		return 0
	}
	points := 0
	for _, line := range realInput {
		points += computeLineScore(line)
	}
	return int64(points)
}

func part2() int64 {
	var pointsByDelim = map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	solve := func(line string) int {
		var stack bytestack
		for i := 0; i < len(line); i++ {
			switch c := line[i]; c {
			case '(', '[', '{', '<':
				stack.push(c)
			case ')', ']', '}', '>':
				expect := invertDelim(c)
				if stack.top() != expect {
					return -1 // corrupted
				}
				stack.pop()
			}
		}
		score := 0
		var autocomplete []byte
		for stack.size() != 0 {
			c := invertDelim(stack.pop())
			autocomplete = append(autocomplete, c)
			score = (score * 5) + pointsByDelim[c]
		}
		return score
	}
	var scores []int
	for _, line := range realInput {
		if n := solve(line); n >= 0 {
			scores = append(scores, n)
		}
	}
	sort.Ints(scores)
	return int64(scores[len(scores)/2])
}

type bytestack struct{ a []byte }

func (s *bytestack) size() int   { return len(s.a) }
func (s *bytestack) push(b byte) { s.a = append(s.a, b) }
func (s *bytestack) top() byte {
	if s.size() == 0 {
		return '?'
	}
	return s.a[s.size()-1]
}
func (s *bytestack) pop() byte {
	n := s.size()
	if n == 0 {
		panic("pop on empty stack")
	}
	c := s.a[n-1]
	s.a = s.a[:n-1]
	return c
}

func invertDelim(c byte) byte {
	switch c {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	}
	panic(fmt.Sprintf("invalid delim: %+v", c))
}
