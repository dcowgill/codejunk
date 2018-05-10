/*

In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:

1£1 + 150p + 220p + 15p + 12p + 31p
How many different ways can £2 be made using any number of coins?

*/

package p031

func combinations(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}
	n := 0
	for i := 0; ; i++ {
		newAmount := amount - i*coins[0]
		if newAmount < 0 {
			break
		}
		n += combinations(newAmount, coins[1:])
	}
	return n
}

func solve() int {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	return combinations(200, coins)
}
