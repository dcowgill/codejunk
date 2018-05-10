/*

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of
their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.

*/

package p034

func solve() int {
	facts := make([]int, 10)
	f := 1
	facts[0] = 1 // 0! = 1
	for i := 1; i <= 9; i++ {
		f *= i
		facts[i] = f
	}
	upper := facts[9] * 9
	total := 0
	for i := 10; i < upper; i++ {
		sum := 0
		for n := i; n > 0; n /= 10 {
			sum += facts[n%10]
		}
		if sum == i {
			total += sum
		}
	}
	return total
}
