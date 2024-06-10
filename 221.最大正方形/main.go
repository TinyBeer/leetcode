package main

/*

在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
*/

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	explore := func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			return 0
		}
		return dp[x][y]
	}
	ans := 0
	for i, arr := range matrix {
		for j, v := range arr {
			if v == '1' {
				dp[i][j] = min(explore(i-1, j), explore(i, j-1), explore(i-1, j-1)) + 1
				ans = max(ans, dp[i][j]*dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ans
}
