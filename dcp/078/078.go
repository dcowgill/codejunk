/*

Good morning! Here's your coding interview problem for today.

This problem was asked recently by Google.

Given k sorted singly linked lists, write a function to merge all the lists into
one sorted singly linked list.

*/
package dcp078

import (
	"container/heap"
)

// A node in a linked list.
type node struct {
	next  *node
	value string
}

// Returns a shallow copy of n.
// n and its copy share the same next link.
func (n *node) Copy() *node {
	return &node{n.next, n.value}
}

// Use a min-heap to pick the next list head to merge.
func mergeLists(heads []*node) *node {
	pq := make(listHeap, len(heads))
	copy(pq, heads)
	heap.Init(&pq)
	var head, prev *node
	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*node)
		if prev == nil {
			head = n.Copy()
			prev = head
		} else {
			prev.next = n.Copy()
			prev = prev.next
		}
		if n.next != nil {
			heap.Push(&pq, n.next)
		}
	}
	return head
}

type listHeap []*node

func (h listHeap) Len() int            { return len(h) }
func (h listHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h listHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *listHeap) Push(x interface{}) { *h = append(*h, x.(*node)) }
func (h *listHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
