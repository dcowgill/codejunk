/*

This problem was asked by Google.

Given an array of strictly the characters 'R', 'G', and 'B', segregate the
values of the array so that all the Rs come first, the Gs come second, and the
Bs come last. You can only swap elements of the array.

Do this in linear time and in-place.

For example, given the array ['G', 'B', 'R', 'R', 'B', 'R', 'G'], it should
become ['R', 'R', 'R', 'G', 'G', 'B', 'B'].

*/
package dcp035

func sortRGB(a []rune) {
	var (
		i = 0
		j = 0
		k = len(a) - 1
	)
	for j <= k {
		switch a[j] {
		case 'R': // low
			a[i], a[j] = a[j], a[i]
			i++
			j++
		case 'G': // middle
			j++
		case 'B': // high
			a[k], a[j] = a[j], a[k]
			k--
		}
	}
}
