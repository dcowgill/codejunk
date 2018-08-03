/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given three 32-bit integers x, y, and b, return x if b is 1 and y if b is 0,
using only mathematical or bit operations. You can assume b can only be 1 or 0.

*/
package dcp085

func bitty(x, y, b int32) int32 {
	// bb is 00000000 if b==0, 11111111 if b==1.
	bb := (b << 0) | (b << 1) | (b << 2) | (b << 3) | (b << 4) | (b << 5) | (b << 6) | (b << 7)
	xm := bb | bb<<8 | bb<<16 | bb<<24
	ym := ^xm
	return (x & xm) | (y & ym)
}
