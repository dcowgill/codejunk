package dcp041

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	var tests = []struct {
		flights []flight
		start   airport
		it      itinerary
	}{
		{
			[]flight{{"SFO", "HKO"}, {"YYZ", "SFO"}, {"YUL", "YYZ"}, {"HKO", "ORD"}},
			"YUL",
			itinerary{"YUL", "YYZ", "SFO", "HKO", "ORD"},
		},
		{
			[]flight{{"SFO", "COM"}, {"COM", "YYZ"}},
			"COM",
			nil,
		},
		{
			[]flight{{"A", "B"}, {"A", "C"}, {"B", "C"}, {"C", "A"}},
			"A",
			itinerary{"A", "B", "C", "A", "C"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %+v", tt.start, tt.flights), func(t *testing.T) {
			it := findItinerary(newFlightGraph(tt.flights), tt.start)
			if !reflect.DeepEqual(it, tt.it) {
				t.Fatalf("findItinerary returned %+v, want %+v", it, tt.it)
			}
		})
	}
}
