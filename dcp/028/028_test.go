package dcp028

import (
	"reflect"
	"strconv"
	"testing"
)

func TestJustify(t *testing.T) {
	var tests = []struct {
		words []string
		k     int
		lines []string
	}{
		// The provided example.
		{[]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}, 16, []string{"the  quick brown", "fox  jumps  over", "the   lazy   dog"}},

		// Test a string longer than k and a line that must be right-padded.
		{[]string{"abcd", "a", "b", "c"}, 3, []string{"abcd", "a b", "c  "}},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			lines := justify(tt.words, tt.k)
			if !reflect.DeepEqual(lines, tt.lines) {
				t.Fatalf("justify returned %#v, want %#v", lines, tt.lines)
			}
		})
	}
}
