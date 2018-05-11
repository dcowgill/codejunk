/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Uber.

Given an array of integers, return a new array such that each element at index i
of the new array is the product of all the numbers in the original array except
the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be
[120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be
[2, 3, 6].

Follow-up: what if you can't use division?

*/
package dcp002

// Brute-force approach.
func quadratic(a []int) []int {
	b := make([]int, len(a))
	for i := range a {
		p := 1
		for _, y := range a[:i] {
			p *= y
		}
		for _, y := range a[i+1:] {
			p *= y
		}
		b[i] = p
	}
	return b
}

// Single pass using division.
func linearWithDivision(a []int) []int {
	p := 1
	for _, x := range a {
		p *= x
	}
	b := make([]int, len(a))
	for i, x := range a {
		b[i] = p / x
	}
	return b
}

// Two passes without using division.
func linearNoDivision(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	b := make([]int, len(a))
	p := 1
	b[0] = p
	for i := 1; i < len(a); i++ {
		p *= a[i-1]
		b[i] = p
	}
	p = 1
	for i := len(a) - 2; i >= 0; i-- {
		p *= a[i+1]
		b[i] *= p
	}
	return b
}
