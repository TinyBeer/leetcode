package main

/*
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是质因子只包含 2、3 和 5 的正整数。
*/

func nthUglyNumber(n int) int {
	// if n == 1 {
	// 	return 1
	// }
	dp := make([]int, 0, n)
	dp = append(dp, 1)
	p2, p3, p5 := 0, 0, 0
	for len(dp) < n {
		next := min(dp[p2]*2, dp[p3]*3, dp[p5]*5)
		dp = append(dp, next)
		if next == dp[p2]*2 {
			p2++
		}
		if next == dp[p3]*3 {
			p3++
		}
		if next == dp[p5]*5 {
			p5++
		}
	}
	return dp[n-1]
}
