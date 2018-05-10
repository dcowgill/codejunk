package sets

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/omakasecorp/samurai/pkg/util"
)

// goal: return set(xs) - set(ys)

type empty struct{}

func uniqMapStable(xs []string) []string {
	m := make(map[string]struct{})
	var ys []string
	for _, x := range xs {
		if _, ok := m[x]; !ok {
			m[x] = struct{}{}
			ys = append(ys, x)
		}
	}
	return ys
}

func toSet(xs []string) map[string]empty {
	m := make(map[string]empty, len(xs))
	for _, x := range xs {
		m[x] = empty{}
	}
	return m
}

func minusMapStable(xs, ys []string) []string {
	ym := toSet(ys)
	var a []string
	for _, x := range uniqMapStable(xs) {
		if _, ok := ym[x]; !ok {
			a = append(a, x)
		}
	}
	return a
}

func minusList(xs, ys []string) []string {
	var a []string
outer:
	for _, x := range xs {
		for _, y := range ys {
			if x == y {
				continue outer
			}
		}
		for _, y := range a {
			if x == y {
				continue outer
			}
		}
		a = append(a, x)
	}
	return a
}

var (
	rand1, rand10, rand20, rand50, rand100, rand1000 []string
	same1, same10, same20, same50, same100, same1000 []string
	fill1, fill10, fill20, fill50, fill100, fill1000 []string
)

func randString() string {
	const n = 50
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = byte(rand.Intn(128)) // ASCII
	}
	return string(s)
}

func init() {
	rand1 = []string{randString()}
	same1 = []string{randString()}
	rand10 = make([]string, 10)
	same10 = make([]string, 10)
	for i := 0; i < 10; i++ {
		rand10[i] = randString()
		same10[i] = rand10[0]
	}
	rand20 = make([]string, 20)
	same20 = make([]string, 20)
	for i := 0; i < 20; i++ {
		rand20[i] = randString()
		same20[i] = rand20[0]
	}
	rand50 = make([]string, 50)
	same50 = make([]string, 50)
	for i := 0; i < 50; i++ {
		rand50[i] = randString()
		same50[i] = rand50[0]
	}
	rand100 = make([]string, 100)
	same100 = make([]string, 100)
	for i := 0; i < 100; i++ {
		rand100[i] = randString()
		same100[i] = rand100[0]
	}
	rand1000 = make([]string, 1000)
	same1000 = make([]string, 1000)
	for i := 0; i < 1000; i++ {
		rand1000[i] = randString()
		same1000[i] = rand1000[0]
	}

	source := rand10[:5]
	fill1 = repeatFill(source, 1)
	fill10 = repeatFill(source, 10)
	fill20 = repeatFill(source, 20)
	fill50 = repeatFill(source, 50)
	fill100 = repeatFill(source, 100)
	fill1000 = repeatFill(source, 1000)
}

func repeatFill(source []string, n int) []string {
	xs := make([]string, 0, n)
	k := len(source)
	for i := 0; i < n/k; i++ {
		for _, s := range source {
			xs = append(xs, s)
		}
	}
	for len(xs) < cap(xs) {
		xs = append(xs, source[0])
	}
	return xs
}

//--------------------

func shuffle(a sort.Interface) {
	for i := a.Len() - 1; i >= 1; i-- {
		a.Swap(i, rand.Intn(i+1))
	}
}

func copyAndShuffle(xs []string) []string {
	ys := make([]string, len(xs))
	copy(ys, xs)
	shuffle(sort.StringSlice(ys))
	return ys
}

func TestAll(t *testing.T) {
	var funcs = []struct {
		name string
		fn   func(xs, ys []string) []string
	}{
		{"list", minusList},
		{"mapStable", minusMapStable},
	}
	var cases = []struct {
		name   string
		xs, ys []string
	}{
		{"rand20 - rand10", rand20, rand10},
		{"rand50 - rand20", rand50, rand20},
		{"rand100 - rand50", rand100, rand50},
	}
	for _, fn := range funcs {
		for _, cs := range cases {
			t.Run(fmt.Sprintf("%s(%s)", fn.name, cs.name), func(t *testing.T) {
				xs := copyAndShuffle(cs.xs)
				ys := copyAndShuffle(cs.ys)
				correct := minusList(xs, ys)
				result := fn.fn(xs, ys)
				if !reflect.DeepEqual(result, correct) {
					fmt.Println(util.Dumps(correct, true))
					fmt.Println(util.Dumps(result, true))
					t.Fatalf("%s returns wrong result for test case %s", fn.name, cs.name)
				}
			})
		}
	}
}

func BenchmarkAll(b *testing.B) {
	var funcs = []struct {
		name string
		fn   func(xs, ys []string) []string
	}{
		{"list", minusList},
		{"mapStable", minusMapStable},
	}
	var cases = []struct {
		name   string
		xs, ys []string
	}{
		{"rand20 - rand10", rand20, rand10},
		{"rand50 - rand20", rand50, rand20},
		{"rand100 - rand50", rand100, rand50},
		{"rand50 - rand50", rand50, rand50},
		{"rand100 - rand100", rand100, rand100},
	}
	for _, fn := range funcs {
		for _, cs := range cases {
			b.Run(fmt.Sprintf("%s(%s)", fn.name, cs.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					fn.fn(cs.xs, cs.ys)
				}
			})
		}
	}
}
