package intset64

import "sort"

type Range struct{ low, high, step int }

// Uses a bitset.
func applyBitSet(rs []Range, low, high int) []int {
	var set intSet
	for _, r := range rs {
		l, h := low, high
		if r.low > l {
			l = r.low
		}
		if r.high < h {
			h = r.high
		}
		for i := l; i <= h; i += r.step {
			set.add(i)
		}
	}
	a := make([]int, 0, set.count())
	for i := 0; i < 64; i++ {
		if set.contains(i) {
			a = append(a, i)
		}
	}
	return a
}

// Uses an O(N) set.
func applyLinearSet(rs []Range, low, high int) []int {
	set := make(linearIntSet, 0)
	for _, r := range rs {
		l, h := low, high
		if r.low > l {
			l = r.low
		}
		if r.high < h {
			h = r.high
		}
		for i := l; i <= h; i += r.step {
			set.add(i)
		}
	}
	a := make([]int, 0, set.count())
	for i := low; i <= high; i++ {
		if set.contains(i) {
			a = append(a, i)
		}
	}
	return a
}

// Uses a sorted O(N) set.
func applySortedSet(rs []Range, low, high int) []int {
	set := make(sortedIntSet, 0)
	for _, r := range rs {
		l, h := low, high
		if r.low > l {
			l = r.low
		}
		if r.high < h {
			h = r.high
		}
		for i := l; i <= h; i += r.step {
			set.add(i)
		}
	}
	a := make([]int, 0, set.count())
	for i := low; i <= high; i++ {
		if set.contains(i) {
			a = append(a, i)
		}
	}
	return a
}

// Uses sorting.
func applySortUnique(rs []Range, low, high int) []int {
	a := []int{}
	for _, r := range rs {
		l, h := low, high
		if r.low > l {
			l = r.low
		}
		if r.high < h {
			h = r.high
		}
		for i := l; i <= h; i += r.step {
			a = append(a, i)
		}
	}
	if len(a) == 0 {
		return a
	}
	sort.Ints(a)
	var b []int
	prev := a[0]
	b = append(b, prev)
	for i := 1; i < len(a); i++ {
		if a[i] != prev {
			prev = a[i]
			b = append(b, prev)
		}
	}
	return b
}

type empty struct{}

// Uses a map.
func applyMap(rs []Range, low, high int) []int {
	set := make(map[int]empty)
	for _, r := range rs {
		l, h := low, high
		if r.low > l {
			l = r.low
		}
		if r.high < h {
			h = r.high
		}
		for i := l; i <= h; i += r.step {
			set[i] = empty{}
		}
	}
	if len(set) == 0 {
		return []int{}
	}
	a := make([]int, 0, len(set))
	for k := range set {
		a = append(a, k)
	}
	sort.Ints(a)
	return a
}

// Integer set that only supports elements in the range [0, 63].
type intSet int64

func (x *intSet) add(n int)           { *x |= (1 << uint(n)) }
func (x *intSet) contains(n int) bool { return *x&(1<<uint(n)) != 0 }
func (x *intSet) count() int {
	n := 0
	y := *x
	for y != 0 {
		y &= (y - 1)
		n++
	}
	return n

}

type linearIntSet []int

func (x *linearIntSet) add(n int) {
	if !x.contains(n) {
		*x = append(*x, n)
	}
}
func (x *linearIntSet) contains(n int) bool {
	for _, m := range *x {
		if n == m {
			return true
		}
	}
	return false
}
func (x *linearIntSet) count() int { return len(*x) }

type sortedIntSet []int

func (x *sortedIntSet) add(n int) {
	i := 0
	for i < len(*x) && n >= (*x)[i] {
		if n == (*x)[i] {
			return
		}
		i++
	}
	*x = append(*x, 0)
	for k := len(*x) - 1; k > i; k-- {
		(*x)[k] = (*x)[k-1]
	}
	(*x)[i] = n
}
func (x *sortedIntSet) contains(n int) bool {
	for _, m := range *x {
		if n == m {
			return true
		}
		if n < m {
			break
		}
	}
	return false
}
func (x *sortedIntSet) count() int {
	return len(*x)
}
