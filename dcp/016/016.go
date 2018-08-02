/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Twitter.

You run an e-commerce website and want to record the last N order ids in a log.
Implement a data structure to accomplish this, with the following API:

record(order_id): adds the order_id to the log

get_last(i): gets the ith last element from the log. i is guaranteed to be
smaller than or equal to N.

You should be as efficient with time and space as possible.

*/
package dcp016

import "fmt"

type OrderID uint64

// RingBuffer is a circular buffer of fixed size.
type RingBuffer struct {
	log   []OrderID
	begin int // position of oldest element in log
	len   int // number of elements in log
}

// NewRingBuffer creates a new RingBuffer of the given size.
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{log: make([]OrderID, size)}
}

// Record appends id to the log.
func (rb *RingBuffer) Record(id OrderID) {
	rb.log[(rb.begin+rb.len)%len(rb.log)] = id
	if rb.isFull() {
		rb.begin = (rb.begin + 1) % len(rb.log)
	} else {
		rb.len++
	}
}

// Get reports the ith last element in the log.
func (rb *RingBuffer) Get(i int) OrderID {
	if i < 0 || i >= rb.len {
		panic(fmt.Sprintf("Get(%d): out of range [0,%d]", i, rb.len))
	}
	return rb.log[(rb.begin+rb.len-i-1)%len(rb.log)]
}

// isFull reports whether the buffer is at capacity.
func (rb *RingBuffer) isFull() bool {
	return rb.len == len(rb.log)
}
