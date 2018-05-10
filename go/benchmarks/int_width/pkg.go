package int_width

import (
	"math"
	"strconv"
)

func log10(n int) int {
	if n <= 0 {
		return 1
	}
	return int(math.Floor(math.Log10(float64(n)))) + 1
}

func itoa(n int) int {
	return len(strconv.Itoa(n))
}
