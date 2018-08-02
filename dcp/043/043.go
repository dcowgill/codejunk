/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Amazon.

Implement a stack that has the following methods:

push(val), which pushes an element onto the stack

pop(), which pops off and returns the topmost element of the stack. If there are
no elements in the stack, then it should throw an error or return null.

max(), which returns the maximum value in the stack currently. If there are no
elements in the stack, then it should throw an error or return null.

Each method should run in constant time.

*/
package dcp043

type stack struct {
	values []int // contents of the stack
	maxes  []int // offsets into values
}

// Reports the number of elements in the stack.
func (s *stack) size() int {
	return len(s.values)
}

// Pushes v onto the stack.
func (s *stack) push(v int) {
	if s.size() == 0 || v > s.max() {
		s.maxes = append(s.maxes, s.size())
	}
	s.values = append(s.values, v)
}

// Pops the top of the stack.
func (s *stack) pop() int {
	n := s.size()
	v := s.values[n-1]
	s.values = s.values[:n-1]
	if s.maxIdx() == n-1 {
		s.maxes = s.maxes[:n-1]
	}
	return v
}

// Reports the largest value in the stack.
func (s *stack) max() int {
	return s.values[s.maxIdx()]
}

// Reports the offset of the largest value in s.values.
func (s *stack) maxIdx() int {
	return s.maxes[len(s.maxes)-1]
}
