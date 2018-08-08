/*

Good morning! Here's your coding interview problem for today.

This question was asked by ContextLogic.

Implement division of two positive integers without using the division,
multiplication, or modulus operators. Return the quotient as an integer,
ignoring the remainder.

*/
package dcp088

// O(a)
func slowdiv(a, b int32) int32 {
	var q int32
	for a >= b {
		a -= b
		q++
	}
	return q
}

// Cribbed from https://en.wikipedia.org/wiki/Division_algorithm
func fastdiv(a, b int32) int32 {
	var (
		r int64 = int64(a)
		d int64 = int64(b) << 32
		q int32
	)
	for i := 31; i >= 0; i-- {
		r = r<<1 - d
		if r >= 0 {
			q |= 1 << uint32(i)
		} else {
			r += d
		}
	}
	return q
}
