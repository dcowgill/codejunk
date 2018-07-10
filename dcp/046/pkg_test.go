package dcp046

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		s string
		p string
	}{
		{"", ""},
		{"a", "a"},
		{"aa", "aa"},
		{"ab", "a"},
		{"aba", "aba"},
		{"banana", "anana"},
		{"aabcdcb", "bcdcb"},
		{"bcdcbaa", "bcdcb"},
		{"hello, world", "ll"},
		{"amanaplanacanalpanama", "amanaplanacanalpanama"},
		{"amanaplanacanal", "lanacanal"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			p := longestPalindrome(tt.s)
			if p != tt.p {
				t.Fatalf("got %q, want %q", p, tt.p)
			}
		})
	}
}
