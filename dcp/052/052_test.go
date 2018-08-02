package dcp052

import (
	"reflect"
	"strconv"
	"testing"
)

// What to do to the cache.
type opKind int

const (
	setOp opKind = iota + 1
	getOp
)

// Describes a cache set of get operation.
type op struct {
	op opKind
	kv kvPair
}

// A named pair of strings.
type kvPair struct {
	key, value string
}

func TestLRUCache(t *testing.T) {
	var tests = []struct {
		capacity int
		ops      []op
		final    []kvPair
	}{
		{
			10,
			[]op{
				{setOp, kvPair{"foo", "bar"}},
				{getOp, kvPair{"foo", "bar"}},
			},
			[]kvPair{{"foo", "bar"}},
		},
		{
			10,
			[]op{
				{setOp, kvPair{"foo", "bar"}},
				{setOp, kvPair{"baz", "qux"}},
				{setOp, kvPair{"abc", "xyz"}},
				{getOp, kvPair{"baz", "qux"}},
				{getOp, kvPair{"abc", "xyz"}},
				{getOp, kvPair{"foo", "bar"}},
			},
			[]kvPair{{"foo", "bar"}, {"abc", "xyz"}, {"baz", "qux"}},
		},
		{
			2,
			[]op{
				{setOp, kvPair{"foo", "bar"}},
				{setOp, kvPair{"baz", "qux"}},
				{setOp, kvPair{"abc", "xyz"}}, // ejects foo/bar
				{getOp, kvPair{"foo", ""}},    // "" = not found
				{getOp, kvPair{"abc", "xyz"}},
				{getOp, kvPair{"baz", "qux"}},
			},
			[]kvPair{{"baz", "qux"}, {"abc", "xyz"}},
		},
		// TODO: more tests
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := newLRUCache(tt.capacity)
			for _, op := range tt.ops {
				switch op.op {
				case setOp:
					c.set(op.kv.key, op.kv.value)
				case getOp:
					if value, _ := c.get(op.kv.key); value != op.kv.value {
						t.Fatalf("get(%q) returned %q, want %q", op.kv.key, value, op.kv.value)
					}
				}
			}
			if !cacheEqual(c, tt.final) {
				t.Fatalf("cache does not match %+v", tt.final)
			}
		})
	}
}

// Reports whether c contains (exactly) the set of key-value pairs.
func cacheEqual(c *lruCache, pairs []kvPair) bool {
	if len(c.entries) != len(pairs) {
		return false
	}
	for _, kv := range pairs {
		n := c.entries[kv.key]
		if n == nil || n.value != kv.value {
			return false
		}
	}
	pairs2 := make([]kvPair, 0, len(pairs)) // inelegant but straightforward
	for el := c.head.next; el != c.head; el = el.next {
		pairs2 = append(pairs2, kvPair{el.key, el.value})
	}
	return reflect.DeepEqual(pairs, pairs2)
}
