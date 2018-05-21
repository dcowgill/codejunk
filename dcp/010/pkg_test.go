package dcp010

import (
	"sync"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	now := time.Now()
	var ds []time.Duration
	var expected []time.Duration
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 10; i <= 1000; i *= 2 {
		d := time.Duration(i) * time.Millisecond
		expected = append(expected, d)
		wg.Add(1)
		schedule(d, func() {
			defer wg.Done()
			mu.Lock()
			ds = append(ds, time.Since(now))
			mu.Unlock()
		})
	}
	wg.Wait()
	if len(ds) != len(expected) {
		t.Fatalf("found %d durations, want %d", len(ds), len(expected))
	}
	for i, x := range ds {
		y := expected[i]
		if !closeEnough(x, y) {
			t.Fatalf("duration %d is %s, want %s (delta=%v, ratio=%v)", i, x, y,
				delta(x, y), ratio(x, y))
		}
	}
}

// Within 80% or 5ms.
func closeEnough(a, b time.Duration) bool {
	return delta(a, b) <= 5*time.Millisecond || ratio(a, b) >= 0.80
}

func delta(a, b time.Duration) time.Duration {
	if a > b {
		return a - b
	}
	return b - a
}
func ratio(a, b time.Duration) float64 {
	x := float64(a)
	y := float64(b)
	if x < y {
		return x / y
	}
	return y / x
}
