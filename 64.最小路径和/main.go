package main

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
1 3 1
1 5 1
1 4 1
*/

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j != 0 {
					f[j] = f[j-1]
				}
			} else if j-1 >= 0 {
				f[j] = min(f[j-1], f[j])
			}
			f[j] += grid[i][j]
		}
	}
	return f[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
