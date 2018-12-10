package dcp212

import (
	"strconv"
	"testing"
)

func TestColNumToID(t *testing.T) {
	tests := []struct {
		num int
		id  string
	}{
		{1, "A"},
		{2, "B"},
		{3, "C"},
		{1*26 - 1, "Y"},
		{1*26 + 0, "Z"},
		{1*26 + 1, "AA"},
		{1*26 + 2, "AB"},
		{2*26 - 1, "AY"},
		{2*26 + 0, "AZ"},
		{2*26 + 1, "BA"},
		{2*26 + 2, "BB"},
		{26*26 - 1, "YY"},
		{26*26 + 0, "YZ"},
		{26*26 + 1, "ZA"},
		{26*26 + 2, "ZB"},
		{27*26 - 1, "ZY"},
		{27*26 + 0, "ZZ"},
		{27*26 + 1, "AAA"},
		{27*26 + 2, "AAB"},
		{(27*26+1)*26 - 1, "ZZY"},
		{(27*26+1)*26 + 0, "ZZZ"},
		{(27*26+1)*26 + 1, "AAAA"},
		{(27*26+1)*26 + 2, "AAAB"},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.num), func(t *testing.T) {
			if id := colNumToID(tt.num); id != tt.id {
				t.Fatalf("got %q, want %q", id, tt.id)
			}
		})
	}
}
