/*

Good morning! Here's your coding interview problem for today.

This question was asked by Snapchat.

Given the head to a singly linked list, where each node also has a “random”
pointer that points to anywhere in the linked list, deep clone the list.

*/
package dcp131

type node struct {
	next   *node
	random *node
	value  int
}

func copyList(head *node) *node {
	// Create a mapping from old nodes to their corresponding new nodes, so we
	// can backpatch the "random" links later.
	addrs := make(map[*node]*node)

	// Deep copy the list. Re-use the "random" links; we'll fix them below.
	var copy *node
	q := &copy
	for p := head; p != nil; p = p.next {
		*q = &node{value: p.value, random: p.random}
		addrs[p] = *q
		q = &(*q).next
	}

	// Backpatch the "random" links in the new list, using "addrs" to map them
	// from nodes in the old list to nodes in the new list.
	for p := copy; p != nil; p = p.next {
		p.random = addrs[p.random] // orig -> copy
	}

	// Done.
	return copy
}
