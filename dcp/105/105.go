/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a function f, and N return a debounced f of N milliseconds.

That is, as long as the debounced f continues to be invoked, f itself will not
be called for N milliseconds.

*/
package dcp105

import (
	"sync"
	"time"
)

func debounce(f func(), wait time.Duration) func() {
	var (
		timer *time.Timer
		mu    sync.Mutex
	)
	return func() {
		mu.Lock()
		if timer == nil {
			timer = time.AfterFunc(wait, f)
		} else {
			timer.Stop()
			timer.Reset(wait)
		}
		mu.Unlock()
	}
}
