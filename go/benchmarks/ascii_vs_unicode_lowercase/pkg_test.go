package ascii_vs_unicode_lowercase

import (
	"strings"
	"testing"
)

func ToLowerASCII(s string) string {
	b := make([]byte, len(s))
	for i := range b {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A' // a=97 A=65
		}
		b[i] = c
	}
	return string(b)
}

var (
	s1 = "The quick brown JUMPED over the lazy DOG!"
)

func BenchmarkToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.ToLower(s1)
	}
}

func BenchmarkToLowerASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToLowerASCII(s1)
	}
}
