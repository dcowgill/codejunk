/*

Good morning! Here's your coding interview problem for today.

Given a list of numbers, return whether any two sums to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?

*/
package dcp001

func anyPairSumToMap(a []int, k int) bool {
	type empty struct{}
	m := make(map[int]empty)
	for _, x := range a {
		if _, ok := m[k-x]; ok {
			return true
		}
		m[x] = empty{}
	}
	return false
}

func anyPairSumToNestedLoops(a []int, k int) bool {
	for i, x := range a {
		for j := i + 1; j < len(a); j++ {
			if x+a[j] == k {
				return true
			}
		}
	}
	return false
}
