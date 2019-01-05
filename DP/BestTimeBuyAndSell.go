package main

/*
309. Best Time to Buy and Sell Stock with Cooldown

Say you have an array for which the ith element is the price of a given stock on day i.
Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
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
	buy[0] = 0           //buy[i] means before day i what is the maxProfit for any sequence end with buy.
	sell[0] = -prices[0] //sell[i] means before day i what is the maxProfit for any sequence end with sell.

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
