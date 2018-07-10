package dcp047

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		prices []int
		max    int
	}{
		{nil, 0},
		{[]int{5}, 0},
		{[]int{5, 5}, 0},
		{[]int{10, 5}, 0},
		{[]int{15, 10, 5}, 0},
		{[]int{9, 11, 8, 5, 7, 10}, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.prices), func(t *testing.T) {
			max := maxProfit(tt.prices)
			if max != tt.max {
				t.Fatalf("got %d, want %d", max, tt.max)
			}
		})
	}
}

func TestRandomPrices(t *testing.T) {
	const ntrials = 10000
	for i := 0; i < ntrials; i++ {
		prices := randPrices(rand.Intn(100))
		expected := refImpl(prices)
		actual := maxProfit(prices)
		if actual != expected {
			t.Fatalf("oops: got %d, want %d; prices = %#v\n", actual, expected, prices)
		}
	}
}

func refImpl(prices []int) int {
	best := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > best {
				best = profit
			}
		}
	}
	return best
}

func randPrices(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = 1 + rand.Intn(1000)
	}
	return a
}
