package unique_slice

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func shuffle(a sort.Interface) {
	for i := a.Len() - 1; i >= 1; i-- {
		a.Swap(i, rand.Intn(i+1))
	}
}

func uniqSqueeze(xs []string) []string {
	end := len(xs) - 1
	for i := 0; i < end; i++ {
		for j := i + 1; j <= end; j++ {
			if xs[i] == xs[j] {
				xs[j] = xs[end]
				xs = xs[:end]
				end--
				j--
			}
		}
	}
	return xs
}

func uniqMap(xs []string) []string {
	m := make(map[string]int)
	i := 0
	for _, x := range xs {
		if _, ok := m[x]; !ok {
			m[x] = i
			i++
		}
	}
	ys := make([]string, len(m))
	for s, i := range m {
		ys[i] = s
	}
	return ys
}

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

func uniqList(xs []string) []string {
	// var ys []string
	ys := make([]string, 0, len(xs))
outer:
	for _, x := range xs {
		for _, y := range ys {
			if x == y {
				continue outer
			}
		}
		ys = append(ys, x)
	}
	return ys
}

func uniqSmart(xs []string) []string {
	if len(xs) <= 100 {
		return uniqList(xs)
	}
	return uniqMapStable(xs)
}

func uniqExtraSmart(xs []string) []string {
	if len(xs) <= 1 {
		return xs
	}
	if len(xs) <= 100 {
		for i, x := range xs {
			for j := i + 1; j < len(xs); j++ {
				if x == xs[j] {
					goto notUnique
				}
			}
		}
		return xs
	notUnique:
		return uniqList(xs)
	}
	return uniqMapStable(xs)
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

func TestAll(t *testing.T) {
	var funcs = []struct {
		name string
		fn   func([]string) []string
	}{
		{"list", uniqList},
		{"mapStable", uniqMapStable},
		{"mapIndex", uniqMap},
		{"smart", uniqSmart},
		// {"extraSmart", uniqExtraSmart},
		// {"squeeze", uniqSqueeze},
	}
	var cases = []struct {
		name string
		xs   []string
	}{
		{"rand1", rand1},
		{"rand10", rand10},
		{"rand20", rand20},
		{"rand50", rand50},
		{"rand100", rand100},
		{"rand1000", rand1000},
		{"same1", same1},
		{"same10", same10},
		{"same20", same20},
		{"same50", same50},
		{"same100", same100},
		{"same1000", same1000},
		{"fill1", fill1},
		{"fill10", fill10},
		{"fill20", fill20},
		{"fill50", fill50},
		{"fill100", fill100},
		{"fill1000", fill1000},
	}
	for _, fn := range funcs {
		for _, cs := range cases {
			t.Run(fmt.Sprintf("%s(%s)", fn.name, cs.name), func(t *testing.T) {
				xs := make([]string, len(cs.xs))
				copy(xs, cs.xs)
				shuffle(sort.StringSlice(xs))
				correct := uniqList(xs)
				result := fn.fn(xs)
				if !reflect.DeepEqual(result, correct) {
					t.Errorf("%s returns wrong result for test case %s", fn.name, cs.name)
				}
			})
		}
	}
}

//--------------------

func BenchmarkUnique(b *testing.B) {
	var funcs = []struct {
		name string
		fn   func([]string) []string
	}{
		{"list", uniqList},
		{"mapStable", uniqMapStable},
		{"mapIndex", uniqMap},
		// {"smart", uniqSmart},
		// {"extraSmart", uniqExtraSmart},
		// {"squeeze", uniqSqueeze},
	}
	var cases = []struct {
		name string
		xs   []string
	}{
		// {"rand1", rand1},
		// {"rand10", rand10},
		// {"rand20", rand20},
		// {"rand50", rand50},
		// {"rand100", rand100},
		// {"rand1000", rand1000},
		// {"same1", same1},
		// {"same10", same10},
		// {"same20", same20},
		// {"same50", same50},
		// {"same100", same100},
		// {"same1000", same1000},
		// {"fill1", fill1},
		// {"fill10", fill10},
		{"fill20", fill20},
		{"fill50", fill50},
		{"fill100", fill100},
		// {"fill1000", fill1000},
	}
	for _, cs := range cases {
		for _, fn := range funcs {
			b.Run(fmt.Sprintf("%s(%s)", fn.name, cs.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					fn.fn(cs.xs)
				}
			})
		}
	}
}
