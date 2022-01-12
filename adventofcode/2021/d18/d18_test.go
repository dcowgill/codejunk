package d18

import (
	"adventofcode2021/lib"
	"fmt"
	"testing"
)

func TestParsing(t *testing.T) {
	var tests = []string{
		"[[[[[9,8],1],2],3],4]",
		"[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]",
		"[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}
	for _, s := range tests {
		root := parseNum(s)
		if root.String() != s {
			t.Fatalf("parseNum(%q).String() returned %q", s, root)
		}
	}
}

func TestExplode(t *testing.T) {
	var tests = []struct {
		input, output string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]"},
		{"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[15,[0,13]]],[1,1]]"},
		{"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}
	for _, tt := range tests {
		n := parseNum(tt.input)
		if ok := explode(n); !ok {
			t.Fatalf("explode(%q) returned false, want true", tt.input)
		}
		if n.String() != tt.output {
			t.Fatalf("explode(%q) produced %q, want %q", tt.input, n, tt.output)
		}
	}
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		input, output string
	}{
		{"[[[[0,7],4],[15,[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
		{"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"},
	}
	for _, tt := range tests {
		root := parseNum(tt.input)
		if ok := split(root); !ok {
			t.Fatalf("split(%q) returned false, want true", tt.input)
		}
		if root.String() != tt.output {
			t.Fatalf("split(%q) produced %q, want %q", tt.input, root, tt.output)
		}
	}
}

func TestReduce(t *testing.T) {
	var tests = []struct {
		input, output string
	}{
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}
	for _, tt := range tests {
		root := parseNum(tt.input)
		reduce(root)
		if root.String() != tt.output {
			t.Fatalf("reduce(%q) produced %q, want %q", tt.input, root, tt.output)
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		addends []string
		sum     string
	}{
		{
			[]string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
			},
			"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			[]string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			},
			"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			[]string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			},
			"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			[]string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			},
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			[]string{
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]],",
			},
			"[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			addends := lib.Map(parseNum, tt.addends)
			sum := lib.Foldl(add, addends)
			if sum.String() != tt.sum {
				t.Fatalf("got %q, want %q", sum, tt.sum)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	var tests = []struct {
		input     string
		magnitude int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
		{"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", 4140},
		{"[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]", 3993},
	}
	for _, tt := range tests {
		v := magnitude(parseNum(tt.input))
		if v != tt.magnitude {
			t.Fatalf("magnitude(%q) returned %d, want %d", tt.input, v, tt.magnitude)
		}
	}
}
