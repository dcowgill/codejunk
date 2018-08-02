package dcp025

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	var tests = []struct {
		pattern string
		input   string
		match   bool
	}{
		{".*at", "chat", true},
		{".*at", "chats", false},
		{"a*b*c*", "aaabbc", true},
		{"a*b*c*", "aaabb", true},
		{"a*b*c*", "aaa", true},
		{"a*b*c*", "", true},
		{"a*b*c*", "abcd", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.pattern, tt.input), func(t *testing.T) {
			if b := match(tt.pattern, tt.input); b != tt.match {
				t.Fatalf("match returned %v, want %v", b, tt.match)
			}
		})
	}
}
