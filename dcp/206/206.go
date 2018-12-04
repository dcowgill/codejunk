/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Twitter.

A permutation can be specified by an array P, where P[i] represents the location
of the element at i in the permutation. For example, [2, 1, 0] represents the
permutation where elements at the index 0 and 2 are swapped.

Given an array and a permutation, apply the permutation to the array. For
example, given the array ["a", "b", "c"] and the permutation [2, 1, 0], return
["c", "b", "a"].

*/
package dcp206

func applyPerm(a []string, p []int) []string {
	b := make([]string, len(a))
	for i, j := range p {
		b[i] = a[j]
	}
	return b
}

func applyPermInPlace(a []string, p []int) {
	// next returns the lowest i where p[i] >= 0
	next := func() int {
		for i, j := range p {
			if j >= 0 {
				return i
			}
		}
		return -1
	}
	i := 0
	for {
		if p[i] < 0 {
			i = next()
			if i < 0 {
				return
			}
		}
		j := p[i]
		if i == j || p[j] < 0 {
			p[i] = -1
			continue
		}
		a[i], a[j] = a[j], a[i]
		p[i] = -1
		i = j
	}
}
