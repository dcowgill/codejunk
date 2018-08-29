/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Lyft.

Given a list of integers and a number K, return which contiguous elements of the
list sum to K.

For example, if the list is [1, 2, 3, 4, 5] and K is 9, then it should return
[2, 3, 4].

*/
package dcp102

// Returns a sub-slice of a that sums to k, or nil if so such sub-slice exists.
// Returns nil if k is zero, since the empty slice sums to zero.
func subarraySum(a []int, k int) []int {
	if k == 0 || len(a) == 0 {
		return nil
	}
	i := 0
	j := 1
	sum := a[i]
	for {
		if sum == k {
			return a[i:j]
		}
		if j >= len(a) {
			return nil
		}
		sum += a[j]
		for sum > k && i <= j {
			sum -= a[i]
			i++
		}
		j++
	}
}
