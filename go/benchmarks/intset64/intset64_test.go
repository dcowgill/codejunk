package intset64

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type Range struct{ low, high, step int }

var (
	ranges = []Range{
		{1, 9, 1},
		{1, 31, 4},
		{0, 63, 1},
		{0, 63, 3},
		{8, 9, 1},
	}
)

func (r Range) expand() []int {
	var a []int
	for i := r.low; i <= r.high; i += r.step {
		a = append(a, i)
	}
	return a
}

func (r Range) shuffled() []int {
	a := r.expand()
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func (r Range) String() string {
	return fmt.Sprintf("%d to %d by %d", r.low, r.high, r.step)
}

func TestBitSet(t *testing.T) {
	var set bitSet
	set = set.add(4)
	set = set.add(32)
	set = set.add(33)
	set = set.add(5)
	set = set.add(0)
	set = set.add(60)
	if want := 6; set.count() != want {
		t.Fatalf("set.count() returned %d, want %d", set.count(), want)
	}
	set = set.remove(32)
	if want := 5; set.count() != want {
		t.Fatalf("set.count() returned %d, want %d", set.count(), want)
	}
	{
		var a []int
		it := set.iter(63)
		for x := it.next(); x >= 0; x = it.next() {
			a = append(a, x)
		}
		if want := []int{0, 4, 5, 33, 60}; !reflect.DeepEqual(want, a) {
			t.Fatalf("iterator returned %+v, want %+v", a, want)
		}
	}
}

func TestBitSetCountIsExactlyOne(t *testing.T) {
	// Zero values
	var set bitSet
	if set.countIsExactlyOne() {
		t.Fatalf("bitSet(%v).countIsExactlyOne returned true, want false", set)
	}
	// One value
	for i := 0; i < 64; i++ {
		var set bitSet
		set = set.add(i)
		if !set.countIsExactlyOne() {
			t.Fatalf("bitSet(%v).countIsExactlyOne returned false, want true", set)
		}
	}
	// Two values
	for i := 0; i < 63; i++ {
		var set bitSet
		set = set.add(i).add(i + 1)
		if set.countIsExactlyOne() {
			t.Fatalf("bitSet(%v).countIsExactlyOne returned true, want false", set)
		}
	}
}

func TestLinearSet(t *testing.T) {
	set := newLinearIntSet()
	set = set.add(4)
	set = set.add(32)
	set = set.add(33)
	set = set.add(5)
	set = set.add(0)
	set = set.add(60)
	if want := 6; set.count() != want {
		t.Fatalf("set.count() returned %d, want %d", set.count(), want)
	}
	set = set.remove(32)
	if want := 5; set.count() != want {
		t.Fatalf("set.count() returned %d, want %d", set.count(), want)
	}
	var a []int
	it := set.iter()
	for x := it.next(); x >= 0; x = it.next() {
		a = append(a, x)
	}
	sort.Ints(a)
	if want := []int{0, 4, 5, 33, 60}; !reflect.DeepEqual(want, a) {
		t.Fatalf("iterator returned %+v, want %+v", a, want)
	}
}

var impls = []struct {
	name              string
	addRange          func(low, high, step int) int
	addAndRemoveRange func(add, remove []int) int
}{
	{"bits", addRangeBitSet, addAndRemoveRangeBitSet},
	{"slice", addRangeLinearSet, addAndRemoveRangeLinearSet},
	{"sorted", addRangeSortedSet, addAndRemoveRangeSortedSet},
	{"map", addRangeMapSet, addAndRemoveRangeMapSet},
}

func BenchmarkAddRange(b *testing.B) {
	for _, impl := range impls {
		for _, r := range ranges {
			b.Run(fmt.Sprintf("%s: %d to %d by %d", impl.name, r.low, r.high, r.step), func(b *testing.B) {
				x := 0
				for i := 0; i < b.N; i++ {
					x += impl.addRange(r.low, r.high, r.step)
				}
			})
		}
	}
}

func BenchmarkAddAndRemoveRange(b *testing.B) {
	for _, impl := range impls {
		for _, r := range ranges {
			add := r.expand()
			remove := r.shuffled()
			b.Run(fmt.Sprintf("%s: %s", impl.name, r.String()), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if impl.addAndRemoveRange(add, remove) != 0 {
						panic("expected zero count")
					}
				}
			})
		}
	}
}

func BenchmarkBitSetIterator(b *testing.B) {
	for _, r := range ranges {
		var set bitSet
		for _, x := range r.expand() {
			set = set.add(x)
		}
		n := 0
		b.Run(r.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := set.iter(r.high)
				for x := it.next(); x >= 0; x = it.next() {
					n++
				}
			}
		})
	}
}

func BenchmarkLinearSetIterator(b *testing.B) {
	for _, r := range ranges {
		set := newLinearIntSet()
		for _, x := range r.expand() {
			set = set.add(x)
		}
		n := 0
		b.Run(r.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := set.iter()
				for x := it.next(); x >= 0; x = it.next() {
					n += x
				}
			}
		})
	}
}

func BenchmarkLinearSetIteratorRangeLoop(b *testing.B) {
	for _, r := range ranges {
		set := newLinearIntSet()
		for _, x := range r.expand() {
			set = set.add(x)
		}
		n := 0
		b.Run(r.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, x := range set {
					n += x
				}
			}
		})
	}
}

func BenchmarkBitSetPtr(b *testing.B) {
	newMap := func() map[string]bitSet {
		return map[string]bitSet{
			"abc":                bitSet(0),
			"defghi":             bitSet(0),
			"jklmnopqr":          bitSet(0),
			"stuvwxyz1234567890": bitSet(0),
		}
	}
	newSlice := func() []bitSet {
		return make([]bitSet, 10)
	}
	xs := Range{0, 63, 3}.expand()
	b.Run("map/return", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := newMap()
			for k := range m {
				for _, x := range xs {
					m[k] = m[k].add(x)
				}
			}
		}
	})
	b.Run("map/pointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := newMap()
			for k := range m {
				for _, x := range xs {
					set := m[k]
					set.addptr(x)
					m[k] = set
				}
			}
		}
	})
	b.Run("slice/return", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := newSlice()
			for k := range a {
				for _, x := range xs {
					a[k] = a[k].add(x)
				}
			}
		}
	})
	b.Run("slice/pointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := newSlice()
			for k := range a {
				for _, x := range xs {
					a[k].addptr(x)
				}
			}
		}
	})
}
