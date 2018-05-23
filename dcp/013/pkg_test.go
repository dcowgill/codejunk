package dcp013

import (
	"fmt"
	"math/rand"
	"testing"
)

// Use a brute-force O(N^2) approach as the reference implementation.
func bruteForce(s string, k int) (int, int) {
	var begin, end int
	rs := []rune(s)
	for i := range rs {
		m := make(map[rune]bool)
		for j := i; j < len(rs); j++ {
			m[rs[j]] = true
			if len(m) > k {
				break
			}
			if len := j - i + 1; len > end-begin {
				begin, end = i, j+1
			}
		}
	}
	return begin, end
}

func TestLongestSubstringOfAtMostKDistinctRunes(t *testing.T) {
	var tests = []struct {
		s          string
		k          int
		begin, end int
	}{
		{"abcba", 2, 1, 4},
		{"abbccbbdbbccbba", 3, 1, 14},
		{"abbccbbdbbccbbacacacacab", 3, 8, 24},
		{"skittonkkqnffpopijhdvairmxqirqjgwmvvwvrhmgzsqgzmhuukntvnyavchbelxahmippdkdmekhkvemqhbppqqhashahxaihq", 8, 82, 100},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("k=%d, %s", tt.k, tt.s), func(t *testing.T) {
			begin, end := bruteForce(tt.s, tt.k)
			// begin, end := longestSubstringOfAtMostKDistinctRunes(tt.s, tt.k)
			if begin != tt.begin || end != tt.end {
				t.Fatalf("got (%d, %d), want (%d, %d)", begin, end, tt.begin, tt.end)
			}
		})
	}
}

func TestRandomly(t *testing.T) {
	// Generates an n-length string of random letters. Only uses lowercase a-z.
	randString := func(n int) string {
		a := make([]rune, n)
		for i := 0; i < n; i++ {
			a[i] = rune(rand.Intn(26) + 'a')
		}
		return string(a)
	}
	const numTests = 10000
	for i := 0; i < numTests; i++ {
		s := randString(100)
		k := 2 + rand.Intn(10)
		b1, e1 := bruteForce(s, k)
		b2, e2 := longestSubstringOfAtMostKDistinctRunes(s, k)
		if b1 != b2 || e1 != e2 {
			t.Fatalf("longestSubstringOfAtMostKDistinctRunes(%q, %d) returned (%d, %d), want (%d, %d)", s, k, b2, e2, b1, e1)
		}
	}
}
