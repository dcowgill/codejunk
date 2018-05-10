/*

2520 is the smallest number that can be divided by each of the numbers from 1
to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?

*/
package p005

const N = 20

func solve() int {
	var (
		factors = make([]int, 0, N)
		product = 1
	)

	for i := 2; i <= N; i++ {
		n := i
		for _, p := range factors {
			if n%p == 0 {
				n /= p
			}
		}
		if n > 1 {
			factors = append(factors, n)
			product *= n
		}
	}
	return product
}
