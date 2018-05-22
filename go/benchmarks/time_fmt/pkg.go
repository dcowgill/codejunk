package timefmt

import (
	"fmt"
	"time"
)

func Format_YYYY_MM(t time.Time) string { return t.Format("2006_01") }

func Printf_YYYY_MM(t time.Time) string { return fmt.Sprintf("%d_%02d", t.Year(), int(t.Month())) }

func Format_YYYY_MM_DD(t time.Time) string { return t.Format("2006-01-02") }

func Printf_YYYY_MM_DD_v1(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func Printf_YYYY_MM_DD_v2(t time.Time) string {
	a := []byte{'0', '0', '0', '0', '-', '0', '0', '-', '0', '0'}
	year, month, day := t.Date()
	writeInt(a[0:], year, 4)
	writeInt(a[5:], int(month), 2)
	writeInt(a[8:], day, 2)
	return string(a)
}

// Writes up to w bytes of the base-10 right-aligned string
// representation of n into dst. If fewer than w bytes are required, the
// leftmost bytes in dst are not overwritten. If more than w bytes, the
// most significant digits are truncated.
func writeInt(dst []byte, n, w int) {
	w--
	for n > 0 && w >= 0 {
		c := '0' + (n % 10)
		n /= 10
		dst[w] = byte(c)
		w--
	}
}
