/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given two singly linked lists that intersect at some point, find the
intersecting node. The lists are non-cyclical.

For example, given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10, return the
node with value 8.

In this example, assume nodes with the same value are the exact same node
objects.

Do this in O(M + N) time (where M and N are the lengths of the lists) and
constant space.

*/
package dcp020

// Node in a singly linked list.
type node struct {
	value int
	next  *node
}

// Returns the value of the first node in a and b that have the same value. (It
// is assumed, but not necessary, that they share a pointer, too.)
//
// Strategy: we know a priori that if two singly linked lists share a node, they
// share a suffix. Thus, we can safely ignore the leading part of whichever list
// is longer, since that part can't be part of the shared suffix.
//
// Given two lists of equal length that share a suffix, we can walk the lists in
// lockstep until we find the start of that suffix.
//
func findIntersectingNode(a, b *node) int {
	// Returns the length of the list starting at n.
	listLen := func(n *node) int {
		l := 0
		for n != nil {
			n = n.next
			l++
		}
		return l
	}

	// Compute the length of both lists.
	la := listLen(a)
	lb := listLen(b)

	// Advance abs(la-lb) links in the longer of the two lists.
	var p, q *node
	var d int
	if la > lb {
		p, q, d = a, b, la-lb
	} else {
		p, q, d = b, a, lb-la
	}
	for i := 0; i < d; i++ {
		p = p.next
	}

	// Walk the lists until we find the shared value.
	for p != nil && q != nil {
		if p.value == q.value {
			return p.value
		}
		p = p.next
		q = q.next
	}

	// The lists do not intersect.
	return 0
}
