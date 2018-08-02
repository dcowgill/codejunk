/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given a singly linked list and an integer k, remove the kth last element from
the list. k is guaranteed to be smaller than the length of the list.

The list is very long, so making more than one pass is prohibitively expensive.

Do this in constant space and in one pass.

*/
package dcp026

// node is an element in a singly linked list.
type node struct {
	next  *node
	value string
}

// removeKth removes the kth element in the list. Returns the new list head.
func removeKth(head *node, k int) *node {
	p := &head
	for i := 0; i < k; i++ {
		p = &((*p).next)
	}
	*p = (*p).next
	return head
}
