package utils

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

var funcs = []struct {
	name string
	fn   func([]string, func(string) bool) ([]string, []string)
}{
	{"splitSlice", splitSlice},
	{"splitSliceAlloc", splitSliceAlloc},
}

func TestSplitSlice(t *testing.T) {
	var (
		alwaysTrue  = func(s string) bool { return true }
		alwaysFalse = func(s string) bool { return false }
		hasPrefix   = func(prefix string) func(string) bool {
			return func(s string) bool { return strings.HasPrefix(s, prefix) }
		}
		lenLessThan = func(n int) func(string) bool {
			return func(s string) bool { return len(s) < n }
		}
	)
	var tests = []struct {
		in   []string
		pred func(string) bool
		out1 []string
		out2 []string
	}{
		{[]string{}, alwaysTrue, []string{}, []string{}},
		{nil, alwaysTrue, []string{}, []string{}},
		{[]string{"a", "b", "c"}, alwaysTrue, []string{"a", "b", "c"}, []string{}},
		{[]string{"a", "b", "c"}, alwaysFalse, []string{}, []string{"a", "b", "c"}},
		{
			[]string{"foo", "bar", "baz", "qux", "blah", "blah"},
			hasPrefix("b"),
			[]string{"baz", "bar", "blah", "blah"},
			[]string{"foo", "qux"},
		},
		{
			[]string{"a", "aa", "aaa", "aaaa", "aaa", "aa", "a"},
			lenLessThan(3),
			[]string{"a", "a", "aa", "aa"},
			[]string{"aaa", "aaa", "aaaa"},
		},
	}
	// stringSetsEqual reports whether a and b contains the same strings,
	// ignoring order. N.B. sorts a and b in situ.
	stringSetsEqual := func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		sort.Strings(a)
		sort.Strings(b)
		for i, s := range a {
			if s != b[i] {
				return false
			}
		}
		return true
	}
	for i, tt := range tests {
		for _, fn := range funcs {
			t.Run(fmt.Sprintf("%s %d", fn.name, i), func(t *testing.T) {
				out1, out2 := fn.fn(tt.in, tt.pred)
				if !stringSetsEqual(out1, tt.out1) || !stringSetsEqual(out2, tt.out2) {
					t.Fatalf("got (%#v, %#v), want (%#v, %#v)", out1, out2, tt.out1, tt.out2)
				}
			})
		}
	}
}

func BenchmarkSplitSlice(b *testing.B) {
	var (
		xs1    = randStrings(1, 10)
		xs20   = randStrings(20, 10)
		xs100  = randStrings(100, 10)
		mid20  = xs20[10]
		mid100 = xs100[50]
		m20    = mapEveryOther(xs20)
		m100   = mapEveryOther(xs100)
		blank  = make([]string, 200)
	)
	var n int
	var cases = []struct {
		name string
		xs   []string
		pred func(string) bool
	}{
		{"1 string, all", xs1, func(s string) bool { return true }},
		{"1 string, none", xs1, func(s string) bool { return false }},
		{"20 strings, all", xs20, func(s string) bool { return true }},
		{"20 strings, none", xs20, func(s string) bool { return false }},
		{"100 strings, all", xs100, func(s string) bool { return true }},
		{"100 strings, none", xs100, func(s string) bool { return false }},
		{"20 strings, select one", xs20, func(s string) bool { return s == mid20 }},
		{"100 strings, select one", xs100, func(s string) bool { return s == mid100 }},
		{"20 strings, alternating", xs20, func(s string) bool { return m20[s] }},
		{"100 strings, alternating", xs100, func(s string) bool { return m100[s] }},
	}
	for _, tt := range cases {
		for _, fn := range funcs {
			b.Run(tt.name+" "+fn.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					a := blank[:len(tt.xs)]
					copy(a, tt.xs)
					t, f := fn.fn(a, tt.pred)
					n += len(t) + len(f)
				}
			})
		}
	}
}

func mapEveryOther(xs []string) map[string]bool {
	m := make(map[string]bool, len(xs)/2)
	for i := 0; i < len(xs); i += 2 {
		m[xs[i]] = true
	}
	return m
}

// Generates a slice of n random strings of length l.
func randStrings(n, l int) []string {
	a := make([]string, 0, n)
	for i := 0; i < cap(a); i++ {
		a = append(a, randString(l))
	}
	return a
}

// Generates a string of n random ASCII bytes.
func randString(n int) string {
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = byte(rand.Intn(128)) // ASCII
	}
	return string(s)
}
