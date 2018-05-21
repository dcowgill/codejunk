package dcp011

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestAutocomplete(t *testing.T) {
	var tests = []struct {
		dict    []string // input dictionary
		term    string   // search term
		matches []string // autocomplete results
	}{
		{[]string{"dog", "deer", "deal"}, "de", []string{"deal", "deer"}},
		{[]string{"aaa", "bbb", "ccc"}, "d", nil},
		{[]string{"aaa", "bbb", "ccc"}, "", nil},
		{[]string{"a", "aa", "aaa"}, "a", []string{"a", "aa", "aaa"}},
		{[]string{"你好", "世界", "你好世界", "你界", "世界你好"}, "你好", []string{"你好", "你好世界"}},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			matches := autocomplete(buildTrie(tt.dict), tt.term)
			sort.Strings(matches)
			sort.Strings(tt.matches)
			if !reflect.DeepEqual(matches, tt.matches) {
				t.Fatalf("got %+v, want %+v", matches, tt.matches)
			}
		})
	}
}
