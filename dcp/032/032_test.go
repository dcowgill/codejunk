package dcp032

import (
	"strconv"
	"testing"
)

func TestArbitrage(t *testing.T) {
	var tests = []struct {
		rates [][]float64
		arbit bool
	}{
		// Basic cases involving two currencies.
		{[][]float64{{1.0, 0.9}, {0.9, 1.0}}, false},
		{[][]float64{{1.0, 1.0}, {1.0, 1.0}}, false},
		{[][]float64{{1.0, 1.1}, {1.0, 1.0}}, true},
		{[][]float64{{1.0, 1.0}, {1.1, 1.0}}, true},
		{[][]float64{{1.0, 1.1}, {1.1, 1.0}}, true},

		// A few straightforward A->B->C->A cases.
		{
			[][]float64{
				{1.0, 1.5, 0.0}, // A = 1.5B
				{0.0, 1.0, 1.5}, // B = 1.5C
				{0.5, 0.0, 1.0}, // C = 0.5A (1.5*1.5*0.5 = 1.125)
			},
			true,
		},
		{
			[][]float64{
				{1.0, 1.5, 0.0}, // A = 1.5B
				{0.0, 1.0, 1.5}, // B = 1.5C
				{0.4, 0.0, 1.0}, // C = 0.4A (1.5*1.5*0.4 = 0.900)
			},
			false,
		},

		// Write more tests!
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := arbitrage(tt.rates)
			if b != tt.arbit {
				t.Fatalf("arbitrage returned %v, want %v", b, tt.arbit)
			}
		})
	}
}
