package dcp016

import (
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(4)
	for i := 1; i <= 8; i++ {
		rb.Record(OrderID(i * 2))
	}
	var expect = []OrderID{16, 14, 12, 10}
	for i := 0; i < 4; i++ {
		if id := rb.Get(i); id != expect[i] {
			t.Errorf("rb.Get(%d) returned %d, want %d", i, id, expect[i])
		}
	}
}
