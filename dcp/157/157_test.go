package dcp157

import "testing"

func TestHasPalindromePermutation(t *testing.T) {
	var tests = []struct {
		s string
		b bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aaa", true},
		{"ab", false},
		{"bba", true},
		{"babb", false},
		{"carrace", true},
		{"daily", false},
		{"l e i lbeiabas seewar a w", true},   // able was i ere i saw elba
		{"a nam a alampalc a pannaa n", true}, // a man a plan a canal panama
		{"on nvred ovre dee", true},           // never odd or even
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if hasPalindromePermutation(tt.s) != tt.b {
				t.Fatalf("got %v, want %v", !tt.b, tt.b)
			}
		})
	}
}
