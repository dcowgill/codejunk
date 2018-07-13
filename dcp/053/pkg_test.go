package dcp053

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	const (
		numValues = 1000
		pDequeue  = 0.5
	)

	// Randomly generate a slice of ints to enqueue.
	a := make([]int, numValues)
	for i := range a {
		a[i] = rand.Int()
	}

	// Enqueue all the values in a, occasionally dequeuing one into b.
	var b []int
	q := &queue{}
	for _, v := range a {
		q.enqueue(v)
		if rand.Float64() < pDequeue {
			b = append(b, q.dequeue())
		}
	}

	// Drain the rest of the queue into b.
	for q.size() > 0 {
		b = append(b, q.dequeue())
	}

	// a and b must be identical.
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("dequeued values do not match enqueued values")
	}
}
