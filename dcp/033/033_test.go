package dcp033

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestMedian(t *testing.T) {
	var tests = []struct {
		a []int
		m []float64
	}{
		// Problem description example.
		{[]int{2, 1, 5, 7, 2, 0, 5}, []float64{2, 1.5, 2, 3.5, 2, 2, 2}},

		// Simple cases.
		{[]int{}, []float64{}},
		{[]int{5}, []float64{5}},
		{[]int{5, 10}, []float64{5, 7.5}},
		{[]int{5, 10, 15}, []float64{5, 7.5, 10}},

		// More tests!
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.a), func(t *testing.T) {
			m := wrapper(tt.a)
			if !reflect.DeepEqual(m, tt.m) {
				t.Fatalf("got %+v, want %+v", m, tt.m)
			}
		})
	}
}

func TestRandomInputs(t *testing.T) {
	const (
		ntrials = 10000
		size    = 50
	)
	for i := 0; i < ntrials; i++ {
		// Generate a random sequence of size integers.
		a := make([]int, size)
		for j := range a {
			a[j] = rand.Int()
		}
		// Test both algorithms and compare.
		xs := bruteForce(a)
		ys := wrapper(a)
		if !reflect.DeepEqual(xs, ys) {
			t.Fatalf("median(%+v) returned %+v, want %+v", a, ys, xs)
		}
	}
}

// Wraps streamMedian to feed/drain channels, as a convenience.
func wrapper(a []int) []float64 {
	var (
		in  = make(chan int)
		out = make(chan float64)
	)
	go streamMedian(out, in)
	result := make([]float64, 0, len(a))
	for _, x := range a {
		in <- x
		result = append(result, <-out)
	}
	close(in)
	return result
}

// For comparison testing.
func bruteForce(a []int) []float64 {
	sorted := func(a []int) []int {
		b := make([]int, len(a))
		copy(b, a)
		sort.Ints(b)
		return b
	}
	median := func(a []int) float64 {
		a = sorted(a)
		switch n := len(a); {
		case n == 1:
			return float64(a[0])
		case n%2 == 0: // even
			x := a[n/2]
			y := a[n/2-1]
			return float64(x+y) / 2
		default: // odd
			return float64(a[n/2])
		}
	}
	medians := make([]float64, 0, len(a))
	for i := 1; i <= len(a); i++ {
		medians = append(medians, median(a[:i]))
	}
	return medians
}
