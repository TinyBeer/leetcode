package main

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	df := make([]int, len(prices))
	df[0] = 0
	profit, cur := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if cur < prices[i] {
			profit = max(profit, prices[i]-cur)
		} else {
			cur = prices[i]
		}
		df[i] = profit
	}
	ans := 0
	cur = prices[len(prices)-1]
	for j := len(prices) - 2; j >= 0; j-- {
		if cur > prices[j] {
			ans = max(ans, cur-prices[j]+df[j])
		} else {
			cur = prices[j]
		}
	}
	return ans
}
