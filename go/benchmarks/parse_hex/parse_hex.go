package pkg

import "strconv"

func parseStdlib(s string) int64 {
	i, _ := strconv.ParseInt(s, 16, 64)
	return i
}

func parseHexCustom(s string) int64 {
	var n, p uint64
	for i := len(s) - 1; i >= 0; i-- {
		var x uint64
		switch b := s[i]; {
		case b >= '0' && b <= '9':
			x = uint64(b - '0')
		case b >= 'A' && b <= 'F':
			x = uint64(b - 'A' + 10)
		case b >= 'a' && b <= 'f':
			x = uint64(b - 'a' + 10)
		}
		n += x << p
		p += 4
	}
	return int64(n)
}

var hexByteTab = []uint64{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15,
	'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15,
}

func parseHexByteLookup(s string) int64 {
	var n, p uint64
	for i := len(s) - 1; i >= 0; i-- {
		n += hexByteTab[s[i]] << p
		p += 4
	}
	return int64(n)
}
