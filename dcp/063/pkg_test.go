package dcp063

import (
	"strconv"
	"testing"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		mat  [][]rune // input matrix
		good []string // these words shoould result in a successful search
		bad  []string // ...and these should not
	}{
		{
			[][]rune{
				{'F', 'A', 'C', 'I'},
				{'O', 'B', 'Q', 'P'},
				{'A', 'N', 'O', 'B'},
				{'M', 'A', 'S', 'S'},
			},
			[]string{"FOAM", "MASS", "PBS", "ABN", "CQOS", "M"},
			[]string{"OAMA", "NACQ", "NOBM", "FOOM", "X"},
		},
		{
			[][]rune{
				{'快', '速', '的', '棕', '色'},
				{'狐', '狸', '跳', '了', '起'},
				{'来', '现', '在', '是', '所'},
				{'有', '好', '人', '的', '时'},
			},
			[]string{"速的棕", "在是所", "有好人的时", "速狸现好", "色", "狐"},
			[]string{"的棕速", "人了", "时的人"},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			corpus := mat2str(tt.mat) // preprocess
			for _, word := range tt.good {
				if indexRabinKarp(corpus, word) < 0 {
					t.Fatalf("failed to find %q", word)
				}
			}
			for _, word := range tt.bad {
				if indexRabinKarp(corpus, word) >= 0 {
					t.Fatalf("found %q but did not expect to", word)
				}
			}
		})
	}
}
