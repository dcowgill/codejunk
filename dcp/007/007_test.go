package dcp007

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	var tests = []struct {
		digits string
		ss     []string
	}{
		{
			"111",
			[]string{"aaa", "ak", "ka"},
		},
		{
			"222",
			[]string{"bbb", "bv", "vb"},
		},
		{
			"812341569",
			[]string{"habcdaefi", "habcdofi", "hawdaefi", "hawdofi", "hlcdaefi", "hlcdofi"},
		},
		{
			"11223344",
			[]string{"aabbccdd", "aabwcdd", "aavccdd", "albccdd", "alwcdd", "kbbccdd", "kbwcdd", "kvccdd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.digits, func(t *testing.T) {
			ss := runesToStrings(decode(toDigits(tt.digits)))
			if !reflect.DeepEqual(ss, tt.ss) {
				t.Fatalf("got %+v, want %+v", ss, tt.ss)
			}
		})
	}
}
