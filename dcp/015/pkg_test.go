package dcp015

import (
	"math/rand"
	"testing"
	"time"
)

func TestChoose(t *testing.T) {
	const (
		numTrials = 100000
		numValues = 20
	)
	rand.Seed(time.Now().UTC().UnixNano())
	histogram := make([]int, numValues)
	for i := 0; i < numTrials; i++ {
		ch := make(chan int)
		go func() {
			for j := 0; j < numValues; j++ {
				ch <- j
			}
			close(ch)
		}()
		histogram[choose(ch)]++
	}

	// TODO: Kolmogorov-Smirnov test of uniformity
	// ...
}
