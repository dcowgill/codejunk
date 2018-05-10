package string_joins

import (
	"strconv"
	"strings"
	"testing"
)

func joinIfs(a []string) string {
	s := ""
	for i, t := range a {
		s += t + " = $" + strconv.Itoa(i+1)
		if i < len(a)-1 {
			s += ", "
		}
	}
	return s
}

func joinSep(a []string) string {
	const comma = ", "
	s := ""
	sep := ""
	for i, t := range a {
		s += sep
		s += t + " = $" + strconv.Itoa(i+1)
		sep = comma
	}
	return s
}

func joinLib(a []string) string {
	b := make([]string, len(a))
	for i, t := range a {
		b[i] = t + " = $" + strconv.Itoa(i+1)
	}
	return strings.Join(b, ", ")
}

func TestEqual(t *testing.T) {
	fields := []string{"foo", "bar", "baz"}
	a := joinIfs(fields)
	b := joinSep(fields)
	if expected := "foo = $1, bar = $2, baz = $3"; a != expected {
		t.Fatalf("joinIfs returned %q, want %q", a, expected)
	}
	if a != b {
		t.Fatalf("joinSep returned %q, wang %q", b, a)
	}
}

var fields = []string{
	"some_id = $1",
	"another_field = $2",
	"some_attr = $3",
	"foo_bar = $4",
	"blah_blah = $5",
	"long_field_name = $6",
}

func BenchmarkJoinIfs(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		n += len(joinIfs(fields))
	}
}

func BenchmarkJoinSep(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		n += len(joinSep(fields))
	}
}

func BenchmarkJoinLib(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		n += len(joinLib(fields))
	}
}
