package dcp209

import "testing"

func TestLCS3(t *testing.T) {
	const (
		a        = "epidemiologist"
		b        = "refrigeration"
		c        = "supercalifragilisticexpialodocious"
		expected = "eieio"
	)
	perms := [][]string{
		{a, b, c},
		{a, c, b},
		{b, a, c},
		{b, c, a},
		{c, a, b},
		{c, b, a},
	}
	for _, p := range perms {
		if l := lcs3(p[0], p[1], p[2]); l != expected {
			t.Fatalf("lcs3(%q) returned %q, want %q", p, l, expected)
		}
	}
}
