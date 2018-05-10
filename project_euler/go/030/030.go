/*

Surprisingly there are only three numbers that can be written as the sum of
fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4⁴ ⁼
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers
of their digits.

*/

package p030

func exp(a, b int) int {
	p := 1
	for i := 0; i < b; i++ {
		p *= a
	}
	return p
}

// Calls f for each digit in n. For example, given n=1234, then f(1), f(2),
// f(3), and f(4) will be called in that order
func foreachDigit(n int, f func(digit int)) {
	for n > 0 {
		f(n % 10)
		n /= 10
	}
}

func sumExpDigits(n, e int) int {
	sum := 0
	foreachDigit(n, func(d int) {
		sum += exp(d, e)
	})
	return sum
}

func solve() int {
	// FIXME correct upper bound
	sum := 0
	for i := 10; i < 50000000; i++ {
		if sumExpDigits(i, 5) == i {
			sum += i
		}
	}
	return sum
}
