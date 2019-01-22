package intset64

//
// bitSet
//

type bitSet uint64
type bitSetIter struct {
	set bitSet
	cur int
	end int
}

func (set *bitSet) addptr(x int)           { *set |= (1 << uint64(x)) }
func (set *bitSet) removeptr(x int)        { *set &= ^(1 << uint64(x)) }
func (set bitSet) add(x int) bitSet        { return set | (1 << uint64(x)) }
func (set bitSet) remove(x int) bitSet     { return set & ^(1 << uint64(x)) }
func (set bitSet) contains(x int) bool     { return set&(1<<uint64(x)) != 0 }
func (set bitSet) empty() bool             { return set == 0 }
func (set bitSet) countIsExactlyOne() bool { return set != 0 && (set&(set-1)) == 0 }
func (set bitSet) count() int {
	x := 0
	for set != 0 {
		set &= set - 1
		x++
	}
	return x
}
func (set bitSet) iter(high int) bitSetIter {
	return bitSetIter{set: set, end: high + 1}
}
func (it *bitSetIter) next() int {
	for it.cur < it.end {
		if it.set.contains(it.cur) {
			x := it.cur
			it.cur++
			return x
		}
		it.cur++
	}
	return -1
}

//
// linearIntSet
//

type linearIntSet []int
type linearIntSetIter struct {
	set linearIntSet
	pos int
}

func newLinearIntSet() linearIntSet { return make(linearIntSet, 0) }
func (set linearIntSet) add(x int) linearIntSet {
	for _, y := range set {
		if x == y {
			return set
		}
	}
	return append(set, x)
}
func (set linearIntSet) remove(x int) linearIntSet {
	for i, y := range set {
		if x == y {
			n := len(set)
			set[i], set[n-1] = set[n-1], set[i]
			return set[:n-1]
		}
	}
	return set
}
func (set linearIntSet) contains(x int) bool {
	for _, y := range set {
		if x == y {
			return true
		}
	}
	return false
}
func (set linearIntSet) count() int             { return len(set) }
func (set linearIntSet) iter() linearIntSetIter { return linearIntSetIter{set, 0} }
func (it *linearIntSetIter) next() int {
	if it.pos >= len(it.set) {
		return -1
	}
	x := it.set[it.pos]
	it.pos++
	return x
}

//
// sortedIntSet
//

type sortedIntSet []int

func newSortedIntSet() sortedIntSet { return make(sortedIntSet, 0) }
func (set sortedIntSet) add(x int) sortedIntSet {
	i := 0
	for i < len(set) && x >= set[i] {
		if x == set[i] {
			return set
		}
		i++
	}
	set = append(set, 0)
	for k := len(set) - 1; k > i; k-- {
		set[k] = set[k-1]
	}
	set[i] = x
	return set
}
func (set sortedIntSet) remove(x int) sortedIntSet {
	for i, y := range set {
		if x == y {
			n := len(set)
			for j := i + 1; j < n; j++ {
				set[j-1] = set[j]
			}
			return set[:n-1]
		}
		if x < y {
			break
		}
	}
	return set
}
func (set sortedIntSet) contains(x int) bool {
	for _, y := range set {
		if x == y {
			return true
		}
		if x < y {
			break
		}
	}
	return false
}
func (set sortedIntSet) count() int { return len(set) }

//
// mapIntSet
//

type empty struct{}

type mapIntSet map[int]empty

func newMapIntSet() mapIntSet             { return make(mapIntSet) }
func (set mapIntSet) add(x int)           { set[x] = empty{} }
func (set mapIntSet) remove(x int)        { delete(set, x) }
func (set mapIntSet) contains(x int) bool { _, ok := set[x]; return ok }
func (set mapIntSet) count() int          { return len(set) }

//====================
// test functions
//====================

func addRangeBitSet(low, high, step int) int {
	var set bitSet
	for i := low; i <= high; i += step {
		set = set.add(i)
	}
	return set.count()
}

func addRangeLinearSet(low, high, step int) int {
	set := newLinearIntSet()
	for i := low; i <= high; i += step {
		set = set.add(i)
	}
	return set.count()
}

func addRangeSortedSet(low, high, step int) int {
	set := newSortedIntSet()
	for i := low; i <= high; i += step {
		set = set.add(i)
	}
	return set.count()
}

func addRangeMapSet(low, high, step int) int {
	set := newMapIntSet()
	for i := low; i <= high; i += step {
		set.add(i)
	}
	return set.count()
}

func addAndRemoveRangeBitSet(add, remove []int) int {
	var set bitSet
	for _, x := range add {
		set = set.add(x)
	}
	for _, x := range remove {
		set = set.remove(x)
	}
	return set.count()
}

func addAndRemoveRangeLinearSet(add, remove []int) int {
	set := newLinearIntSet()
	for _, x := range add {
		set = set.add(x)
	}
	for _, x := range remove {
		set = set.remove(x)
	}
	return set.count()
}

func addAndRemoveRangeSortedSet(add, remove []int) int {
	set := newSortedIntSet()
	for _, x := range add {
		set = set.add(x)
	}
	for _, x := range remove {
		set = set.remove(x)
	}
	return set.count()
}

func addAndRemoveRangeMapSet(add, remove []int) int {
	set := newMapIntSet()
	for _, x := range add {
		set.add(x)
	}
	for _, x := range remove {
		set.remove(x)
	}
	return set.count()
}
