package dcp082

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReadN(t *testing.T) {
	const original = "Hello world"
	var tests = []struct {
		n int      // how many runes to read at once
		a []string // resulting strings
	}{
		{1, []string{"H", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d"}},
		{2, []string{"He", "ll", "o ", "wo", "rl", "d"}},
		{3, []string{"Hel", "lo ", "wor", "ld"}},
		{4, []string{"Hell", "o wo", "rld"}},
		{5, []string{"Hello", " worl", "d"}},
		{6, []string{"Hello ", "world"}},
		{7, []string{"Hello w", "orld"}},
		{8, []string{"Hello wo", "rld"}},
		{9, []string{"Hello wor", "ld"}},
		{10, []string{"Hello worl", "d"}},
		{11, []string{"Hello world"}},
		{12, []string{"Hello world"}},
	}
	for _, tt := range tests {
		a := readAll(strings.NewReader(original), tt.n)
		if !reflect.DeepEqual(a, tt.a) {
			t.Fatalf("readAll(%d) returned %#v, want %#v", tt.n, a, tt.a)
		}
		if s := strings.Join(a, ""); s != original {
			t.Fatalf("concat of readAll(%d) is %q, want %q", tt.n, s, original)
		}
	}
}

// Read n chars at a time from r until empty. Return the non-empty strings.
func readAll(r io.Reader, n int) []string {
	readN := makeReadN(makeRead7(r))
	var a []string
	for {
		b := readN(n)
		if len(b) == 0 {
			break
		}
		a = append(a, string(b))
	}
	return a
}
