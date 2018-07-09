/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

We can determine how "out of order" an array A is by counting the number of
inversions it has. Two elements A[i] and A[j] form an inversion if A[i] > A[j]
but i < j. That is, a smaller element appears after a larger element.

Given an array, count the number of inversions it has. Do this faster than
O(N^2) time.

You may assume each element in the array is distinct.

For example, a sorted list has zero inversions. The array [2, 4, 1, 3, 5] has
three inversions: (2, 1), (4, 1), and (4, 3). The array [5, 4, 3, 2, 1] has ten
inversions: every distinct pair forms an inversion.

*/
package dcp044

func numInversions(a []int) int {
	return splitMerge(copyInts(a), a) // throw away temporary storage
}

// Mergesorts src into dst. Returns count of pairwise inversions.
func splitMerge(dst, src []int) int {
	if len(src) < 2 {
		return 0
	}
	n := 0
	m := len(src) / 2
	n += splitMerge(src[:m], dst[:m])
	n += splitMerge(src[m:], dst[m:])
	n += merge(dst, src, 0, m, len(dst))
	return n
}

// Merges a[x..y] and a[y..z] into b.
// Returns the number of pairwise inversions.
// Assumption: both subranges are sorted.
func merge(b []int, a []int, x, y, z int) int {
	i := x
	j := y
	k := 0
	n := 0
	for {
		l := a[i]
		r := a[j]
		if l < r {
			b[k] = l
			i++
			if i == y {
				copy(b[k+1:], a[j:z])
				return n
			}
		} else {
			b[k] = r
			j++
			n += y - i // one inversion per lefthand value
			if j == z {
				copy(b[k+1:], a[i:y])
				return n
			}
		}
		k++
	}
}

func copyInts(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
