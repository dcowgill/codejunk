package p019

// See http://www.tondering.dk/claus/cal/chrweek.php
func dayOfWeek(y, m, d int) int {
	return (d + y + y/4 - y/100 + y/400 + 31*m/12) % 7
}

func solve() int {
	n := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			if dayOfWeek(y, m, 1) == 0 {
				n++
			}
		}
	}
	return n
}
