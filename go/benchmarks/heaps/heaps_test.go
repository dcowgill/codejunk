package heaps

import (
	"container/heap"
	"math/rand"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	N = 10
	M = 100
)

func BenchmarkTimeZHeap(b *testing.B) {
	var a []time.Time
	now := time.Now()
	for i := 0; i < N; i++ {
		a = append(a, now.Add(time.Duration(i)*time.Minute))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var h zheap
		for _, x := range a {
			h.push(x)
		}
		for j := 0; j < M; j++ {
			t := h.pop()
			h.push(t.Add(time.Duration(rand.Intn(10)) * time.Minute))
		}
	}
}

func BenchmarkTimeHeap(b *testing.B) {
	var a []time.Time
	now := time.Now()
	for i := 0; i < N; i++ {
		a = append(a, now.Add(time.Duration(i)*time.Minute))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := newTimeHeap(a)
		for j := 0; j < M; j++ {
			t := heap.Pop(h).(time.Time)
			heap.Push(h, t.Add(time.Duration(rand.Intn(10))*time.Minute))
		}
	}
}

func BenchmarkWorker1Heap(b *testing.B) {
	var a []worker1
	now := time.Now()
	for i := 0; i < N; i++ {
		a = append(a, worker1{bson.NewObjectId(), now.Add(time.Duration(i) * time.Minute)})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := newWorker1Heap(a)
		for j := 0; j < M; j++ {
			w := heap.Pop(h).(worker1)
			w.t = w.t.Add(time.Duration(rand.Intn(10)) * time.Minute)
			heap.Push(h, w)
		}
	}
}

func BenchmarkWorker2Heap(b *testing.B) {
	var a []worker2
	now := time.Now()
	for i := 0; i < N; i++ {
		a = append(a, worker2{i, now.Add(time.Duration(i) * time.Minute)})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := newWorker2Heap(a)
		for j := 0; j < M; j++ {
			w := heap.Pop(h).(worker2)
			w.t = w.t.Add(time.Duration(rand.Intn(10)) * time.Minute)
			heap.Push(h, w)
		}
	}
}
