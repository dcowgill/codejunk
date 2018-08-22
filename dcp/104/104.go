/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Determine whether a doubly linked list is a palindrome. What if it’s singly linked?

For example, 1 -> 4 -> 3 -> 4 -> 1 returns true while 1 -> 4 returns false.

*/
package dcp104

// Node in a singly linked list.
type Node struct {
	next  *Node
	value int
}

// Reports whether a singly linked list is a palindrome.
// Requires O(N) time and O(1) space. Doesn't allocate.
func isPalindrome(head *Node) bool {
	// Determine the length of the list so we can find the middle node.
	length := 0
	for p := head; p != nil; p = p.next {
		length++
	}

	// Reverse the list starting at the middle node.
	middle := head
	for i := 0; i < length/2; i++ {
		middle = middle.next
	}
	middle = reverseList(middle)
	defer reverseList(middle) // undo on return

	// Scan the two halves of the list, starting at the front and middle. If the
	// halves are equal, the overall list is a palindrome.
	for p, q := head, middle; p != nil && q != nil; p, q = p.next, q.next {
		if p.value != q.value {
			return false
		}
	}
	return true
}

// Reverses the list. Returns the new head, which was previously the last node.
func reverseList(list *Node) *Node {
	p := list
	var prev *Node
	for p.next != nil {
		p.next, p, prev = prev, p.next, p
	}
	p.next = prev
	return p
}
