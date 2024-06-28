package main

import "math"

/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，
其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
*/

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

// func numSquares(n int) int {
// 	dp := make([]int, n+1)

// 	for x, p := 0, 0; p <= n; x++ {
// 		for i := 0; i+p <= n; i++ {
// 			if x == 0 && i > 0 {
// 				dp[i] = n
// 			} else {
// 				dp[i+p] = min(dp[i+p], dp[i]+1)
// 			}
// 		}
// 		p = (x + 1) * (x + 1)
// 	}
// 	return dp[n]
// }
