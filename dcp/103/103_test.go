package dcp103

import (
	"fmt"
	"testing"
)

func TestShortestSubstringContainingRunes(t *testing.T) {
	var tests = []struct {
		source, chars, substring string
	}{
		{"figehaeci", "aei", "aeci"},

		{"", "", ""},
		{"abcd", "", ""},
		{"abcd", "abcde", ""},
		{"abababcbaba", "cab", "abc"},
		{"abbbbbbbc", "abc", "abbbbbbbc"},
		{"abbbbbbbac", "abc", "bac"},
	}
	for _, tt := range tests {
		desc := fmt.Sprintf("%s %s", tt.source, tt.chars)
		t.Run(desc, func(t *testing.T) {
			substring := shortestSubstringContainingRunes(tt.source, tt.chars)
			if substring != tt.substring {
				t.Fatalf("got %q, want %q", substring, tt.substring)
			}
		})
	}
}
