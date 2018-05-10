package pkg_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type thing struct {
	x int
	y int
	s string
	t string
	b bool
}

func newThing() *thing {
	return &thing{
		x: 0,
		y: 42,
		s: "foo bar",
		t: "baz qux",
		b: true,
	}
}

var (
	cache atomic.Value
	mu    sync.Mutex
)

func userOnce() *user {
	u := cache.Load()
	if u == nil {
		mu.Lock()
		defer mu.Unlock()
		u = cache.Load() // load again
		if u == nil {
			x := newUser()
			cache.Store(x)
			return x
		}
	}
	return u.(*thing)
}

var xx = newThing()

func thingInit() *thing {
	return xx
}

var global *thing

func BenchmarkNew(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var l *thing
		for pb.Next() {
			l = newThing()
		}
		global = l
	})
}

func BenchmarkOnce(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var l *thing
		for pb.Next() {
			l = thingOnce()
		}
		global = l
	})
}

func BenchmarkInit(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var l *thing
		for pb.Next() {
			l = thingInit()
		}
		global = l
	})
}
