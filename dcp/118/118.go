/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

Given a sorted list of integers, square the elements and give the output in sorted order.

For example, given [-9, -2, 0, 2, 3], return [0, 4, 4, 9, 81].

*/
package dcp118

// Given a sorted array of integers, return a sorted array of their squares.
func sortedSquares(a []int) []int {
	// Find the first non-negative integer in a.
	i := 0
	for i < len(a) && a[i] < 0 {
		i++
	}

	// We can now treat a[:i-1] and a[i:] as two sorted arrays to be merged. The
	// first half is negative, though, so we must compare absolute values.
	result := make([]int, 0, len(a))
	j := i - 1
	for i < len(a) && j >= 0 {
		x, y := a[i], a[j]
		if x < -y {
			result = append(result, square(x))
			i++
		} else {
			result = append(result, square(y))
			j--
		}
	}

	// Append the remainder of whichever half did not get exhausted.
	for ; i < len(a); i++ {
		result = append(result, square(a[i]))
	}
	for ; j >= 0; j-- {
		result = append(result, square(a[j]))
	}

	// Done.
	return result
}

func square(x int) int { return x * x }
