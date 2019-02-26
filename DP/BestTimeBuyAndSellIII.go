package main

/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
*/

//k-th transaction on i-th day
//dp[k,i]=max(dp[k,i-1],prices[i]-prices[j]+dp[k-1,j-1]) (j=[0...i-1])
//prices[i] - prices[j] + dp[k-1, j-1]: bought it on j-th day where j=[0..i-1], then sell it on i-th day

func maxProfitIII(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	dp := make([][]int, 3, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, len(prices), len(prices))
	}
	for k := 1; k < 3; k++ {
		minPrice := prices[0]
		for i := 1; i < len(prices); i++ {
			minPrice = min(minPrice, prices[i]-dp[k-1][i-1])
			dp[k][i] = max(dp[k][i-1], prices[i]-minPrice)
		}
	}
	return dp[2][len(prices)-1]
}
