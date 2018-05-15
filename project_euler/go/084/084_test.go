package p084

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	var tests = []struct {
		params   simulationParams
		solution string
	}{
		{simulationParams{numDice: 2, dieSize: 4}, "101524"},
		{simulationParams{numDice: 2, dieSize: 6}, "102400"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%dd%d", tt.params.numDice, tt.params.dieSize), func(t *testing.T) {
			if answer := solve(tt.params, 3); answer != tt.solution {
				t.Fatalf("solve() returned %q, want %q", answer, tt.solution)
			}
		})
	}
}
