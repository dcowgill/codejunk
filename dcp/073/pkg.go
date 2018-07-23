/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given the head of a singly linked list, reverse it in-place.

*/
package dcp073

// A node in a linked list.
type node struct {
	next  *node
	value string
}

// Reverses the list in place.
func reverse(head **node) {
	var prev *node
	p := *head
	for p != nil {
		p.next, prev, p = prev, p, p.next
	}
	*head = prev
}
