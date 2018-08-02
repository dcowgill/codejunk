/*

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a array of numbers representing the stock prices of a company in
chronological order, write a function that calculates the maximum profit you
could have made from buying and selling that stock once. You must buy before you
can sell it.

For example, given [9, 11, 8, 5, 7, 10], you should return 5, since you could
buy the stock at 5 dollars and sell it at 10 dollars.

*/
package dcp047

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	maxsofar := prices[n-1]
	minsofar := maxsofar
	best := 0
	for i := n - 1; i >= 0; i-- {
		if prices[i] > maxsofar {
			maxsofar = prices[i]
			minsofar = maxsofar
		} else if prices[i] < minsofar {
			minsofar = prices[i]
		}
		best = max(best, maxsofar-minsofar)
	}
	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
