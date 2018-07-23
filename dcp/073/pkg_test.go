package dcp073

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	list := &node{value: "one", next: &node{value: "two", next: &node{value: "three", next: &node{
		value: "four", next: &node{value: "five", next: &node{value: "six"}}}}}}
	const (
		fwd = "one,two,three,four,five,six,"
		rev = "six,five,four,three,two,one,"
	)
	if s := dump(list); s != fwd {
		t.Fatalf("list is %q, want %q", s, fwd)
	}
	reverse(&list)
	if s := dump(list); s != rev {
		t.Fatalf("reversed list is %q, want %q", s, rev)
	}
	reverse(&list)
	if s := dump(list); s != fwd {
		t.Fatalf("doubly reversed list is %q, want %q", s, fwd)
	}
}

// Converts the list to a string.
func dump(head *node) string {
	var b strings.Builder
	for p := head; p != nil; p = p.next {
		b.WriteString(p.value)
		b.WriteByte(',')
	}
	return b.String()
}
