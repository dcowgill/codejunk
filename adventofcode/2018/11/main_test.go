package main

import "testing"

func TestPowerLevel(t *testing.T) {
	var tests = []struct {
		x         int
		y         int
		serialNum int
		power     int
	}{
		{3, 5, 8, 4},
		{122, 79, 57, -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},
	}
	for _, tt := range tests {
		power := powerLevel(tt.x, tt.y, tt.serialNum)
		if power != tt.power {
			t.Errorf("powerLevel(%d, %d, %d) returned %d, want %d",
				tt.x, tt.y, tt.serialNum, power, tt.power)
		}
	}
}
