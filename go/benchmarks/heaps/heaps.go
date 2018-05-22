package heaps

import (
	"container/heap"
	"time"
)

type timeHeap []time.Time

func newTimeHeap(a []time.Time) *timeHeap {
	h := timeHeap(make([]time.Time, len(a)))
	copy(h, a)
	heap.Init(&h)
	return &h
}

func (h timeHeap) Len() int            { return len(h) }
func (h timeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h timeHeap) Less(i, j int) bool  { return h[i].Before(h[j]) }
func (h *timeHeap) Push(x interface{}) { *h = append(*h, x.(time.Time)) }
func (h *timeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type worker1 struct {
	id string
	t  time.Time
}

type worker1Heap []worker1

func newWorker1Heap(a []worker1) *worker1Heap {
	h := worker1Heap(make([]worker1, len(a)))
	copy(h, a)
	heap.Init(&h)
	return &h
}

func (h worker1Heap) Len() int            { return len(h) }
func (h worker1Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h worker1Heap) Less(i, j int) bool  { return h[i].t.Before(h[j].t) }
func (h *worker1Heap) Push(x interface{}) { *h = append(*h, x.(worker1)) }
func (h *worker1Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type worker2 struct {
	id int
	t  time.Time
}

type worker2Heap []worker2

func newWorker2Heap(a []worker2) *worker2Heap {
	h := worker2Heap(make([]worker2, len(a)))
	copy(h, a)
	heap.Init(&h)
	return &h
}

func (h worker2Heap) Len() int            { return len(h) }
func (h worker2Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h worker2Heap) Less(i, j int) bool  { return h[i].t.Before(h[j].t) }
func (h *worker2Heap) Push(x interface{}) { *h = append(*h, x.(worker2)) }
func (h *worker2Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type zheap struct {
	a []time.Time
}

func (h *zheap) push(t time.Time) {
	h.a = append(h.a, t)
	up(h.a, len(h.a)-1)
}

func (h *zheap) pop() time.Time {
	n := len(h.a) - 1
	h.a[0], h.a[n] = h.a[n], h.a[0]
	down(h.a, 0, n)
	x := h.a[n]
	h.a = h.a[0:n]
	return x
}

func up(a []time.Time, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !a[j].Before(a[i]) {
			break
		}
		a[i], a[j] = a[j], a[i]
		j = i
	}
}

func down(a []time.Time, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && !a[j1].Before(a[j2]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !a[j].Before(a[i]) {
			break
		}
		a[i], a[j] = a[j], a[i]
		i = j
	}
}
