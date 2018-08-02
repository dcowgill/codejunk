package dcp022

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		words  []string
		term   string
		result []string
	}{
		{
			[]string{"quick", "brown", "the", "fox"},
			"thequickbrownfox",
			[]string{"the", "quick", "brown", "fox"},
		},
		{
			[]string{"bed", "bath", "bedbath", "and", "beyond"},
			"bedbathandbeyond",
			[]string{"bed", "bath", "and", "beyond"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("words=%+v term=%s", tt.words, tt.term), func(t *testing.T) {
			var root *Trie
			for _, w := range tt.words {
				root = root.insert([]rune(w))
			}
			result := search(root, tt.term)
			if !reflect.DeepEqual(result, tt.result) {
				t.Fatalf("search returned %+v, want %+v", result, tt.result)
			}
		})
	}
}
