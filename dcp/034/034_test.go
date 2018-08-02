package dcp034

import "testing"

func TestMinPalindrome(t *testing.T) {
	var tests = []struct {
		q string // input
		p string // result
	}{
		// Provided examples.
		{"race", "ecarace"},
		{"google", "elgoogle"},

		// Edge cases.
		{"", ""},
		{"a", "a"},
		{"aa", "aa"},

		// Simple cases.
		{"abcbb", "abbcbba"},
		{"helloworld", "dhellroworllehd"},
		// ...
	}
	for _, tt := range tests {
		t.Run(tt.q, func(t *testing.T) {
			p := minPalindrome(tt.q)
			if p != tt.p {
				t.Fatalf("minPalindrome returned %q, want %q", p, tt.p)
			}
		})
	}
}
