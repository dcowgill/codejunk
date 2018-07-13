/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Apple.

Implement a queue using two stacks. Recall that a queue is a FIFO (first-in,
first-out) data structure with the following methods: enqueue, which inserts an
element into the queue, and dequeue, which removes it.

*/
package dcp053

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() int {
	old := *s
	n := len(old)
	v := old[n-1]
	*s = old[:n-1]
	return v
}

func (s *stack) size() int {
	return len(*s)
}

type queue struct {
	out stack
	rev stack
}

// func newQueue() *queue {
// 	return &queue{out:new(stack),rev:new(stack)}
// }

func (q *queue) enqueue(v int) {
	q.rev.push(v)
}

func (q *queue) dequeue() int {
	if q.out.size() == 0 {
		for q.rev.size() > 0 {
			q.out.push(q.rev.pop())
		}
	}
	return q.out.pop()
}

func (q *queue) size() int {
	return q.out.size() + q.rev.size()
}
