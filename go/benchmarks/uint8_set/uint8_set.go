package uint8_set

type set1 []bool

func newSet1() set1             { return set1(make([]bool, 256)) }
func (s set1) set(n uint8)      { s[n] = true }
func (s set1) has(n uint8) bool { return s[n] }

type set2 [256]bool

func newSet2() *set2 {
	var a set2
	return &a
}

func (s *set2) set(n uint8)      { (*s)[n] = true }
func (s *set2) has(n uint8) bool { return (*s)[n] }

type set3 [4]uint64

func newSet3() *set3 {
	var a set3
	return &a
}

func (s *set3) set(n uint8) {
	switch {
	case n >= 192:
		(*s)[3] |= 1 << (n - 192)
	case n >= 128:
		(*s)[2] |= 1 << (n - 128)
	case n >= 64:
		(*s)[1] |= 1 << (n - 64)
	default:
		(*s)[0] |= 1 << n
	}
}

func (s *set3) has(n uint8) bool {
	switch {
	case n >= 192:
		return ((*s)[3] & (1 << (n - 192))) != 0
	case n >= 128:
		return ((*s)[2] & (1 << (n - 128))) != 0
	case n >= 64:
		return ((*s)[1] & (1 << (n - 64))) != 0
	default:
		return ((*s)[0] & (1 << n)) != 0
	}
}

type set4 []uint64

func newSet4() *set4 {
	a := set4(make([]uint64, 4))
	return &a
}

func (s *set4) set(n uint8) {
	for i := 3; ; i-- {
		if d := uint8(64 * i); n >= d {
			(*s)[i] |= 1 << (n - d)
			return
		}
	}
}

func (s *set4) has(n uint8) bool {
	for i := 3; ; i-- {
		if d := uint8(64 * i); n >= d {
			return ((*s)[i] & (1 << (n - d))) != 0
		}
	}
}

type set5 map[uint8]bool

func newSet5() set5             { return set5(make(map[uint8]bool)) }
func (s set5) set(n uint8)      { s[n] = true }
func (s set5) has(n uint8) bool { return s[n] }
