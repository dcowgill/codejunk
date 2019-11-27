package main

import (
	"math/rand"
	"sort"
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

func uniqListWithEqFunc(xs []string, eq func(a, b string) bool) []string {
	ys := make([]string, 0, len(xs))
outer:
	for _, x := range xs {
		for _, y := range ys {
			if eq(x, y) {
				continue outer
			}
		}
		ys = append(ys, x)
	}
	return ys
}

func uniqListWithEqFuncOp(xs []string) []string {
	return uniqListWithEqFunc(xs, func(a, b string) bool { return a == b })
}

func uniqSort(xs []string) []string {
	if len(xs) < 2 {
		return xs
	}
	ys := make([]string, len(xs))
	copy(ys, xs)
	sort.Strings(ys)
	j := 1
	for i := 1; i < len(ys); i++ {
		if ys[i] != ys[i-1] {
			ys[j] = ys[i]
			j++
		}
	}
	return ys[:j]
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
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const n = 24
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = chars[rand.Intn(len(chars))]
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
	fill1 = shuffled(repeatFill(source, 1))
	fill10 = shuffled(repeatFill(source, 10))
	fill20 = shuffled(repeatFill(source, 20))
	fill50 = shuffled(repeatFill(source, 50))
	fill100 = shuffled(repeatFill(source, 100))
	fill1000 = shuffled(repeatFill(source, 1000))
}

// Shuffles the slice in-place. Returns it as a convenience.
func shuffled(a []string) []string {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

// Returns a n-length slice of n strings containing the strings in source,
// repeated until the target length is reached.
func repeatFill(source []string, n int) []string {
	xs := make([]string, 0, n)
	k := len(source)
	for i := 0; i < n/k; i++ {
		xs = append(xs, source...)
	}
	i := 0
	for len(xs) < cap(xs) {
		xs = append(xs, source[i])
		i = (i + 1) % len(source)
	}
	return xs
}

func main() {
	// data sets:
	//	N random
	//	N random values, repeated K times
	//	N identical values

	// for each kind of data set
	//	for each N in [1, ... K]
	//		for each algorithm
	//			print times
	//			print number of allocs
	//			print number of bytes allocated
}
