package main

/*
309. Best Time to Buy and Sell Stock with Cooldown
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]

*/
/*
s0[i] = max(s0[i - 1], s2[i - 1]); // Stay at s0, or rest from s2
s1[i] = max(s1[i - 1], s0[i - 1] - prices[i]); // Stay at s1, or buy from s0
s2[i] = s1[i - 1] + prices[i]; // Only one way from s1
*/
//State Machine
func bestTime_stateMachine(prices []int) int {

	if prices == nil || len(prices) < 2 {
		return 0
	}
	length := len(prices)
	buy := make([]int, length)
	sell := make([]int, length)
	cooldown := make([]int, length)
	buy[0] = 0           //At the start,no stock,
	sell[0] = -prices[0] // After buy, you should have -prices[0] profit,

	cooldown[0] = -999 //cooldown[i] means before day i what is the maxProfit for any sequence end with rest.
	for i := 1; i < length; i++ {
		buy[i] = max(buy[i-1], cooldown[i-1])
		sell[i] = max(sell[i-1], buy[i-1]-prices[i])
		cooldown[i] = sell[i-1] + prices[i]
	}
	return max(buy[length-1], cooldown[length-1])
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
