package dcp104

import (
	"strconv"
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		list *Node
		p    bool // palindrome?
	}{
		{makeList(1), true},
		{makeList(1, 1), true},
		{makeList(1, 2), false},
		{makeList(1, 2, 1), true},
		{makeList(1, 2, 3), false},
		{makeList(1, 2, 2, 1), true},
		{makeList(1, 2, 3, 2, 1), true},
		{makeList(1, 2, 3, 4, 2, 1), false},
		{makeList(1, 4, 3, 4, 1), true},
	}
	for _, tt := range tests {
		desc := listToString(tt.list)
		t.Run(desc, func(t *testing.T) {
			if p := isPalindrome(tt.list); p != tt.p {
				t.Fatalf("isPalindrome returned %v, want %v", p, tt.p)
			}
			// Ensure that isPalindrome did not munge the list.
			newDesc := listToString(tt.list)
			if desc != newDesc {
				t.Fatalf("list was altered from %q to %q", desc, newDesc)
			}
		})
	}
}

// Shorthand for making new lists.
func makeList(first int, rest ...int) *Node {
	head := &Node{value: first}
	curr := head
	for _, value := range rest {
		curr.next = &Node{value: value}
		curr = curr.next
	}
	return head
}

// Converts the list to a string with spaces between values.
// Used to print which test is being run.
func listToString(list *Node) string {
	var b strings.Builder
	sep := ""
	for p := list; p != nil; p = p.next {
		b.WriteString(sep)
		b.WriteString(strconv.Itoa(p.value))
		sep = " "
	}
	return b.String()
}
