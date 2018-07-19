/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Microsoft.

A number is considered perfect if its digits sum up to exactly 10.

Given a positive integer n, return the n-th perfect number.

For example, given 1, you should return 19. Given 2, you should return 28.

*/
package dcp070

// Note: this problem's definition of a "perfect number" is idiosyncratic:
// https://en.wikipedia.org/wiki/Perfect_number

// Returns the nth number whose digits sum to 10.
func nth(n int) int {
	i := 0
	for x := 19; ; x += 9 {
		if digitsSumTo10(x) {
			i++
			if i == n {
				return x
			}
		}
	}
}

// Reports whether the digits of n sum to 10.
func digitsSumTo10(n int) bool {
	sum := 0
	for n > 0 && sum <= 10 {
		sum += n % 10
		n /= 10
	}
	return sum == 10
}
