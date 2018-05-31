/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Snapchat.

Given an array of time intervals (start, end) for classroom lectures (possibly
overlapping), find the minimum number of rooms required.

For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

*/
package dcp021

import (
	"container/heap"
)

// An interval represents a meeting with a begin and end time.
type interval [2]int

func (i interval) begin() int { return i[0] }
func (i interval) end() int   { return i[1] }

// Heap of intervals; ordered by h.less.
type intervalHeap struct {
	heap []interval
	less func(x, y interval) bool
}

func (h *intervalHeap) Min() interval      { return h.heap[0] }
func (h *intervalHeap) Len() int           { return len(h.heap) }
func (h *intervalHeap) Swap(i, j int)      { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }
func (h *intervalHeap) Less(i, j int) bool { return h.less(h.heap[i], h.heap[j]) }
func (h *intervalHeap) Push(x interface{}) { h.heap = append(h.heap, x.(interval)) }
func (h *intervalHeap) Pop() interface{} {
	n := len(h.heap)
	x := h.heap[n-1]
	h.heap = h.heap[:n-1]
	return x
}

// minRooms reports the minimum number of rooms necessary to accommodate all of
// the meetings represented by intervals.
func minRooms(intervals []interval) int {
	if len(intervals) == 0 {
		return 0
	}

	// Create two heaps: one for meetings that haven't started yet, and one for
	// meetings that are currently in progress.
	var (
		upcoming = &intervalHeap{
			less: func(x, y interval) bool {
				return x.begin() < y.begin()
			},
		}
		current = &intervalHeap{
			less: func(x, y interval) bool {
				return x.end() < y.end()
			},
		}
	)

	// Add all meetings to the upcoming heap.
	for _, v := range intervals {
		heap.Push(upcoming, v)
	}

	// Push the first meeting onto the current heap.
	heap.Push(current, heap.Pop(upcoming).(interval))
	maxRooms := 1

	// As long as we have more meetings:
	//
	// If the next upcoming meeting begins before any current meeting ends, or
	// if no meeting is in progress, move it from the upcoming heap to the
	// current heap, increasing our maximum room count if necessary.
	//
	// Otherwise, remove the next meeting to end from the current heap.
	for upcoming.Len() > 0 {
		if current.Len() == 0 || upcoming.Min().begin() < current.Min().end() {
			heap.Push(current, heap.Pop(upcoming).(interval))
			if current.Len() > maxRooms {
				maxRooms = current.Len()
			}
		} else {
			heap.Pop(current)
		}
	}
	return maxRooms
}
