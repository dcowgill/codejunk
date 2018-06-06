package dcp027

import "testing"

func TestIsBalanced(t *testing.T) {
	var tests = []struct {
		input string
		b     bool
	}{
		{"([])[]({})", true},
		{"([)]", false},
		{"((()", false},
		{"", true},
		{"()", true},
		{"[]", true},
		{"{}", true},
		{")(", false},
		{"][", false},
		{"}{", false},
		{"(", false},
		{"[", false},
		{"{", false},
		{"((]]", false},
		{"{{))", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			b := isBalanced(tt.input)
			if b != tt.b {
				t.Fatalf("isBalanced returned %v, want %v", b, tt.b)
			}
		})
	}
}
