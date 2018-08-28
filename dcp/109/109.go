/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Cisco.

Given an unsigned 8-bit integer, swap its even and odd bits. The 1st and 2nd bit
should be swapped, the 3rd and 4th bit should be swapped, and so on.

For example, 10101010 should be 01010101. 11100010 should be 11010001.

Bonus: Can you do this in one line?

*/
package dcp109

func swapBitPairs(x uint8) uint8 {
	const (
		b10101010 = 0xaa
		b01010101 = 0x55
	)
	return (x&b10101010)>>1 | (x&b01010101)<<1
}
