/*

Starting with the number 1 and moving to the right in a clockwise direction a
5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral
formed in the same way?

*/

package p028

func sumOfDiagonalsInSpiral(n int) int {
	// In an n x n spiral as defined by this problem, the upper right
	// corner of the outermost ring is equal to n^2. Moving
	// counterclockwise, each corner is is (n-1) smaller than the previous
	// one, ∴ the sum of the corners is 4n^2 - 6n + 6

	// We have to sum that quadratic for every odd number between 3 and 1001.
	// Given that the sum of the first n odd numbers is n^2 and the sum of the
	// squares of the first n odd numbers is n(2n-1)(2n+1)/3

	// k = (n+1)/2
	// x = k^2
	// y = k(2k-1)(2k+1)/3
	// s = 4(y-1) - 6(x-1) + 6(k-1) + 1

	k := (n + 1) / 2
	x := k * k
	y := k * (2*k - 1) * (2*k + 1) / 3
	return 4*(y-1) - 6*(x-1) + 6*(k-1) + 1
}

func solve() int {
	return sumOfDiagonalsInSpiral(1001)
}
