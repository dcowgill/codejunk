package dcp105

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	// Tests debounce() by debouncing a function which atomically increments a
	// zero-initialized counter, then calling it in a loop with varying delays.
	// The final value of the counter is returned. Based on the given durations,
	// we expect the counter to be incremented an exact number of times.
	run := func(debounceWait, loopWait time.Duration, loopIters int) int {
		var numCalls int32
		f := debounce(func() {
			atomic.AddInt32(&numCalls, 1)
		}, debounceWait)

		var wg sync.WaitGroup
		wg.Add(loopIters)
		for i := 1; i <= loopIters; i++ {
			go func(delay time.Duration) {
				time.Sleep(delay)
				f()
				wg.Done()
			}(time.Duration(i) * loopWait)
		}
		wg.Wait()

		time.Sleep(2 * debounceWait)
		return int(atomic.LoadInt32(&numCalls))
	}

	t.Run("called once for all debounced calls", func(t *testing.T) {
		const (
			debounceWait = 50 * time.Millisecond
			loopWait     = 40 * time.Millisecond
			loopIters    = 5
		)
		if n := run(debounceWait, loopWait, loopIters); n != 1 {
			t.Fatalf("numCalls is %d, expected 1", n)
		}
	})

	t.Run("called once per debounced call", func(t *testing.T) {
		const (
			debounceWait = 50 * time.Millisecond
			loopWait     = 60 * time.Millisecond
			loopIters    = 5
		)
		if n := run(debounceWait, loopWait, loopIters); n != loopIters {
			t.Fatalf("numCalls is %d, expected %d", n, loopIters)
		}
	})
}
