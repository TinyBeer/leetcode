package main

/*
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
*/

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var ans int
	dir := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		cur := matrix[i][j]
		if dp[i][j] == 0 {
			big := 0
			for _, p := range dir {
				x, y := i+p[0], j+p[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				if matrix[x][y] > cur {
					big = max(big, dfs(x, y)+1)
				}
			}
			dp[i][j] = big
			return big
		}
		return dp[i][j]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans + 1
}
