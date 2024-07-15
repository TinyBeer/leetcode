package main

/*

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。
*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i >= c && dp[i-c] != -1 {
				if dp[i] == -1 {
					dp[i] = dp[i-c] + 1
				} else {
					dp[i] = min(dp[i-c]+1, dp[i])
				}
			}
		}
	}
	return dp[amount]
}
