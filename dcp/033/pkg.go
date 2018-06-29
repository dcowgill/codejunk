/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

Compute the running median of a sequence of numbers. That is, given a stream of
numbers, print out the median of the list so far on each new element.

Recall that the median of an even-numbered list is the average of the two middle numbers.

For example, given the sequence [2, 1, 5, 7, 2, 0, 5], your algorithm should print out:

2
1.5
2
3.5
2
2
2

*/
package dcp033

/*

Thoughts:

All data must be preserved; if there is a constant-space algorithm, it's not
obvious to me. First idea is insertion sort, which should be O(N^2): N elements
are read from the stream, and insertion into a sorted list requires M/2 swaps
where M is the number of elements read so far (which approximates N).

Alternative: keep two priority heaps, one for items less than the current median
and one for values greater than the current median. This should be O(NlogN):
push/pop on a heap is O(logN) and we have to perform those operations 2N times.

*/

import (
	"container/heap"
)

func streamMedian(out chan<- float64, in <-chan int) {
	// l contains values less than or equal to the current median.
	l := &intHeap{less: func(x, y int) bool { return x > y }} // max heap

	// r contains values greater than the current median.
	r := &intHeap{less: func(x, y int) bool { return x < y }} // min heap

	// Handle the first value as a special case.
	v, ok := <-in
	if !ok {
		return
	}
	out <- float64(v)
	heap.Push(l, v)

	// Process all remaining values in the stream.
	for v := range in {
		if l.Len() == r.Len() {
			// Heaps are of equal size; always push onto l.
			if v < r.top() {
				heap.Push(l, v)
			} else {
				heap.Push(l, heap.Pop(r).(int))
				heap.Push(r, v)
			}
			out <- float64(l.top()) // count is now odd
		} else {
			// l has one more item than r; always push onto r.
			if v < l.top() {
				heap.Push(r, heap.Pop(l).(int))
				heap.Push(l, v)
			} else {
				heap.Push(r, v)
			}
			out <- float64(l.top()+r.top()) / 2 // count is now even
		}
	}
}

// A heap of ints with a configurable comparison function.
type intHeap struct {
	a    []int
	less func(x, y int) bool
}

// top panics if the heap is empty.
func (h *intHeap) top() int { return h.a[0] }

func (h *intHeap) Len() int           { return len(h.a) }
func (h *intHeap) Less(i, j int) bool { return h.less(h.a[i], h.a[j]) }
func (h *intHeap) Swap(i, j int)      { h.a[i], h.a[j] = h.a[j], h.a[i] }

func (h *intHeap) Push(x interface{}) {
	h.a = append(h.a, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := h.a
	n := len(old)
	x := old[n-1]
	h.a = old[0 : n-1]
	return x
}

func main() {}
