/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given an array of integers and a number k, where 1 <= k <= length of the array,
compute the maximum values of each subarray of length k.

For example, given array = [10, 5, 2, 7, 8, 7] and k = 3, we should
get: [10, 7, 8, 8], since:

10 = max(10, 5, 2)
7 = max(5, 2, 7)
8 = max(2, 7, 8)
8 = max(7, 8, 7)

Do this in O(n) time and O(k) space. You can modify the input array in-place and
you do not need to store the results. You can simply print them out as you
compute them.

*/
package dcp018

// Uses a sliding window, represented as a double-ended queue (so that
// enqueue/dequeue operations at either end at O(1)), which contains indices in
// the current window, in ascending order, such that they are also ordered by
// their corresponding values. This invariant is maintained by always popping
// the front value if it is outside the window, and popping from the back all
// lesser values before enqueuing any new index.
func subarrayMaxes(a []int, k int, emit func(int)) {
	q := newDequeue(k)
	for i := 0; i < k; i++ {
		for !q.empty() && a[i] >= a[q.back()] {
			q.popBack()
		}
		q.pushBack(i)
	}
	for i := k; i < len(a); i++ {
		emit(a[q.front()])
		if !q.empty() && q.front() <= i-k {
			q.popFront()
		}
		for !q.empty() && a[i] >= a[q.back()] {
			q.popBack()
		}
		q.pushBack(i)
	}
	emit(a[q.front()])
}

// A simple implementation of a fixed-capacity dequeue.
type dequeue struct {
	vals  []int
	begin int
	size  int
}

func newDequeue(k int) *dequeue { return &dequeue{vals: make([]int, k)} }

func (q *dequeue) empty() bool { return q.size == 0 }
func (q *dequeue) front() int  { return q.vals[q.begin] }
func (q *dequeue) back() int   { return q.vals[(q.begin+q.size-1)%len(q.vals)] }

func (q *dequeue) pushFront(x int) {
	if q.size == len(q.vals) {
		panic("pushFront on full dequeue")
	}
	q.begin = (q.begin + len(q.vals) - 1) % len(q.vals)
	q.vals[q.begin] = x
	q.size++
}

func (q *dequeue) popFront() int {
	if q.size == 0 {
		panic("popFront on empty dequeue")
	}
	x := q.vals[q.begin]
	q.begin = (q.begin + 1) % len(q.vals)
	q.size--
	return x
}

func (q *dequeue) pushBack(x int) {
	if q.size == len(q.vals) {
		panic("pushBack on full dequeue")
	}
	back := (q.begin + q.size) % len(q.vals)
	q.vals[back] = x
	q.size++
}

func (q *dequeue) popBack() int {
	if q.size == 0 {
		panic("popBack on empty dequeue")
	}
	back := (q.begin + q.size - 1) % len(q.vals)
	x := q.vals[back]
	q.size--
	return x
}

func (q *dequeue) dump() []int {
	d := make([]int, q.size)
	for i := 0; i < q.size; i++ {
		d[i] = q.vals[(q.begin+i)%len(q.vals)]
	}
	return d
}
